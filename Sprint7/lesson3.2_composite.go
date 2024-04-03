package main

import "fmt"

// Component — общий интерфейс для файлов и директорий.
type Component interface {
	Print(string)
	GetSize() int
}

type File2 struct {
	Name string
	Size int
}

func (f *File2) Print(prefix string) {
	fmt.Println(prefix+f.Name, f.Size)
}

func (f *File2) GetSize() int {
	return f.Size
}

type Dir struct {
	Name     string
	Children []Component
}

// Print печатает имя директории, её размер и содержимое.
func (d *Dir) Print(prefix string) {
	fmt.Println(prefix+d.Name, d.GetSize())
	for _, v := range d.Children {
		v.Print(prefix + "  ")
	}
}

// GetSize возвращает общий размер всех файлов в директории.
func (d *Dir) GetSize() int {
	var sum int
	for _, v := range d.Children {
		sum += v.GetSize()
	}
	return sum
}

func main() {
	root := &Dir{
		Children: []Component{
			&File2{Name: "file1", Size: 778},
			&File2{Name: "file2", Size: 222},
			&Dir{
				Children: []Component{
					&File2{Name: "file3", Size: 64},
					&File2{Name: "file4", Size: 36},
				},
				Name: "subfolder",
			},
		},
		Name: "root",
	}
	root.Print("")
}
