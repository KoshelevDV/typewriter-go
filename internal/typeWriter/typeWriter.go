package typeWriter

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/koshelevdv/typewriter-go/utils"
)

func TypeWriter(file string, spd string, width int) {
	speed, _ := strconv.Atoi(spd)
	lines := utils.FileToSlice(file)
	for _, str := range lines {
		start := time.Now()

		fmt.Print(strings.Repeat(" ", width))
		for _, j := range str {
			fmt.Print(string(j))
			wait := time.Millisecond * time.Duration(speed/utf8.RuneCountInString(str))
			time.Sleep(wait)
		}

		fmt.Printf("\r%-10v ->\n", time.Since(start))
	}
}
