// Copyright 2021 The Ebiten Authors
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

package audio

import (
	"io"

	"github.com/ebitengine/oto/v3"
)

func newContext(sampleRate int) (context, chan struct{}, error) {
	ctx, ready, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   sampleRate,
		ChannelCount: channelCount,
		Format:       oto.FormatFloat32LE,
	})
	err = addErrorInfo(err)
	return &contextProxy{ctx}, ready, err
}

// contextProxy is a proxy between otoContext and context.
type contextProxy struct {
	*oto.Context
}

// NewPlayer implements context.
func (c *contextProxy) NewPlayer(r io.Reader) player {
	return c.Context.NewPlayer(r)
}
