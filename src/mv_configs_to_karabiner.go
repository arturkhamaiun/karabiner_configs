package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	panicIfErr(err)

	configsPath, err := filepath.Abs(path + "/../configs")
	panicIfErr(err)

	files, err := ioutil.ReadDir(configsPath)
	panicIfErr(err)

	homeDir, err := os.UserHomeDir()
	panicIfErr(err)

	srcPath := configsPath + "/"
	dstPath := homeDir + "/.config/karabiner/assets/complex_modifications/"

	for _, file := range files {
		src := srcPath + file.Name()
		dst := dstPath + file.Name()

		data, err := ioutil.ReadFile(src)
		panicIfErr(err)

		cmd := exec.Command(
			"/Library/Application Support/org.pqrs/Karabiner-Elements/bin/karabiner_cli",
			"--lint-complex-modifications",
			src,
		)

		if err := cmd.Run(); err != nil {
			fmt.Println("Invalid karbiner config for file: " + src)
			fmt.Println("")
			continue
		}

		dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		panicIfErr(err)

		fmt.Println("Writing...")
		fmt.Println("Source: " + src)
		fmt.Println("Destination: " + dst)
		fmt.Println("")

		_, err = dstFile.Write(data)
		panicIfErr(err)

		err = dstFile.Close()
		panicIfErr(err)
	}

	fmt.Println("Finished!")
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
