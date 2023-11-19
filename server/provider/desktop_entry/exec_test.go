package desktop_entry

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	FILE0 = "/home/user/file0"
	FILE1 = "/home/user/file1"
	URL0  = "http://example0.com"
	URL1  = "http://example1.com"
)

type ExecSuite struct {
	suite.Suite
}

func (s *ExecSuite) checkFull(execStr string, ep *ExecParams, expected []string) {
	t := s.T()

	actual, err := buildArgs(execStr, ep)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func (s *ExecSuite) check(needTerminal bool, execStr string, expected []string) {
	ep := NewExecParams(needTerminal, "bash")
	s.checkFull(execStr, ep, expected)
}

func (s *ExecSuite) TestExecStr() {
	ep := NewExecParams(false, "bash")

	s.checkFull(`vim`, ep, []string{"vim"})
	s.checkFull(`vim test`, ep, []string{"vim", "test"})

	ep.SetFiles(FILE0)
	s.checkFull(`vim`, ep, []string{"vim"})
	s.checkFull(`vim test`, ep, []string{"vim", "test"})

	ep.SetFiles(FILE0, FILE1)
	s.checkFull(`vim`, ep, []string{"vim"})
	s.checkFull(`vim test`, ep, []string{"vim", "test"})

	ep.SetFiles()
	ep.SetUrls(URL0)
	s.checkFull(`vim`, ep, []string{"vim"})
	s.checkFull(`vim test`, ep, []string{"vim", "test"})

	ep.SetUrls(URL0, URL1)
	s.checkFull(`vim`, ep, []string{"vim"})
	s.checkFull(`vim test`, ep, []string{"vim", "test"})
}

func (s *ExecSuite) TestExecStrWithFileParam() {
	ep := NewExecParams(false, "bash")

	s.checkFull(`vim %f`, ep, []string{"vim"})
	s.checkFull(`vim test %f`, ep, []string{"vim", "test"})

	ep.SetFiles(FILE0)
	s.checkFull(`vim %f`, ep, []string{"vim", FILE0})
	s.checkFull(`vim test %f`, ep, []string{"vim", "test", FILE0})

	ep.SetFiles(FILE0, FILE1)
	s.checkFull(`vim %f`, ep, []string{"vim", FILE0})
	s.checkFull(`vim test %f`, ep, []string{"vim", "test", FILE0})

	ep.SetFiles()
	ep.SetUrls(URL0)
	s.checkFull(`vim %f`, ep, []string{"vim"})
	s.checkFull(`vim test %f`, ep, []string{"vim", "test"})

	ep.SetUrls(URL0, URL1)
	s.checkFull(`vim %f`, ep, []string{"vim"})
	s.checkFull(`vim test %f`, ep, []string{"vim", "test"})
}

func (s *ExecSuite) TestExecStrWithFilesParam() {
	ep := NewExecParams(false, "bash")

	s.checkFull(`vim %F`, ep, []string{"vim"})
	s.checkFull(`vim test %F`, ep, []string{"vim", "test"})

	ep.SetFiles(FILE0)
	s.checkFull(`vim %F`, ep, []string{"vim", FILE0})
	s.checkFull(`vim test %F`, ep, []string{"vim", "test", FILE0})

	ep.SetFiles(FILE0, FILE1)
	s.checkFull(`vim %F`, ep, []string{"vim", FILE0, FILE1})
	s.checkFull(`vim test %F`, ep, []string{"vim", "test", FILE0, FILE1})

	ep.SetFiles()
	ep.SetUrls(URL0)
	s.checkFull(`vim %F`, ep, []string{"vim"})
	s.checkFull(`vim test %F`, ep, []string{"vim", "test"})

	ep.SetUrls(URL0, URL1)
	s.checkFull(`vim %F`, ep, []string{"vim"})
	s.checkFull(`vim test %F`, ep, []string{"vim", "test"})
}

func (s *ExecSuite) TestExecStrWithUrlParam() {
	ep := NewExecParams(false, "bash")

	s.checkFull(`vim %u`, ep, []string{"vim"})
	s.checkFull(`vim test %u`, ep, []string{"vim", "test"})

	ep.SetFiles(FILE0)
	s.checkFull(`vim %u`, ep, []string{"vim"})
	s.checkFull(`vim test %u`, ep, []string{"vim", "test"})

	ep.SetFiles(FILE0, FILE1)
	s.checkFull(`vim %u`, ep, []string{"vim"})
	s.checkFull(`vim test %u`, ep, []string{"vim", "test"})

	ep.SetFiles()
	ep.SetUrls(URL0)
	s.checkFull(`vim %u`, ep, []string{"vim", URL0})
	s.checkFull(`vim test %u`, ep, []string{"vim", "test", URL0})

	ep.SetUrls(URL0, URL1)
	s.checkFull(`vim %u`, ep, []string{"vim", URL0})
	s.checkFull(`vim test %u`, ep, []string{"vim", "test", URL0})
}

func (s *ExecSuite) TestExecStrWithUrlsParam() {
	ep := NewExecParams(false, "bash")

	s.checkFull(`vim %U`, ep, []string{"vim"})
	s.checkFull(`vim test %U`, ep, []string{"vim", "test"})

	ep.SetFiles(FILE0)
	s.checkFull(`vim %U`, ep, []string{"vim"})
	s.checkFull(`vim test %U`, ep, []string{"vim", "test"})

	ep.SetFiles(FILE0, FILE1)
	s.checkFull(`vim %U`, ep, []string{"vim"})
	s.checkFull(`vim test %U`, ep, []string{"vim", "test"})

	ep.SetFiles()
	ep.SetUrls(URL0)
	s.checkFull(`vim %U`, ep, []string{"vim", URL0})
	s.checkFull(`vim test %U`, ep, []string{"vim", "test", URL0})

	ep.SetUrls(URL0, URL1)
	s.checkFull(`vim %U`, ep, []string{"vim", URL0, URL1})
	s.checkFull(`vim test %U`, ep, []string{"vim", "test", URL0, URL1})
}

func (s *ExecSuite) TestNonValidFieldCodes() {
	ep := NewExecParams(false, "bash")

	s.checkFull(`vim test %x`, ep, []string{"vim", "test", "%x"})
	s.checkFull(`vim %x test`, ep, []string{"vim", "%x", "test"})
	s.checkFull(`vim %x/.vimrc`, ep, []string{"vim", "%x/.vimrc"})

	ep.SetFiles(FILE0)
	s.checkFull(`vim test %x`, ep, []string{"vim", "test", "%x"})
	s.checkFull(`vim %x test`, ep, []string{"vim", "%x", "test"})
	s.checkFull(`vim %x/.vimrc`, ep, []string{"vim", "%x/.vimrc"})

	ep.SetFiles(FILE0, FILE1)
	s.checkFull(`vim test %x`, ep, []string{"vim", "test", "%x"})
	s.checkFull(`vim %x test`, ep, []string{"vim", "%x", "test"})
	s.checkFull(`vim %x/.vimrc`, ep, []string{"vim", "%x/.vimrc"})

	ep.SetFiles()
	ep.SetUrls(URL0)
	s.checkFull(`vim test %x`, ep, []string{"vim", "test", "%x"})
	s.checkFull(`vim %x test`, ep, []string{"vim", "%x", "test"})
	s.checkFull(`vim %x/.vimrc`, ep, []string{"vim", "%x/.vimrc"})

	ep.SetUrls(URL0, URL1)
	s.checkFull(`vim test %x`, ep, []string{"vim", "test", "%x"})
	s.checkFull(`vim %x test`, ep, []string{"vim", "%x", "test"})
	s.checkFull(`vim %x/.vimrc`, ep, []string{"vim", "%x/.vimrc"})
}

func (s *ExecSuite) TestQuotes() {
	// TODO: https://specifications.freedesktop.org/desktop-entry-spec/latest/ar01s07.html
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

func (s *ExecSuite) TestEscapeValidFieldCodes() {
	s.check(false, `vim "\\%u" ~/.vimrc`, []string{"vim", "%u", "~/.vimrc"})
}

func TestExecSuite(t *testing.T) {
	suite.Run(t, new(ExecSuite))
}
