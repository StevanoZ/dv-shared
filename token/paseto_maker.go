package shrd_token

import (
	"fmt"
	"time"

	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(config *shrd_utils.BaseConfig) (Maker, error) {
	key := config.TokenSymmetricKey
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(key),
	}
	return maker, nil
}

func (m *PasetoMaker) CreateToken(params PayloadParams, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(params, duration)
	token, err := m.paseto.Encrypt(m.symmetricKey, payload, nil)

	return token, payload, err

}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := m.paseto.Decrypt(token, m.symmetricKey, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
