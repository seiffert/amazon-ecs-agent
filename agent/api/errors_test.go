// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package api

import (
	"errors"
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
)

func TestNewAPIError(t *testing.T) {
	retriable := []error{
		aws.APIError{Code: "ServerException"},
		errors.New("Error"),
	}
	unretriable := []error{
		aws.APIError{Code: "ClientException"},
		aws.APIError{StatusCode: 400, Code: "BadRequestException"},
	}

	for i, err := range retriable {
		sce := NewAPIError(err)
		if !sce.Retry() {
			t.Errorf("Expected error to be retriable: #%v: %v", i, err)
		}
	}

	for i, err := range unretriable {
		sce := NewAPIError(err)
		if sce.Retry() {
			t.Errorf("Expected error to be unretriable: #%v: %v", i, err)
		}
	}
}
