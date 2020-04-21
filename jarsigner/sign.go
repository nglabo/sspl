package jarsigner

import (
	"github.com/weslenng/sspl/utils"
)

func Sign(file string, store string) error {
	cmd := []string{
		"jarsigner",
		"-keypass",
		"android",
		"-keystore",
		store,
		"-storepass",
		"android",
		file,
		"androiddebugkey",
	}

	return utils.Spawn(cmd...)
}
