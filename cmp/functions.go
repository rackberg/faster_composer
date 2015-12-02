package cmp

import (
	"io/ioutil"
	"github.com/clbanning/mxj"
	"os"
	"fmt"
	"os/exec"
	"bufio"
	"strconv"
	"regexp"
	"strings"
)

func GetPackagistInfo (repository RepositoryInfo) (packagistInfo PackagistInfo) {
	content, err := GetHttpResponseBody(repository.Url + "/packages.json")
	if err != nil {
		panic(err)
	}

	CreateStructFromByteArray(content, &packagistInfo)

	return packagistInfo
}

func CreateStructFromByteArray (bytes []byte, structPtr interface{}) {
	m, err := mxj.NewMapJson(bytes)
	if err != nil {
		panic(err)
	}

	m.Struct(&structPtr)
}

func ReadComposerJson(filename string) (composerInfo ComposerInfo, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	CreateStructFromByteArray(bytes, &composerInfo)

	return composerInfo, err
}

func GetCursorPosition() (position int, err error) {
	// Set the terminal to raw mode (to be undone with `-raw`)
	rawMode := exec.Command("/bin/stty", "raw")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()

	fmt.Printf("%c[6n", 27)
	// capture keyboard output from echo command
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadSlice('R')

	// Set the terminal back from raw mode to 'cooked'
	rawModeOff := exec.Command("/bin/stty", "-raw")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()

	col := ""
	// check for the desired output
	if strings.Contains(string(text), ";") {
		re := regexp.MustCompile(`\d+`)
		col = re.FindString(string(text))
	}

	return strconv.Atoi(col)
}

func EchoAtLine(text string, lineNumber int) {
	fmt.Printf("\033[%d;0H\033[2K%s", lineNumber, text)
}