package main

import (
	"fmt"
)

// 1. Definisi struktur data
type Pasien struct {
	ID     string
	Nama   string
	Umur   int
	Gender string
}

type Dokter struct {
	ID   string
	Nama string
}

type Riwayat struct {
	PasienID string
	DokterID string
}

// 2. Fungsi untuk menampilkan data pasien
func tampilkanPasien(pasien []Pasien) {
	fmt.Println("\nDaftar Pasien:")
	fmt.Println("==========================================")
	fmt.Printf("%-6s %-15s %-5s %-10s\n", "ID", "Nama", "Umur", "Gender")
	fmt.Println("==========================================")
	for i := range pasien {
		p := pasien[i]
		fmt.Printf("%-6s %-15s %-5d %-10s\n", p.ID, p.Nama, p.Umur, p.Gender)
	}
	fmt.Println()
}

// 3. Selection Sort berdasarkan umur
func urutkanByUmur(pasien []Pasien) []Pasien {
	for i := range pasien {
		minIndex := i
		for j := i; j < len(pasien); j++ {
			if pasien[j].Umur < pasien[minIndex].Umur {
				minIndex = j
			}
		}
		if minIndex != i {
			pasien[i], pasien[minIndex] = pasien[minIndex], pasien[i]
		}
	}
	return pasien
}

// 4. Binary Insertion Sort berdasarkan nama
func urutkanByNama(pasien []Pasien) []Pasien {
	for i := range pasien {
		key := pasien[i]
		left := 0
		right := i - 1
		
		for left <= right {
			mid := (left + right) / 2
			if pasien[mid].Nama < key.Nama {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		
		for j := i - 1; j >= left; j-- {
			pasien[j+1] = pasien[j]
		}
		pasien[left] = key
	}
	return pasien
}

// 5. Sequential Search by ID
func cariByID(pasien []Pasien, id string) *Pasien {
	for i := range pasien {
		if pasien[i].ID == id {
			return &pasien[i]
		}
	}
	return nil
}

// 6. Binary Search by Nama (harus data sudah terurut)
func cariByNama(pasien []Pasien, nama string) *Pasien {
	left := 0
	right := len(pasien) - 1
	
	for left <= right {
		mid := (left + right) / 2
		if pasien[mid].Nama == nama {
			return &pasien[mid]
		} else if pasien[mid].Nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nil
}

func main() {
	// 7. Data contoh pasien
	pasien := []Pasien{
		{"P001", "Budi Santoso", 35, "Laki-laki"},
		{"P002", "Ani Wijaya", 28, "Perempuan"},
		{"P003", "Citra Dewi", 45, "Perempuan"},
	}

	// 8. Tampilkan data awal
	fmt.Println("=== Data Awal Pasien ===")
	tampilkanPasien(pasien)

	// 9. Sorting dan tampilkan hasil
	fmt.Println("=== Pasien Terurut Berdasarkan Umur ===")
	pasienByUmur := urutkanByUmur(pasien)
	tampilkanPasien(pasienByUmur)

	fmt.Println("=== Pasien Terurut Berdasarkan Nama ===")
	pasienByNama := urutkanByNama(pasien)
	tampilkanPasien(pasienByNama)

	// 10. Searching dan tampilkan hasil
	fmt.Println("=== Cari Pasien by ID (P002) ===")
	hasilCariID := cariByID(pasien, "P002")
	if hasilCariID != nil {
		tampilkanPasien([]Pasien{*hasilCariID})
	} else {
		fmt.Println("Pasien tidak ditemukan")
	}

	fmt.Println("=== Cari Pasien by Nama (Ani Wijaya) ===")
	hasilCariNama := cariByNama(pasienByNama, "Ani Wijaya")
	if hasilCariNama != nil {
		tampilkanPasien([]Pasien{*hasilCariNama})
	} else {
		fmt.Println("Pasien tidak ditemukan")
	}
}