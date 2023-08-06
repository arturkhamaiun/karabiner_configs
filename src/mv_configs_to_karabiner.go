package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	panicIfErr(err)

	configsPath, err := filepath.Abs(path + "/../configs")
	panicIfErr(err)

	files, err := ioutil.ReadDir(configsPath)
	panicIfErr(err)

	for _, file := range files {
		homeDir, err := os.UserHomeDir()
		panicIfErr(err)

		src := configsPath + "/" + file.Name()
		dst := homeDir + "/.config/karabiner/assets/complex_modifications/" + file.Name()

		data, err := ioutil.ReadFile(src)
		panicIfErr(err)

		dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0644)
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
