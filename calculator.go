package main

import "fmt"

func main() {

	for {

		var num1 float64
		var num2 float64

		fmt.Println("Birinchi raqamni kiriting:")
		fmt.Scan(&num1)

		fmt.Println("Ikkinchi raqamni kiriting:")
		fmt.Scan(&num2)

		var inn int
		fmt.Println("Arifmetik operatsiyani tanlang")
		fmt.Println("1.'+' Qo'shish")
		fmt.Println("2.'-' Ayirish")
		fmt.Println("3.'*' Ko'paytirish")
		fmt.Println("4.'/' Bo'lish")
		fmt.Println("5.Exit")

		fmt.Scan(&inn)

		if inn == 5 {
			fmt.Println("Dastur yakunlandi.")
			break
		}

		switch inn {
		case 1:
			fmt.Println("Natija:", num1+num2)
		case 2:
			fmt.Println("Natija:", num1-num2)
		case 3:
			fmt.Println("Natija:", num1*num2)
		case 4:
			if num2 == 0 {
				fmt.Println("0 ga bo'lish mumkin emas")
			} else {
				fmt.Println("Natija:", num1/num2)
			}
		default:
			fmt.Println("Noto'g'ri tanlov!")
			break
		}
		var input string
		fmt.Println("Davom ettirishni hohlaysizmi?")
		fmt.Println("1.Yes")
		fmt.Println("2.NO")
		fmt.Scan(&input)

		var answer int
		if input == "1" || input == "Yes" || input == "YES" {

			answer = 1

		} else if input == "2" || input == "No" || input == "NO" {
			answer = 2

		} else {
			fmt.Println("Noto'g'ri tanlov.")
			return
		}
		switch answer {
		case 1:
			fmt.Println("Davom etamiz!")
			continue
		case 2:
			fmt.Println("Yakunlandi.")
		}
		break
	}
}
