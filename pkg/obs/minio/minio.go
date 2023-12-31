package minio

import (
	"context"
	"os"

	log "github.com/eric-tech01/simple-log"
	"github.com/minio/minio-go/v7"
)

type MinioClient struct {
	config Config
	client *minio.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *MinioClient) UploadFile(filePath string, bucket string, keyName string) error {
	log.Info("Start to upload file " + filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	info, err := c.client.PutObject(c.ctx, bucket, keyName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Infof("Success upload bytes: %d", info.Size)
	log.Debugf("Upload Details: %v", info)
	return nil
}

func (c MinioClient) PrefixUrl(bucket string) string {
	return "https://" + c.config.Domain + "/" + bucket + "/"
}

func (c MinioClient) DeleteObject(bucket string, objectName string, opts minio.RemoveObjectOptions) error {
	return c.client.RemoveObject(c.ctx, bucket, objectName, opts)
}

func (c MinioClient) DeleteAllObjectsByPrefix(bucket string, prefix string, opts minio.ListObjectsOptions) {
	objectsCh := c.client.ListObjects(c.ctx, bucket, opts)
	// Send object names that are needed to be removed to objectsCh
	for rErr := range c.client.RemoveObjects(c.ctx, bucket, objectsCh, minio.RemoveObjectsOptions{}) {
		log.Errorf("Remove Error. ObjectName:%s, VersionID:%s, Err:%s", rErr.ObjectName, rErr.VersionID, rErr.Err)
	}
}
