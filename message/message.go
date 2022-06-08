package message

import (
	"time"

	"github.com/google/uuid"
)

// TOPIC
const EMAIL_TOPIC = "EMAIL"
const USER_TOPIC = "USER"
const USER_IMAGE_TOPIC = "USER-IMAGE"

// KEY
const CREATED_KEY = "created"
const UPDATED_KEY = "updated"
const DELETED_KEY = "deleted"
const SUCCESS = "success"
const FAILED = "failed"

// SPECIFIC KEY
const SEND_OTP_KEY = "send-otp"
const UPDATED_USER_MAIN_IMAGE_KEY = "updated-user-main-image"

// PAYLOAD
type OtpPayload struct {
	Email   string
	OtpCode int
}

type CreatedUserPayload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	OtpCode   int64     `json:"otp_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatedUserPayload struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	AttemptLeft int32     `json:"attempt_left"`
	OtpCode     int64     `json:"otp_code"`
	Status      string    `json:"status"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdatedUserMainImagePayload struct {
	MainImageUrl  string    `json:"main_image_url"`
	MainImagePath string    `json:"main_image_path"`
	ID            uuid.UUID `json:"id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreatedUserImagePayload struct {
	ID        uuid.UUID `json:"id"`
	ImageUrl  string    `json:"image_url"`
	ImagePath string    `json:"image_path"`
	IsMain    bool      `json:"is_main"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatedUserImagePayload struct {
	ID        uuid.UUID `json:"id"`
	IsMain    bool      `json:"is_main"`
	UserID    uuid.UUID `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletedUserImagePayload struct {
	ID uuid.UUID `json:"id"`
}
