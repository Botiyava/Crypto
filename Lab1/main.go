package main

import (
	"fmt"
	"strings"
)
const alphabet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encrypt(text, key string) (cipherText string){
	//----- key modification -----
	modifiedKey := key
	for len(modifiedKey) < len(text){
		modifiedKey += key
	}
	key = modifiedKey[:len(text)]
	for i := 0; i < len(text); i++{
		cipherText += string(alphabet[(len(alphabet) + strings.Index(alphabet, string(key[i])) - strings.Index(alphabet, string(text[i]))) % 26])
	}
	return cipherText
}
func main(){

	plainText := "HELLOWORLD"
	key := "KEY"

	fmt.Println(encrypt(plainText, key))

}
