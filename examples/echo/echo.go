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

package main

import (
	"github.com/Deathtales/Cami/channels"
	"github.com/Deathtales/Cami/core"
)

func main() {
	messages := make(chan core.Message)

	go channels.StdInChannel.Receive(messages)

	for m := range messages {
		channels.StdOutchannel.Send(m)
	}

}
