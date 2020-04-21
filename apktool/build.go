package apktool

import "github.com/weslenng/sspl/utils"

func Build(src string) error {
	cmd := []string{"apktool", "empty-framework-dir", "--force", src}
	if err := utils.Spawn(cmd...); err != nil {
		return err
	}

	cmd = []string{"apktool", "build", src}
	return utils.Spawn(cmd...)
}
