package files

import (
	"os"

	"github.com/ReanGD/runify/server/gtk"
)

type IconResolver struct {
	theme *gtk.IconTheme
}

func NewIconResolver(theme *gtk.IconTheme) *IconResolver {
	return &IconResolver{theme: theme}
}

func (r *IconResolver) Resolve(icon string, size int) string {
	if len(icon) == 0 {
		return ""
	}

	if _, err := os.Stat(icon); err == nil {
		return icon
	}

	iconPath := r.theme.LookupIcon(icon, size, gtk.ICON_LOOKUP_NO_SVG)
	if len(iconPath) == 0 {
		iconPath = r.theme.LookupIcon(icon, size, 0)
	}

	return iconPath
}
