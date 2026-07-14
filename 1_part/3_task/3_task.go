package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	re := regexp.MustCompile(`^[0-9]+$`)

	fmt.Print("Введите число: ")
	
	celsi, _ := reader.ReadString('\n')
	
	celsi = strings.TrimSpace(celsi)

	if re.MatchString(celsi) {
			number, _ := strconv.Atoi(celsi)
			fmt.Printf("Успешно введено число: %d\n", number)
			forengeit := (float64(number)*9/5) + float64(32)

			fmt.Printf("%d = %f по фаренгейту", number, forengeit)
		} else {
			fmt.Println("Ошибка: вводить можно только числа. Попробуйте снова.")
		}
}