/*
Copyright 2021 The Skaffold Authors

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
package validate

import (
	"testing"

	latestV2 "github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest/v2"
	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestValidatorInit(t *testing.T) {
	tests := []struct {
		description string
		config      []latestV2.Validator
		shouldErr   bool
	}{
		{
			description: "no validation",
			config:      []latestV2.Validator{},
			shouldErr:   false,
		},
		{
			description: "kubeval validator",
			config: []latestV2.Validator{
				{Name: "kubeval"},
			},
			shouldErr: false,
		},
		{
			description: "invalid validator",
			config: []latestV2.Validator{
				{Name: "bad-validator"},
			},
			shouldErr: true,
		},
	}
	for _, test := range tests {
		testutil.Run(t, test.description, func(t *testutil.T) {
			_, err := NewValidator(test.config)
			t.CheckError(test.shouldErr, err)
		})
	}
}
