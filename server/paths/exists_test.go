package paths

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ReanGD/runify/server/test/utils/fsh"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ExistsSuite struct {
	suite.Suite

	rootDir  string
	rootItem *fsh.FSItem
}

func (s *ExistsSuite) SetupSuite() {
	New()
	s.rootDir = filepath.Join(os.TempDir(), "exists_suite")
	s.removeAll()
	s.rootItem = fsh.CreateRoot(s.T(), s.rootDir,
		fsh.CreateDir("exists_dir",
			fsh.CreateFile("exists.file"),
		),
		fsh.CreateLink("exists_symlink_dir", "exists_dir"),
		fsh.CreateLink("no_exists_symlink_dir", "no_exists_dir"),
		fsh.CreateLink("exists_symlink_to_symlink_dir", "exists_symlink_dir"),
		fsh.CreateLink("no_exists_symlink_to_symlink_dir", "no_exists_symlink_dir"),
		fsh.CreateDir("with_symlinks_dir",
			fsh.CreateLink("exists_symlink.file", "exists_dir/exists.file"),
			fsh.CreateLink("no_exists_symlink.file", "exists_dir/no_exists.file"),
			fsh.CreateLink("exists_symlink_to_symlink.file", "with_symlinks_dir/exists_symlink.file"),
			fsh.CreateLink("no_exists_symlink_to_symlink.file", "with_symlinks_dir/no_exists_symlink.file"),
		),
	)
}

func (s *ExistsSuite) TearDownSuite() {
	s.removeAll()
}

func (s *ExistsSuite) removeAll() {
	require.NoError(s.T(), os.RemoveAll(s.rootDir))
}

func (s *ExistsSuite) checkExists(path string, expectedIsDir bool, expectedIsFile bool, expectedIsSymlink bool, expectedIsAny bool) {
	t := s.T()

	var err error
	var actual bool

	actual, err = ExistsDir(path)
	require.NoError(t, err)
	require.Equal(t, expectedIsDir, actual, "ExistsDir return wrong value")

	actual, err = ExistsFile(path)
	require.NoError(t, err)
	require.Equal(t, expectedIsFile, actual, "ExistsFile return wrong value")

	actual, err = ExistsSymlink(path)
	require.NoError(t, err)
	require.Equal(t, expectedIsSymlink, actual, "ExistsSymlink return wrong value")

	actual, err = Exists(path)
	require.NoError(t, err)
	require.Equal(t, expectedIsAny, actual, "Exists return wrong value")
}

func (s *ExistsSuite) checkRead(path string, expectedData []byte) {
	t := s.T()

	actualData, err := os.ReadFile(path)
	if len(expectedData) != 0 {
		require.NoError(t, err)
		require.Equal(t, expectedData, actualData)
	} else {
		require.Error(t, err)
	}
}

func (s *ExistsSuite) TestExistsDir() {
	path := s.rootItem.Get("exists_dir").FullPath
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := true

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsDir() {
	path := s.rootItem.Join("no_exists_dir")
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := false

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsSymlinkDir() {
	path := s.rootItem.Get("exists_symlink_dir").FullPath
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsSymlinkDir() {
	path := s.rootItem.Get("no_exists_symlink_dir").FullPath
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsSymlinkToSymlinkDir() {
	path := s.rootItem.Get("exists_symlink_to_symlink_dir").FullPath
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsSymlinkToSymlinkDir() {
	path := s.rootItem.Get("no_exists_symlink_to_symlink_dir").FullPath
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsFile() {
	item := s.rootItem.Get("exists_dir").Get("exists.file")
	path := item.FullPath
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := false
	expectedIsAny := true

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(path, item.Data())
}

func (s *ExistsSuite) TestNoExistsFile() {
	path := s.rootItem.Get("exists_dir").Join("no_exists.file")
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := false

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(path, []byte{})
}

func (s *ExistsSuite) TestExistsSymlinkFile() {
	path := s.rootItem.Get("with_symlinks_dir").Get("exists_symlink.file").FullPath
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(path, s.rootItem.Get("exists_dir").Get("exists.file").Data()) // link to s.existsFile
}

func (s *ExistsSuite) TestNoExistsSymlinkFile() {
	path := s.rootItem.Get("with_symlinks_dir").Get("no_exists_symlink.file").FullPath
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(path, []byte{})
}

func (s *ExistsSuite) TestExistsSymlinkToSymlinkFile() {
	path := s.rootItem.Get("with_symlinks_dir").Get("exists_symlink_to_symlink.file").FullPath
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(path, s.rootItem.Get("exists_dir").Get("exists.file").Data()) // link to s.existsFile
}

func (s *ExistsSuite) TestNoExistsSymlinkToSymlinkFile() {
	path := s.rootItem.Get("with_symlinks_dir").Get("no_exists_symlink_to_symlink.file").FullPath
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(path, []byte{})
}

func (s *ExistsSuite) TestHomeDir() {
	path := "~"
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := true

	s.checkExists(path, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func TestExistsSuite(t *testing.T) {
	suite.Run(t, new(ExistsSuite))
}
