package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func FeistelCipher(plaintext, method string, seed int64) string{
	if method == "encrypt"{
		return encrypt(plaintext, seed)
	}
	return decrypt(plaintext,seed)
	return ""
}
func encrypt(plaintext string, seed int64) string{
	keys := randomKeyGenerator(seed, "normal")
	binString := getBinStrFromString(plaintext)
	L := binString[:16]
	R := binString[16:]
	LNumber, err := strconv.ParseInt(L, 2, 64)
	if err != nil{
		fmt.Println(err)
	}
	RNumber, err := strconv.ParseInt(R, 2, 64)
	if err != nil{
		fmt.Println(err)
	}
	for i := 0; i < 4; i++{
		tmp := LNumber
		LNumber = RNumber ^ f(LNumber, int64(keys[i]))
		RNumber = tmp
	}
	L = strconv.FormatInt(LNumber,2)
	L = fmt.Sprintf("%016s",L)
	R = strconv.FormatInt(RNumber,2)
	R = fmt.Sprintf("%016s",R)
	str := L + R
	res := make([]string,0)
	res1 := ""
	for j := 0; j < len(str); j += 8{
		res = append(res, str[j:j+8])
	}
	for h := 0; h < 4; h++{
		res1 += string(binaryToDecimal(res[h]))
	}
return res1
}
func decrypt(plaintext string, seed int64) string{
	keys := randomKeyGenerator(seed, "reversed")
	binString := getBinStrFromString(plaintext)
	L := binString[:16]
	R := binString[16:]
	LNumber, err := strconv.ParseInt(L, 2, 64)
	if err != nil{
		fmt.Println(err)
	}
	RNumber, err := strconv.ParseInt(R, 2, 64)
	if err != nil{
		fmt.Println(err)
	}
	for i := 0; i < 4; i++{

		tmp := RNumber
		RNumber = LNumber ^ f(RNumber, int64(keys[i]))
		LNumber = tmp
	}
	L = strconv.FormatInt(LNumber,2)
	L = fmt.Sprintf("%016s",L)
	R = strconv.FormatInt(RNumber,2)
	R = fmt.Sprintf("%016s",R)
	str := L + R
	res := make([]string,0)
	res1 := ""
	for j := 0; j < len(str); j += 8{
		res = append(res, str[j:j+8])
	}
	for h := 0; h < 4; h++{
		res1 += string(binaryToDecimal(res[h]))
	}

	return res1
}
//Генерируем 32 битный ключ для каждого раунда
func randomKeyGenerator(K int64, mode string) []int{
	rand.Seed(K)
	if mode == "normal" {
	keys := make([]int,0)
		for i := 0; i < 4; i++ {
			val := getRandomKey()
			keys = append(keys, val)
		}
		return keys
	}
	keys := make([]int,4)
	for i := 3; i >= 0; i-- {
		val := getRandomKey()
		keys[i] = val
	}
	return keys
}
//Генерирует один 32 битный ключ
func getRandomKey() int{
	return 2147483648 + rand.Intn(4294967295 - 2147483648)
}

func getBinStrFromString(plaintext string) string{
	binString := ""
	for _, c := range plaintext{
		binString =fmt.Sprintf("%s%08b",binString, c)
	}
	return binString
}


func f(B, K int64) int64 {
	result1 := math.Pow(2, 64 * ((math.Sqrt(math.Abs(math.Sin(float64(B)))))/(1 + math.Abs(math.Tan(float64(K)))))) - 1
	return  int64( result1) % 65535
}
func binaryToDecimal(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil{
		fmt.Println(err)
	}
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0{
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder * int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
func main(){
enc := FeistelCipher("supa", "encrypt", 123)
dec := FeistelCipher(enc, "decrypt", 123)
fmt.Printf("encrypted:%s\ndecrypted: %s",enc, dec)
/*res := encrypt("hell", 123)
fmt.Println("encrypt:",res)
fmt.Println("decrypt: ",decrypt(res, 123))*/

}


