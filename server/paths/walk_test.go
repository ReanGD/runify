package paths

import (
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
	t := s.T()

	s.rootDir = filepath.Join(os.TempDir(), "walk_suite")
	s.removeAll()
	s.rootItem = fsh.CreateRoot(t, s.rootDir,
		fsh.CreateFile("file_01"),
		fsh.CreateFile("file_02"),
		fsh.CreateDir("dir_1",
			fsh.CreateFile("file_11"),
			fsh.CreateLink("file_link_01", "file_01"),
			fsh.CreateLink("file_link_02", "file_02"),
			fsh.CreateLink("file_link_link_01", "file_link_01"),
			fsh.CreateLink("file_link_link_02", "file_link_02"),
			fsh.CreateLink("file_link_no_exists", "file_no_exists"),
			fsh.CreateLink("file_link_link_no_exists", "file_link_no_exists"),
			fsh.CreateDir("dir_2",
				fsh.CreateFile("file_21"),
				fsh.CreateFile("file_22"),
			),
		),
		fsh.CreateLink("link_dir_1", "dir_1"),
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
	entries, err := readDir(path)
	s.Require().NoError(err)
	for _, entry := range entries {
		child := item.Get(entry.Name())
		s.Require().NotNil(child)
		switch child.ItemType {
		case fsh.FSItemFile:
			s.Require().Equal(fs.FileMode(0x0), entry.Type())
		case fsh.FSItemDir:
			s.Require().Equal(os.ModeDir, entry.Type())
		case fsh.FSItemLink:
			s.Require().Equal(os.ModeSymlink, entry.Type())
		}
	}

	s.Require().Equal(item.CountChildren(), len(entries))
}

func (s *WalkSuite) checkReadDir(item *fsh.FSItem) {
	s.checkReadLinkDir(item.FullPath, item)
}

func (s *WalkSuite) TestReadDir() {
	s.checkReadDir(s.rootItem)
	s.checkReadDir(s.rootItem.Get("dir_1"))
	s.checkReadDir(s.rootItem.Get("dir_1").Get("dir_2"))
	s.checkReadLinkDir(s.rootItem.Get("link_dir_1").FullPath, s.rootItem.Get("dir_1"))
	s.checkReadLinkDir(s.rootItem.Get("link_link_dir_1").FullPath, s.rootItem.Get("dir_1"))

	_, err := readDir(s.rootItem.Get("link_dir_no_exists").FullPath)
	s.Require().Error(err)

	_, err = readDir(s.rootItem.Get("file_link_link_no_exists").FullPath)
	s.Require().Error(err)

	_, err = readDir("~")
	s.Require().Error(err)
}

func TestWalkSuite(t *testing.T) {
	suite.Run(t, new(WalkSuite))
}
