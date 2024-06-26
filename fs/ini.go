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
	"path"

	"github.com/gookit/ini/v2"
)

const (
	iniFilename    = "go-repo-scrum.ini"
	defaultInifile = `
[user]
email.enable = false
	`
)

func readConfig(root *Root) *ini.Ini {
	i := ini.NewWithOptions(func(o *ini.Options) {
		o.Readonly = false
		o.ParseEnv = true
		o.ParseVar = true
		o.IgnoreCase = false
	})
	i.LoadStrings(defaultInifile)
	i.LoadExists(path.Join(root.Dirname, iniFilename))
	root.Ini = i
	return i
}

func writeConfig(root *Root) {
	root.Ini.WriteToFile(path.Join(root.Dirname, iniFilename))
}
