# Backup and Restore SDK BOSH release

The Backup and Restore SDK BOSH release is used for two distinct things:

1. enabling release authors to incorporate database backup & restore functionality in their releases (See _[Database Backup and Restore](docs/database-backup-restore.md)_)
1. enabling operators to configure their deployments which use external blobstores to be backed up and restored by [BBR](https://github.com/cloudfoundry-incubator/bosh-backup-and-restore) (See _[Blobstore Backup and Restore](docs/blobstore-backup-restore.md)_)

**Slack**: #bbr on https://slack.cloudfoundry.org

**Pivotal Tracker**: https://www.pivotaltracker.com/n/projects/1662777

## CI Status

Backup and Restore SDK Release status [![Build SDK Release Badge](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/bbr/badge)](https://ci.cryo.cf-app.com/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release)

## Database Backup and Restore

| Name     | Versions                 |
|:---------|:-------------------------|
| MariaDB  | 10.2.x            |
| MySQL    | 5.7.x, 8.0.x      |
| Postgres | 9.6.x, 10.x, 11.x |

The SDK can use used against Postgres 9.4, but is not supported upstream by the Postgres community.

CI Status:
* GCP: [![GCP Test Badge](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/system-tests-external-dbs-gcp/badge)](https://ci.cryo.cf-app.com/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/system-tests-external-dbs-gcp/)
* AWS (RDS): [![AWS Test Badge](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/system-tests-external-dbs-rds/badge)](https://ci.cryo.cf-app.com/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/system-tests-external-dbs-rds/)
* Bosh Deployed: [![GCP Test Badge](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/system-tests-internal-dbs/badge)](https://ci.cryo.cf-app.com/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/system-tests-internal-dbs/)

## Blobstore Backup and Restore

### Supported Blobstores

CI Status: [![Blobstore test
badge](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/system-tests-blobstore-backuper/badge)](https://ci.cryo.cf-app.com/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/system-tests-blobstore-backuper/)

| Name                 | Status                                                                                                                                                                                                                                                                                                          |
|:---------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| S3-Compatible        | [![S3 Badge](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/s3-blobstore-backuper-system-tests/badge)](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/s3-blobstore-backuper-system-tests)          |
| Azure                | [![Azure Badge](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/azure-blobstore-backuper-system-tests/badge)](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/azure-blobstore-backuper-system-tests) |
| Google Cloud Storage | [![GCS Badge](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/s3-blobstore-backuper-system-tests/badge)](https://ci.cryo.cf-app.com/api/v1/teams/bosh-backup-restore/pipelines/backup-and-restore-sdk-release/jobs/gcs-blobstore-backuper-system-tests)        |

## Developing

This repository using develop as the main branch, tested releases are tagged with their versions, and master branch represents the latest release.

## Testing

### Unit tests

You can run all unit tests by running

```
make unit
```

Individual targets exist for individual unit tests, like `make unit-s3`. Check
the Makefile for all available targets

### Contract tests

To run the Blobstore contract tests, you'll need to export the environment
variables the particular test requires. Check the [sdk-unit-blobstore pipeline
task](ci/tasks/sdk-unit-blobstore/task.yml) for details.

### System tests

To run the system tests, you'll need to export the necessary environment
variables that the particular test requires.

See the [sdk-system-blobstore pipeline
task](ci/tasks/sdk-system-blobstore/task.yml) and [sdk-system-db pipeline
task](ci/tasks/sdk-system-db/task.yml) for more details.

You'll also need a bosh director with a few particular bosh releases deployed in
it. Given a bosh director exists, the [infrastructure pipeline](ci/pipelines/bbr-sdk-test-infrastructure/pipeline.yml)
can be set to deploy the necessary releases.
