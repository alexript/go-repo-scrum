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

package fs

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/gookit/ini/v2"
	"golang.org/x/sys/unix"
)

const (
	rootdirName      = ".reposcrum"
	rootReadmeMdName = "README.md"
	rootReadmeMd     = `
# Reposcrum

This is directory with Scrum data of the repo
	`
)

type Root struct {
	Dirname string
	Ini     *ini.Ini
	Repodir string
}

func GetIn(dirname string) *Root {

	abspath, err := filepath.Abs(dirname)
	if err != nil {
		panic(err)
	}

	root := path.Join(abspath, rootdirName)

	return &Root{
		Dirname: root,
		Repodir: abspath,
	}
}
func Get() *Root {
	var dirname string
	if len(os.Args) > 1 {
		args := os.Args[1:]
		if len(args) > 0 {
			dirname = args[0]
		}
	}

	if len(dirname) == 0 {
		d, err := os.Getwd()
		if err != nil {
			exec, err := os.Executable()
			if err != nil {
				panic(err)
			}
			dirname = path.Dir(exec)
		} else {
			dirname = d
		}
	}
	return GetIn(dirname)
}

// TODO: ensure correct behavior under windows
func (root *Root) Prepare() error {
	stat, err := os.Stat(root.Dirname)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			err := os.MkdirAll(root.Dirname, 0775)
			if err != nil {
				return err
			}

			stat, err = os.Stat(root.Dirname)
			if err != nil {
				return err
			}
		}
	}

	if stat.IsDir() {
		if unix.Access(root.Dirname, unix.W_OK) != nil {
			return os.ErrPermission
		}

		readConfig(root)

		return root.checkReadmeMd()
	}

	return os.ErrInvalid
}

func Close(root *Root) {
	writeConfig(root)
}

func (root *Root) checkReadmeMd() error {
	filename := path.Join(root.Dirname, rootReadmeMdName)
	_, err := os.Stat(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if err != nil {
		tpl, err := template.New(rootReadmeMdName).Parse(rootReadmeMd)
		if err != nil {
			return err
		}
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		buffer := bufio.NewWriter(f)
		defer buffer.Flush()

		if err := tpl.Execute(buffer, nil); err != nil {
			return err
		}
	}
	return nil
}
