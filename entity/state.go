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

type State struct {
	Number uint32 `json:"id"`
	Weight uint32 `json:"weight"`
	Name   string `json:"name"`
}

func defaultStateArray() []State {
	return []State{
		{
			Number: 0,
			Weight: 0,
			Name:   "Draft",
		},
		{
			Number: 1,
			Weight: 100,
			Name:   "New",
		},
		{
			Number: 2,
			Weight: 200,
			Name:   "Open",
		},
		{
			Number: 3,
			Weight: 300,
			Name:   "InProgress",
		},
		{
			Number: 4,
			Weight: 400,
			Name:   "Done",
		},
		{
			Number: 5,
			Weight: 500,
			Name:   "Cancel",
		},
		{
			Number: 6,
			Weight: 10000,
			Name:   "Complete",
		},
	}
}
