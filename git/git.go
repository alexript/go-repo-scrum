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

package git

import (
	"os"
	"os/exec"
	"strings"
)

type Git struct {
	Dir    string
	gitCmd string
}

func New(gitpath string, repopath string) *Git {
	return &Git{
		Dir:    repopath,
		gitCmd: gitpath,
	}
}

func Find(repopath string) (*Git, error) {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return nil, err
	}
	return New(gitPath, repopath), nil
}

func (g *Git) exec(args []string) (string, error) {
	cmd := exec.Command(g.gitCmd, args...)
	cmd.Dir = g.Dir
	cmd.Env = os.Environ()
	out, err := cmd.Output()
	if err != nil {
		return "", err
	} else if out == nil {
		return "", exec.ErrNotFound
	}
	str := string(out)
	return strings.TrimSpace(str), nil

}

func (g *Git) GetUsername() (string, error) {
	return g.exec([]string{"config", "--get", "user.name"})
}

func (g *Git) GetUseremail() (string, error) {
	return g.exec([]string{"config", "--get", "user.email"})
}
