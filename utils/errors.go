// Copyright 2022 Lea WIllame
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import "fmt"

type CompositeError []error

func (c CompositeError) Error() string {

	errors := make([]interface{}, len(c))

	for i, err := range c {
		errors[i] = err.Error()
	}
	return fmt.Sprint(errors...)
}

func (c CompositeError) Check(err error) CompositeError {
	if err != nil {
		return append(c, err)
	} else {
		return c
	}
}
