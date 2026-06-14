package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"source/pkg/loader"
	"source/pkg/printer"
	"time"
)

const DATA_PATH = "data/data.txt"

func registration() string {

	agree := []string{"y", "Y", "\n"}

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

func main() {
	fmt.Print("Are you ready to start hacking?(Y/n) ")

	res := registration()

	switch res {
	case "ok":
		ok, lines := loader.Loader(DATA_PATH)
		if ok {
			// for _, line := range lines {
			// 	fmt.Println(line)
			// }
			loading()
		}
		printer.Print_lines(lines)
		return

	case "no":
		fmt.Print("Aborting!")
		return
	case "err":
		fmt.Print("Invalid imput.")
		return
	}

}
