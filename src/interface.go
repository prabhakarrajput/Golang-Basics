package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'o' || rune == 'i' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Ram Seige")
	var v VowelsFinder
	v = name 
	fmt.Printf("Vowels are %c",v.FindVowels())
}
