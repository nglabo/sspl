package main

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/weslenng/sspl/apktool"
	"github.com/weslenng/sspl/jarsigner"
	"github.com/weslenng/sspl/keytool"
	"github.com/weslenng/sspl/utils"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("Usage: sspl <file>\n")
		return
	}

	apkFile := os.Args[1]
	apkInfo, err := os.Stat(apkFile)
	if err != nil {
		utils.NewError("getting file info", err)
	}

	apkName := apkInfo.Name()
	if path.Ext(apkName) != ".apk" {
		utils.NewError("file must be an *.apk", nil)
	}

	src, err := apktool.Decode(apkFile)
	if err != nil {
		utils.NewError("decompiling file", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := utils.ReplaceConfig(src); err != nil {
			utils.NewError("replacing network_security_config.xml", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := utils.ReplaceManifest(src); err != nil {
			utils.NewError("replacing AndroidManifest.xml", err)
		}
	}()

	wg.Wait()
	if err := apktool.Build(src); err != nil {
		utils.NewError("compiling project", err)
	}

	store := path.Join(os.Getenv("HOME"), ".android/debug.keystore")
	_, err = os.Stat(store)
	if err != nil {
		if os.IsNotExist(err) {
			if err := keytool.NewKeyStore(store); err != nil {
				utils.NewError("creating keystore", err)
			}
		} else {
			utils.NewError("getting store info", err)
		}
	}

	dist := path.Join(src, "dist", apkName)
	if err := jarsigner.Sign(dist, store); err != nil {
		utils.NewError("signing *.apk file", err)
	}

	fmt.Printf("\033[1;36m%v\033[0m %v\n", "[success]", dist)
}
