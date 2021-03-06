package incremental

import (
	"errors"
	"fmt"
)

type Restorer struct {
	bucketPairs map[string]RestoreBucketPair
	artifact    Artifact
}

func NewRestorer(bucketPairs map[string]RestoreBucketPair, artifact Artifact) Restorer {
	return Restorer{
		bucketPairs: bucketPairs,
		artifact:    artifact,
	}
}

func (b Restorer) Run() error {
	backups, err := b.artifact.Load()
	if err != nil {
		return err
	}

	for bucketID, backup := range backups {
		_, exists := b.bucketPairs[bucketID]
		if !exists {
			return fmt.Errorf(
				"restore config does not mention bucket: %s, but is present in the artifact",
				bucketID,
			)
		}

		if backup.SameBucketAs != "" {
			continue
		}

		backupBlobs, _ := b.bucketPairs[bucketID].ArtifactBackupBucket.ListBlobs(backup.SrcBackupDirectoryPath)
		if missingBlobs := validateArtifact(backup.SrcBackupDirectoryPath, backupBlobs, backup.Blobs); len(missingBlobs) > 0 {
			return formatError(fmt.Sprintf("found blobs in artifact that are not present in backup directory for bucket %s:", backup.BucketName), missingBlobs)
		}
	}

	for key, pair := range b.bucketPairs {
		backup, exists := backups[key]
		if !exists {
			return fmt.Errorf("cannot restore bucket %s, not found in backup artifact", key)
		}

		if len(backup.Blobs) != 0 {
			err = pair.Restore(backup)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func validateArtifact(srcBackupDirectoryPath string, backupBlobs []Blob, artifactBlobs []string) []string {
	var missingBlobs []string
	for _, artifactBlobPath := range artifactBlobs {
		if !contains(srcBackupDirectoryPath, artifactBlobPath, backupBlobs) {
			missingBlobs = append(missingBlobs, artifactBlobPath)
		}
	}

	return missingBlobs
}

func contains(srcBackupDirectoryPath, key string, blobs []Blob) bool {
	for _, blob := range blobs {
		if key == joinBlobPath(srcBackupDirectoryPath, blob.Path()) {
			return true
		}
	}
	return false
}

func formatError(errorMsg string, blobs []string) error {
	for _, blob := range blobs {
		errorMsg = errorMsg + "\n" + blob
	}
	return errors.New(errorMsg)
}
