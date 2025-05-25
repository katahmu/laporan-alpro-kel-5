package main

import (
	"fmt"
	"strconv"
)

// Struktur data untuk menyimpan informasi pasien
type Pasien struct {
	id     string
	nama   string
	umur   int
	gender string
}

// Struktur data untuk menyimpan informasi dokter
type Dokter struct {
	id           string
	nama         string
	spesialisasi string
	gender       string
	jadwal       string
}

// Struktur data untuk menyimpan informasi riwayat medis
type Riwayat struct {
	id       string
	pasienID string
	dokterID string
	tanggal  string
	diagnosa string
}

// Konstanta untuk menentukan ukuran maksimum data
const maxData = 100

// Variabel global untuk menyimpan data
var (
	daftarPasien  [maxData]Pasien
	daftarDokter  [maxData]Dokter
	daftarRiwayat [maxData]Riwayat
	jumlahPasien  int
	jumlahDokter  int
	jumlahRiwayat int
)

func main() {
	// Inisialisasi data contoh
	initData()

	for {
		fmt.Println("\n=== MEDICORE - SISTEM INFORMASI RUMAH SAKIT ===")
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

// Fungsi untuk menginisialisasi data awal
func initData() {
	// Inisialisasi pasien
	daftarPasien[0] = Pasien{"P001", "Assyifa Dwi Safitri", 35, "P"}
	daftarPasien[1] = Pasien{"P002", "Amirah Sujuthi", 28, "P"}
	daftarPasien[2] = Pasien{"P003", "Keisha Hananta", 45, "P"}
	daftarPasien[3] = Pasien{"P004", "Dit Adit", 20, "L"}
	jumlahPasien = 4

	// Inisialisasi dokter
	daftarDokter[0] = Dokter{"D001", "Dr. Ahmad", "Kardiologi", "L", "Senin-Jumat, 08.00-12.00"}
	daftarDokter[1] = Dokter{"D002", "Dr. Budi", "Kedokteran Umum", "L", "Senin-Jumat, 13.00-17.00"}
	daftarDokter[2] = Dokter{"D003", "Dr. Candra", "Kedokteran Gigi", "L", "Senin-Jumat, 08.00-12.00"}
	daftarDokter[3] = Dokter{"D004", "Dr. Dwi", "Kedokteran Gigi", "P", "Senin-Jumat, 13.00-17.00"}
	daftarDokter[4] = Dokter{"D005", "Dr. Eka", "Kedokteran Mata", "P", "Senin-Jumat, 08.00-12.00"}
	jumlahDokter = 5

	// Inisialisasi riwayat
	daftarRiwayat[0] = Riwayat{"R001", "P001", "D001", "2023-07-01", "Cephalgia"}
	daftarRiwayat[1] = Riwayat{"R002", "P002", "D002", "2023-07-02", "Nausea"}
	daftarRiwayat[2] = Riwayat{"R003", "P003", "D003", "2023-07-03", "Odontalgia"}
	daftarRiwayat[3] = Riwayat{"R004", "P004", "D004", "2023-07-04", "Karies Gigi"}
	daftarRiwayat[4] = Riwayat{"R005", "P001", "D005", "2023-07-05", "Miopia"}
	jumlahRiwayat = 5
}



// MANAJEMEN PASIEN
func menuPasien() {
	for {
		fmt.Println("\n=== MENU PASIEN ===")
		fmt.Println("1. Tampilkan Semua Pasien")
		fmt.Println("2. Tambah Pasien Baru")
		fmt.Println("3. Edit Data Pasien")
		fmt.Println("4. Hapus Data Pasien")
		fmt.Println("5. Urutkan Pasien Berdasarkan Umur (Selection Sort)")
		fmt.Println("6. Urutkan Pasien Berdasarkan Nama (Insertion Sort)")
		fmt.Println("7. Cari Pasien Berdasarkan ID (Sequential Search)")
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

// Fungsi untuk menampilkan daftar pasien
func tampilkanPasien() {
	fmt.Println("\nDaftar Pasien:")
	fmt.Println("==========================================")
	fmt.Printf("%-6s %-25s %-5s %-10s\n", "ID", "Nama", "Umur", "Gender")
	fmt.Println("==========================================")
	for i := 0; i < jumlahPasien; i++ {
		p := daftarPasien[i]
		fmt.Printf("%-6s %-25s %-5d %-10s\n", p.id, p.nama, p.umur, p.gender)
	}
}

// Fungsi untuk menghasilkan ID pasien baru secara otomatis
func generatePasienID() string {
	if jumlahPasien == 0 {
		return "P001"
	}
	lastID := daftarPasien[jumlahPasien-1].id
	num, _ := strconv.Atoi(lastID[1:])
	return fmt.Sprintf("P%03d", num+1)
}

// Fungsi untuk menambah pasien baru
func tambahPasien() {
	if jumlahPasien >= maxData {
		fmt.Println("Kapasitas pasien penuh!")
		return
	}

	var pasien Pasien

	pasien.id = generatePasienID() // Menghasilkan ID otomatis
	
	fmt.Println("\nTambah Pasien Baru")

	fmt.Println("ID Pasien:", pasien.id) // Menampilkan ID yang dihasilkan
	fmt.Print("Nama Pasien: ")
	fmt.Scanln(&pasien.nama)
	fmt.Print("Umur Pasien: ")
	fmt.Scanln(&pasien.umur)
	fmt.Print("Gender (L/P): ")
	fmt.Scanln(&pasien.gender)

	daftarPasien[jumlahPasien] = pasien
	jumlahPasien++
	fmt.Println("Pasien berhasil ditambahkan!")
}

// Fungsi untuk mengedit data pasien
func editPasien() {
	fmt.Print("Masukkan ID Pasien yang akan diedit: ")
	var id string
	fmt.Scanln(&id)

	found := false
	for i := 0; i < jumlahPasien && !found; i++ {
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
			found = true
		}
	}

	if !found {
		fmt.Println("Pasien dengan ID", id, "tidak ditemukan.")
	}
}

// Fungsi untuk menghapus data pasien
func hapusPasien() {
	fmt.Print("Masukkan ID Pasien yang akan dihapus: ")
	var id string
	fmt.Scanln(&id)

	found := false
	for i := 0; i < jumlahPasien && !found; i++ {
		if daftarPasien[i].id == id {
			for j := 0; j < jumlahRiwayat; {
				if daftarRiwayat[j].pasienID == id {
					for k := j; k < jumlahRiwayat-1; k++ {
						daftarRiwayat[k] = daftarRiwayat[k+1]
					}
					jumlahRiwayat--
				} else {
					j++
				}
			}

			for j := i; j < jumlahPasien-1; j++ {
				daftarPasien[j] = daftarPasien[j+1]
			}
			jumlahPasien--

			fmt.Println("Pasien dan riwayat terkait berhasil dihapus!")
			found = true
		}
	}

	if !found {
		fmt.Println("Pasien dengan ID", id, "tidak ditemukan.")
	}
}

// Fungsi untuk mengurutkan pasien berdasarkan umur menggunakan Selection Sort
func urutkanPasienByUmur() {

	// Selection Sort Ascending (urutan naik)
	for i := 0; i < jumlahPasien-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlahPasien; j++ {
			if daftarPasien[j].umur < daftarPasien[minIndex].umur {
				minIndex = j
			}
		}
		if minIndex != i {
			daftarPasien[i], daftarPasien[minIndex] = daftarPasien[minIndex], daftarPasien[i]
		}
	}
}

// Fungsi untuk mengurutkan pasien berdasarkan nama menggunakan Insertion Sort
func urutkanPasienByNama() {

	// Insertion Sort Ascending (urutan naik)
	for i := 1; i < jumlahPasien; i++ {
		key := daftarPasien[i]
		j := i - 1
		for j >= 0 && daftarPasien[j].nama > key.nama {
			daftarPasien[j+1] = daftarPasien[j]
			j--
		}
		daftarPasien[j+1] = key
	}
}

// Fungsi untuk mencari pasien berdasarkan ID
func cariPasienById(id string) {
	found := false
	for i := 0; i < jumlahPasien && !found; i++ {
		if daftarPasien[i].id == id {
			fmt.Println("\nData Pasien Ditemukan:")
			fmt.Println("==========================================")
			fmt.Printf("%-6s %-25s %-5s %-10s\n", "ID", "Nama", "Umur", "Gender")
			fmt.Println("==========================================")
			fmt.Printf("%-6s %-25s %-5d %-10s\n", daftarPasien[i].id, daftarPasien[i].nama, daftarPasien[i].umur, daftarPasien[i].gender)
			found = true
		}
	}
	if !found {
		fmt.Println("Pasien dengan ID", id, "tidak ditemukan.")
	}
}

// MANAJEMEN DOKTER
func menuDokter() {
	for {
		fmt.Println("\n=== MENU DOKTER ===")
		fmt.Println("1. Tampilkan Semua Dokter")
		fmt.Println("2. Tambah Dokter Baru")
		fmt.Println("3. Edit Data Dokter")
		fmt.Println("4. Hapus Data Dokter")
		fmt.Println("5. Urutkan Dokter Berdasarkan Nama (Selection Sort)")
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

// Fungsi untuk menampilkan daftar dokter
func tampilkanDokter() {
	fmt.Println("\nDaftar Dokter:")
	fmt.Println("=============================================================")
	fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", "ID", "Nama", "Spesialisasi", "Gender", "Jadwal")
	fmt.Println("=============================================================")
	for i := 0; i < jumlahDokter; i++ {
		d := daftarDokter[i]
		fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", d.id, d.nama, d.spesialisasi, d.gender, d.jadwal)
	}
}

// Fungsi untuk menghasilkan ID dokter secara otomatis
func generateDokterID() string {
	if jumlahDokter == 0 {
		return "D001"
	}
	lastID := daftarDokter[jumlahDokter-1].id
	num, _ := strconv.Atoi(lastID[1:])
	return fmt.Sprintf("D%03d", num+1)
}

// Fungsi untuk menambah data dokter baru
func tambahDokter() {
	if jumlahDokter >= maxData {
		fmt.Println("Kapasitas dokter penuh!")
		return
	}

	var dokter Dokter
	// Membuat ID secara otomatis
	dokter.id = generateDokterID()
	
	fmt.Println("\nTambah Dokter Baru")
	fmt.Println("ID Dokter:", dokter.id) // Menampilkan ID yang dibuat
	fmt.Print("Nama Dokter: ")
	fmt.Scanln(&dokter.nama)
	fmt.Print("Spesialisasi: ")
	fmt.Scanln(&dokter.spesialisasi)
	fmt.Print("Gender (L/P): ")
	fmt.Scanln(&dokter.gender)
	fmt.Print("Jadwal Praktek: ")
	fmt.Scanln(&dokter.jadwal)

	daftarDokter[jumlahDokter] = dokter
	jumlahDokter++
	fmt.Println("Dokter berhasil ditambahkan!")
}

// Fungsi untuk mengedit data dokter
func editDokter() {
	fmt.Print("Masukkan ID Dokter yang akan diedit: ")
	var id string
	fmt.Scanln(&id)

	found := false
	for i := 0; i < jumlahDokter && !found; i++ {
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
			found = true
		}
	}

	if !found {
		fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
	}
}

// Fungsi untuk menghapus data dokter
func hapusDokter() {
	fmt.Print("Masukkan ID Dokter yang akan dihapus: ")
	var id string
	fmt.Scanln(&id)

	found := false
	for i := 0; i < jumlahDokter && !found; i++ {
		if daftarDokter[i].id == id {
			// Menghapus riwayat terkait dengan dokter
			for j := 0; j < jumlahRiwayat; {
				if daftarRiwayat[j].dokterID == id {
					for k := j; k < jumlahRiwayat-1; k++ {
						daftarRiwayat[k] = daftarRiwayat[k+1]
					}
					jumlahRiwayat--
				} else {
					j++
				}
			}

			// Menghapus data dokter
			for j := i; j < jumlahDokter-1; j++ {
				daftarDokter[j] = daftarDokter[j+1]
			}
			jumlahDokter--

			fmt.Println("Dokter dan riwayat terkait berhasil dihapus!")
			found = true
		}
	}

	if !found {
		fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
	}
}

// Fungsi untuk mengurutkan dokter berdasarkan nama menggunakan Selection Sort
func urutkanDokterByNama() {
	// Pengurutan naik (ascending)
	for i := 0; i < jumlahDokter-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlahDokter; j++ {
			if daftarDokter[j].nama < daftarDokter[minIndex].nama {
				minIndex = j
			}
		}
		if minIndex != i {
			daftarDokter[i], daftarDokter[minIndex] = daftarDokter[minIndex], daftarDokter[i]
		}
	}
}

// Fungsi untuk mencari dokter berdasarkan ID menggunakan Binary Search
func cariDokterByIdBinary(id string) {
	// Pengurutan menggunakan Bubble Sort
	for i := 0; i < jumlahDokter-1; i++ {
		for j := 0; j < jumlahDokter-i-1; j++ {
			if daftarDokter[j].id > daftarDokter[j+1].id {
				daftarDokter[j], daftarDokter[j+1] = daftarDokter[j+1], daftarDokter[j]
			}
		}
	}

	// Pencarian menggunakan Binary Search
	low := 0
	high := jumlahDokter - 1
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
		fmt.Println("\nData Dokter Ditemukan (Pencarian Biner):")
		fmt.Println("=============================================================")
		fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", "ID", "Nama", "Spesialisasi", "Gender", "Jadwal")
		fmt.Println("=============================================================")
		fmt.Printf("%-6s %-20s %-15s %-10s %-20s\n", d.id, d.nama, d.spesialisasi, d.gender, d.jadwal)
	} else {
		fmt.Println("Dokter dengan ID", id, "tidak ditemukan.")
	}
}
// Manajemen Riwayat
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

// Fungsi untuk menampilkan semua data riwayat
func tampilkanRiwayat() {
	fmt.Println("\nDaftar Riwayat:")
	fmt.Println("====================================================================")
	fmt.Printf("%-6s %-10s %-10s %-12s %-30s\n", "ID", "Pasien ID", "Dokter ID", "Tanggal", "Diagnosa")
	fmt.Println("====================================================================")
	for i := 0; i < jumlahRiwayat; i++ {
		r := daftarRiwayat[i]
		fmt.Printf("%-6s %-10s %-10s %-12s %-30s\n", r.id, r.pasienID, r.dokterID, r.tanggal, r.diagnosa)
	}
}

// Fungsi untuk menghasilkan ID riwayat secara otomatis
func generateRiwayatID() string {
	if jumlahRiwayat == 0 {
		return "R001"
	}
	lastID := daftarRiwayat[jumlahRiwayat-1].id
	num, _ := strconv.Atoi(lastID[1:])
	return fmt.Sprintf("R%03d", num+1)
}

// Fungsi untuk menambah data riwayat baru
func tambahRiwayat() {
	if jumlahRiwayat >= maxData {
		fmt.Println("Kapasitas riwayat penuh!")
		return
	}

	var riwayat Riwayat
	riwayat.id = generateRiwayatID() // Menghasilkan ID secara otomatis
	
	fmt.Println("\nTambah Riwayat Baru")
	fmt.Println("ID Riwayat:", riwayat.id) // Menampilkan ID yang dihasilkan
	
	// Menampilkan daftar pasien yang tersedia dan mendapatkan ID pasien
	fmt.Println("\nDaftar Pasien Tersedia:")
	tampilkanPasien()
	fmt.Print("\nID Pasien: ")
	fmt.Scanln(&riwayat.pasienID)
	
	// Memvalidasi keberadaan pasien
	foundPasien := false
	for i := 0; i < jumlahPasien; i++ {
		if daftarPasien[i].id == riwayat.pasienID {
			foundPasien = true
			break
		}
	}
	if !foundPasien {
		fmt.Println("Pasien dengan ID", riwayat.pasienID, "tidak ditemukan.")
		return
	}
	
	// Menampilkan daftar dokter yang tersedia dan mendapatkan ID dokter
	fmt.Println("\nDaftar Dokter Tersedia:")
	tampilkanDokter()
	fmt.Print("\nID Dokter: ")
	fmt.Scanln(&riwayat.dokterID)
	
	// Memvalidasi keberadaan dokter
	foundDokter := false
	for i := 0; i < jumlahDokter; i++ {
		if daftarDokter[i].id == riwayat.dokterID {
			foundDokter = true
			break
		}
	}
	if !foundDokter {
		fmt.Println("Dokter dengan ID", riwayat.dokterID, "tidak ditemukan.")
		return
	}
	
	// Mendapatkan detail lainnya
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scanln(&riwayat.tanggal)
	fmt.Print("Diagnosa: ")
	fmt.Scanln(&riwayat.diagnosa)

	daftarRiwayat[jumlahRiwayat] = riwayat
	jumlahRiwayat++
	fmt.Println("Riwayat berhasil ditambahkan!")
}

// Fungsi untuk mengedit data riwayat
func editRiwayat() {
	fmt.Print("Masukkan ID Riwayat yang akan diedit: ")
	var id string
	fmt.Scanln(&id)

	found := false
	for i := 0; i < jumlahRiwayat && !found; i++ {
		if daftarRiwayat[i].id == id {
			fmt.Println("\nData Riwayat Ditemukan:")
			fmt.Printf("%-6s %-10s %-10s %-12s %-30s\n", "ID", "Pasien ID", "Dokter ID", "Tanggal", "Diagnosa")
			fmt.Printf("%-6s %-10s %-10s %-12s %-30s\n", daftarRiwayat[i].id, daftarRiwayat[i].pasienID, daftarRiwayat[i].dokterID, daftarRiwayat[i].tanggal, daftarRiwayat[i].diagnosa)

			fmt.Print("\nMasukkan data baru (ID Riwayat tidak dapat diubah):\n")
			
			// Menampilkan daftar pasien yang tersedia dan mendapatkan ID pasien
			fmt.Println("\nDaftar Pasien Tersedia:")
			tampilkanPasien()
			fmt.Print("\nID Pasien: ")
			fmt.Scanln(&daftarRiwayat[i].pasienID)
			
			// Memvalidasi keberadaan pasien
			foundPasien := false
			for j := 0; j < jumlahPasien; j++ {
				if daftarPasien[j].id == daftarRiwayat[i].pasienID {
					foundPasien = true
					break
				}
			}
			if !foundPasien {
				fmt.Println("Pasien dengan ID", daftarRiwayat[i].pasienID, "tidak ditemukan.")
				return
			}
			
			// Menampilkan daftar dokter yang tersedia dan mendapatkan ID dokter
			fmt.Println("\nDaftar Dokter Tersedia:")
			tampilkanDokter()
			fmt.Print("\nID Dokter: ")
			fmt.Scanln(&daftarRiwayat[i].dokterID)
			
			// Memvalidasi keberadaan dokter
			foundDokter := false
			for j := 0; j < jumlahDokter; j++ {
				if daftarDokter[j].id == daftarRiwayat[i].dokterID {
					foundDokter = true
					break
				}
			}
			if !foundDokter {
				fmt.Println("Dokter dengan ID", daftarRiwayat[i].dokterID, "tidak ditemukan.")
				return
			}
			
			// Mendapatkan detail lainnya
			fmt.Print("Tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&daftarRiwayat[i].tanggal)
			fmt.Print("Diagnosa: ")
			fmt.Scanln(&daftarRiwayat[i].diagnosa)

			fmt.Println("Data riwayat berhasil diperbarui!")
			found = true
		}
	}

	if !found {
		fmt.Println("Riwayat dengan ID", id, "tidak ditemukan.")
	}
}

// Fungsi untuk menghapus data riwayat
func hapusRiwayat() {
	fmt.Print("Masukkan ID Riwayat yang akan dihapus: ")
	var id string
	fmt.Scanln(&id)

	found := false
	for i := 0; i < jumlahRiwayat && !found; i++ {
		if daftarRiwayat[i].id == id {
			for j := i; j < jumlahRiwayat-1; j++ {
				daftarRiwayat[j] = daftarRiwayat[j+1]
			}
			jumlahRiwayat--

			fmt.Println("Riwayat berhasil dihapus!")
			found = true
		}
	}

	if !found {
		fmt.Println("Riwayat dengan ID", id, "tidak ditemukan.")
	}
}
