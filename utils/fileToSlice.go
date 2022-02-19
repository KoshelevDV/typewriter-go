package utils

import (
	"bufio"
	"os"
	"time"
)

func FileToSlice(file string) (lines []string) {
	f, _ := os.Open(file)
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return
}

func MaxString(strs []string) (strLen int) {
	for _, str := range strs {
		tmp := 0
		for _, i := range str {
			bts := len(string(i))
			if bts == 3 {
				tmp += bts / 3 * 2
			} else {
				tmp += bts
			}
		}
		if strLen < tmp {
			strLen = tmp
		}
	}
	return
}

func Duration(t time.Time) time.Duration {
	return time.Since(t)
}
