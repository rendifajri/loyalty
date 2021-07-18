package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var n, m int
	fmt.Println("Masukkan jumlah pemain: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Masukkan jumlah dadu: ")
	_, err = fmt.Scan(&m)
	if err != nil {
		fmt.Println(err)
		return
	}

	poin := make([]int, n)
	for i := 0; i < n; i++ {
		poin[i] = 0
	}
	dadu := make([][]string, n, m)
	for i := 0; i < n; i++ {
		dadu[i] = make([]string, m)
		for j := 0; j < m; j++ {
			dadu[i][j] = "0" //inisiasi dadu
		}
	}
	rand.Seed(time.Now().UTC().UnixNano())
	giliran := 1
	for len(cekJmlDaduPemain(dadu, n)) != 1 {
		fmt.Println("==================")
		fmt.Printf("Giliran %d lempar dadu:\n", giliran)
		temp_dadu := make([]int, n)
		for i, po := range poin {
			temp_dadu[i] = 0
			temp_i := i + 1
			fmt.Printf("Pemain #%d (%d): ", temp_i, po)
			for j := 0; j < len(dadu[i]); j++ {
				rand_dadu := rand.Intn(6) + 1
				dadu[i][j] = fmt.Sprint(rand_dadu)
			}
			if len(dadu[i]) > 0 {
				fmt.Printf(strings.Join(dadu[i], ", "))
			} else {
				fmt.Printf("_ (Berhenti bermain karena tidak memiliki dadu)")
			}
			fmt.Printf("\n")
		}
		for i := 0; i < n; i++ {
			for j := 0; j < len(dadu[i]); j++ {
				if dadu[i][j] == "1" {
					dadu[i] = append(dadu[i][:j], dadu[i][j+1:]...)
					j -= 1
					temp_i := i + 1
					if temp_i >= n {
						temp_i = 0
					}
					for len(dadu[temp_i]) == 0 {
						temp_i = temp_i + 1
						if temp_i >= n {
							temp_i = 0
						}
					}
					temp_dadu[temp_i]++
				} else if dadu[i][j] == "6" {
					dadu[i] = append(dadu[i][:j], dadu[i][j+1:]...)
					j -= 1
					poin[i]++
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < temp_dadu[i]; j++ {
				dadu[i] = append(dadu[i], "1")
			}
		}
		fmt.Println("Setelah evaluasi:")
		for i, po := range poin {
			temp_i := i + 1
			fmt.Printf("Pemain #%d (%d): ", temp_i, po)
			if len(dadu[i]) > 0 {
				fmt.Printf(strings.Join(dadu[i], ", "))
			} else {
				fmt.Printf("_ (Berhenti bermain karena tidak memiliki dadu)")
			}
			fmt.Printf("\n")
		}
		giliran++
	}
	fmt.Println("==================")
	nolDaduPemain := cekJmlDaduPemain(dadu, n)
	fmt.Printf("Game berakhir karena pemain hanya %s yang memiliki dadu.\n", strings.Join(nolDaduPemain, ", "))
	poinTerbanyak := 0
	for i := 0; i < n; i++ {
		if poin[i] > poinTerbanyak {
			poinTerbanyak = poin[i]
		}
	}
	var poinPemain []string
	for i := 0; i < n; i++ {
		if poin[i] == poinTerbanyak {
			temp_i := i + 1
			poinPemain = append(poinPemain, fmt.Sprint("#", temp_i))
		}
	}
	fmt.Printf("Game dimenangkan oleh pemain %s.", strings.Join(poinPemain, ", "))
}

func cekJmlDaduPemain(dadu [][]string, n int) []string {
	var nolDaduPemain []string
	for i := 0; i < n; i++ {
		if len(dadu[i]) > 0 {
			temp_i := i + 1
			nolDaduPemain = append(nolDaduPemain, fmt.Sprint("#", temp_i))
		}
	}
	return nolDaduPemain
}
