package gcs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"

	"errors"

	"github.com/cloudfoundry-incubator/gcs-blobstore-backup-restore"
	"github.com/cloudfoundry-incubator/gcs-blobstore-backup-restore/fakes"
)

var _ = Describe("Restorer", func() {
	var firstBucket *fakes.FakeBucket
	var secondBucket *fakes.FakeBucket
	var thirdBucket *fakes.FakeBucket

	var restorer gcs.Restorer

	const firstBucketName = "first-bucket-name"
	const secondBucketName = "second-bucket-name"
	const thirdBucketName = "third-bucket-name"

	var executionStrategy = gcs.NewParallelStrategy()

	BeforeEach(func() {
		firstBucket = new(fakes.FakeBucket)
		secondBucket = new(fakes.FakeBucket)
		thirdBucket = new(fakes.FakeBucket)

		firstBucket.NameReturns(firstBucketName)
		secondBucket.NameReturns(secondBucketName)
		thirdBucket.NameReturns(thirdBucketName)

		firstBucket.VersioningEnabledReturns(true, nil)
		secondBucket.VersioningEnabledReturns(true, nil)
		thirdBucket.VersioningEnabledReturns(true, nil)

		restorer = gcs.NewRestorer(map[string]gcs.Bucket{
			"first":  firstBucket,
			"second": secondBucket,
			"third":  thirdBucket,
		},
			executionStrategy)
	})

	It("restores the version of every blob", func() {
		backups := map[string]gcs.BucketBackup{
			"first": {
				Name: firstBucketName,
				Blobs: []gcs.Blob{
					{Name: "blob1", GenerationID: 123},
					{Name: "blob2", GenerationID: 234},
				},
			},
			"second": {
				Name: secondBucketName,
				Blobs: []gcs.Blob{
					{Name: "blob3", GenerationID: 345},
				},
			},
		}
		var expectedBlobs []gcs.Blob
		restorer.Restore(backups)

		Expect(firstBucket.VersioningEnabledCallCount()).To(Equal(1))
		Expect(firstBucket.CopyVersionCallCount()).To(Equal(2))

		expectedBlob, expectedRestoreBucket := firstBucket.CopyVersionArgsForCall(0)
		expectedBlobs = append(expectedBlobs, expectedBlob)
		Expect(expectedRestoreBucket).To(Equal(firstBucketName))

		expectedBlob, expectedRestoreBucket = firstBucket.CopyVersionArgsForCall(1)
		expectedBlobs = append(expectedBlobs, expectedBlob)
		Expect(expectedRestoreBucket).To(Equal(firstBucketName))

		Expect(expectedBlobs).To(ConsistOf(
			gcs.Blob{Name: "blob2", GenerationID: 234},
			gcs.Blob{Name: "blob1", GenerationID: 123},
		))

		Expect(secondBucket.VersioningEnabledCallCount()).To(Equal(1))
		Expect(secondBucket.CopyVersionCallCount()).To(Equal(1))

		expectedBlob, expectedRestoreBucket = secondBucket.CopyVersionArgsForCall(0)
		Expect(expectedBlob).To(Equal(gcs.Blob{Name: "blob3", GenerationID: 345}))
		Expect(expectedRestoreBucket).To(Equal(secondBucketName))
	})

	Context("when restoring to a different bucket", func() {
		var firstRestoreBucketName string
		var secondRestoreBucketName string
		It("restores successfully", func() {
			firstRestoreBucketName = "restoreBucket1"
			secondRestoreBucketName = "restoreBucket2"

			backups := map[string]gcs.BucketBackup{
				"first": {
					Name: firstRestoreBucketName,
					Blobs: []gcs.Blob{
						{Name: "blob1", GenerationID: 123},
						{Name: "blob2", GenerationID: 234},
					},
				},
				"second": {
					Name: secondRestoreBucketName,
					Blobs: []gcs.Blob{
						{Name: "blob3", GenerationID: 345},
					},
				},
			}

			restorer.Restore(backups)
			var expectedBlobs []gcs.Blob

			Expect(firstBucket.VersioningEnabledCallCount()).To(Equal(1))
			Expect(firstBucket.CopyVersionCallCount()).To(Equal(2))

			expectedBlob, expectedRestoreBucket := firstBucket.CopyVersionArgsForCall(0)
			expectedBlobs = append(expectedBlobs, expectedBlob)
			Expect(expectedRestoreBucket).To(Equal(firstRestoreBucketName))

			expectedBlob, expectedRestoreBucket = firstBucket.CopyVersionArgsForCall(1)
			expectedBlobs = append(expectedBlobs, expectedBlob)
			Expect(expectedRestoreBucket).To(Equal(firstRestoreBucketName))

			Expect(expectedBlobs).To(ConsistOf(
				gcs.Blob{Name: "blob2", GenerationID: 234},
				gcs.Blob{Name: "blob1", GenerationID: 123},
			))

			Expect(secondBucket.VersioningEnabledCallCount()).To(Equal(1))
			Expect(secondBucket.CopyVersionCallCount()).To(Equal(1))

			expectedBlob, expectedRestoreBucket = secondBucket.CopyVersionArgsForCall(0)
			Expect(expectedBlob).To(Equal(gcs.Blob{Name: "blob3", GenerationID: 345}))
			Expect(expectedRestoreBucket).To(Equal(secondRestoreBucketName))

		})
	})

	Context("when versioning is turned off on a bucket", func() {
		It("returns an error", func() {
			firstBucket.VersioningEnabledReturns(false, nil)

			err := restorer.Restore(nil)

			Expect(err).To(MatchError(fmt.Sprintf("versioning is not enabled on bucket '%s'", firstBucketName)))
		})
	})

	Context("when the bucket versioning check fails", func() {
		It("returns an error", func() {
			firstBucket.VersioningEnabledReturns(false, errors.New("ooops!"))

			err := restorer.Restore(nil)

			Expect(err).To(MatchError(SatisfyAll(
				ContainSubstring(fmt.Sprintf("failed to check if versioning is enabled on bucket '%s'", firstBucketName)),
				ContainSubstring("ooops!"),
			)))
		})
	})

	Context("when a backup bucket is not configured", func() {
		It("returns an error", func() {
			bucketIdentifier := "fourth"
			backups := map[string]gcs.BucketBackup{
				bucketIdentifier: {
					Name:  "fourth-bucket-name",
					Blobs: []gcs.Blob{},
				},
			}

			err := restorer.Restore(backups)

			Expect(err).To(MatchError(fmt.Sprintf("bucket identifier '%s' not found in bucketPairs configuration", bucketIdentifier)))
		})

		It("does not check versioning or restore blobs", func() {
			bucketIdentifier := "fourth"
			backups := map[string]gcs.BucketBackup{
				"first": {
					Name: firstBucketName,
					Blobs: []gcs.Blob{
						{Name: "blob1", GenerationID: 123},
						{Name: "blob2", GenerationID: 234},
					},
				},
				"second": {
					Name: secondBucketName,
					Blobs: []gcs.Blob{
						{Name: "blob3", GenerationID: 345},
					},
				},
				bucketIdentifier: {
					Name:  "fourth-bucket-name",
					Blobs: []gcs.Blob{},
				},
			}

			err := restorer.Restore(backups)

			Expect(err).To(MatchError(fmt.Sprintf("bucket identifier '%s' not found in bucketPairs configuration", bucketIdentifier)))
			Expect(firstBucket.CopyVersionCallCount()).To(Equal(0))
			Expect(secondBucket.CopyVersionCallCount()).To(Equal(0))
			Expect(firstBucket.VersioningEnabledCallCount()).To(Equal(0))
			Expect(secondBucket.VersioningEnabledCallCount()).To(Equal(0))
		})
	})

	Context("when copying a blob version fails", func() {
		It("returns an error", func() {
			blobName := "blob1"
			backups := map[string]gcs.BucketBackup{
				"first": {
					Name: firstBucketName,
					Blobs: []gcs.Blob{
						{Name: blobName, GenerationID: 123},
					},
				},
			}
			firstBucket.CopyVersionReturns(errors.New("ooops!"))

			err := restorer.Restore(backups)

			Expect(err).To(MatchError(SatisfyAll(
				ContainSubstring(fmt.Sprintf("failed to restore bucket '%s'", firstBucketName)),
				ContainSubstring("ooops!"),
			)))
		})
	})
})
