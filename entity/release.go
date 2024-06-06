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

import "math"

type Release struct {
	Number    Id        `json:"id"`
	StartDate Timestamp `json:"start"`
	EndDate   Timestamp `json:"end"`
	Done      bool      `json:"done"`
	Title     string    `json:"title"`
}

func (r *Release) GetID() Id {
	return r.Number
}

func defaultReleaseArray() []Release {
	return []Release{
		{ // this is special Release: Backlog
			Number:    0,
			StartDate: math.MinInt64,
			EndDate:   math.MaxInt64,
			Done:      false,
			Title:     "Backlog",
		},
		{ // this is special Release: Canceled
			Number:    math.MaxUint32,
			StartDate: math.MinInt64,
			EndDate:   math.MaxInt64,
			Done:      false,
			Title:     "Canceled",
		},
	}
}

func (r *Release) Backlog() bool {
	return r.Number == 0
}

func (r *Release) Canceled() bool {
	return r.Number == math.MaxUint32
}

func (r *Release) Predefined() bool {
	return r.Backlog() || r.Canceled()
}
