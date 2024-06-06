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

type Issue struct {
	ReleaseNuber   uint32 `json:"releaseid"`
	SprintNumber   uint32 `json:"sprintid"`
	TypeNumber     uint32 `json:"typeid"`
	ParentNumber   uint32 `json:"parentid"`
	Number         uint32 `json:"id"`
	Date           int64  `json:"date"` // NOTE: no date until 'New' state
	Title          string `json:"title"`
	ReporterNumber uint32 `json:"personid"`
	Description    string `json:"text"`
	PriorityNumber uint32 `json:"priorityid"`
	StateNumber    uint32 `json:"stateid"`
}
