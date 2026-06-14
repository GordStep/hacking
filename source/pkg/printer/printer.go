package printer

import (
	"fmt"
	"math/rand"
	color "source/pkg/color"
	"time"
)

func Print_lines(lines []string) {
	status := []string{
		"[  Load resusrs  ]",
		"[    Stopping    ]",
		"[  Progressing   ]",
		"[    Starting    ]",
	}
	sleep_time := 200
	for {
		line_ind := rand.Int() % len(lines)
		r_line := lines[line_ind]
		col := rand.Int() % 8

		switch col {
		case 1:
			color.White.Print(status[0])
			sleep_time = 600
			// time.Sleep(600 * time.Millisecond)

		case 4:
			color.Red.Print(status[1])
			sleep_time = 400
			// time.Sleep(400 * time.Millisecond)

		case 6:
			color.Yellow.Print(status[2])
			sleep_time = 500
			// time.Sleep(time.Duration(line_ind) * time.Millisecond)?
		default:
			color.Green.Print(status[3])
			sleep_time = 300

		}
		fmt.Println(" ", r_line)
		// fmt.Println(line_ind, r_line)
		time.Sleep(time.Duration(sleep_time) * time.Millisecond)
	}
}
