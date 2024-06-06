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

type Id uint32

type Identable interface {
	GetID() Id
}

type Identables []Identable

func (a Identables) First() Identable {
	var first Identable
	for _, v := range a {
		if v == nil {
			continue
		}
		if first == nil {
			first = v

		} else if v.GetID() < first.GetID() {
			first = v
		}
	}
	return first
}

func (a Identables) Last() Identable {
	var last Identable
	for _, v := range a {
		if v == nil {
			continue
		}
		if last == nil {
			last = v
		} else if v.GetID() > last.GetID() {
			last = v
		}
	}
	return last
}

func (a Identables) CreateID() Id {
	last := a.Last()
	id := last.GetID()
	return id + 1
}

func (a Identables) Contain(id Id) bool {

	for _, v := range a {
		if v == nil {
			continue
		}
		if id == v.GetID() {
			return true
		}
	}
	return false
}

func (a Identables) Get(id Id) Identable {

	for _, v := range a {
		if v == nil {
			continue
		}
		if id == v.GetID() {
			return v
		}
	}
	return nil
}
