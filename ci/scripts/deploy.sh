#!/usr/bin/env bash

# Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
#
# This program and the accompanying materials are made available under
# the terms of the under the Apache License, Version 2.0 (the "License”);
# you may not use this file except in compliance with the License.
#
# You may obtain a copy of the License at
# http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#
# See the License for the specific language governing permissions and
# limitations under the License.

set -eu

./bosh-backup-and-restore-meta/unlock-ci.sh
export BOSH_CLIENT
export BOSH_CLIENT_SECRET
export BOSH_ENVIRONMENT
export BOSH_CA_CERT="./bosh-backup-and-restore-meta/certs/${BOSH_ENVIRONMENT}.crt"
export OPTIONAL_BOSH_VARS_db_password=${DB_PASSWORD}
export OPTIONAL_BOSH_VARS_db_host=${DB_HOST}

local vars-store-argument=""
if [-z "$VARS_STORE_PATH" ]; then
  vars-store-argument="--vars-store bosh-backup-and-restore-meta/${VARS_STORE_PATH}"
fi

bosh-cli --non-interactive \
  --deployment ${BOSH_DEPLOYMENT} \
  deploy "backup-and-restore-sdk-release/ci/manifests/${MANIFEST_NAME}" \
  --var=backup-and-restore-sdk-release-version=$(cat release-tarball/version) \
  --var=backup-and-restore-sdk-release-url=$(cat release-tarball/url) \
  --vars-env=OPTIONAL_BOSH_VARS "${vars-store-argument}" \
  --var=deployment-name=${BOSH_DEPLOYMENT}
