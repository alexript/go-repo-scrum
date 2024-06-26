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

package runtime

import (
	"errors"
	"fmt"
	rt "runtime"

	"github.com/alexript/go-repo-scrum/entity"
	"github.com/alexript/go-repo-scrum/fs"
	"github.com/alexript/go-repo-scrum/user"
)

var (
	catalog *entity.Catalog
	root    *fs.Root
	uid     entity.Id
)

func startUI(root *fs.Root) {
	fmt.Println("UI is not implemented yet.")
}

func Start(withUI bool) {
	root = fs.Get()
	onStart(root)
	if withUI {
		startUI(root)
	}
}

func StartIn(dirname string, withUI bool) {
	root = fs.GetIn(dirname)
	onStart(root)
	if withUI {
		startUI(root)
	}
}

func Stop() {
	fs.Close(root)
}

func onStart(root *fs.Root) {

	if err := root.Prepare(); err != nil {
		panic(err)
	}

	catalog, err := root.OpenCatalog()
	if err != nil {
		panic(err)
	}
	defer root.SaveCatalog(catalog)

	if catalog == nil {
		panic(errors.New("Something strange is happens now..."))
	}
	usr := user.CurrentUser(root)
	uid = catalog.EnsurePersone(usr.Name, usr.Login, usr.Email)
	rt.GC()
}

func Catalog() *entity.Catalog {
	return catalog
}

func UID() entity.Id {
	return uid
}
