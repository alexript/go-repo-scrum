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

import "testing"

type parent struct {
	id  Id
	idx uint
}

func (p *parent) GetID() Id {
	return p.id
}

type child struct {
	id     Id
	parent Id
}

func (c *child) GetID() Id {
	return c.id
}

func (c *child) Parent() Id {
	return c.parent
}

func createTestCase() (Identables, Children) {
	p := make(Identables, 4)

	p[0] = &parent{
		id:  1,
		idx: 0,
	}
	p[1] = &parent{
		id:  4,
		idx: 1,
	}
	p[2] = &parent{
		id:  3,
		idx: 2,
	}
	p[3] = &parent{
		id:  2,
		idx: 3,
	}

	c := make(Children, 4)
	c[0] = &child{
		id:     1,
		parent: 2,
	}
	c[1] = &child{
		id:     2,
		parent: 1,
	}
	c[2] = &child{
		id:     3,
		parent: 2,
	}
	c[3] = &child{
		id:     4,
		parent: 5,
	}

	return p, c
}

func TestParent(t *testing.T) {
	_, c := createTestCase()
	id := c[2].Parent()
	if 2 != id {
		t.Fatalf("Unexpected Parent id %d", id)
	}
}

func TestBrothers(t *testing.T) {
	_, c := createTestCase()

	brothers := c.Brothers(2)
	if len(brothers) != 2 {
		t.Fatalf("This numbr of brothers is unexpected: %d", len(brothers))
	}
}

func TestGetParent(t *testing.T) {
	p, c := createTestCase()
	existed := p.Get(c[2].Parent())
	notexisted := p.Get(c[3].Parent())
	if existed == nil {
		t.Fatalf("Existed not found")
	}
	if notexisted != nil {
		t.Fatalf("Found not existed %v+", notexisted)
	}
	if 3 != existed.(*parent).idx {
		t.Fatalf("Found wrong parent: %d", existed.GetID())
	}

}
