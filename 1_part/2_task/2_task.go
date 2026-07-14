package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите цело число")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parsedInt, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Вы ввели не целое число")
		return
	}

	fmt.Printf("Вы ввели: %d\n", parsedInt)
}

