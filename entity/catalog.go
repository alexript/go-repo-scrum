// MIT License
//
// Copyright (c) 2024 Alex 'Ript' Malyshev
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package entity

type Catalog struct {
	changed  bool
	Type     []Type     `json:"type"`
	State    []State    `json:"state"`
	Priority []Priority `json:"priority"`
	Persone  []Persone  `json:"persone"`
	Release  []Release  `json:"release"`
	Sprint   []Sprint   `json:"sprint"`
}

func CreateDefaultCatalog() *Catalog {
	return &Catalog{
		changed:  true,
		Type:     defaultTypeArray(),
		State:    defaultStateArray(),
		Priority: defaultPriorityArray(),
		Persone:  defaultPersoneArray(),
		Release:  defaultReleaseArray(),
		Sprint:   defaultSprintArray(),
	}
}

func (c *Catalog) Changed() bool {
	return c.changed
}

func (c *Catalog) MarkSaved() {
	c.changed = false
}
