// Copyright 2022 lea
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

//package core contains the core data tyes and functions for the library
package core

type Channel interface {
}

type InputChannel interface {
	//Receive messages as long as thy come, if an error interrupts the process, return it
	Receive(chan Message) error
}

type OutputChannel interface {
	Send(Message) error
}

type IOChannel interface {
	InputChannel
	OutputChannel
}
