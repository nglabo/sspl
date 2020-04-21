package keytool

import (
	"github.com/weslenng/sspl/utils"
)

func NewKeyStore(store string) error {
	cmd := []string{
		"keytool",
		"-alias",
		"androiddebugkey",
		"-deststoretype",
		"pkcs12",
		"-genkey",
		"-keyalg",
		"RSA",
		"-keypass",
		"android",
		"-keysize",
		"2048",
		"-keystore",
		store,
		"-storepass",
		"android",
		"-v",
		"-validity",
		"10000",
	}

	return utils.Spawn(cmd...)
}
