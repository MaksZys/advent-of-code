package _021

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func GetPath(relativePath string) string {
	rootPath, pathError := os.Getwd()
	if pathError != nil {
		log.Fatalf("Cannot get rooted path name: %s", pathError)
	}

	return rootPath + relativePath
}

func GetFolderPath() string {
	_, filePath, _, _ := runtime.Caller(0)

	return filepath.Dir(filePath)
}

func GetFile(relativePath string) *os.File {
	path := GetFolderPath() + relativePath

	file, fileError := os.Open(path)
	if fileError != nil {
		log.Fatalf("Cannot open file: %s \n%s", path, fileError)
	}

	return file
}

/*
*
 */
func ScanLines(file *os.File) []string {
	data := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func ParseToInt(strVal string) int {
	val, err := strconv.Atoi(strVal)
	if err != nil {
		panic(err)
	}

	return val
}

func Print2dArray(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%3d", arr[j][i])
		}
		fmt.Println()
	}

	fmt.Println()
}

func MeasureTime() func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s \n", time.Since(start))
	}
}
