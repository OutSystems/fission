/*
Copyright 2021 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

// fission CRD API version
const FISSION_API_VERSION = "fission.io/v1"

// fission CRD Kind
const (
	FISSION_ENVIRONMENT            = "Environment"
	FISSION_CANARYCONFIG           = "CanaryConfig"
	FISSION_FUNCTION               = "Function"
	FISSION_HTTPTRIGGER            = "HTTPTrigger"
	FISSION_KUBERNETESWATCHTRIGGER = "KubernetesWatchTrigger"
	FISSION_MESSAGEQUEUETRIGGER    = "MessageQueueTrigger"
	FISSION_PACKAGE                = "Package"
	FISSION_TIMETRIGGER            = "TimeTrigger"
)

// fission-cli options
const (
	SPEC_IGNORE_FILE   = ".specignore"
	COMMIT_LABEL       = "commit"
	FISSION_AUTH_URI   = "/auth/login"
	FISSION_AUTH_TOKEN = "FISSION_AUTH_TOKEN"
)
