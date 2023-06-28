package main

import (
	"fmt"
	"gg/Funcs"
	"log"
	"os"
	"strings"
)

func main() {
	err1 := "There must be 3 arguments!"
	err2 := "Needed characters related to ASCII"
	err3 := "Do not change Shadow.txt"
	if len(os.Args) != 4 {
		log.Fatal(err1)
	}
	clr := ReadArgs(os.Args[1])
	ltr := ReadArgs(os.Args[2])
	for _, i := range os.Args[3] {
		if i > 126 || i < 32 {
			log.Fatal(err2)
		}
	}
	k := Funcs.ReadFile("shadow.txt")
	k2 := Funcs.Md5(k)
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	if hashShadow != k2 {
		log.Fatal(err3)
	}
	args := os.Args[3]
	argss := strings.Split(args, "\\n")

	mapa := Funcs.Start(k)
	if args == "" {
		fmt.Print("")
	} else if args == "\n" {
		fmt.Println("")
	} else {
		for i := range argss {
			if argss[i] == "" {
				fmt.Println("")
			} else {
				Funcs.Print(argss[i], mapa, ltr, clr)
			}
		}
	}
	// Funcs.PrintEverything(args, mapa)
}

func ReadArgs(s string) string {
	var result string
	var count int
	for i := range s {
		if string(s[i]) == "<" {
			count = i
		} else if string(s[i]) == ">" {
			result += s[count+1 : i]
		}
	}
	return result
}
