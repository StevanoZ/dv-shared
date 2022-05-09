package message

// TOPIC
const EMAIL_TOPIC = "EMAIL"

// KEY
const SEND_OTP_KEY = "send-otp"

// PAYLOAD
type OtpPayload struct {
	Email   string
	OtpCode int
}
