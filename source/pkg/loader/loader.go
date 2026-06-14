package loader

import (
	"bufio"
	"fmt"
	"os"
)

func Loader(path string) (bool, []string) {
	text_data := make([]string, 0, 10)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return false, []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		// fmt.Printf("%d: %s\n", lineNumber, line)

		text_data = append(text_data, line)
	}

	if len(text_data) < 1 {
		return false, []string{}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
		return false, []string{}
	}

	return true, text_data
}
