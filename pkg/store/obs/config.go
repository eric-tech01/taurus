package obs

import (
	"context"

	log "github.com/eric-tech01/simple-log"
	"github.com/eric-tech01/taurus/pkg/conf"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Key             string `json:"key"`
	Domain          string `json:"domain" toml:"Domain"`
	Endpoint        string `json:"endpoint" toml:"Endpoint"`
	AccessKeyID     string `json:"accessKeyID" toml:"AccessKeyID"`
	SecretAccessKey string `json:"secretAccessKey" toml:"SecretAccessKey"`
	UseSSL          bool   `json:"useSSL" toml:"UseSSL"`
}

func StdConfig(key string) *Config {
	//unmarsh from toml
	config := DefaultConfig()
	config.Key = key
	err := conf.UnmarshalKey(key, config)
	if err != nil {
		log.Panic(err)
	}

	return config
}

func (config Config) Build() *MinioClient {
	c := &MinioClient{}
	var err error
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.client, err = minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		log.Panic(err)
	}

	return c
}

func DefaultConfig() *Config {
	c := &Config{
		AccessKeyID:     "",
		SecretAccessKey: "",
		UseSSL:          true,
	}
	return c
}
