package main

import (
	"fmt"
	"os"

	"github.com/ReanGD/runify/server/files"
	"github.com/ReanGD/runify/server/gtk"
	"github.com/ReanGD/runify/server/providers"
)

func main() {
	gtk.Init()
	defaultIconTheme, err := gtk.IconThemeGetDefault()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get default icon theme error: %s\n", err)
		return
	}

	iconResolver := files.NewIconResolver(defaultIconTheme)

	providers.Get(iconResolver)
}
