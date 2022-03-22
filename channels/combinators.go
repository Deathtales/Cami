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

package channels

import (
	"github.com/Deathtales/Cami/core"
	"github.com/Deathtales/Cami/utils"
)

type BroadCastChannel []core.OutputChannel

func (b BroadCastChannel) Send(m core.Message) error {
	var errors utils.CompositeError = nil

	for _, c := range b {
		errors.Check(c.Send(m))
	}
	return errors
}

func Tee(o1, o2 core.OutputChannel) core.OutputChannel {
	return BroadCastChannel{o1, o2}
}
