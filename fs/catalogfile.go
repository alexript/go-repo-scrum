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
	"encoding/json"
	"os"
	"path"

	"github.com/alexript/go-repo-scrum/entity"
)

const (
	catalogversion  = 1
	catalogFilename = ".catalog.json"
)

type CatalogFile struct {
	Version uint8          `json:"schemaversion"`
	Catalog entity.Catalog `json:"catalog"`
}

func (root *Root) SaveCatalog(catalog *entity.Catalog) error {
	if catalog == nil {
		return nil // TODO: assert?
	}

	catalogFile := CatalogFile{
		Version: catalogversion,
		Catalog: *catalog,
	}

	bytes, err := json.Marshal(catalogFile)
	if err != nil {
		return err
	}
	filename := path.Join(root.Dirname, catalogFilename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	buffer := bufio.NewWriter(f)
	defer buffer.Flush()
	_, err = buffer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (root *Root) OpenCatalog() (*entity.Catalog, error) {
	var catalog *entity.Catalog
	filename := path.Join(root.Dirname, catalogFilename)
	stats, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		catalog = entity.CreateDefaultCatalog()
		err := root.SaveCatalog(catalog)
		if err != nil {
			return catalog, err
		}
		return catalog, nil
	} else if err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(f)
	_, err = bufr.Read(bytes)
	if err != nil {
		return nil, err
	}

	catalogfile := &CatalogFile{}

	err = json.Unmarshal(bytes, catalogfile)
	if err != nil {
		return nil, err
	}
	return &catalogfile.Catalog, nil
}
