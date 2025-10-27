package main

import (
	"fmt"
	"strings"
)

type Karyawan struct {
	Nama string
	Usia int
	Gaji float64
}

// Fungsi untuk menambah karyawan
func addKaryawan(data *[]Karyawan, nama string, usia int, gaji float64) {
	*data = append(*data, Karyawan{Nama: nama, Usia: usia, Gaji: gaji})
}

// Fungsi untuk menampilkan semua karyawan
func showKaryawan(data []Karyawan) {
	fmt.Println("\nDaftar Karyawan:")
	for i, k := range data {
		fmt.Printf("%d. Nama: %s | Usia: %d | Gaji: Rp %.2f\n", i+1, k.Nama, k.Usia, k.Gaji)
	}
}

// Fungsi untuk mencari karyawan berdasarkan nama (pakai map)
func findKaryawanByNama(data []Karyawan, nama string) *Karyawan {
	karyawanMap := make(map[string]Karyawan)
	for _, k := range data {
		karyawanMap[strings.ToLower(k.Nama)] = k
	}

	if k, ada := karyawanMap[strings.ToLower(nama)]; ada {
		return &k
	}
	return nil
}

// Fungsi untuk mengubah data karyawan (pakai pointer)
func editGajiKaryawan(k *Karyawan, gajiBaru float64) {
	k.Gaji = gajiBaru
}

// Fungsi untuk menghapus karyawan berdasarkan index
func deleteKaryawan(data *[]Karyawan, index int) {
	if index < 0 || index >= len(*data) {
		fmt.Println("Index tidak valid!")
		return
	}
	*data = append((*data)[:index], (*data)[index+1:]...)
}

func main() {
	var listKaryawan []Karyawan

	addKaryawan(&listKaryawan, "Mulyono", 25, 6000000)
	addKaryawan(&listKaryawan, "Budi Arie", 30, 7000000)
	addKaryawan(&listKaryawan, "Rakabuming", 28, 6500000)

	showKaryawan(listKaryawan)

	fmt.Println("\nCari Karyawan: Budi Arie")
	k := findKaryawanByNama(listKaryawan, "Budi Arie")
	if k != nil {
		fmt.Printf("Ditemukan: %s dengan gaji Rp %.2f\n", k.Nama, k.Gaji)
	}

	fmt.Println("\nUbah gaji karyawan pertama (Mulyono)...")
	editGajiKaryawan(&listKaryawan[0], 6500000)
	showKaryawan(listKaryawan)

	fmt.Println("\nHapus karyawan kedua (Budi Arie)...")
	deleteKaryawan(&listKaryawan, 1)
	showKaryawan(listKaryawan)
}
