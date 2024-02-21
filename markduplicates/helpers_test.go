// Copyright 2019 Grail Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package markduplicates

import (
	"testing"

	"github.com/Schaudge/hts/sam"
	"github.com/stretchr/testify/assert"
)

func TestClearDupFlagTags(t *testing.T) {
	r := sam.GetFromFreePool()
	r.Name = "A"
	r.Ref = chr1
	r.Pos = 10
	r.MatePos = 20
	r.MateRef = chr1
	r.Flags = r1F | sam.Duplicate
	r.AuxFields = sam.AuxFields{}

	// Insert duplicate tags, separated by other tags.
	for i, tag := range []string{"RG", "DI", "VN", "DS", "SM", "DT", "PU", "DU", "XM"} {
		aux, err := sam.NewAux(sam.NewTag(tag), i)
		assert.Nil(t, err)
		r.AuxFields = append(r.AuxFields, aux)
	}

	clearDupFlagTags(r)

	// Verify flag 1024 has been cleared.
	assert.Equal(t, r1F, r.Flags)

	// Verify that duplicate tags are absent.
	expected := []struct {
		tagCode string
		value   int
	}{{"RG", 0}, {"VN", 2}, {"SM", 4}, {"PU", 6}, {"XM", 8}}
	assert.Equal(t, len(expected), len(r.AuxFields))
	for i, e := range expected {
		aux, err := sam.NewAux(sam.NewTag(e.tagCode), e.value)
		assert.Nil(t, err)
		assert.Equal(t, aux, r.AuxFields[i])
	}
}
