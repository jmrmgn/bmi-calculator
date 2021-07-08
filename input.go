package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jmrmgn/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {
	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)

	return
}

func getUserInput(prompText string) (value float64) {
	fmt.Print(prompText)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)

	return
}
