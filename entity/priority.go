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

type Priority struct {
	Number Id     `json:"id"`
	Name   string `json:"name"`
	Weight Weight `json:"weight"`
}

func (p *Priority) GetID() Id {
	return p.Number
}

func defaultPriorityArray() []Priority {
	return []Priority{
		{
			Number: 0,
			Name:   "Trivial",
			Weight: 0,
		},

		{
			Number: 1,
			Name:   "Minor",
			Weight: 1000,
		},

		{
			Number: 2,
			Name:   "Normal",
			Weight: 2000,
		},
		{
			Number: 3,
			Name:   "Major",
			Weight: 3000,
		},
		{
			Number: 4,
			Name:   "Critical",
			Weight: 10000,
		},
		{
			Number: 5,
			Name:   "Blocker",
			Weight: 1000000,
		},
	}
}
