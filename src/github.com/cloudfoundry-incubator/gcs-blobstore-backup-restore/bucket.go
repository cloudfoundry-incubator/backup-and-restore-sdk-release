package gcs

import (
	"errors"
	"strings"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const readWriteScope = "https://www.googleapis.com/auth/devstorage.read_write"

//go:generate counterfeiter -o fakes/fake_bucket.go . Bucket
type Bucket interface {
	Name() string
	ListBlobs() ([]Blob, error)
	ListLastBackupBlobs() ([]Blob, error)
	CopyBlobWithinBucket(string, string) (int64, error)
	CopyBlobBetweenBuckets(Bucket, string, string) (int64, error)
	DeleteBlob(string) error
	CreateFile(name string, content []byte) (int64, error)
}

type BucketPair struct {
	Bucket       Bucket
	BackupBucket Bucket
}

func BuildBuckets(gcpServiceAccountKey string, config map[string]Config) (map[string]BucketPair, error) {
	buckets := map[string]BucketPair{}

	for bucketId, bucketConfig := range config {
		bucket, err := NewSDKBucket(gcpServiceAccountKey, bucketConfig.BucketName)
		if err != nil {
			return nil, err
		}

		backupBucket, err := NewSDKBucket(gcpServiceAccountKey, bucketConfig.BackupBucketName)
		if err != nil {
			return nil, err
		}

		buckets[bucketId] = BucketPair{
			Bucket:       bucket,
			BackupBucket: backupBucket,
		}
	}

	return buckets, nil
}

type Blob struct {
	Name         string `json:"name"`
	GenerationID int64  `json:"generation_id"`
}

type BucketBackup struct {
	Name  string `json:"name"`
	Blobs []Blob `json:"blobs"`
}

type SDKBucket struct {
	name   string
	handle *storage.BucketHandle
	ctx    context.Context
	client *storage.Client
}

func NewSDKBucket(serviceAccountKeyJson string, name string) (SDKBucket, error) {
	ctx := context.Background()

	creds, err := google.CredentialsFromJSON(ctx, []byte(serviceAccountKeyJson), readWriteScope)
	if err != nil {
		return SDKBucket{}, err
	}

	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		return SDKBucket{}, err
	}

	handle := client.Bucket(name)

	return SDKBucket{name: name, handle: handle, ctx: ctx, client: client}, nil
}

func (b SDKBucket) Name() string {
	return b.name
}

func (b SDKBucket) ListBlobs() ([]Blob, error) {
	var blobs []Blob

	objectsIterator := b.handle.Objects(b.ctx, nil)
	for {
		objectAttributes, err := objectsIterator.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		blobs = append(blobs, Blob{Name: objectAttributes.Name, GenerationID: objectAttributes.Generation})
	}

	return blobs, nil
}

func (b SDKBucket) ListLastBackupBlobs() ([]Blob, error) {
	var lastBackupBlobs []Blob
	allBackupBlobs, err := b.ListBlobs()
	if err != nil {
		return lastBackupBlobs, err
	}

	if len(allBackupBlobs) == 0 {
		return lastBackupBlobs, nil
	}

	lastBackupTimestamp := strings.Split(allBackupBlobs[len(allBackupBlobs)-1].Name, "/")[0]

	for _, blob := range allBackupBlobs {
		if strings.HasPrefix(blob.Name, lastBackupTimestamp) {
			lastBackupBlobs = append(lastBackupBlobs, blob)
		}
	}

	return lastBackupBlobs, nil
}

func (b SDKBucket) CopyBlobWithinBucket(srcBlob, dstBlob string) (int64, error) {
	return b.CopyBlobBetweenBuckets(b, srcBlob, dstBlob)
}

func (b SDKBucket) CopyBlobBetweenBuckets(dstBucket Bucket, srcBlob, dstBlob string) (int64, error) {
	if dstBucket == nil {
		return 0, errors.New("destination bucket does not exist")
	}

	src := b.client.Bucket(b.name).Object(srcBlob)
	dst := b.client.Bucket(dstBucket.Name()).Object(dstBlob)
	attr, err := dst.CopierFrom(src).Run(b.ctx)
	if err != nil {
		return 0, err
	}

	return attr.Generation, nil
}

func (b SDKBucket) DeleteBlob(blob string) error {
	return b.client.Bucket(b.name).Object(blob).Delete(b.ctx)
}

func (b SDKBucket) CreateFile(name string, content []byte) (int64, error) {
	writer := b.client.Bucket(b.name).Object(name).NewWriter(b.ctx)
	writer.Write(content)
	err := writer.Close()
	if err != nil {
		return -1, err
	}
	return writer.Attrs().Generation, nil
}
