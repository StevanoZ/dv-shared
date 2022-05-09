package kafka_client

// TOPIC
const EMAIL_TOPIC = "EMAIL"

// KEY
const SEND_OTP = "send-otp"

// PAYLOAD
type OtpPayload struct {
	Email   string
	OtpCode int
}
