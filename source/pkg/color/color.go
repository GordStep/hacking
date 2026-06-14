package color

import "fmt"

type Color string

// Константы цветов
const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	White  Color = "\033[37m"
	Bold   Color = "\033[1m"
)

func (c Color) Sprint(s string) string {
	return string(c) + s + string(Reset)
}

func (c Color) Print(s string) {
	fmt.Print(c.Sprint(s))
}

func (c Color) Println(s string) {
	fmt.Println(c.Sprint(s))
}
