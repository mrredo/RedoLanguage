package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	pathReplaceRegex = regexp.MustCompile("")
)

func main() {
	var arguments = os.Args
	var path = arguments[0]
	var fileName = arguments[1]
	var pathMod = strings.Split(path, "\\")
	pathMod[len(pathMod)-1] = fileName
	path = strings.Join(pathMod, "\\")
	fmt.Println(path)

}
