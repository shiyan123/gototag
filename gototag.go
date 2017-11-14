package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fileName := os.Args[1]
	structName := os.Args[2]

	obj := getObjectStruct(fileName, structName)
	WriteJsonTag(obj, fileName)
}

func getObjectStruct(fileName, structName string) (object map[int]string) {
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)

	object = make(map[int]string, 0)
	count := 0
	temp := false
	object[count] = ""

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if string(a) == fmt.Sprintf("type %s struct {", structName) {
			object[count] = string(a)
			temp = true
		}
		if temp {
			if _, ok := object[count]; !ok {
				object[count] = string(a)
			}
			if string(a) == "}" {
				break
			}
		}
		count = count + 1
	}
	return
}
func WriteJsonTag(object map[int]string, fileName string) {

	var keys []int
	for k := range object {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fi, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	fi.Seek(0, 2)

	w := bufio.NewWriter(fi)
	for k, v := range keys {
		if k == 0 || k == 1 || object[v] == "}" {
			lineStr := fmt.Sprintf("%s", object[v])
			fmt.Fprintln(w, lineStr)
		} else {
			currName := strings.Fields(object[v])[0]
			lineStr := fmt.Sprintf("%s", fmt.Sprintf("%s    %s", object[v], getTagName(currName, "json")))
			fmt.Fprintln(w, lineStr)
		}
	}
	w.Flush()
}

func getTagName(currName, tag string) (newName string) {
	first := true
	for _, r := range currName {
		if unicode.IsUpper(r) {
			if first {
				newName = fmt.Sprintf("%s%s", newName, strings.ToLower(string(r)))
				first = false
			} else {
				newName = fmt.Sprintf("%s_%s", newName, strings.ToLower(string(r)))
			}
		} else {
			newName = fmt.Sprintf("%s%s", newName, string(r))
		}
	}
	newName = fmt.Sprintf("`%s:\"%s\"`", tag, newName)
	return
}