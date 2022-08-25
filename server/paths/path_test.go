package paths

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PathSuite struct {
	suite.Suite
}

func (s *PathSuite) SetupSuite() {
	New()
}

func (s *PathSuite) TestGetSysTmp() {
	t := s.T()

	require.Equal(t, os.TempDir(), GetSysTmp())
}

func (s *PathSuite) TestGetUserHome() {
	t := s.T()

	homeDir, err := os.UserHomeDir()
	require.NoError(t, err)
	require.Equal(t, homeDir, GetUserHome())
}

func (s *PathSuite) TestExpandUser() {
	t := s.T()

	homeDir, err := os.UserHomeDir()
	require.NoError(t, err)

	require.Equal(t, homeDir, ExpandUser("~"))
	require.Equal(t, homeDir+"/dir/dir/file.name", ExpandUser("~/dir/dir/file.name"))
}

func TestPathSuite(t *testing.T) {
	suite.Run(t, new(PathSuite))
}
