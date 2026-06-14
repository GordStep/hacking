package printer

import (
	"fmt"
	"math/rand"
	color "source/pkg/color"
	status "source/pkg/statuses"
	"time"
)

func Print_lines(lines []string) {
	sleep_time := 200
	for {
		line_ind := rand.Int() % len(lines)
		r_line := lines[line_ind]
		col := rand.Int() % 8
		fmt.Print("[")
		switch col {
		case 1:
			color.White.Print(status.GetByNum(col))
			sleep_time = 600
			// time.Sleep(600 * time.Millisecond)

		case 4:
			color.Red.Print(status.GetByNum(col))
			sleep_time = 400
			// time.Sleep(400 * time.Millisecond)

		case 6:
			color.Yellow.Print(status.GetByNum(col))
			sleep_time = 500
			// time.Sleep(time.Duration(line_ind) * time.Millisecond)?
		default:
			color.Green.Print(status.GetByNum(col))
			sleep_time = 300

		}
		fmt.Println("] ", r_line)
		// fmt.Println(line_ind, r_line)
		time.Sleep(time.Duration(sleep_time) * time.Millisecond)
	}
}
