package problem1

import (
	f "fmt"
	"strconv"
)

func main() {
	f.Println(strToMath("1+2-5"))
}

func strToMath(s string) int {
	angka := []int{}
	simbol := []string{}

	all := []string{}

	var sum int

	for _, v := range s {
		if v == 43 || v == 45 {
			simbol = append(simbol, string(v))
		} else {
			input, _ := strconv.Atoi(string(v))
			angka = append(angka, input)
		}
	}

	for _, v := range s {
		all = append(all, string(v))
	}

	sum, _ = strconv.Atoi(string(all[0]))

	for i := 1; i < len(s); i++ {

		if all[i] == "+" {
			b, _ := strconv.Atoi(string(all[i+1]))
			sum += b
		}

		if all[i] == "-" {
			b, _ := strconv.Atoi(string(all[i+1]))
			sum -= b
		}

	}
	return sum

}
