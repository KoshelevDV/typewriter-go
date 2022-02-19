package typeWriter

import (
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/koshelevdv/typewriter-go/utils"
)

func TypeWriter(file string, spd string, width int) {
	speed, _ := strconv.Atoi(spd)
	lines := utils.FileToSlice(file)
	for _, str := range lines {
		for _, j := range str {
			fmt.Print(string(j))
			wait := time.Millisecond * time.Duration(speed/utf8.RuneCountInString(str))
			time.Sleep(wait)
		}

		fmt.Println()
	}
}
