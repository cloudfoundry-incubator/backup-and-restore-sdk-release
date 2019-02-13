package contract_test

import (
	"fmt"

	"github.com/cloudfoundry-incubator/gcs-blobstore-backup-restore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bucket", func() {
	Describe("BuildBucketPairs", func() {
		It("builds bucket pairs", func() {
			config := map[string]gcs.Config{
				"droplets": {
					BucketName:       "droplets-bucket",
					BackupBucketName: "backup-droplets-bucket",
				},
			}

			buckets, err := gcs.BuildBucketPairs(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), config)
			Expect(err).NotTo(HaveOccurred())

			Expect(buckets).To(HaveLen(1))
			Expect(buckets["droplets"].LiveBucket.Name()).To(Equal("droplets-bucket"))
			Expect(buckets["droplets"].BackupBucket.Name()).To(Equal("backup-droplets-bucket"))
			Expect(buckets["droplets"].IDs).To(Equal([]string{"droplets"}))
		})

		Context("when providing invalid service account key", func() {
			It("returns an error", func() {
				config := map[string]gcs.Config{
					"droplets": {
						BucketName:       "droplets-bucket",
						BackupBucketName: "backup-droplets-bucket",
					},
				}

				_, err := gcs.BuildBucketPairs("not-valid-json", config)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when providing two configs with the same live and backup bucket setup", func() {
			It("de-duplicates the bucket pairs", func() {
				config := map[string]gcs.Config{
					"droplets": {
						BucketName:       "droplets-bucket",
						BackupBucketName: "backup-droplets-bucket",
					},
					"buildpacks": {
						BucketName:       "droplets-bucket",
						BackupBucketName: "backup-droplets-bucket",
					},
				}

				buckets, err := gcs.BuildBucketPairs(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), config)
				Expect(err).NotTo(HaveOccurred())

				Expect(buckets).To(HaveLen(1))
				Expect(buckets["buildpacks-droplets"].LiveBucket.Name()).To(Equal("droplets-bucket"))
				Expect(buckets["buildpacks-droplets"].BackupBucket.Name()).To(Equal("backup-droplets-bucket"))
				Expect(buckets["buildpacks-droplets"].IDs).To(HaveLen(2))
				Expect(buckets["buildpacks-droplets"].IDs).To(ContainElement("droplets"))
				Expect(buckets["buildpacks-droplets"].IDs).To(ContainElement("buildpacks"))
			})
		})

		Context("when providing three configs with the same live and backup bucket setup", func() {
			It("de-duplicates the bucket pairs", func() {
				config := map[string]gcs.Config{
					"a": {
						BucketName:       "a-bucket",
						BackupBucketName: "backup-a-bucket",
					},
					"b": {
						BucketName:       "a-bucket",
						BackupBucketName: "backup-a-bucket",
					},
					"c": {
						BucketName:       "a-bucket",
						BackupBucketName: "backup-a-bucket",
					},
				}

				buckets, err := gcs.BuildBucketPairs(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), config)
				Expect(err).NotTo(HaveOccurred())

				Expect(buckets).To(HaveLen(1))
				fmt.Println(buckets)
				Expect(buckets["a-b-c"].LiveBucket.Name()).To(Equal("a-bucket"))
				Expect(buckets["a-b-c"].BackupBucket.Name()).To(Equal("backup-a-bucket"))
				Expect(buckets["a-b-c"].IDs).To(HaveLen(3))
				Expect(buckets["a-b-c"].IDs).To(ContainElement("a"))
				Expect(buckets["a-b-c"].IDs).To(ContainElement("b"))
				Expect(buckets["a-b-c"].IDs).To(ContainElement("c"))
			})
		})
	})

	Describe("ListBlobs", func() {
		var bucketName string
		var bucket gcs.Bucket
		var err error

		JustBeforeEach(func() {
			bucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), bucketName)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when the bucket contains multiple files and some match the prefix", func() {
			BeforeEach(func() {
				bucketName = CreateBucketWithTimestampedName("list_blobs")
				UploadFileWithDir(bucketName, "my/prefix", "file1", "file-content")
				UploadFileWithDir(bucketName, "not/my/prefix", "file2", "file-content")
				UploadFile(bucketName, "file3", "file-content")
			})

			It("lists all files that have the prefix", func() {
				blobs, err := bucket.ListBlobs("my/prefix")
				Expect(err).NotTo(HaveOccurred())
				Expect(blobs).To(ConsistOf(
					gcs.NewBlob("my/prefix/file1"),
				))
			})

			AfterEach(func() {
				DeleteBucket(bucketName)
			})
		})

		Context("when providing a non-existing bucket", func() {
			It("returns an error", func() {
				config := map[string]gcs.Config{
					"droplets": {
						BucketName:       "I-am-not-a-bucket",
						BackupBucketName: "definitely-not-a-bucket",
					},
				}

				bucketPair, err := gcs.BuildBucketPairs(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), config)
				Expect(err).NotTo(HaveOccurred())
				_, err = bucketPair["droplets"].LiveBucket.ListBlobs("")
				Expect(err).To(MatchError("storage: bucket doesn't exist"))
			})
		})
	})

	Describe("CopyBlobToBucket", func() {
		var srcBucketName string
		var dstBucketName string
		var srcBucket gcs.Bucket
		var dstBucket gcs.Bucket
		var err error

		Context("copying an existing file", func() {
			BeforeEach(func() {
				srcBucketName = CreateBucketWithTimestampedName("src")
				dstBucketName = CreateBucketWithTimestampedName("dst")
				UploadFile(srcBucketName, "file1", "file-content")

				srcBucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), srcBucketName)
				Expect(err).NotTo(HaveOccurred())

				dstBucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), dstBucketName)
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				DeleteBucket(srcBucketName)
				DeleteBucket(dstBucketName)
			})

			It("copies the blob to the specified location", func() {
				blob := gcs.NewBlob("file1")

				err := srcBucket.CopyBlobToBucket(dstBucket, blob.Name(), "copydir/file1")
				Expect(err).NotTo(HaveOccurred())

				blobs, err := dstBucket.ListBlobs("")
				Expect(err).NotTo(HaveOccurred())
				Expect(blobs).To(ConsistOf(
					gcs.NewBlob("copydir/file1"),
				))
			})
		})

		Context("copying a file that doesn't exist", func() {
			BeforeEach(func() {
				srcBucketName = CreateBucketWithTimestampedName("src")
				dstBucketName = CreateBucketWithTimestampedName("dst")

				srcBucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), srcBucketName)
				Expect(err).NotTo(HaveOccurred())

				dstBucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), dstBucketName)
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				DeleteBucket(srcBucketName)
				DeleteBucket(dstBucketName)
			})

			It("errors with a useful message", func() {
				err := srcBucket.CopyBlobToBucket(dstBucket, "foobar", "copydir/file1")
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring("failed to copy object: ")))
			})
		})

		Context("copying to a bucket that doesn't exist", func() {
			BeforeEach(func() {
				srcBucketName = CreateBucketWithTimestampedName("src")
				UploadFile(srcBucketName, "file1", "file-content")

				srcBucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), srcBucketName)
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				DeleteBucket(srcBucketName)
			})

			It("errors", func() {
				err := srcBucket.CopyBlobToBucket(nil, "file1", "copydir/file1")
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("destination bucket does not exist"))
			})
		})
	})

	Describe("CopyBlobsToBucket", func() {
		var (
			srcBucketName string
			dstBucketName string
			srcBucket     gcs.Bucket
			dstBucket     gcs.Bucket
			badBucket     gcs.Bucket
			err           error
		)

		BeforeEach(func() {
			srcBucketName = CreateBucketWithTimestampedName("src")
			dstBucketName = CreateBucketWithTimestampedName("dst")
			UploadFile(srcBucketName, "notInSourcePath", "file-content")
			UploadFile(dstBucketName, "alreadyInDstBucket", "file-content")
			UploadFileWithDir(srcBucketName, "sourcePath", "file1", "file-content1")
			UploadFileWithDir(srcBucketName, "sourcePath", "file2", "file-content2")

			srcBucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), srcBucketName)
			Expect(err).NotTo(HaveOccurred())

			dstBucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), dstBucketName)
			Expect(err).NotTo(HaveOccurred())

			badBucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), "badBucket")
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			DeleteBucket(srcBucketName)
			DeleteBucket(dstBucketName)
		})

		It("copies only the blobs from source/sourcePath to destination/destinationPath", func() {
			err := srcBucket.CopyBlobsToBucket(dstBucket, "sourcePath")
			Expect(err).NotTo(HaveOccurred())

			blobs, err := dstBucket.ListBlobs("")
			Expect(err).NotTo(HaveOccurred())
			Expect(blobs).To(ConsistOf(
				gcs.NewBlob("file1"),
				gcs.NewBlob("file2"),
				gcs.NewBlob("alreadyInDstBucket"),
			))
		})

		It("returns an error if the destination bucket does not exist", func() {
			err := srcBucket.CopyBlobsToBucket(nil, "sourcePath")
			Expect(err).To(MatchError("destination bucket does not exist"))
		})

		It("returns an error when the source bucket does not exist", func() {
			err := badBucket.CopyBlobsToBucket(dstBucket, "path")
			Expect(err).To(HaveOccurred())
		})

		It("returns an error when the destination bucket does not exist", func() {
			err := srcBucket.CopyBlobsToBucket(badBucket, "sourcePath")
			Expect(err).To(HaveOccurred())
		})

	})

	Describe("DeleteBlob", func() {
		var (
			bucketName string
			bucket     gcs.Bucket
			err        error
			dirName    = "mydir"
			fileName1  = "file1"
			fileName2  = "file2"
		)

		AfterEach(func() {
			DeleteBucket(bucketName)
		})

		Context("when deleting a file that doesn't exist", func() {
			BeforeEach(func() {
				bucketName = CreateBucketWithTimestampedName("src")
				UploadFile(bucketName, fileName1, "file-content")

				bucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), bucketName)
				Expect(err).NotTo(HaveOccurred())
			})

			It("errors", func() {
				err := bucket.DeleteBlob(fileName1 + "idontexist")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when deleting existing files", func() {
			BeforeEach(func() {
				bucketName = CreateBucketWithTimestampedName("src")
				UploadFileWithDir(bucketName, dirName, fileName1, "file-content")
				UploadFileWithDir(bucketName, dirName, fileName2, "file-content")

				bucket, err = gcs.NewSDKBucket(MustHaveEnv("GCP_SERVICE_ACCOUNT_KEY"), bucketName)
				Expect(err).NotTo(HaveOccurred())
			})

			It("deletes all files and the folder", func() {
				err := bucket.DeleteBlob(fmt.Sprintf("%s/%s", dirName, fileName1))
				Expect(err).NotTo(HaveOccurred())

				err = bucket.DeleteBlob(fmt.Sprintf("%s/%s", dirName, fileName2))
				Expect(err).NotTo(HaveOccurred())

				blobs, err := bucket.ListBlobs("")
				Expect(err).NotTo(HaveOccurred())
				Expect(blobs).To(BeNil())
			})
		})
	})

})
