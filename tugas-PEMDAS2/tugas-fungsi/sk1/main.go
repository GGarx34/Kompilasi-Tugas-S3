package main

import (
	"fmt"
	"math"
)
// Fungsi untuk menghitung total harga dari barang yang dibeli
func countTotalHarga(hargaBarang []float64, jumlahBarang []int) (float64, error) {
	if len(hargaBarang) != len(jumlahBarang) {
		return 0, fmt.Errorf("jumlah barang dan harga tidak sesuai")
	}

	total := 0.0
	for i := 0; i < len(hargaBarang); i++ {
		total += hargaBarang[i] * float64(jumlahBarang[i])
	}
	return total, nil
}

// Fungsi untuk menerapkan diskon tetap
func diskonTetap(total float64, potongan float64) float64 {
	return math.Max(total-potongan, 0) // Tidak boleh kurang dari nol
}

// Fungsi untuk menerapkan diskon persentase
func diskonPersentase(total float64, persentase float64) float64 {
	potongan := total * (persentase / 100)
	return math.Max(total-potongan, 0) // Tidak boleh kurang dari nol
}

// Closure untuk menerapkan diskon kombinasi
func diskonKombinasi(
	diskon1 func(float64) float64,
	diskon2 func(float64) float64,
) func(float64) float64 {
	return func(total float64) float64 {
		totalSetelahDiskon1 := diskon1(total)
		totalSetelahDiskon2 := diskon2(totalSetelahDiskon1)
		return totalSetelahDiskon2
	}
}

func main() {
	// Daftar harga barang dan jumlah yang dibeli
	hargaBarang := []float64{100000, 50000, 75000} 
	jumlahBarang := []int{2, 1, 3}                

	// Menghitung total harga barang
	totalHarga, err := countTotalHarga(hargaBarang, jumlahBarang)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total Harga Sebelum Diskon: Rp %.2f\n", totalHarga)

	// Menerapkan diskon tetap Rp 20.000
	totalSetelahDiskonTetap := diskonTetap(totalHarga, 20000)
	fmt.Printf("Total Setelah Diskon Tetap Rp 20.000: Rp %.2f\n", totalSetelahDiskonTetap)

	// Menerapkan diskon persentase 10%
	totalSetelahDiskonPersen := diskonPersentase(totalHarga, 10)
	fmt.Printf("Total Setelah Diskon Persen 10%%: Rp %.2f\n", totalSetelahDiskonPersen)

	// Menggabungkan diskon tetap dan diskon persentase (Rp 20.000 + 10%)
	diskonGabungan := diskonKombinasi(
		func(total float64) float64 { return diskonTetap(total, 20000) },
		func(total float64) float64 { return diskonPersentase(total, 10) },
	)

	totalSetelahDiskonGabungan := diskonGabungan(totalHarga)
	fmt.Printf("Total Setelah Diskon Kombinasi (Rp 20.000 + 10%%): Rp %.2f\n", totalSetelahDiskonGabungan)
}
