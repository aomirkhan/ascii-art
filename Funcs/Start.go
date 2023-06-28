package Funcs

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[90m"
)

func Start(s string) []string {
	List := strings.Split(s, "\n")
	b := 0
	var alph []string
	for i := 1; len(List) > i; i++ {
		if i%9 == 0 {
			alph = append(alph, strings.Join(List[b:i+1], "\n"))
			b = i + 1
		}
	}

	return alph
}

func Print(s string, alph []string, ltr string, clr string) {
	var obj [][]string
	for i := range s {
		if strings.Contains(ltr, string(s[i])) {
			for j := 32; 126 > j; j++ {
				if rune(s[i]) == rune(j) {
					obj = append(obj, Colorize(strings.Split(alph[j-32], "\n"), clr))
				}
			}
		} else {
			for j := 32; 126 > j; j++ {
				if rune(s[i]) == rune(j) {
					obj = append(obj, strings.Split(alph[j-32], "\n"))
				}
			}
		}
	}

	ind := 0
	for ind < 8 {
		var line string
		for i := range obj {
			line = line + strings.ReplaceAll(obj[i][ind], "\n", "")
		}
		fmt.Println(line)
		ind++
	}
}

func ReadFile(fileName string) string {
	body, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body)
}

func Md5(s string) string {
	h := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}

func Colorize(s []string, color string) []string {
	if color == "red" {
		color = Red
	} else if color == "green" {
		color = Green
	} else if color == "yellow" {
		color = Yellow
	} else if color == "blue" {
		color = Blue
	} else if color == "purple" {
		color = Purple
	} else if color == "cyan" {
		color = Cyan
	} else if color == "gray" {
		color = Gray
	} else {
		log.Fatal("Color doesn't exists or Creator haven't add this color!")
	}
	for i := range s {
		s[i] = Reset + color + s[i] + Reset
	}

	return s
}
