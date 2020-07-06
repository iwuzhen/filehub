package seaweeds3

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/iwuzhen/filehub"
)

type SeaweedS3FileInfo struct {
	key     *s3.Object
	filehub filehub.Filehub
}

func (a *SeaweedS3FileInfo) Path() string {
	return *a.key.Key
}

func (a *SeaweedS3FileInfo) Size() int64 {
	return *a.key.Size
}

func (a *SeaweedS3FileInfo) ModTime() time.Time {
	return *a.key.LastModified
}

func (a *SeaweedS3FileInfo) Filehub() filehub.Filehub {
	return a.filehub
}

func (a *SeaweedS3FileInfo) String() string {
	return fmt.Sprintf("%s %s %d", a.Path(), a.ModTime(), a.Size())
}
