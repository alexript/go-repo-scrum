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

import (
	"slices"
)

type Persone struct {
	Number Id     `json:"id"`
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Email  string `json:"email"`
}

func defaultPersoneArray() []Persone {
	return []Persone{}
}

func (p *Persone) GetID() Id {
	return p.Number
}

func (c *Catalog) FindPersoneByAlias(alias string) *Persone {

	idx := slices.IndexFunc(c.Persone, func(pers Persone) bool {
		return pers.Alias == alias
	})
	if idx == -1 {
		return nil
	}
	return &c.Persone[idx]
}

func (c *Catalog) getLastPersone() Persone {
	return slices.MaxFunc(c.Persone, func(p1 Persone, p2 Persone) int {
		return int(p1.GetID() - p2.GetID())
	})
}

func (c *Catalog) newPersone(name string, alias string, email string) ([]Persone, Id) {
	rec := c.FindPersoneByAlias(alias)
	if rec != nil {
		return nil, rec.GetID()
	}
	var id Id
	if len(c.Persone) < 1 {
		id = 0
	} else {
		last := c.getLastPersone()
		id = last.GetID() + 1
	}
	newp := append(c.Persone, Persone{
		Number: id,
		Name:   name,
		Alias:  alias,
		Email:  email,
	})
	return newp, id
}

func (c *Catalog) EnsurePersone(name string, login string, email string) Id {
	slise, id := c.newPersone(name, login, email)
	if slise != nil {
		c.Persone = slise
		c.changed = true
	}
	return id
}

func (c *Catalog) FindPersone(id Id) *Persone {
	idx := slices.IndexFunc(c.Persone, func(pers Persone) bool { return pers.GetID() == id })
	if idx == -1 {
		return nil
	}
	return &c.Persone[idx]
}
