#!/usr/bin/env bash

# Copyright 2020 The Tekton Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script runs the presubmit tests; it is started by prow for each PR.
# For convenience, it can also be executed manually.
# Running the script without parameters, or with the --all-tests
# flag, causes all tests to be executed, in the right order.
# Use the flags --build-tests, --unit-tests and --integration-tests
# to run a specific set of tests.

# Markdown linting failures don't show up properly in Gubernator resulting
# in a net-negative contributor experience.
export DISABLE_MD_LINTING=1
export DISABLE_MD_LINK_CHECK=1
export DISABLE_YAML_LINTING=1

source "$(git rev-parse --show-toplevel)"/vendor/github.com/tektoncd/plumbing/scripts/presubmit-tests.sh

function post_build_tests() {
  header "copyright licenses check"
  addlicense -ignore "vendor/**" -ignore "third_party/**"  -l apache -c 'The Tekton Authors' -v *
  git diff --exit-code
}

# We use the default build, unit and integration test runners.

main "$@"
