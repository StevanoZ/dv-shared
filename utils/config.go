package shrd_utils

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type BaseConfig struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerPort           string        `mapstructure:"SERVER_PORT"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	SGKey                string        `mapstructure:"SG_KEY"`
	SenderEmail          string        `mapstructure:"SENDER_EMAIL"`
	AWSSecretKey         string        `mapstructure:"AWS_SECRET_KEY"`
	AWSAccessKey         string        `mapstructure:"AWS_ACCESS_KEY"`
	AWSRegion            string        `mapstructure:"AWS_REGION"`
	S3PublicBucketName   string        `mapstructure:"S3_PUBLIC_BUCKET_NAME"`
	S3PrivateBucketName  string        `mapstructure:"S3_PRIVATE_BUCKET_NAME"`
	PreSignUrlDuration   time.Duration `mapstructure:"PRE_SIGN_URL_DURATION"`
}

func LoadBaseConfig(path string, configName string) (config *BaseConfig) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("failed when reading config " + err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln("failed when unmarshal config " + err.Error())
	}

	return
}
