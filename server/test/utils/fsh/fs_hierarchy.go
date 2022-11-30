package fsh

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type FSItemType uint

const (
	FSItemDir FSItemType = iota
	FSItemFile
	FSItemLink

	fileDataPrefix = "123ABCabcАБВабв"
)

type FSItem struct {
	FullPath string
	ItemName string
	ItemType FSItemType
	linkPath string
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
		linkPath: link,
		ItemType: FSItemLink,
	}
}

func (ci *FSItem) Join(name string) string {
	return filepath.Join(ci.FullPath, name)
}

func (ci *FSItem) Data() []byte {
	return []byte(fileDataPrefix + ci.FullPath)
}

func (ci *FSItem) Get(name string) *FSItem {
	if item, ok := ci.children[name]; ok {
		return item
	}

	return nil
}

func (ci *FSItem) getChildrenRecursive(root *FSItem, parent *FSItem, fullpath string, items map[string]FSItemType) {
	if ci.ItemType != FSItemLink {
		items[fullpath] = ci.ItemType
		for _, item := range ci.children {
			item.getChildrenRecursive(root, ci, filepath.Join(fullpath, item.ItemName), items)
		}
		return
	}

	it := root
	linkPath := ci.linkPath
	if strings.HasPrefix(ci.linkPath, "./") {
		it = parent
		linkPath = ci.linkPath[2:]
	}

	for _, name := range strings.Split(linkPath, "/") {
		it = it.Get(name)
		if it == nil {
			// link no exists
			return
		}
	}

	it.getChildrenRecursive(root, parent, fullpath, items)
}

func (ci *FSItem) GetExistChildrenRecursive(root *FSItem) map[string]FSItemType {
	items := make(map[string]FSItemType)
	ci.getChildrenRecursive(root, ci, ci.FullPath, items)

	return items
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
		require.NoError(t, os.WriteFile(ci.FullPath, []byte(fileDataPrefix+ci.FullPath), 0o777))
		return
	}

	if ci.ItemType == FSItemDir {
		ci.FullPath = filepath.Join(parentPath, ci.ItemName)
		require.NoError(t, os.Mkdir(ci.FullPath, 0o777))
		for _, child := range ci.children {
			child.create(t, rootPath, ci.FullPath)
		}
		return
	}

	if ci.ItemType == FSItemLink {
		ci.FullPath = filepath.Join(parentPath, ci.ItemName)

		srcPath := ci.linkPath
		if !strings.HasPrefix(ci.linkPath, "./") {
			srcPath = filepath.Join(rootPath, ci.linkPath)
		}
		require.NoError(t, os.Symlink(srcPath, ci.FullPath))
		return
	}
}
