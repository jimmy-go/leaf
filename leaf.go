package leaf

// TODO; read file and fill matrix or map.

// Open reads the specified file.
func Open(name string) (*Leaf, error) {
	// TODO; proposal: read file and fill data.
	return nil, nil
}

// Leaf type.
type Leaf struct {
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
