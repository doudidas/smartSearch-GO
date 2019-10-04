package main

import "fmt"

func customLog(s string) {
	fmt.Println("[LOG]" + s)
}

func customWarn(s string) {
	fmt.Println("[WARNING]" + s)
}
func customErr(s string) {
	fmt.Println("[ERROR]" + s)
}
