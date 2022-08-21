package paths

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/ReanGD/runify/server/test/utils/fsh"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type WalkSuite struct {
	suite.Suite

	rootDir  string
	rootItem *fsh.FSItem
}

func (s *WalkSuite) SetupSuite() {
	Init()
	s.rootDir = filepath.Join(os.TempDir(), "walk_suite")
	s.removeAll()
	s.rootItem = fsh.CreateRoot(s.T(), s.rootDir,
		fsh.CreateFile("file_01"),
		fsh.CreateFile("file_02"),
		fsh.CreateDir("dir_1",
			fsh.CreateFile("file_11"),
			fsh.CreateLink("file_link_01", "file_01"),
			fsh.CreateLink("file_link_02", "file_02"),
			fsh.CreateLink("file_link_11", "./file_11"),
			fsh.CreateLink("file_link_link_01", "dir_1/file_link_01"),
			fsh.CreateLink("file_link_link_02", "dir_1/file_link_02"),
			fsh.CreateLink("file_link_no_exists", "file_no_exists"),
			fsh.CreateLink("file_link_no_exists_relative", "./file_no_exists"),
			fsh.CreateLink("file_link_link_no_exists", "dir_1/file_link_no_exists"),
			fsh.CreateDir("dir_2",
				fsh.CreateFile("file_21"),
				fsh.CreateFile("file_22"),
			),
		),
		fsh.CreateLink("link_dir_1", "dir_1"),
		fsh.CreateLink("link_dir_2", "dir_1/dir_2"),
		fsh.CreateLink("link_link_dir_1", "link_dir_1"),
		fsh.CreateLink("link_dir_no_exists", "dir_no_exists"),
		fsh.CreateLink("file_link_link_no_exists", "link_dir_no_exists"),
	)
}

func (s *WalkSuite) TearDownSuite() {
	s.removeAll()
}

func (s *WalkSuite) removeAll() {
	require.NoError(s.T(), os.RemoveAll(s.rootDir))
}

func (s *WalkSuite) checkReadLinkDir(path string, item *fsh.FSItem) {
	t := s.T()

	entries, err := readDir(path)
	require.NoError(t, err)
	for _, entry := range entries {
		child := item.Get(entry.Name())
		require.NotNil(t, child)
		switch child.ItemType {
		case fsh.FSItemFile:
			require.Equal(t, fs.FileMode(0x0), entry.Type())
		case fsh.FSItemDir:
			require.Equal(t, os.ModeDir, entry.Type())
		case fsh.FSItemLink:
			require.Equal(t, os.ModeSymlink, entry.Type())
		}
	}

	require.Equal(t, item.CountChildren(), len(entries))
}

func (s *WalkSuite) checkReadDir(item *fsh.FSItem) {
	s.checkReadLinkDir(item.FullPath, item)
}

func (s *WalkSuite) checkWalkFiles(startItem *fsh.FSItem) {
	t := s.T()

	actualCnt := 0
	expected := startItem.GetExistChildrenRecursive(s.rootItem)
	Walk(startItem.FullPath, func(path string, mode PathMode) {
		_, ok := expected[path]
		require.True(t, ok, fmt.Sprintf("not found actual path %s inside expected", path))
		actualCnt++
	})

	require.Equal(t, len(expected), actualCnt)
}

func (s *WalkSuite) TestReadExistDirs() {
	s.checkReadDir(s.rootItem)
	s.checkReadDir(s.rootItem.Get("dir_1"))
	s.checkReadDir(s.rootItem.Get("dir_1").Get("dir_2"))
	s.checkReadLinkDir(s.rootItem.Get("link_dir_1").FullPath, s.rootItem.Get("dir_1"))
	s.checkReadLinkDir(s.rootItem.Get("link_link_dir_1").FullPath, s.rootItem.Get("dir_1"))
}

func (s *WalkSuite) TestReadNoExistDirs() {
	t := s.T()

	_, err := readDir(s.rootItem.Get("link_dir_no_exists").FullPath)
	require.Error(t, err)

	_, err = readDir(s.rootItem.Get("file_link_link_no_exists").FullPath)
	require.Error(t, err)
}

func (s *WalkSuite) TestReadHomeDir() {
	t := s.T()

	_, err := readDir("~")
	require.Error(t, err)
}

func (s *WalkSuite) TestWalkFiles() {
	s.checkWalkFiles(s.rootItem)
	s.checkWalkFiles(s.rootItem.Get("dir_1"))
	s.checkWalkFiles(s.rootItem.Get("dir_1").Get("dir_2"))
}

func TestWalkSuite(t *testing.T) {
	suite.Run(t, new(WalkSuite))
}
