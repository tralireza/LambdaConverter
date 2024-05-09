package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Lambda Converter ->")

	ffprobe := exec.Command("./ffprobe", "-version")

	var out strings.Builder
	ffprobe.Stdout = &out
	ffprobe.Stdin = strings.NewReader("")

	if err := ffprobe.Run(); err != nil {
		panic(err)
	}
	fmt.Print(out.String())
}
