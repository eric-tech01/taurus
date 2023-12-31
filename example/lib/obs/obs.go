package obs

import "github.com/eric-tech01/taurus/pkg/obs/minio"

var objectStore *ObjectStore

type ObjectStore struct {
	*minio.MinioClient
}

func SetObjectStore(c *minio.MinioClient) {
	objectStore = &ObjectStore{MinioClient: c}
}

func (o ObjectStore) UploadFile() {

}
func (o ObjectStore) DeleteFile() {

}
