package main

import (
	"fmt"
	"strings"
)

type Barang struct {
	Nama   string
	Harga  float64
	Stok   int
}

// Fungsi menambah barang
func addBarang(data *[]Barang, nama string, harga float64, stok int) {
	*data = append(*data, Barang{Nama: nama, Harga: harga, Stok: stok})
}

// Fungsi menampilkan semua barang
func showBarang(data []Barang) {
	fmt.Println("\nDaftar Inventaris Barang:")
	for i, b := range data {
		fmt.Printf("%d. Nama: %s | Harga: Rp %.2f | Stok: %d\n", i+1, b.Nama, b.Harga, b.Stok)
	}
}

// Fungsi mengubah stok barang (menggunakan pointer)
func updateStokBarang(b *Barang, stokBaru int) {
	b.Stok = stokBaru
}

// Fungsi mencari barang berdasarkan nama (pakai map)
func findBarangByNama(data []Barang, nama string) *Barang {
	barangMap := make(map[string]Barang)
	for _, b := range data {
		barangMap[strings.ToLower(b.Nama)] = b
	}

	if b, ada := barangMap[strings.ToLower(nama)]; ada {
		return &b
	}
	return nil
}

// Fungsi menghapus barang dari slice
func deleteBarang(data *[]Barang, index int) {
	if index < 0 || index >= len(*data) {
		fmt.Println("Index tidak valid!")
		return
	}
	*data = append((*data)[:index], (*data)[index+1:]...)
}

func main() {
	var inventaris []Barang

	addBarang(&inventaris, "Laptop", 8500000, 10)
	addBarang(&inventaris, "Mouse", 150000, 50)
	addBarang(&inventaris, "Keyboard", 300000, 30)

	showBarang(inventaris)

	fmt.Println("\nCari Barang: Mouse")
	b := findBarangByNama(inventaris, "Mouse")
	if b != nil {
		fmt.Printf("Ditemukan: %s | Harga: Rp %.2f | Stok: %d\n", b.Nama, b.Harga, b.Stok)
	}

	fmt.Println("\nUbah stok barang pertama (Laptop)...")
	updateStokBarang(&inventaris[0], 8)
	showBarang(inventaris)

	fmt.Println("\nHapus barang kedua (Mouse)...")
	deleteBarang(&inventaris, 1)
	showBarang(inventaris)
}
