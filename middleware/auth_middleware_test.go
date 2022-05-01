package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	shrd_helper "github.com/StevanoZ/dv-shared/shared/helper"
	shrd_token "github.com/StevanoZ/dv-shared/shared/token"
	shrd_utils "github.com/StevanoZ/dv-shared/shared/utils"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	UNAUTHORIZE = "unauthorize"
	FORBIDDEN   = "forbidden"

	TOKEN   = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	SUCCESS = "success"
)

func testHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(SUCCESS))
	})
}
func initAuthMiddleware(ctrl *gomock.Controller) (AuthMiddleware, *shrd_token.MockMaker) {
	token := shrd_token.NewMockMaker(ctrl)
	authMiddleware := NewAuthMiddleware(token)

	return authMiddleware, token
}

func createTokenPayload(status string) *shrd_token.Payload {
	return &shrd_token.Payload{ID: uuid.New(),
		UserId:   uuid.New(),
		Email:    shrd_utils.RandomEmail(),
		Status:   status,
		IssuedAt: time.Now()}
}

func TestAuthMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authMiddleware, tokenMaker := initAuthMiddleware(ctrl)
	handler := authMiddleware.CheckIsAuthenticated(testHandler())

	t.Run("Success Request (passed auth middleware)", func(t *testing.T) {
		recorder, req := setUpRecorder()
		shrd_helper.SetHeaderApplicationJson(req)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", TOKEN))
		tokenMaker.EXPECT().VerifyToken(TOKEN).Return(createTokenPayload("active"), nil).Times(1)

		assert.NotPanics(t, func() {
			handler.ServeHTTP(recorder, req)

			assert.Equal(t, SUCCESS, recorder.Body.String())
		})
	})
	t.Run("Do not carry token (status code 401)", func(t *testing.T) {
		recorder, req := setUpRecorder()
		shrd_helper.SetHeaderApplicationJson(req)
		req.Header.Add("Authorization", "Bearer")

		tokenMaker.EXPECT().VerifyToken(TOKEN).Return(createTokenPayload("active"), nil).Times(0)

		assert.PanicsWithValue(t, shrd_utils.AppError{
			Message:    "|invalid token",
			StatusCode: 401,
		}, func() {
			handler.ServeHTTP(recorder, req)
		})
	})
	t.Run("Invalid token (status code 401)", func(t *testing.T) {
		recorder, req := setUpRecorder()
		shrd_helper.SetHeaderApplicationJson(req)
		req.Header.Add("Authorization", "Bearer xxxxx")

		tokenMaker.EXPECT().VerifyToken("xxxxx").Return(nil, errors.New(UNAUTHORIZE)).
			Times(1)

		assert.PanicsWithValue(t, shrd_utils.AppError{
			Message:    fmt.Sprintf("%s|invalid token", UNAUTHORIZE),
			StatusCode: 401,
		}, func() {
			handler.ServeHTTP(recorder, req)
		})
	})

	t.Run("Inactive user (status code 403)", func(t *testing.T) {
		recorder, req := setUpRecorder()
		shrd_helper.SetHeaderApplicationJson(req)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", TOKEN))

		tokenMaker.EXPECT().VerifyToken(TOKEN).Return(createTokenPayload("not-active"), nil).
			Times(1)

		assert.PanicsWithValue(t, shrd_utils.AppError{
			Message:    "|inactive user can't access this route",
			StatusCode: 403,
		}, func() {
			handler.ServeHTTP(recorder, req)
		})
	})
}
