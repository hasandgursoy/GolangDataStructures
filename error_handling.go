package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Go da hata ayiklama konusu uzerinde calisirken bir fonksiyon yazdigimizda error degeri donebilecegini hesap etmemiz gerekiyor.
	// Detayli bir sekilde tasarlarsak gelecekte o kadar rahat ederiz. Yoksa NaN donuyor

	result, err := evenNum(20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	if value, err := nSqurt(64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	file, err := os.Create("test.txt")
	if err == nil {
		stringArray := []byte("selamlar")
		file.Write(stringArray)
		os.Open("text.txt")
	}

	// vals, err := getGrade()

	// if err == nil {
	// 	fmt.Println(vals)
	// }

	guesTheNum()
}

func evenNum(num int) (string, error) {
	if num%2 != 0 {
		return "", errors.New("error : ciftSayı girmediniz")
	}
	return "Cift sayi girdiniz", nil

}

func nSqurt(num float64) (float64, error) {

	if num < 0 {
		return 0, errors.New("error : Negatif sayilarin koku alınamaz")
	} else {
		return math.Sqrt(num), nil
	}

}

func getGrade() (int, error) {
	fmt.Println("Lütfen bir sayi giriniz : ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	
	return num, nil
}

func guesTheNum() {

	target := numRand(1, 100)

	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın :")

	reader := bufio.NewReader(os.Stdin)

	for attemps := 0; attemps < 10; attemps++ {
		fmt.Println(10-attemps, " hakkınız kaldı")
		fmt.Print("Lütfen Tahmininizi Giriniz : ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		num, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println(err)
		}

		if num == target {
			fmt.Println("Tebrikler, ", 10-attemps, ". hakkınızda cevabı buldunuz : ", num)
			break
		}

		if num < target {
			fmt.Println("Sayıyı arttırın !")
		} else {
			fmt.Println("Sayıyı azaltın !")
		}

	}

}

func numRand(min, max int) int {
	return rand.Intn(max - min)
}
