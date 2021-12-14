package adventOfCode2021

import (
	"bufio"
	"log"
	"os"
)

func GetPath(relativePath string) string {
	rootPath, pathError := os.Getwd()
	if pathError != nil {
		log.Fatalf("Cannot get rooted path name: %s", pathError)
	}

	return rootPath + relativePath
}

/**

 */
func GetFile(relativePath string) *os.File {
	path := GetPath(relativePath)

	file, fileError := os.Open(path)
	if fileError != nil {
		log.Fatalf("Cannot open file: %s \n%s", path, fileError)
	}

	return file
}

/**

 */
func ScanLines(file *os.File) []string {
	data := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
