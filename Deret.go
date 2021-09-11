package main

import (
	"fmt"
	"strings"
)

func main(){
	// Risky Kurniawan - ARS University
	var jml int
	fmt.Println("========================================")
	fmt.Println("            DERET REKURSIF")
	fmt.Println("========================================")
	fmt.Print("Masukan jumlah deret : ")
	fmt.Scan(&jml)
	fmt.Println("========================================")
	fmt.Println(getDeret(jml, 20))
	fmt.Println(getDeret(jml, 100))
	fmt.Println(getDeret(jml, 1))
	fmt.Println(getDeretPecahan(jml, "penyebut"))
	fmt.Println(getDeretPecahan(jml, "pembilang"))
	fmt.Println("========================================")
	fmt.Println("           ** Terimakasih **")
	fmt.Println("   Risky Kurniawan - ARS University")
	fmt.Println("========================================")
}

func getDeret(jml int, awal float64) string {
	if jml <= 0 {
		return "0 = 0"
	}

	var deret string = runDeret(1, jml, awal, 0, "")
	result := strings.Split(deret, " = ")
	return result[1] + " = " + result[0]
}

func runDeret(n int, jml int, awal float64, sum float64, temp string) string{
	temp = roundComma(fmt.Sprintf("%.3f", awal))
	if (n%3) == 0 {
		var n3 float64 = ((awal*2) * 66.66) / 100
		sum += n3
		temp = roundComma(fmt.Sprintf("%.3f", n3))
	}else{
		sum += awal
		awal = awal / 2
	}
	
	if n != jml {
		temp += fmt.Sprintf(" + ")
	}

	if n == jml {
		return fmt.Sprintf("%s = %s", temp, roundComma(fmt.Sprintf("%.3f", sum)))
	}

	temp += runDeret(n + 1, jml, awal, sum, temp)

	return temp
}

func roundComma(number string) string {
	value := strings.Split(number, ".")

	value[1] = string(value[1][0]) + string(value[1][1])

	if value[1] == "00" {
		return value[0]
	}

	if string(value[1][1]) == "0" {
		return value[0] + "." +string(value[1][0])
	}

	return value[0] + "." + value[1]
}

func getDeretPecahan(jml int, x string) string {
	if jml <= 0 {
		return "0 = 0"
	}

	var deret string
	
	if x == "pembilang" {
		deret = runDeretPecahan(1, jml, 2, 1, x, 0, 0, "")
	}else if x == "penyebut"{
		deret = runDeretPecahan(1, jml, 1, 2, x, 0, 0, "")
	}else{
		return "error code x | set parameter to 'pembilang' or 'penyebut'"
	}

	result := strings.Split(deret, " = ")
	return result[1] + " = " + result[0]
}

func runDeretPecahan(n int, jml int, pembilang int, penyebut int, x string, sum_pembilang int, sum_penyebut int, temp string) string{
	sum_pembilang, sum_penyebut = sumPecahan(pembilang, penyebut, sum_pembilang, sum_penyebut)
	if x == "pembilang" {
		temp = fmt.Sprintf("%dx/%d", pembilang, penyebut)
		penyebut += 1
		pembilang = penyebut * penyebut
	}else if x == "penyebut" {
		temp = fmt.Sprintf("%d/%dx", pembilang, penyebut)
		pembilang += 1
		penyebut = pembilang * pembilang
	}
	
	if n != jml {
		temp += fmt.Sprintf(" + ")
	}

	if n == jml {
		var perkecil string = perkecilPecahan(sum_pembilang, sum_penyebut, x)
		var sum_temp string = ""
		if x == "pembilang" {
			sum_temp = fmt.Sprintf("%dx/%d", sum_pembilang, sum_penyebut)
		}else if x == "penyebut" {
			sum_temp = fmt.Sprintf("%d/%dx", sum_pembilang, sum_penyebut)
		}

		temp = fmt.Sprintf("%s = %s", temp, sum_temp)

		if perkecil != ("-> "+sum_temp) {
			temp += " " + perkecil
		}

		return temp
	}
	temp += runDeretPecahan(n + 1, jml, pembilang, penyebut, x, sum_pembilang, sum_penyebut, temp)

	return temp
}

func sumPecahan(pembilang int, penyebut int, sum_pembilang int, sum_penyebut int) (int, int){
	if sum_pembilang == 0 && sum_penyebut == 0 {
		return pembilang, penyebut
	}

	new_penyebut := penyebut * sum_penyebut

	new_pembilang := ((new_penyebut / penyebut) * pembilang) + ((new_penyebut / sum_penyebut) * sum_pembilang)

	return new_pembilang,new_penyebut
}

func perkecilPecahan(sum_pembilang int, sum_penyebut int, x string) string {
	if sum_pembilang != 1 && sum_penyebut != 1 {
		if (sum_pembilang%2) == 0 && (sum_penyebut%2) == 0 {
			sum_pembilang /= 2
			sum_penyebut /= 2
	
			return perkecilPecahan(sum_pembilang, sum_penyebut, x)
		}else if (sum_pembilang%3) == 0 && (sum_penyebut%3) == 0 {
			sum_pembilang /= 3
			sum_penyebut /= 3
	
			return perkecilPecahan(sum_pembilang, sum_penyebut, x)
		}else if (sum_pembilang%5) == 0 && (sum_penyebut%5) == 0 {
			sum_pembilang /= 5
			sum_penyebut /= 5
	
			return perkecilPecahan(sum_pembilang, sum_penyebut, x)
		}else if (sum_pembilang%7) == 0 && (sum_penyebut%7) == 0 {
			sum_pembilang /= 5
			sum_penyebut /= 5
	
			return perkecilPecahan(sum_pembilang, sum_penyebut, x)
		}
	}

	if x == "pembilang" {
		return fmt.Sprintf("-> %dx/%d", sum_pembilang, sum_penyebut)
	}else if x == "penyebut" {
		return fmt.Sprintf("-> %d/%dx", sum_pembilang, sum_penyebut)
	}

	return ""
}