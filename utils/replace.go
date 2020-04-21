package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

func ReplaceConfig(src string) error {
	xml := path.Join(src, "res/xml")

	_, err := os.Stat(xml)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(xml, 0644); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	b, err := ioutil.ReadFile(path.Join(
		os.Getenv("HOME"),
		".sspl/network_security_config.xml",
	))

	if err != nil {
		return err
	}

	return ioutil.WriteFile(
		path.Join(xml, "network_security_config.xml"),
		b,
		0644,
	)
}

func ReplaceManifest(src string) error {
	p := path.Join(src, "AndroidManifest.xml")
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	pa := `networkSecurityConfig="(.*?)"`
	matched, err := regexp.Match(pa, b)
	if err != nil {
		return err
	}

	if !matched {
		b = bytes.Replace(
			b,
			[]byte("<application"),
			[]byte(`<application networkSecurityConfig=""`),
			1,
		)
	}

	re := regexp.MustCompile(pa)
	b = re.ReplaceAll(
		b,
		[]byte(`networkSecurityConfig="@xml/network_security_config"`),
	)

	return ioutil.WriteFile(p, b, 0644)
}
