package fsh

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

type FSItemType uint

const (
	FSItemDir FSItemType = iota
	FSItemFile
	FSItemLink

	FileDataPrefix = "123ABCabcАБВабв"
)

type FSItem struct {
	FullPath string
	ItemName string
	LinkPath string
	ItemType FSItemType
	children map[string]*FSItem
}

func CreateRoot(t *testing.T, rootPath string, children ...*FSItem) *FSItem {
	res := &FSItem{
		FullPath: rootPath,
		ItemType: FSItemDir,
	}

	res.addChildren(children)
	res.create(t, rootPath, rootPath)

	return res
}

func CreateDir(name string, children ...*FSItem) *FSItem {
	res := &FSItem{
		ItemName: name,
		ItemType: FSItemDir,
	}
	res.addChildren(children)

	return res
}

func CreateFile(name string) *FSItem {
	return &FSItem{
		ItemName: name,
		ItemType: FSItemFile,
	}
}

func CreateLink(name string, link string) *FSItem {
	return &FSItem{
		ItemName: name,
		LinkPath: link,
		ItemType: FSItemLink,
	}
}

func (ci *FSItem) Get(name string) *FSItem {
	if item, ok := ci.children[name]; ok {
		return item
	}

	return nil
}

func (ci *FSItem) CountChildren() int {
	return len(ci.children)
}

func (ci *FSItem) addChildren(items []*FSItem) {
	ci.children = make(map[string]*FSItem)
	for _, it := range items {
		ci.children[it.ItemName] = it
	}
}

func (ci *FSItem) create(t *testing.T, rootPath string, parentPath string) {
	if ci.ItemType == FSItemFile {
		ci.FullPath = filepath.Join(parentPath, ci.ItemName)
		require.NoError(t, os.WriteFile(ci.FullPath, []byte(FileDataPrefix+ci.FullPath), 0777))
		return
	}

	if ci.ItemType == FSItemDir {
		ci.FullPath = filepath.Join(parentPath, ci.ItemName)
		require.NoError(t, os.Mkdir(ci.FullPath, 0777))
		for _, child := range ci.children {
			child.create(t, rootPath, ci.FullPath)
		}
		return
	}

	if ci.ItemType == FSItemLink {
		ci.FullPath = filepath.Join(parentPath, ci.ItemName)
		srcPath := filepath.Join(rootPath, ci.LinkPath)
		require.NoError(t, os.Symlink(srcPath, ci.FullPath))
		return
	}
}
