package paths

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type itemType uint

const (
	itemDir itemType = iota
	itemFile
	itemLink

	fileDataPrefix = "123ABCabcАБВабв"
)

type createItem struct {
	fullPath string
	itemName string
	link     string
	itemType itemType
	children map[string]*createItem
}

func createRoot(t *testing.T, rootpath string, children ...*createItem) *createItem {
	res := &createItem{
		fullPath: rootpath,
		itemType: itemDir,
	}

	res.addChildren(children)
	res.create(t, rootpath, rootpath)

	return res
}

func createDir(name string, children ...*createItem) *createItem {
	res := &createItem{
		itemName: name,
		itemType: itemDir,
	}
	res.addChildren(children)

	return res
}

func createFile(name string) *createItem {
	return &createItem{
		itemName: name,
		itemType: itemFile,
	}
}

func createLink(name string, link string) *createItem {
	return &createItem{
		itemName: name,
		link:     link,
		itemType: itemLink,
	}
}

func (ci *createItem) get(name string) *createItem {
	return ci.children[name]
}

func (ci *createItem) addChildren(items []*createItem) {
	ci.children = make(map[string]*createItem)
	for _, it := range items {
		ci.children[it.itemName] = it
	}
}

func (ci *createItem) create(t *testing.T, rootPath string, parentPath string) {
	if ci.itemType == itemFile {
		ci.fullPath = filepath.Join(parentPath, ci.itemName)
		require.NoError(t, os.WriteFile(ci.fullPath, []byte(fileDataPrefix+ci.fullPath), 0777))
		return
	}

	if ci.itemType == itemDir {
		ci.fullPath = filepath.Join(parentPath, ci.itemName)
		require.NoError(t, os.Mkdir(ci.fullPath, 0777))
		for _, child := range ci.children {
			child.create(t, rootPath, ci.fullPath)
		}
		return
	}

	if ci.itemType == itemLink {
		ci.fullPath = filepath.Join(parentPath, ci.itemName)
		srcPath := filepath.Join(rootPath, ci.link)
		require.NoError(t, os.Symlink(srcPath, ci.fullPath))
		return
	}
}

type WalkSuite struct {
	suite.Suite

	rootDir  string
	rootItem *createItem
}

func (s *WalkSuite) SetupSuite() {
	t := s.T()

	s.rootDir = filepath.Join(os.TempDir(), "walk_suite")
	s.removeAll()
	s.rootItem = createRoot(t, s.rootDir,
		createFile("file_1"),
		createFile("file_2"),
		createDir("dir_1",
			createFile("file_11"),
			createLink("file_link_2", "file_2"),
			createLink("file_link_link_2", "file_link_2"),
			createLink("file_link_no_exists", "file_no_exists"),
			createLink("file_link_link_no_exists", "file_link_no_exists"),
		),
		createLink("link_dir_1", "dir_1"),
		createLink("link_link_dir_1", "link_dir_1"),
		createLink("link_dir_no_exists", "dir_no_exists"),
		createLink("file_link_link_no_exists", "link_dir_no_exists"),
	)
}

func (s *WalkSuite) TearDownSuite() {
	s.removeAll()
}

func (s *WalkSuite) removeAll() {
	require.NoError(s.T(), os.RemoveAll(s.rootDir))
}

func (s *WalkSuite) checkReadLinkDir(path string, item *createItem) {
	entries, err := readDir(path)
	s.Require().NoError(err)
	for _, entry := range entries {
		child, ok := item.children[entry.Name()]
		s.Require().True(ok)
		switch child.itemType {
		case itemFile:
			s.Require().Equal(fs.FileMode(0x0), entry.Type())
		case itemDir:
			s.Require().Equal(os.ModeDir, entry.Type())
		case itemLink:
			s.Require().Equal(os.ModeSymlink, entry.Type())
		}
	}
}

func (s *WalkSuite) checkReadDir(item *createItem) {
	s.checkReadLinkDir(item.fullPath, item)
}

func (s *WalkSuite) TestReadDir() {
	s.checkReadDir(s.rootItem)
	s.checkReadDir(s.rootItem.get("dir_1"))
	s.checkReadLinkDir(s.rootItem.get("link_dir_1").fullPath, s.rootItem.get("dir_1"))
	s.checkReadLinkDir(s.rootItem.get("link_link_dir_1").fullPath, s.rootItem.get("dir_1"))

	_, err := readDir(s.rootItem.get("link_dir_no_exists").fullPath)
	s.Require().Error(err)

	_, err = readDir(s.rootItem.get("file_link_link_no_exists").fullPath)
	s.Require().Error(err)

	_, err = readDir("~")
	s.Require().Error(err)
}

func TestWalkSuite(t *testing.T) {
	suite.Run(t, new(WalkSuite))
}
