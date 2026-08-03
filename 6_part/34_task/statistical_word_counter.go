package main

import (
	"fmt"
	"regexp"
	"strings"
)

var text string = `Go — простой и быстрый язык программирования. Go позволяет писать быстрые и надёжные программы. Язык Go отлично подходит для разработки серверов, микросервисов и сетевых приложений. Многие разработчики выбирают Go за простоту, скорость и удобство. Простота языка Go помогает быстрее изучать программирование и создавать качественные программы.`
var	re = regexp.MustCompile(`[.,!?;:—\-'""()\[\]]`)


func main() {
	wordsMap := make(map[string]int) 

	listOfWords := textPreparation(text)
	counter(listOfWords, wordsMap)
	fmt.Println(wordsMap)
}

func counter(text []string, wordsMap map[string]int) {
	for _, v := range text {
		wordsMap[v]++
	}
}

func textPreparation(text string) []string {

	text = re.ReplaceAllString(text, "")
	text = strings.ToLower(text)
	listOfWords := strings.Fields(text)
	return listOfWords
}
