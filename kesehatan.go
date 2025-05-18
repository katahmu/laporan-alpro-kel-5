package main

import (
	"fmt"
)

type Pasien struct {
	id     string
	nama   string
	umur   int
	gender string
}

type Dokter struct {
	id           string
	nama         string
	spesialisasi string
	gender       string
	jadwal       string
}

type Riwayat struct {
	id       string
	pasienID string
	dokterID string
	tanggal  string
	diagnosa string
}

var (
	daftarPasien  []Pasien
	daftarDokter  []Dokter
	daftarRiwayat []Riwayat
)

func main() {
	// Inisialisasi beberapa data contoh
	daftarPasien = []Pasien{
		{"P001", "Assyifa Dwi Safitri", 35, "P"},
		{"P002", "Amirah Essary Yunsarah Sujuthi", 28, "P"},
		{"P003", "Keisha Hananta", 45, "P"},
		{"P004", "Dit Adit", 20, "L"},
	}

	daftarDokter = []Dokter{
		{"D001", "Dr. Ahmad", "Kardiologi", "L", "Senin-Jumat, 08.00-12.00"},
		{"D002", "Dr. Budi", "Kedokteran Umum", "L", "Senin-Jumat, 13.00-17.00"},
		{"D003", "Dr. Candra", "Kedokteran Gigi", "L", "Senin-Jumat, 08.00-12.00"},
		{"D004", "Dr. Dwi", "Kedokteran Gigi", "P", "Senin-Jumat, 13.00-17.00"},
		{"D005", "Dr. Eka", "Kedokteran Mata", "P", "Senin-Jumat, 08.00-12.00"},
	}

	daftarRiwayat = []Riwayat{
		{"R001", "P001", "D001", "2023-07-01", "Cephalgia"},
		{"R002", "P002", "D002", "2023-07-02", "Nausea"},
		{"R003", "P003", "D003", "2023-07-03", "Odontalgia"},
		{"R004", "P004", "D004", "2023-07-04", "Karies Gigi"},
		{"R005", "P005", "D005", "2023-07-05", "Miopia"},
	}

	for {
		fmt.Println("\n=== SISTEM INFORMASI RUMAH SAKIT ===")
		fmt.Println("1. Kelola Data Pasien")
		fmt.Println("2. Kelola Data Dokter")
		fmt.Println("3. Kelola Data Riwayat")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu (1-4): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			menuPasien()
		case 2:
			menuDokter()
		case 3:
			menuRiwayat()
		case 4:
			fmt.Println("Terima kasih telah menggunakan sistem ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid! Silakan pilih 1-4.")
		}
	}
}

// Manajemen Pasien
func menuPasien() {
	for {
		fmt.Println("\n=== MENU PASIEN ===")
		fmt.Println("1. Tampilkan Semua Pasien")
		fmt.Println("2. Tambah Pasien Baru")
		fmt.Println("3. Edit Data Pasien")
		fmt.Println("4. Hapus Data Pasien")
		fmt.Println("5. Urutkan Pasien Berdasarkan Umur")
		fmt.Println("6. Urutkan Pasien Berdasarkan Nama")
		fmt.Println("7. Cari Pasien Berdasarkan ID")
		fmt.Println("8. Kembali ke Menu Utama")
		fmt.Print("Pilih menu (1-8): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilkanPasien()
		case 2:
			tambahPasien()
		case 3:
			editPasien()
		case 4:
			hapusPasien()
		case 5:
			urutkanPasienByUmur()
			tampilkanPasien()
		case 6:
			urutkanPasienByNama()
			tampilkanPasien()
		case 7:
			fmt.Print("Masukkan ID Pasien: ")
			var id string
			fmt.Scanln(&id)
			cariPasienById(id)
		case 8:
			return
		default:
			fmt.Println("Pilihan tidak valid! Silakan pilih 1-8.")
		}
	}
}

func tampilkanPasien() {
	fmt.Println("\nDaftar Pasien:")
	fmt.Println("==========================================")
	fmt.Printf("%-6s %-25s %-5s %-10s\n", "ID", "Nama", "Umur", "Gender")
	fmt.Println("==========================================")
	for i := 0; i < len(daftarPasien); i++ {
		p := daftarPasien[i]
		fmt.Printf("%-6s %-25s %-5d %-10s\n", p.id, p.nama, p.umur, p.gender)
	}
}

func tambahPasien() {
	var pasien Pasien
	fmt.Println("\nTambah Pasien Baru")
	fmt.Print("ID Pasien: ")
	fmt.Scanln(&pasien.id)
	fmt.Print("Nama Pasien: ")
	fmt.Scanln(&pasien.nama)
	fmt.Print("Umur Pasien: ")
	fmt.Scanln(&pasien.umur)
	fmt.Print("Gender (L/P): ")
	fmt.Scanln(&pasien.gender)

	daftarPasien = append(daftarPasien, pasien)
	fmt.Println("Pasien berhasil ditambahkan!")
}

func editPasien() {
	fmt.Print("Masukkan ID Pasien yang akan diedit: ")
	var id string
	fmt.Scanln(&id)

	for i := 0; i < len(daftarPasien); i++ {
		if daftarPasien[i].id == id {
			fmt.Println("\nData Pasien Ditemukan:")
			fmt.Printf("%-6s %-25s %-5s %-10s\n", "ID", "Nama", "Umur", "Gender")
			fmt.Printf("%-6s %-25s %-5d %-10s\n", daftarPasien[i].id, daftarPasien[i].nama, daftarPasien[i].umur, daftarPasien[i].gender)

			fmt.Print("\nMasukkan data baru:\n")
			fmt.Print("Nama Pasien: ")
			fmt.Scanln(&daftarPasien[i].nama)
			fmt.Print("Umur Pasien: ")
			fmt.Scanln(&daftarPasien[i].umur)
			fmt.Print("Gender (L/P): ")
			fmt.Scanln(&daftarPasien[i].gender)

			fmt.Println("Data pasien berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Pasien dengan ID", id, "tidak ditemukan.")
}

func hapusPasien() {
	fmt.Print("Masukkan ID Pasien yang akan dihapus: ")
	var id string
	fmt.Scanln(&id)

	for i := 0; i < len(daftarPasien); i++ {
		if daftarPasien[i].id == id {
			// Hapus riwayat yang terkait dengan pasien ini
			for j := 0; j < len(daftarRiwayat); j++ {
				if daftarRiwayat[j].pasienID == id {
					daftarRiwayat = append(daftarRiwayat[:j], daftarRiwayat[j+1:]...)
					j-- 
				}
			}

			// Hapus pasien
			daftarPasien = append(daftarPasien[:i], daftarPasien[i+1:]...)
			fmt.Println("Pasien dan riwayat terkait berhasil dihapus!")
			return
		}
	}
	fmt.Println("Pasien dengan ID", id, "tidak ditemukan.")
}

func urutkanPasienByUmur() {
	for i := 0; i < len(daftarPasien); i++ {
		minIndex := i
		for j := i + 1; j < len(daftarPasien); j++ {
			if daftarPasien[j].umur < daftarPasien[minIndex].umur {
				minIndex = j
			}
		}
		if minIndex != i {
			daftarPasien[i], daftarPasien[minIndex] = daftarPasien[minIndex], daftarPasien[i]
		}
	}
}

func urutkanPasienByNama() {
	for i := 1; i < len(daftarPasien); i++ {
		key := daftarPasien[i]
		j := i - 1
		for j >= 0 && daftarPasien[j].nama > key.nama {
			daftarPasien[j+1] = daftarPasien[j]
			j--
		}
		daftarPasien[j+1] = key
	}
}

func cariPasienById(id string) {
	found := false
	for i := 0; i < len(daftarPasien); i++ {
		if daftarPasien[i].id == id {
			fmt.Println("\nData Pasien Ditemukan:")
			fmt.Println("==========================================")
			fmt.Printf("%-6s %-25s %-5s %-10s\n", "ID", "Nama", "Umur", "Gender")
			fmt.Println("==========================================")
			fmt.Printf("%-6s %-25s %-5d %-10s\n", daftarPasien[i].id, daftarPasien[i].nama, daftarPasien[i].umur, daftarPasien[i].gender)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Pasien dengan ID", id, "tidak ditemukan.")
	}
}
func menuDokter() {
	for {
		fmt.Println("\n=== MENU DOKTER ===")
		fmt.Println("1. Tampilkan Semua Dokter")
		fmt.Println("2. Tambah Dokter Baru")
		fmt.Println("3. Edit Data Dokter")
		fmt.Println("4. Hapus Data Dokter")
		fmt.Println("5. Urutkan Dokter Berdasarkan Nama")
		fmt.Println("6. Cari Dokter Berdasarkan ID (Binary Search)")
		fmt.Println("7. Kembali ke Menu Utama")
		fmt.Print("Pilih menu (1-7): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilkanDokter()
		case 2:
			tambahDokter()
		case 3:
			editDokter()
		case 4:
			hapusDokter()
		case 5:
			urutkanDokterByNama()
			tampilkanDokter()
		case 6:
			fmt.Print("Masukkan ID Dokter: ")
			var id string
			fmt.Scanln(&id)
			cariDokterByIdBinary(id)
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid! Silakan pilih 1-7.")
		}
	}
}

func tampilkanDokter() {
	fmt.Println("\nDaftar Dokter:")
	fmt.Println("=============================================================")
	fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", "ID", "Nama", "Spesialisasi", "Gender", "Jadwal")
	fmt.Println("=============================================================")
	for i := 0; i < len(daftarDokter); i++ {
		d := daftarDokter[i]
		fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", d.id, d.nama, d.spesialisasi, d.gender, d.jadwal)
	}
}

func tambahDokter() {
	var dokter Dokter
	fmt.Println("\nTambah Dokter Baru")
	fmt.Print("ID Dokter: ")
	fmt.Scanln(&dokter.id)
	fmt.Print("Nama Dokter: ")
	fmt.Scanln(&dokter.nama)
	fmt.Print("Spesialisasi: ")
	fmt.Scanln(&dokter.spesialisasi)
	fmt.Print("Gender (L/P): ")
	fmt.Scanln(&dokter.gender)
	fmt.Print("Jadwal Praktek: ")
	fmt.Scanln(&dokter.jadwal)

	daftarDokter = append(daftarDokter, dokter)
	fmt.Println("Dokter berhasil ditambahkan!")
}

func editDokter() {
	fmt.Print("Masukkan ID Dokter yang akan diedit: ")
	var id string
	fmt.Scanln(&id)

	for i := 0; i < len(daftarDokter); i++ {
		if daftarDokter[i].id == id {
			fmt.Println("\nData Dokter Ditemukan:")
			fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", "ID", "Nama", "Spesialisasi", "Gender", "Jadwal")
			fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", daftarDokter[i].id, daftarDokter[i].nama, daftarDokter[i].spesialisasi, daftarDokter[i].gender, daftarDokter[i].jadwal)

			fmt.Print("\nMasukkan data baru:\n")
			fmt.Print("Nama Dokter: ")
			fmt.Scanln(&daftarDokter[i].nama)
			fmt.Print("Spesialisasi: ")
			fmt.Scanln(&daftarDokter[i].spesialisasi)
			fmt.Print("Gender (L/P): ")
			fmt.Scanln(&daftarDokter[i].gender)
			fmt.Print("Jadwal Praktek: ")
			fmt.Scanln(&daftarDokter[i].jadwal)

			fmt.Println("Data dokter berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
}

func hapusDokter() {
	fmt.Print("Masukkan ID Dokter yang akan dihapus: ")
	var id string
	fmt.Scanln(&id)

	for i := 0; i < len(daftarDokter); i++ {
		if daftarDokter[i].id == id {
			// Hapus riwayat yang terkait dengan dokter ini
			for j := 0; j < len(daftarRiwayat); j++ {
				if daftarRiwayat[j].dokterID == id {
					daftarRiwayat = append(daftarRiwayat[:j], daftarRiwayat[j+1:]...)
					j-- 
				}
			}

			// Hapus dokter
			daftarDokter = append(daftarDokter[:i], daftarDokter[i+1:]...)
			fmt.Println("Dokter dan riwayat terkait berhasil dihapus!")
			return
		}
	}
	fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
}

func urutkanDokterByNama() {
	for i := 0; i < len(daftarDokter); i++ {
		minIndex := i
		for j := i + 1; j < len(daftarDokter); j++ {
			if daftarDokter[j].nama < daftarDokter[minIndex].nama {
				minIndex = j
			}
		}
		if minIndex != i {
			daftarDokter[i], daftarDokter[minIndex] = daftarDokter[minIndex], daftarDokter[i]
		}
	}
}

func cariDokterByIdBinary(id string) {
	// First sort the doctors by ID using bubble sort
	n := len(daftarDokter)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if daftarDokter[j].id > daftarDokter[j+1].id {
				daftarDokter[j], daftarDokter[j+1] = daftarDokter[j+1], daftarDokter[j]
			}
		}
	}

	// Perform binary search
	low := 0
	high := len(daftarDokter) - 1
	found := false
	var mid int

	for low <= high {
		mid = (low + high) / 2

		if daftarDokter[mid].id == id {
			found = true
			break
		} else if daftarDokter[mid].id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		d := daftarDokter[mid]
		fmt.Println("\nData Dokter Ditemukan (Binary Search):")
		fmt.Println("=============================================================")
		fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", "ID", "Nama", "Spesialisasi", "Gender", "Jadwal")
		fmt.Println("=============================================================")
		fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", d.id, d.nama, d.spesialisasi, d.gender, d.jadwal)
	} else {
		fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
	}
}

// Manajemen Riwayat Medis
func menuRiwayat() {
	for {
		fmt.Println("\n=== MENU RIWAYAT ===")
		fmt.Println("1. Tampilkan Semua Riwayat")
		fmt.Println("2. Tambah Riwayat Baru")
		fmt.Println("3. Edit Data Riwayat")
		fmt.Println("4. Hapus Data Riwayat")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilih menu (1-5): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilkanRiwayat()
		case 2:
			tambahRiwayat()
		case 3:
			editRiwayat()
		case 4:
			hapusRiwayat()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid! Silakan pilih 1-5.")
		}
	}
}

func tampilkanRiwayat() {
	fmt.Println("\nDaftar Riwayat:")
	fmt.Println("====================================================================")
	fmt.Printf("%-6s %-10s %-10s %-12s %-30s\n", "ID", "Pasien ID", "Dokter ID", "Tanggal", "Diagnosa")
	fmt.Println("====================================================================")
	for i := 0; i < len(daftarRiwayat); i++ {
		r := daftarRiwayat[i]
		fmt.Printf("%-6s %-10s %-10s %-12s %-30s\n", r.id, r.pasienID, r.dokterID, r.tanggal, r.diagnosa)
	}
}

func tambahRiwayat() {
	var riwayat Riwayat
	fmt.Println("\nTambah Riwayat Baru")
	fmt.Print("ID Riwayat: ")
	fmt.Scanln(&riwayat.id)
	fmt.Print("ID Pasien: ")
	fmt.Scanln(&riwayat.pasienID)
	fmt.Print("ID Dokter: ")
	fmt.Scanln(&riwayat.dokterID)
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scanln(&riwayat.tanggal)
	fmt.Print("Diagnosa: ")
	fmt.Scanln(&riwayat.diagnosa)

	daftarRiwayat = append(daftarRiwayat, riwayat)
	fmt.Println("Riwayat berhasil ditambahkan!")
}

func editRiwayat() {
	fmt.Print("Masukkan ID Riwayat yang akan diedit: ")
	var id string
	fmt.Scanln(&id)

	for i := 0; i < len(daftarRiwayat); i++ {
		if daftarRiwayat[i].id == id {
			fmt.Println("\nData Riwayat Ditemukan:")
			fmt.Printf("%-6s %-10s %-10s %-12s %-30s\n", "ID", "Pasien ID", "Dokter ID", "Tanggal", "Diagnosa")
			fmt.Printf("%-6s %-10s %-10s %-12s %-30s\n", daftarRiwayat[i].id, daftarRiwayat[i].pasienID, daftarRiwayat[i].dokterID, daftarRiwayat[i].tanggal, daftarRiwayat[i].diagnosa)

			fmt.Print("\nMasukkan data baru:\n")
			fmt.Print("ID Pasien: ")
			fmt.Scanln(&daftarRiwayat[i].pasienID)
			fmt.Print("ID Dokter: ")
			fmt.Scanln(&daftarRiwayat[i].dokterID)
			fmt.Print("Tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&daftarRiwayat[i].tanggal)
			fmt.Print("Diagnosa: ")
			fmt.Scanln(&daftarRiwayat[i].diagnosa)

			fmt.Println("Data riwayat berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Riwayat dengan ID", id, "tidak ditemukan.")
}

func hapusRiwayat() {
	fmt.Print("Masukkan ID Riwayat yang akan dihapus: ")
	var id string
	fmt.Scanln(&id)

	for i := 0; i < len(daftarRiwayat); i++ {
		if daftarRiwayat[i].id == id {
			daftarRiwayat = append(daftarRiwayat[:i], daftarRiwayat[i+1:]...)
			fmt.Println("Riwayat berhasil dihapus!")
			return
		}
	}
	fmt.Println("Riwayat dengan ID", id, "tidak ditemukan.")
}
