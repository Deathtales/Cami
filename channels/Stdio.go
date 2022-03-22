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

//package channels contin example/basic channels for use within CAMI
package channels

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Deathtales/Cami/core"
	"github.com/Deathtales/Cami/utils"
)

type StreamInputChannel interface {
	core.InputChannel
	io.Closer
}

type StreamOutputChannel interface {
	core.OutputChannel
	io.Closer
}

//
type StreamIOChannel struct {
	input  io.Reader
	output io.Writer
}

//Function Close will close the underlying channels if able
func (s StreamIOChannel) Close() error {
	var errors utils.CompositeError
	if i, ok := s.input.(io.Closer); ok {
		errors.Check(i.Close())
	}
	if o, ok := s.output.(io.Closer); ok {
		errors.Check(o.Close())
	}
	return errors
}

func (s StreamIOChannel) Receive(output chan core.Message) error {

	if s.input == nil {
		return fmt.Errorf("this StreamIO channel doesn't accept Inputs")
	}

	scanner := bufio.NewScanner(s.input)

	for scanner.Scan() {
		line := scanner.Text()

		output <- core.Message{
			Origin:  s,
			Payload: line,
		}
	}

	return scanner.Err()
}

func (s StreamIOChannel) Send(m core.Message) error {
	if s.output == nil {
		return fmt.Errorf("this StreamIO channel doesn't accept outputs")
	}
	fmt.Fprint(s.output, m.Payload)
	return nil
}

var (
	StdInChannel StreamInputChannel = StreamIOChannel{
		input: os.Stdin,
	}
	StdOutchannel StreamOutputChannel = StreamIOChannel{
		output: os.Stdout,
	}
	StdErrChannel StreamOutputChannel = StreamIOChannel{
		output: os.Stdout,
	}
	StdIOChannel StreamIOChannel = StreamIOChannel{
		input:  os.Stdin,
		output: os.Stdout,
	}
)

func FileInputChannel(filename string) (core.InputChannel, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return StreamIOChannel{
		input: f,
	}, nil

}

func FileOutputChannel(filename string) (core.OutputChannel, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return StreamIOChannel{
		output: f,
	}, nil

}
