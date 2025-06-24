package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	text := flag.String("text", "", "Text to encrypt or decrypt")
	encrypt := flag.Bool("encrypt", false, "Encrypt the text")
	decrypt := flag.Bool("decrypt", false, "Decrypt the text")

	flag.Parse()

	if *encrypt && *decrypt {
		panic("You cannot encrypt and decrypt at the same time")
	}

	if *text == "" {
		panic("Text cannot be empty")
	}

	if len(os.Args) < 3 {
		panic("You must provide a text and either encrypt or decrypt flag")
	}

	fmt.Println("Text:", *text)
	fmt.Println("Encrypt:", *encrypt)
	fmt.Println("Decrypt:", *decrypt)
	if *encrypt {
		encryptedText := Encrypt(*text)
		fmt.Println("Encrypted Text:", encryptedText)
	} else {
		decryptedText := Decrypt(*text)
		fmt.Println("Decrypted Text:", decryptedText)
	}

}

var final = "VWXYZABCDEFGHIJKLMNOPQRSTU"
var original = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encrypt(value string) string {
	mymap := make(map[rune]rune)

	customMap := CustomMap(original)

	for _, v := range value {
		mymap[v] = rune(final[(customMap[v]+7)%26])
	}

	result := ""

	for _, value := range value {
		result += string(mymap[value])
	}
	return result
}

func Decrypt(value string) string {
	mymap := make(map[rune]rune)

	custom := CustomMap(final)

	for _, v := range value {
		mymap[v] = rune(original[(custom[v]-7+26)%26])
	}
	result := ""
	for _, v := range value {
		result += string(mymap[v])
	}
	return result
}

func CustomMap(item string) map[rune]int {
	m := make(map[rune]int)

	for i, v := range item {
		m[v] = i
	}
	return m
}
