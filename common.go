package adventOfCode2021

import (
	"bufio"
	"log"
	"os"
)

/**

 */
func GetFile(relativePath string) *os.File {
	rootPath, pathError := os.Getwd()
	if pathError != nil {
		log.Fatalf("Cannot get rooted path name: %s", pathError)
	}

	file, fileError := os.Open(rootPath + relativePath)
	if fileError != nil {
		log.Fatalf("Cannot open file: %s \n%s", rootPath+relativePath, fileError)
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
