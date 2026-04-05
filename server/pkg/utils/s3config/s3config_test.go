package s3config

import (
	"testing"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewS3Config(t *testing.T) {
	viper.Reset()
	viper.Set("s3.b2-eu-cen.key", "my-key")
	viper.Set("s3.b2-eu-cen.secret", "my-secret")
	viper.Set("s3.b2-eu-cen.endpoint", "http://localhost:9000")
	viper.Set("s3.b2-eu-cen.region", "us-east-1")
	viper.Set("s3.b2-eu-cen.bucket", "my-bucket")

	config := NewS3Config()

	s3Config := config.GetS3Config("b2-eu-cen")
	assert.NotNil(t, s3Config)

	creds, err := s3Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "my-key", creds.AccessKeyID)
	assert.Equal(t, "my-secret", creds.SecretAccessKey)
	assert.Equal(t, "", creds.SessionToken)
}

func TestNewS3ConfigWithToken(t *testing.T) {
	viper.Reset()
	viper.Set("s3.b2-eu-cen.key", "my-key")
	viper.Set("s3.b2-eu-cen.secret", "my-secret")
	viper.Set("s3.b2-eu-cen.token", "my-token")
	viper.Set("s3.b2-eu-cen.endpoint", "http://localhost:9000")
	viper.Set("s3.b2-eu-cen.region", "us-east-1")
	viper.Set("s3.b2-eu-cen.bucket", "my-bucket")

	config := NewS3Config()

	s3Config := config.GetS3Config("b2-eu-cen")
	assert.NotNil(t, s3Config)

	creds, err := s3Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "my-key", creds.AccessKeyID)
	assert.Equal(t, "my-secret", creds.SecretAccessKey)
	assert.Equal(t, "my-token", creds.SessionToken)
}
