module s3-blobstore-backup-restore

go 1.14

require (
	github.com/aws/aws-sdk-go v1.36.29
	github.com/cloudfoundry-incubator/bosh-backup-and-restore v1.9.0
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.4
	system-tests v0.0.0
)

replace system-tests => ../system-tests

replace s3-blobstore-backup-restore => ./
