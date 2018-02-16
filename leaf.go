package leaf

import (
	"io"
	"net/url"
)

// TODO; read file and fill matrix or map.

// Open reads the specified file.
func Open(name string) (*Leaf, error) {
	var f io.ReadCloser
	var err error

	// Validate file is remote.
	if u, er := url.Parse(name); er == nil && u.Scheme != "" {
		f, err = openRemoteFile(name)
		if err != nil {
			return nil, err
		}
		return &Leaf{file: f}, nil
	}

	// Validate file is local.
	f, err = openLocalFile(name)
	if err != nil {
		return nil, err
	}
	return &Leaf{file: f}, nil
}

// Leaf type.
type Leaf struct {
	file io.ReadCloser
	book *Book
}

// SheetByName returns sheet by name if exists.
func (le *Leaf) SheetByName(name string) (Sheeter, error) {
	return nil, nil
}

// SheetAt returns sheet by name if exists.
func (le *Leaf) SheetAt(name string) (Sheeter, error) {
	return nil, nil
}

// Book type.
type Book struct {
	// sheet index, cell column:row, value.
	data   map[uint16]map[string]string
	sheets []Sheeter
}

// Sheeter interface.
type Sheeter interface {
	Row(name string) ([]string, error)
	RowAt(i int) ([]string, error)

	Col(name string) ([]string, error)
	ColAt(i int) ([]string, error)
}
