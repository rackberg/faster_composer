package main

import (
	"fmt"
	"os"
	"cmp"
)

func main() {
	currentPath, _ := os.Getwd()
	filePath := currentPath + "/testfiles/composer.json"

	composerInfo, err := cmp.ReadComposerJson(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("read composer.json: %#v\n", composerInfo)

	for _, repository := range composerInfo.Repositories {
		packagistInfo := cmp.GetPackagistInfo(repository)

		fmt.Printf("%#v\n", packagistInfo)
	}
}