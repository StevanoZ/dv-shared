package shrd_utils

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type BaseConfig struct {
	GCP_PROJECT_ID          string        `mapstructure:"GCP_PROJECT_ID"`
	DLQ_TOPIC               string        `mapstructure:"DLQ_TOPIC"`
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBSource                string        `mapstructure:"DB_SOURCE"`
	ServerPort              string        `mapstructure:"SERVER_PORT"`
	TokenSymmetricKey       string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration    time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	SGKey                   string        `mapstructure:"SG_KEY"`
	SenderEmail             string        `mapstructure:"SENDER_EMAIL"`
	AWSSecretKey            string        `mapstructure:"AWS_SECRET_KEY"`
	AWSAccessKey            string        `mapstructure:"AWS_ACCESS_KEY"`
	AWSRegion               string        `mapstructure:"AWS_REGION"`
	S3PublicBucketName      string        `mapstructure:"S3_PUBLIC_BUCKET_NAME"`
	S3PrivateBucketName     string        `mapstructure:"S3_PRIVATE_BUCKET_NAME"`
	PreSignUrlDuration      time.Duration `mapstructure:"PRE_SIGN_URL_DURATION"`
	Environment             string        `mapstructure:"ENVIRONMENT"`
	KafkaBroker             string        `mapstructure:"KAFKA_BROKER"`
	KafkaUsername           string        `mapstructure:"KAFKA_USERNAME"`
	KafkaPassword           string        `mapstructure:"KAFKA_PASSWORD"`
	IsRemoteBroker          bool          `mapstructure:"IS_REMOTE_BROKER"`
	ServiceName             string        `mapstructure:"SERVICE_NAME"`
	REDIS_HOST              string        `mapstructure:"REDIS_HOST"`
	REDIS_USERNAME          string        `mapstructure:"REDIS_USERNAME"`
	REDIS_PASSWORD          string        `mapstructure:"REDIS_PASSWORD"`
	CACHE_DURATION          time.Duration `mapstructure:"CACHE_DURATION"`
	MIGRATION_URL           string        `mapstructure:"MIGRATION_URL"`
	PS_PUBSUB_EMULATOR_HOST string        `mapstructure:"PS_PUBSUB_EMULATOR_HOST"`
	RETRY_TIME              time.Duration `mapstructure:"RETRY_TIME"`
	DATA_DOG_AGENT_HOST     string        `mapstructure:"DATA_DOG_AGENT_HOST"`
}

func LoadBaseConfig(path string, configName string) (config *BaseConfig) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	LogAndPanicIfError(err, "failed when reading config")

	err = viper.Unmarshal(&config)
	LogAndPanicIfError(err, "failed when unmarshal config")

	return
}

func CheckAndSetConfig(path string, configName string) *BaseConfig {
	config := LoadBaseConfig(path, configName)
	if config.Environment == TEST {
		os.Setenv("PUBSUB_EMULATOR_HOST", config.PS_PUBSUB_EMULATOR_HOST)
		config = LoadBaseConfig(path, "test")
	}

	if config.Environment == LOCAL {
		os.Setenv("PUBSUB_EMULATOR_HOST", config.PS_PUBSUB_EMULATOR_HOST)
	}

	return config
}
