package apktool

import (
	"io/ioutil"

	"github.com/weslenng/sspl/utils"
)

func Decode(file string) (string, error) {
	src, err := ioutil.TempDir("", "sspl-")
	if err != nil {
		return "", err
	}

	cmd := []string{
		"apktool",
		"decode",
		"--force",
		"--no-src",
		"--output",
		src,
		file,
	}

	if err := utils.Spawn(cmd...); err != nil {
		return "", err
	}

	return src, nil
}
