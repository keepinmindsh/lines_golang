package main

import (
	"fmt"
	"testing"
)

type Inode interface {
	print(string)
	clone() Inode
}

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Printf("%s %s %p\n", indentation, f.name, &f.name)
}

func (f File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Printf("%s %s %p \n", indentation, f.name, &f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

// clone 프로토 타입의 핵심이 되는 부분임.
// 그 이유는 clone 복제를 하는 방식이 으로 얕은 복사가 아닌 깊은 복사가 이뤄져야함!
func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func TestProtoTypeMain(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("   ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("   ")
}
