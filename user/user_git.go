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

package user

import (
	"errors"
	"net/mail"
	"os"
	"path"

	"github.com/alexript/go-repo-scrum/fs"
	"github.com/alexript/go-repo-scrum/git"
)

var (
	errNoGit    = errors.New("Git client not found")
	errNoDotgit = errors.New("This is not a git repo")
	errGitError = errors.New("Git error")
)

func gitCurrentUser(root *fs.Root, user *User) error {
	dotgitdir := path.Join(root.Repodir, ".git")
	stat, err := os.Stat(dotgitdir)
	if err != nil {
		return errors.Join(errNoDotgit, err)
	}
	if !stat.IsDir() {
		return errNoDotgit
	}

	gitcli, err := git.Find(root.Repodir)
	if err != nil {
		return errors.Join(errNoGit, err)
	}

	username, err := gitcli.GetUsername()
	if err != nil {
		return errors.Join(errGitError, err)
	}
	user.Name = username

	email, err := gitcli.GetUseremail()
	if err != nil {
		return errors.Join(errGitError, err)
	}

	address, err := mail.ParseAddress(email)
	if err == nil {
		user.Login = address.Name
	}

	if root.Ini.Bool("user.email.enable", false) {
		user.Email = email
	} else {
		user.Email = ""
	}

	return errGitError
}
