package paths

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ExistsSuite struct {
	suite.Suite

	data                         string
	rootDir                      string
	existsDir                    string
	noExistsDir                  string
	existsSymlinkDir             string
	noExistsSymlinkDir           string
	existsSymlinkToSymlinkDir    string
	noExistsSymlinkToSymlinkDir  string
	withSymlinksDir              string
	existsFile                   string
	noExistsFile                 string
	existsSymlinkFile            string
	noExistsSymlinkFile          string
	existsSymlinkToSymlinkFile   string
	noExistsSymlinkToSymlinkFile string
}

func (s *ExistsSuite) SetupSuite() {
	t := s.T()

	s.data = "123ABCabcАБВабв"

	s.rootDir = filepath.Join(os.TempDir(), "exists_suite")
	s.removeAll()
	require.NoError(t, os.Mkdir(s.rootDir, 0777))

	s.existsDir = filepath.Join(s.rootDir, "exists_dir")
	require.NoError(t, os.Mkdir(s.existsDir, 0777))

	s.noExistsDir = filepath.Join(s.rootDir, "no_exists_dir")

	s.existsSymlinkDir = filepath.Join(s.rootDir, "exists_symlink_dir")
	require.NoError(t, os.Symlink(s.existsDir, s.existsSymlinkDir))

	s.noExistsSymlinkDir = filepath.Join(s.rootDir, "no_exists_symlink_dir")
	require.NoError(t, os.Symlink(s.noExistsDir, s.noExistsSymlinkDir))

	s.existsSymlinkToSymlinkDir = filepath.Join(s.rootDir, "exists_symlink_to_symlink_dir")
	require.NoError(t, os.Symlink(s.existsSymlinkDir, s.existsSymlinkToSymlinkDir))

	s.noExistsSymlinkToSymlinkDir = filepath.Join(s.rootDir, "no_exists_symlink_to_symlink_dir")
	require.NoError(t, os.Symlink(s.noExistsSymlinkDir, s.noExistsSymlinkToSymlinkDir))

	s.withSymlinksDir = filepath.Join(s.rootDir, "with_symlinks_dir")
	require.NoError(t, os.Mkdir(s.withSymlinksDir, 0777))

	s.existsFile = filepath.Join(s.existsDir, "exists.file")
	require.NoError(t, os.WriteFile(s.existsFile, []byte(s.data+s.existsFile), 0777))

	s.noExistsFile = filepath.Join(s.rootDir, "no_exists.file")

	s.existsSymlinkFile = filepath.Join(s.withSymlinksDir, "exists_symlink.file")
	require.NoError(t, os.Symlink(s.existsFile, s.existsSymlinkFile))

	s.noExistsSymlinkFile = filepath.Join(s.withSymlinksDir, "no_exists_symlink.file")
	require.NoError(t, os.Symlink(s.noExistsFile, s.noExistsSymlinkFile))

	s.existsSymlinkToSymlinkFile = filepath.Join(s.withSymlinksDir, "exists_symlink_to_symlink.file")
	require.NoError(t, os.Symlink(s.existsSymlinkFile, s.existsSymlinkToSymlinkFile))

	s.noExistsSymlinkToSymlinkFile = filepath.Join(s.withSymlinksDir, "no_exists_symlink_to_symlink.file")
	require.NoError(t, os.Symlink(s.noExistsSymlinkFile, s.noExistsSymlinkToSymlinkFile))
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

func (s *ExistsSuite) checkRead(name string, expectedData string) {
	t := s.T()

	actualData, err := os.ReadFile(name)
	if len(expectedData) != 0 {
		require.NoError(t, err)
		require.Equal(t, []byte(expectedData), actualData)
	} else {
		require.Error(t, err)
	}
}

func (s *ExistsSuite) TestExistsDir() {
	name := s.existsDir
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsDir() {
	name := s.noExistsDir
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsSymlinkDir() {
	name := s.existsSymlinkDir
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsSymlinkDir() {
	name := s.noExistsSymlinkDir
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsSymlinkToSymlinkDir() {
	name := s.existsSymlinkToSymlinkDir
	expectedIsDir := true
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestNoExistsSymlinkToSymlinkDir() {
	name := s.noExistsSymlinkToSymlinkDir
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
}

func (s *ExistsSuite) TestExistsFile() {
	name := s.existsFile
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := false
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, s.data+s.existsFile)
}

func (s *ExistsSuite) TestNoExistsFile() {
	name := s.noExistsFile
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := false
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, "")
}

func (s *ExistsSuite) TestExistsSymlinkFile() {
	name := s.existsSymlinkFile
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, s.data+s.existsFile) // link to s.existsFile
}

func (s *ExistsSuite) TestNoExistsSymlinkFile() {
	name := s.noExistsSymlinkFile
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, "")
}

func (s *ExistsSuite) TestExistsSymlinkToSymlinkFile() {
	name := s.existsSymlinkToSymlinkFile
	expectedIsDir := false
	expectedIsFile := true
	expectedIsSymlink := true
	expectedIsAny := true

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, s.data+s.existsFile) // link to s.existsFile
}

func (s *ExistsSuite) TestNoExistsSymlinkToSymlinkFile() {
	name := s.noExistsSymlinkToSymlinkFile
	expectedIsDir := false
	expectedIsFile := false
	expectedIsSymlink := true
	expectedIsAny := false

	s.checkExists(name, expectedIsDir, expectedIsFile, expectedIsSymlink, expectedIsAny)
	s.checkRead(name, "")
}

func TestExistsSuite(t *testing.T) {
	suite.Run(t, new(ExistsSuite))
}
