package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unsafe"
)

var dic map[string]string

// * 初始化緩存，從 JSON 文件載入數據
func loadDictionary(filename string) error {
	// * 初始化空字典，確保無錯誤
	dic = make(map[string]string)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Warning: Could not read dictionary file: %v\n", err)
		return nil
	}

	err = json.Unmarshal(data, &dic)
	if err != nil {
		fmt.Printf("Warning: Could not parse dictionary file: %v\n", err)
		dic = make(map[string]string)
		return nil
	}

	return nil
}

//export convert
func convert(input *C.char) *C.char {
	text := C.GoString(input)

	for cn, zh := range dic {
		text = strings.ReplaceAll(text, cn, zh)
	}

	patterns := []string{
		`#{1,6}[\x{0020}\x{00A0}\x{0009}]+([^\n]+)`,
		`\*{2}([^\n\*]*)\*{2}`,
		`_{2}([^\n_]*)_{2}`,
		`\*([^\n\*]*)\*`,
		`_([^\n_]*)_`,
		`~{2}([^\n~]*)~{2}`,
		`={2}([^\n=]*)={2}`,
		`~([^\n~]*)~`,
		`\^([^\n^]*)\^`,
	}

	for _, pattern := range patterns {
		r := regexp.MustCompile(pattern)
		text = r.ReplaceAllString(text, "$1")
	}

	return C.CString(text)
}

//export freeString
func freeString(ptr *C.char) {
	defer C.free(unsafe.Pointer(ptr))
}

//export Init
func Init() {
	dictPath := "./TextConvertorDictionary.json"
	loadDictionary(dictPath)
}

func main() {
	Init()
}
