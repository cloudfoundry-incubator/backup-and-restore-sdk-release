package versioned

import "github.com/cloudfoundry-incubator/backup-and-restore-sdk/s3-blobstore-backup-restore/s3bucket"

//go:generate counterfeiter -o fakes/fake_bucket.go . Bucket
type Bucket interface {
	Name() string
	Region() string
	CopyVersion(blobKey, versionId, originBucketName, originBucketRegion string) error
	ListVersions() ([]s3bucket.Version, error)
	IsVersioned() (bool, error)
}
