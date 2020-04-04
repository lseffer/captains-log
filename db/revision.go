package db

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type revision struct {
	id  uint16
	sql string
}

type byID []revision

func (s byID) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func parseRevisions() []revision {
	files, _ := ioutil.ReadDir("sql")
	var result []revision
	for _, file := range files {
		fileName := file.Name()
		id, _ := strconv.ParseUint(strings.Replace(fileName, ".sql", "", -1), 10, 16)
		fileContents, _ := ioutil.ReadFile("sql/" + fileName)
		result = append(result, revision{uint16(id), string(fileContents)})
	}
	return result
}
