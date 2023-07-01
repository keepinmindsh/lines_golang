package pattern

import (
	"fmt"
	"testing"
)

type Inode interface {
	print(string)
	clone() Inode
}

type FileProtoType struct {
	name string
}

func (f *FileProtoType) print(indentation string) {
	fmt.Printf("%s %s %p\n", indentation, f.name, &f.name)
}

func (f FileProtoType) clone() Inode {
	return &FileProtoType{name: f.name + "_clone"}
}

type FolderProtoType struct {
	children []Inode
	name     string
}

func (f *FolderProtoType) print(indentation string) {
	fmt.Printf("%s %s %p \n", indentation, f.name, &f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

// clone 프로토 타입의 핵심이 되는 부분임.
// 그 이유는 clone 복제를 하는 방식이 으로 얕은 복사가 아닌 깊은 복사가 이뤄져야함!
func (f *FolderProtoType) clone() Inode {
	cloneFolder := &FolderProtoType{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func TestProtoTypeMain(t *testing.T) {
	file1 := &FileProtoType{name: "File1"}
	file2 := &FileProtoType{name: "File2"}
	file3 := &FileProtoType{name: "File3"}

	folder1 := &FolderProtoType{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &FolderProtoType{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("   ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("   ")
}
