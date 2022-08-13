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

func (s *ExistsSuite) checkExists(name string, expectedIsDir bool, expectedIsFile bool, expectedIsSymlink bool, expectedIsAny bool) {
	t := s.T()

	var err error
	var actual bool

	actual, err = ExistsDir(name)
	require.NoError(t, err)
	require.Equal(t, expectedIsDir, actual)

	actual, err = ExistsFile(name)
	require.NoError(t, err)
	require.Equal(t, expectedIsFile, actual)

	actual, err = ExistsSymlink(name)
	require.NoError(t, err)
	require.Equal(t, expectedIsSymlink, actual)

	actual, err = Exists(name)
	require.NoError(t, err)
	require.Equal(t, expectedIsAny, actual)
}

func (s *ExistsSuite) checkRead(name string, expectedData []byte) {
	t := s.T()

	actualData, err := os.ReadFile(name)
	if len(expectedData) != 0 {
		require.NoError(t, err)
		require.Equal(t, expectedData, actualData)
	} else {
		require.Error(t, err)
	}
}

func (s *ExistsSuite) TestExistsDir() {
	name := s.rootItem.Get("exists_dir").FullPath
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsDir() {
	name := s.rootItem.Join("no_exists_dir")
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsSymlinkDir() {
	name := s.rootItem.Get("exists_symlink_dir").FullPath
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsSymlinkDir() {
	name := s.rootItem.Get("no_exists_symlink_dir").FullPath
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsSymlinkToSymlinkDir() {
	name := s.rootItem.Get("exists_symlink_to_symlink_dir").FullPath
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsSymlinkToSymlinkDir() {
	name := s.rootItem.Get("no_exists_symlink_to_symlink_dir").FullPath
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsFile() {
	item := s.rootItem.Get("exists_dir").Get("exists.file")
	name := item.FullPath
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := false
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, item.Data())
}

func (s *ExistsSuite) TestNoExistsFile() {
	name := s.rootItem.Get("exists_dir").Join("no_exists.file")
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, []byte{})
}

func (s *ExistsSuite) TestExistsSymlinkFile() {
	name := s.rootItem.Get("with_symlinks_dir").Get("exists_symlink.file").FullPath
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, s.rootItem.Get("exists_dir").Get("exists.file").Data()) // link to s.existsFile
}

func (s *ExistsSuite) TestNoExistsSymlinkFile() {
	name := s.rootItem.Get("with_symlinks_dir").Get("no_exists_symlink.file").FullPath
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, []byte{})
}

func (s *ExistsSuite) TestExistsSymlinkToSymlinkFile() {
	name := s.rootItem.Get("with_symlinks_dir").Get("exists_symlink_to_symlink.file").FullPath
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, s.rootItem.Get("exists_dir").Get("exists.file").Data()) // link to s.existsFile
}

func (s *ExistsSuite) TestNoExistsSymlinkToSymlinkFile() {
	name := s.rootItem.Get("with_symlinks_dir").Get("no_exists_symlink_to_symlink.file").FullPath
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, []byte{})
}

func TestExistsSuite(t *testing.T) {
	suite.Run(t, new(ExistsSuite))
}
