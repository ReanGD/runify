package desktop_entry

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ExecSuite struct {
	suite.Suite
}

func (s *ExecSuite) check(withTerminal bool, execStr string, expected []string) {
	t := s.T()

	actual, err := buildArgs(execStr, withTerminal, "bash")
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func (s *ExecSuite) TestArgs() {
	s.check(false, `gvim`, []string{"gvim"})
	s.check(false, `gvim test`, []string{"gvim", "test"})
}

func (s *ExecSuite) TestQuotes() {
	s.check(false, `"gvim" test`, []string{"gvim", "test"})
	s.check(false, `"gvim test"`, []string{"gvim test"})
	s.check(false, `vim ~/.vimrc`, []string{"vim", "~/.vimrc"})
	s.check(false, `vim '~/.vimrc test'`, []string{"vim", "~/.vimrc test"})
	s.check(false, `vim '~/.vimrc " test'`, []string{"vim", `~/.vimrc " test`})
}

func (s *ExecSuite) TestEscapeSequences() {
	s.check(false, `"gvim test" test2 "test \\" 3"`, []string{"gvim test", "test2", `test " 3`})
	s.check(false, `"test \\\\\\\\ \\" moin" test`, []string{`test \\ " moin`, "test"})
	s.check(false, `"gvim \\\\\\\\ \\`+"`"+`test\\$"`, []string{`gvim \\ ` + "`" + `test$`})
}

func (s *ExecSuite) TestTerminal() {
	s.check(true, `vim ~/.vimrc`, []string{"bash", "-e", "vim", "~/.vimrc"})
	s.check(false, `sh -c 'vim ~/.vimrc " test'`, []string{"sh", "-c", `vim ~/.vimrc " test`})
	s.check(false, `sh -c 'vim ~/.vimrc " test"'`, []string{"sh", "-c", `vim ~/.vimrc " test"`})
}

func (s *ExecSuite) TestValidFieldCodesWithRemoveThem() {
	s.check(false, `vim %u`, []string{"vim"})
	s.check(false, `vim ~/.vimrc %u`, []string{"vim", "~/.vimrc"})
	s.check(false, `vim '%u' ~/.vimrc`, []string{"vim", "%u", "~/.vimrc"})
	s.check(false, `vim %u ~/.vimrc`, []string{"vim", "~/.vimrc"})
	s.check(false, `vim /%u/.vimrc`, []string{"vim", "//.vimrc"})
	s.check(false, `vim %u/.vimrc`, []string{"vim", "/.vimrc"})
	s.check(false, `vim %U/.vimrc`, []string{"vim", "/.vimrc"})
	s.check(false, `vim /%U/.vimrc`, []string{"vim", "//.vimrc"})
	s.check(false, `vim %U .vimrc`, []string{"vim", ".vimrc"})
}

func (s *ExecSuite) TestEscapeValidFieldCodes() {
	s.check(false, `vim "\\%u" ~/.vimrc`, []string{"vim", "%u", "~/.vimrc"})
}

func (s *ExecSuite) TestNonValidFieldCodes() {
	s.check(false, `vim %x .vimrc`, []string{"vim", "%x", ".vimrc"})
	s.check(false, `vim %x/.vimrc`, []string{"vim", "%x/.vimrc"})
}

func TestExecSuite(t *testing.T) {
	suite.Run(t, new(ExecSuite))
}
