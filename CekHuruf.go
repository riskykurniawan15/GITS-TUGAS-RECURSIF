package main

import (
	"fmt"
)

func main(){
	// Risky Kurniawan - ARS University
	var character, filter_type string
	fmt.Println("======================================")
	fmt.Println("            CEK HURUF")
	fmt.Println("======================================")
	fmt.Print("Masukan sebuah kata : ")
	fmt.Scan(&character)
	fmt.Print("Pilih jenis pengecekan ('number' or 'lowercase') : ")
	fmt.Scan(&filter_type)
	fmt.Println("======================================")
	if filter_type != "number" && filter_type != "lowercase" {
		fmt.Println("Pilihan pengecekan tidak tersedia!")
	}else{
		fmt.Println("Jumlah karakter", filter_type, "adalah =", check(character, filter_type))
	}
	fmt.Println("======================================")
	fmt.Println("         ** Terimakasih **")
	fmt.Println("   Risky Kurniawan - ARS University")
	fmt.Println("======================================")
}

func check(character string, filter_type string) int {
	if character == "" {
		return 0
	}
	
	var filter string
	if filter_type == "number" {
		filter = "1234567890"
	}else if filter_type == "lowercase" {
		filter = "abcdefghijklmnopqrstuvwxyz"
	}

	return count_filter(character, filter, 0, 0)
}

func count_filter(character string, filter string, i int, counter int) int {
	if i == len(character) {
		return counter
	}

	var result bool = filter_check(string(character[i]), filter, 0)

	if result == true {
		counter += 1
	}

	return count_filter(character, filter, i+1, counter)
}

func filter_check(character string, filter string, i int) bool {
	if string(character[0]) == string(filter[i]) {
		return true
	}

	if (i+1) == len(filter) {
		return false
	}

	return filter_check(character, filter, i+1)
}