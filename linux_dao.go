package penv

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	linuxProfileShell = &shell{
		configFileName: filepath.Join(os.Getenv("HOME"), ".profile"),
		commentSigil:   " #",
		quote: func(value string) string {
			r := strings.NewReplacer(
				"\\", "\\\\",
				"'", "\\'",
				"\n", `'"\n"'`,
				"\r", `'"\r"'`,
			)
			return "'" + r.Replace(value) + "'"
		},
		mkSet: func(sh *shell, nv NameValue) string {
			return fmt.Sprintf(
				"export %s=%s",
				nv.Name, sh.quote(nv.Value),
			)
		},
		mkAppend: func(sh *shell, nv NameValue) string {
			return fmt.Sprintf(
				"export %s=${%s}${%s:+:}%s",
				nv.Name, nv.Name, nv.Name, sh.quote(nv.Value),
			)
		},
		mkUnset: func(sh *shell, nv NameValue) string {
			return fmt.Sprintf(
				"unset %s",
				nv.Name,
			)
		},
	}
)

// LinuxDAO is the data access object for Linux
type LinuxDAO struct{}

func init() {
	RegisterDAO(1000, func() bool {
		return runtime.GOOS == "linux"
	}, linuxProfileShell)
}
