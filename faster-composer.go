package main

import (
	"fmt"
	"os"
	"cmp"
	"encoding/json"
	"os/exec"
)

func routine(repository cmp.RepositoryInfo, c chan cmp.PackagistInfo) {
	packagistInfo := cmp.GetPackagistInfo(repository)
	c <- packagistInfo
}

func downloadTest(url string, c chan DownloadProgress, lineNumber int) {
	text := fmt.Sprintf("Cloning package: %s", url)
	cmp.EchoAtLine(text, lineNumber)

	cmd := exec.Command("git", "clone", url)
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd.Wait()

	progress := new(DownloadProgress)
	progress.id = lineNumber
	c <- *progress
}

type DownloadProgress struct {
	id int
}

func main() {
	currentPath, _ := os.Getwd()
	filePath := currentPath + "/testfiles/composer.json"

	composerInfo, err := cmp.ReadComposerJson(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("read composer.json: %#v\n", composerInfo)

	var parsedTestData []string
	testData := []byte(`
		[
			"https://github.com/symfony/symfony.git",
			"https://github.com/symfony/web-profiler-bundle.git",
			"https://github.com/doctrine/doctrine2.git",
			"https://github.com/symfony/yaml.git",
			"https://github.com/artdarek/ga-4-laravel.git",
			"https://github.com/doctrine/instantiator.git",
			"https://github.com/doctrine/collections.git"
		]
	`)

	json.Unmarshal(testData, &parsedTestData)

	downloadChannel := make(chan DownloadProgress)

	lineNumber, _ := cmp.GetCursorPosition()

	runningProcedures := 0
	for _, value := range parsedTestData {
		go downloadTest(value, downloadChannel, lineNumber)
		runningProcedures = runningProcedures + 1
		lineNumber = lineNumber + 1
	}

	for {
		if runningProcedures <= 0 {
			break
		}

		downloadProgress := <- downloadChannel
		if downloadProgress.id > 0 {
			text := fmt.Sprintf("Download finished at: %d", downloadProgress.id)
			cmp.EchoAtLine(text, downloadProgress.id)
			runningProcedures = runningProcedures - 1
		}
	}

	fmt.Println("\nDownloading complete!")
}