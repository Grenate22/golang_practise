package main

import (
	"fmt"
	"strings"
)

const (
	french     = "french"
	spanish    = "spanish"
	russian    = "russian"
	chinese    = "chinese"
	italian    = "italian"
	japanese   = "japanese"
	german     = "german"
	portuguese = "portuguese"
	korean     = "korean"
	arabic     = "arabic"
	danish     = "danish"
	swahili    = "swahili"
	dutch      = "dutch"
	greek      = "greek"
	polish     = "polish"
	indonesian = "indonesian"
	hindi      = "hindi"
	norwegian  = "norwegian"
	turkish    = "turkish"
	hebrew     = "hebrew"
	swedish    = "swedish"

	englishprefix    = "hello, "
	frenchprefix     = "Bonjour, "
	spanishprefix    = "Hola, "
	russianprefix    = "Zdravstvuyte, "
	chineseprefix    = "nin hao, "
	italianprefix    = "salve, "
	japaneseprefix   = "konnichiwa, "
	germanprefix     = "Guten Tag, "
	portugueseprefix = "ola, "
	koreanrefix      = "Anyounghaseyo, "
	arabicprefix     = "Asalaam alaikum, "
	danishprefix     = "Goddag, "
	swahiliprefix    = "Shikamoo, "
	dutchprefix      = "salve"
	greekprefix      = "salve"
	polishprefix     = "salve"
	indonesianprefix = "salve"
	hindiprefix      = "salve"
	norwegianprefix  = "salve"
	turkishprefix    = "salve"
	hebrewprefix     = "salve"
	swedishprefix    = "salve"
)

func Greet(name string, language string) string {
	// i must have a return data type else i will get error
	if name == "" {
		name = "World"
	}

	return greetingPrefix(strings.ToLower(language)) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchprefix
	case spanish:
		prefix = spanishprefix
	case russian:
		prefix = russianprefix
	case chinese:
		prefix = chineseprefix
	case italian:
		prefix = italianprefix
	case japanese:
		prefix = japaneseprefix
	case german:
		prefix = germanprefix
	case portuguese:
		prefix = portugueseprefix
	case korean:
		prefix = koreanrefix
	case arabic:
		prefix = arabicprefix
	case swahili:
		prefix = swahiliprefix
	case dutch:
		prefix = danishprefix
	case greek:
		prefix = greekprefix
	case polish:
		prefix = polishprefix
	case indonesian:
		prefix = indonesianprefix
	case hindi:
		prefix = hindiprefix
	case norwegian:
		prefix = norwegianprefix
	case turkish:
		prefix = turkishprefix
	case hebrew:
		prefix = hebrewprefix
	case swedish:
		prefix = swedishprefix
	case danish:
		prefix = danishprefix
	default:
		prefix = englishprefix
	}
	return
}

func main() {
	fmt.Println("Input your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Input your language")
	var language string
	fmt.Scanln(&language)
	fmt.Println(Greet(name, language))

}
