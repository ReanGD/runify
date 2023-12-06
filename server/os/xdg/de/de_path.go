package de

import (
	"path/filepath"
	"strings"

	"github.com/ReanGD/runify/server/paths"
)

func getXDGAppDirs() []string {
	var xdgAppDirs []string
	for _, xdgDataDir := range paths.GetAllDataDirs() {
		xdgAppDir := filepath.Join(xdgDataDir, "applications")
		if ok, _ := paths.ExistsDir(xdgAppDir); ok {
			xdgAppDirs = append(xdgAppDirs, xdgAppDir)
		}
	}

	return xdgAppDirs
}

func getXDGMimeAppFiles() []string {
	var fileNames []string
	if currentDesktops, ok := paths.GetEnv("XDG_CURRENT_DESKTOP"); ok {
		for _, currentDesktop := range strings.Split(currentDesktops, ":") {
			currentDesktop = strings.ToLower(strings.TrimSpace(currentDesktop))
			if len(currentDesktop) == 0 {
				fileNames = append(fileNames, currentDesktop+"-mimeapps.list")
			}
		}
	}
	fileNames = append(fileNames, "mimeapps.list")

	var baseDirs []string
	baseDirs = append(baseDirs, paths.GetAllConfigDirs()...)
	for _, xdgDataDir := range paths.GetAllDataDirs() {
		baseDirs = append(baseDirs, filepath.Join(xdgDataDir, "applications"))
	}

	var res []string
	for _, baseDir := range baseDirs {
		if ok, _ := paths.ExistsDir(baseDir); ok {
			for _, fileName := range fileNames {
				dir := filepath.Join(baseDir, fileName)
				if ok, _ := paths.ExistsFile(dir); ok {
					res = append(res, dir)
				}
			}
		}
	}

	return res
}
