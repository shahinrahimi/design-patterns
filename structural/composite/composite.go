package composite

import "fmt"

type FileSystemItem interface {
  Show(indent string)
}

type File struct {
  name string
}

func (f *File) Show (indent string) {
  fmt.Println(indent + "- File:", f.name)
}

type Folder struct {
  name string
  children []FileSystemItem
}

func (f *Folder) Add(item FileSystemItem) {
  f.children = append(f.children, item)
}

func (f *Folder) Show(indent string) {
  fmt.Println(indent + "+ Folder:", f.name)
  for _,child := range f.children {
    child.Show(indent + " ")
  }
}

func Run() {
  file1 := &File{name: "file1.txt"}
  file2 := &File{name: "file2.txt"}
  file3 := &File{name: "file3.txt"}

  rootFolder := &Folder{name : "root"}
  subFolder := &Folder{name : "subFolder"}

  subFolder.Add(file2)
  subFolder.Add(file3)
  rootFolder.Add(file1)
  rootFolder.Add(subFolder)

  rootFolder.Show("")
}

