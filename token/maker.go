package shrd_token

import (
	"time"
)

type Maker interface {
	CreateToken(params PayloadParams, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
