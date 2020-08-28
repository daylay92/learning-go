package primitive

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"
	"fmt"
)

/*
ParseHTMLOneProp recieves fileName, propName, propValue and returns 
a string whose value is read from the file specified.
*/
func ParseHTMLOneProp(fileName string, propName string, propValue string) (string, error) {
	root, err := os.Getwd()
	fmt.Println(root)
	if err != nil {
		return "", err
	}
	absFileName := fileName
	if !strings.Contains(absFileName, ".html") {
		absFileName = fileName + ".html"
	}
	fileURL := filepath.Join(root, "./html_templates", absFileName)
	var bs []byte
	bs, err = ioutil.ReadFile(fileURL)
	if err != nil {
		return "", err
	}
	strValue := string(bs)
	if propName != "" {
		strValue = strings.ReplaceAll(strValue, "{{" + propName + "}}", propValue)
	}
	return strValue, err
}