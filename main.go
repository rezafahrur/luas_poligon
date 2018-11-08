package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var jumlahKoordinat int
	var totalArea, area, isiX, isiY float64

	fmt.Print("Masukkan Jumlah Koordinat: ")
	if _, err := fmt.Scan(&jumlahKoordinat); err != nil {
		log.Print("  Error: ", err, " For Jumlah Koordinat")
		return
	} else if jumlahKoordinat >= 3 {
		koordinatX := make([]float64, jumlahKoordinat) //slice koordinat x
		koordinatY := make([]float64, jumlahKoordinat) //slice koordinat y

		//loop untuk nginput koordinat
		for poin := 0; poin < jumlahKoordinat; poin++ {
			fmt.Printf("\nmasukkan koordinat X ke-%d: ", poin+1)
			//jika bukan integer return error
			if _, err := fmt.Scan(&isiX); err != nil {
				log.Print("  Error: ", err, " For Koordinat X number ", poin+1)
				return
			}
			//tambahkan inputan ke slice koordinatX
			koordinatX = append(koordinatX[:poin], isiX)

			fmt.Printf("masukkan koordinat Y ke-%d: ", poin+1)
			//jika bukan integer return error
			if _, err := fmt.Scan(&isiY); err != nil {
				log.Print("  Error: ", err, " For Koordinat Y number ", poin+1)
				return
			}
			//tambah inputan ke slice koordinatY
			koordinatY = append(koordinatY[:poin], isiY)
		}

		// //menentukan titik awal
		// fmt.Print("Masukkan Koordinat Awal Berdasarkan elemen Di Atas: ")
		// if _, err := fmt.Scan(&koordinatAwal); err != nil {
		// 	log.Print("  Error: ", err, " For Koordinat Awal")
		// 	return
		// }

		// //hitung totalArea jika poligon
		// //fix index
		// koordinatAwal--
		// //tampilkan koordinat yang dipilih
		// fmt.Printf("\nKoordinat Awal: (%d,%d)\n", koordinatX[koordinatAwal], koordinatY[koordinatAwal])

		//tampil seluruh koordinat & luas
		fmt.Println("\n\nKoordinat Keseluruhan")
		for elemen := 0; elemen < jumlahKoordinat; elemen++ {
			nextEl := elemen + 1
			fmt.Printf("Titik %d: (%.1f,%.1f)\n", nextEl, koordinatX[elemen], koordinatY[elemen])
			koordinatX = append(koordinatX, koordinatX[0])
			koordinatY = append(koordinatY, koordinatY[0])
			area = (koordinatX[elemen] * koordinatY[elemen+1]) - (koordinatX[elemen+1] * koordinatY[elemen])
			totalArea += area
		}
		// selalu positif
		totalArea = math.Abs(totalArea)
		totalArea /= 2
		fmt.Printf("\nAsumsi Titik Awal adalah Titik 1\nMaka Luas Bidang Ini adalah: %.2f\n", totalArea)

	} else {
		log.Print("Untuk Menghitung Luas Jumlah Koordinat Harus Lebih dari 2")
		return
	}

}
