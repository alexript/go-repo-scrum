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

type testaleId struct {
	Id Id
}

func (id *testaleId) GetID() Id {
	return id.Id
}

func TestPointers(t *testing.T) {
	var p *testaleId = new(testaleId)
	var o testaleId = *p
	o.Id = 42

	var i1 Identable = &o
	var i2 Identable = &o

	if i1.GetID() != i2.GetID() {
		t.Fatalf("WTF %d!=%d", i1.GetID(), i2.GetID())
	}

	if i1 != i2 {
		t.Fatalf("Different pointer values %v+!=%v+", i1, i2)
	}
}

func createTestIdentables() Identables {

	idArray := make(Identables, 4)
	idArray[0] = &testaleId{
		Id: 2,
	}
	idArray[1] = &testaleId{
		Id: 1,
	}
	idArray[2] = &testaleId{
		Id: 4,
	}
	idArray[3] = &testaleId{
		Id: 3,
	}
	return idArray
}

func TestFirst(t *testing.T) {
	idArray := createTestIdentables()
	first := idArray.First()
	if first == nil {
		t.Fatal("NIL result")
	}

	if first.GetID() != 1 {
		t.Fatalf("Expected first with id==1, but result is %d", first.GetID())
	}

	var p1 Identable = first
	var p2 Identable = idArray[1]

	if p1 != p2 {
		t.Fatalf("Different pointers %v+!=%v+", p1, p2)
	}
}
func TestLast(t *testing.T) {
	idArray := createTestIdentables()
	last := idArray.Last()
	if last == nil {
		t.Fatal("NIL result")
	}

	if last.GetID() != 4 {
		t.Fatalf("Expected first with id==1, but result is %d", last.GetID())
	}

	var p1 Identable = last
	var p2 Identable = idArray[2]

	if p1 != p2 {
		t.Fatalf("Different pointers %v+!=%v+", p1, p2)
	}
}

func TestCreateID(t *testing.T) {
	idArray := createTestIdentables()
	newID := idArray.CreateID()
	if 5 != newID {
		t.Fatalf("Unexpected new id: %d", newID)
	}
}

func TestContain(t *testing.T) {
	idArray := createTestIdentables()
	if !idArray.Contain(1) {
		t.Fatalf("Can't find id of existed object")
	}

	if idArray.Contain(42) {
		t.Fatalf("Found unexisted object")
	}
}
