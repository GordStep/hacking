package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

func registration() string {

	agree := []string{"y", "Y"}

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		user_input := scanner.Text()

		if slices.Contains(agree, user_input) {
			return "ok"
		} else {
			return "no"
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
		return "err"
	}

	return "err"
}

func loading() {
	spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

	fmt.Print("Loading ")
	for i := 0; i < 20; i++ {
		fmt.Printf("\rLoading %s", spinner[i%len(spinner)])
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("\rLoaded!    \n")
}

func loader(path string) []string {
	return []string{}
}

func printer() {

}

func main() {
	fmt.Print("Are you ready to start hacking?(Y/n) ")

	res := registration()

	if res == "ok" {
		loading()

	}

}
