package main

import (
	"fmt"
)

const nmax int = 1024 // Maksimun jumlah data pakaian

type DaftarPakaian struct {
	Id, Formalitas, Kategori, list int // Atribut numerik pakaian
	Nama, Warna                    string // Atribut string pakaian
	Aktif                          bool // Flag status untuk data
}

//TabPakaian adalah array dari DaftarPakaian untuk menampung seluruh data pakaian
type TabPakaian [nmax]DaftarPakaian

//	type Riwayat struct {
//		Id   string
//		Used int
//	}
//
// type TabRiwayat Riwayat

//NextId menyimpan nilai ID pakaian berikutnya
var NextId int

// Welcome menampilkan header tampilan berdasarkan tipe
func Welcome(tipe string) {
	//Tampilan Pertama Ketika Memulai Program
	switch tipe {
	case "Start":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", " ")
		fmt.Printf("\n| %-7s%-33s |", " ", "Selamat Datang Di Aplikasi")
		fmt.Printf("\n| %s |", "AI Stylist dan Manajemen Pakaian Digital")
		fmt.Printf("\n| %-40s |", " ")
		fmt.Printf("\n+------------------------------------------+")
	case "DaftarPakaian":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", " ")
		fmt.Printf("\n| %-11s%-29s |", " ", "Manajemen Pakaian")
		fmt.Printf("\n| %-40s |", " ")
		fmt.Printf("\n+------------------------------------------+")
	}
}

// Menu menampilkan menu utama atau submenu berdasarkan tipe
func Menu(tipe string) { 
	fmt.Printf("\n\n+------------------------------------------+")
	fmt.Printf("\n| %-18s%-22s |", " ", "MENU")
	fmt.Printf("\n+------------------------------------------+")
	switch tipe {
	case "Welcome": //Tampilan opsi menu awal saat program dijalankan
		fmt.Printf("\n| %-40s |", "[1] Manajemen Pakaian")
		fmt.Printf("\n| %-40s |", "[2] Rekomendasi AI")
		fmt.Printf("\n| %-40s |", "[0] Exit")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/0)?")
	case "ManajemenPakaian": // tampilan opsi setelah memilih Manajemen Pakaian
		fmt.Printf("\n| %-40s |", "[1] Daftar Pakaian")
		fmt.Printf("\n| %-40s |", "[2] Tambah Pakaian")
		fmt.Printf("\n| %-40s |", "[3] Edit Pakaian")
		fmt.Printf("\n| %-40s |", "[4] Hapus Pakaian")
		fmt.Printf("\n| %-40s |", "[5] Cari Pakaian")
		fmt.Printf("\n| %-40s |", "[6] Sortir Pakaian")
		fmt.Printf("\n| %-40s |", "[0] Home")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Search": //tampilan saat memilih Cari Pakaian
		fmt.Printf("\n| %-40s |", "[1] Search Berdasarkan ID")
		fmt.Printf("\n| %-40s |", "[2] Search Berdasarkan Nama")
		fmt.Printf("\n| %-40s |", "[3] Search Berdasarkan Warna")
		fmt.Printf("\n| %-40s |", "[4] Search Berdasarkan Kategori")
		fmt.Printf("\n| %-40s |", "[5] Search Berdasarkan Formalitas")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Sort": //tampilan saat memilih sotir Pakaian
		fmt.Printf("\n| %-40s |", "[1] Sortir Berdasarkan ID")
		fmt.Printf("\n| %-40s |", "[2] Sortir Berdasarkan Nama")
		fmt.Printf("\n| %-40s |", "[3] Sortir Berdasarkan Warna")
		fmt.Printf("\n| %-40s |", "[4] Sortir Berdasarkan Kategori")
		fmt.Printf("\n| %-40s |", "[5] Sortir Berdasarkan Formalitas")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Ascending": // Tampilan menu sorting ascending/descending
		fmt.Printf("\n| %-40s |", "[1] Ascending")
		fmt.Printf("\n| %-40s |", "[2] Descending")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/0)?")
	}
}

// Header menampilkan tabel data palaian
func Header() { 
	//digunakan sebagai bagian atas yabel daftar pakaian
	fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	fmt.Printf("\n| %-6s | %-30s | %-15s | %-25s | %-12s |", "Id", "Nama Pakaian", "Warna", "Kategori", "Formalitas")
	fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
}

// MenuTabPakaian menampulkan daftar pakaian (jika ada)
func MenuTabPakaian(Pakaian TabPakaian, n int) {
	//Tampilan Menu Daftar Pakaian
	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-43s%-14s%-43s |", " ", "Daftar Pakaian", " ")
	if n == -1 { //Jika array blm diisi, maka akan mengeluarkan informasi bahwa data kosong
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-40s%-20s%-40s |", " ", "Data Pakaian Kosong!", " ")
		fmt.Printf("\n| %-36s%-28s%-36s |", " ", "Silahkan Tambah Data Pakaian", " ")
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	} else { //Jika array sudah diisi, maka akan menulis semua data array
		Header()
		for i := 0; i <= n; i++ { //Looping untuk menulis semua data array
			WriteData(Pakaian, i) // Menampilkan data pakaian ke i
		}
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	}
}

func Arlert(Tipe string) {//tampilan jika memasukan data yang tidak ada
	switch Tipe {
	case "ErrorInput":
		fmt.Printf("\n ______   _____    _____     ____    _____  ")
		fmt.Printf("\n|  ____| |  __ \\  |  __ \\   / __ \\  |  __ \\ ")
		fmt.Printf("\n| |__    | |__) | | |__) | | |  | | | |__) |")
		fmt.Printf("\n|  __|   |  _  /  |  _  /  | |  | | |  _  / ")
		fmt.Printf("\n| |____  | | \\ \\  | | \\ \\  | |__| | | | \\ \\ ")
		fmt.Printf("\n|______| |_|  \\_\\ |_|  \\_\\  \\____/  |_|  \\_\\")
		fmt.Printf("\n%-9s%s", " ", "Masukan Input Yang Sesuai!")
	case "DataTidakAda1":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-10s%-30s |", " ", "Data Tidak Ditemukan")
		fmt.Printf("\n+------------------------------------------+")
	case "DataTidakAda2":
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-40s%-60s |", " ", "Data Tidak Ditemukan")
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	}
}

func AturanInput(Tipe string) { // tampilan menu untuk menambahkan data
	switch Tipe {
	case "Nama": //tampilan menu Nama saat akan menambah data
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-14s%-26s |", " ", "NAMA PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Nama Pakaian")
		fmt.Printf("\n| %-40s |", "#2 Gunakan \"_\" sebagai pengganti spasi")
		fmt.Printf("\n| %-40s |", "EX. Kemeja_Teknik")
		fmt.Printf("\n+------------------------------------------+")
	case "Warna": //tampilan menu warna saat akan menambah data
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-13s%-27s |", " ", "WARNA PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Warna Pakaian")
		fmt.Printf("\n| %-40s |", "EX. Merah")
		fmt.Printf("\n+------------------------------------------+")
	case "Kategori": //tampilan menu kategori saat akan menambah data
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-12s%-28s |", " ", "KATEGORI PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Kategori Pakaian")
		fmt.Printf("\n| %-40s |", "#2 Masukan berupa angka")
		fmt.Printf("\n| %-40s |", "EX. 4")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "01 : Kemeja Lengan Panjang")
		fmt.Printf("\n| %-40s |", "02 : Kemeja Lengan Pendek")
		fmt.Printf("\n| %-40s |", "03 : Kaos Lengan Panjang")
		fmt.Printf("\n| %-40s |", "04 : Kaos Lengan Pendek")
		fmt.Printf("\n| %-40s |", "05 : Celana Panjang")
		fmt.Printf("\n| %-40s |", "06 : Celana Pendek")
		fmt.Printf("\n| %-40s |", "07 : Jaket")
		fmt.Printf("\n| %-40s |", "08 : Luaran")
		fmt.Printf("\n| %-40s |", "09 : Sandal")
		fmt.Printf("\n| %-40s |", "10 : Sepatu")
		fmt.Printf("\n+------------------------------------------+")
	case "Formalitas": //tampilan menu formalitas saat akan menambah data
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-11s%-29s |", " ", "TINGKAT FORMALITAS")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Tingkat Formalitas")
		fmt.Printf("\n| %-40s |", "#1 Masukan berupa angka")
		fmt.Printf("\n| %-40s |", "EX. 2")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "1 : Santai")
		fmt.Printf("\n| %-40s |", "2 : Semi Formal")
		fmt.Printf("\n| %-40s |", "3 : Formal")
		fmt.Printf("\n+------------------------------------------+")
	}
}

//WriteData menampilkan satu baris data pakaian jika statusnya aktif
func WriteData(Pakaian TabPakaian, i int) {
	if Pakaian[i].Aktif { // buat cek data aktif atau kagak
		fmt.Printf("\n| %-6d ", Pakaian[i].Id) // Tampilkan ID
		fmt.Printf("| %-30s ", Pakaian[i].Nama) // Tampilkan Nama
		fmt.Printf("| %-15s ", Pakaian[i].Warna) // Tampilkan Warna

		// Menampilkan kategori berdasarkan kode angka
		switch Pakaian[i].Kategori { 
		case 1: 
			fmt.Printf("| %-25s ", "Kemeja Lengan Panjang")
		case 2:
			fmt.Printf("| %-25s ", "Kemeja Lengan Pendek")
		case 3:
			fmt.Printf("| %-25s ", "Kaos Lengan Panjang")
		case 4:
			fmt.Printf("| %-25s ", "Kaos Lengan Pendek")
		case 5:
			fmt.Printf("| %-25s ", "Celana Panjang")
		case 6:
			fmt.Printf("| %-25s ", "Celana Pendek")
		case 7:
			fmt.Printf("| %-25s ", "Jaket")
		case 8:
			fmt.Printf("| %-25s ", "Luaran")
		case 9:
			fmt.Printf("| %-25s ", "Sendal")
		case 10:
			fmt.Printf("| %-25s ", "Sepatu")
		}

		// Menampilkan tingkat formalitas berdasarkan kode angka
		switch Pakaian[i].Formalitas { 
		case 1:
			fmt.Printf("| %-12s |", "Santai")
		case 2:
			fmt.Printf("| %-12s |", "Semi Formal")
		case 3:
			fmt.Printf("| %-12s |", "Formal")
		}
	}
}


// add menambahkan satu data pakaian baru kedalam array
func add(Pakaian *TabPakaian, n int) {
	//Fitur Tambah Pakaian
	/*
		Menambahkan data pakaian kedalam array dengan ketentuan yang sudah diatur
		Generate Pakaian.Id secara otomatis
	*/
	AturanInput("Nama")
	fmt.Printf("\n| %-10s%-2s ", "Nama", ":")
	fmt.Scan(&Pakaian[n].Nama)

	AturanInput("Warna")
	fmt.Printf("\n| %-10s%-2s ", "Warna", ":")
	fmt.Scan(&Pakaian[n].Warna)

	AturanInput("Kategori")
	fmt.Printf("\n| %-10s%-2s ", "Kategori", ":")
	fmt.Scan(&Pakaian[n].Kategori)

	AturanInput("Formalitas")
	fmt.Printf("\n| %-10s%-2s ", "Formalitas", ":")
	fmt.Scan(&Pakaian[n].Formalitas)

	// Set ID dan aktifkan data
	Pakaian[n].Id = NextId + 1
	Pakaian[n].Aktif = true
	NextId++ //Update NextID untuk ID berikutnya
}

// edit memperbarui data pakaian berdasarkan ID
func edit(Pakaian *TabPakaian, n int, KeyInt int) {
	//Fitur edit
	/*
		Berupa data aktif benilai true, masukan berupa data string nama, warna, dan int kategori, formalitas
		Update Data dari data yang sudah dimasukan

	 */

	var ganti int = BinarySearch(*Pakaian, n, KeyInt) // Cari index data berdasarkan ID
	var TempEdit, backup DaftarPakaian // Data sementara serta backup sebelum diubah
	var konfirmasi bool // Untuk menyimpan konfirmasi user

	if ganti > -1 {
		TempEdit.Id = (*Pakaian)[ganti].Id
		TempEdit.Aktif = true // Set data aktif

		// Tampilkan form edit data
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-15s%-25s |", " ", "EDIT DATA")
		fmt.Printf("\n| %-15s%-3s%6d%15s |", " ", "ID: ", Pakaian[ganti].Id, " ")
		fmt.Printf("\n+------------------------------------------+")

		 // Edit nama
		AturanInput("Nama")
		fmt.Printf("\nData Sebelumnya\t: %s", Pakaian[ganti].Nama)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&TempEdit.Nama)

		 // Edit warna
		AturanInput("Warna")
		fmt.Printf("\nData Sebelumnya\t: %s", Pakaian[ganti].Warna)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&TempEdit.Warna)

		// Edit kategori
		AturanInput("Kategori")
		fmt.Printf("\nData Sebelumnya\t: %d", Pakaian[ganti].Kategori)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&TempEdit.Kategori)

		// Edit formalitas
		AturanInput("Formalitas")
		fmt.Printf("\nData Sebelumnya\t: %d", Pakaian[ganti].Formalitas)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&TempEdit.Formalitas)

		// Menampilkan perbandingan sebelum dan sesudah
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-43s%-14s%-43s |", " ", "Perubahan Data", " ")
		Header()
		WriteData(*Pakaian, ganti) // data Sebelum
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		
		backup = (*Pakaian)[ganti]			// Simpan data lama
		(*Pakaian)[ganti] = TempEdit		// Ganti dengan data baru

		// Visualisasi perubahan data
		fmt.Printf("\n%-51s%-53s", " ", "||")
		fmt.Printf("\n%-50s%-54s", " ", "\\||/")
		fmt.Printf("\n%-51s%-53s", " ", "\\/")

		Header()
		WriteData(*Pakaian, ganti) // data Sesudah

		// form Konfirmasi akhir
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-10s%-30s |", " ", "Konfirmasi Perubahan")
		fmt.Printf("\n| %-40s |", "[1] Accept")
		fmt.Printf("\n| %-40s |", "[0] Reject")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n Pilih (1/0)?")
		fmt.Scan(&konfirmasi)
		if !konfirmasi {
			(*Pakaian)[ganti] = backup // Batalkan perubahan jika ditolak
		}
	} else {
		Arlert("DataTidakAda1") // Tampilkan pesan jika ID tidak ditemukan
	}
}

// min mencari indeks dengan nilai Id minimum dari posisi i sampai n
func min(Pakaian TabPakaian, n, i int) int {
	//Mencari nilai minimum dalam rentang tertentu
	var min = i
	for j := i + 1; j <= n; j++ {
		if Pakaian[min].Id > Pakaian[j].Id {
			min = j // Perbarui indeks minimum jika ditemukan nilai lebih kecil
		}
	}
	return min // kembalikan indeks minimum
}

// max mencari indeks dengan nilai Id maksimum dari posisi i sampai n
func max(Pakaian TabPakaian, n, i int) int {
	//Mencari nilai maksimum dalam rentang tertentu
	var max = i
	for j := i + 1; j <= n; j++ {
		if Pakaian[max].Id < Pakaian[j].Id {
			max = j // Perbarui indeks maksimum jika ditemukan nilai lebih besar
		}
	}
	return max // Kembalikan indeks maksimum
}
 
// SortById mengurutkan data pakaian berdasarkan Id
func SortById(Pakaian *TabPakaian, n int, Ascending bool) {
	//Mengurutkan Index array berdasarkan Id
	var temp DaftarPakaian
	var Idx int
	for i := 0; i <= n; i++ {
		if Ascending {
			Idx = min(*Pakaian, n, i)  // Cari indeks terkecil
		} else {
			Idx = max(*Pakaian, n, i) // Cari indeks terbesar
		}
		temp = (*Pakaian)[i]				// Simpan sementara
		(*Pakaian)[i] = (*Pakaian)[Idx]		// Tukar posisi
		(*Pakaian)[Idx] = temp				// Selesaikan pertukaran
	}
}

func BinarySearch(Pakaian TabPakaian, n int, KeyInt int) int {
	//Pencarian Id menggunakan binary search
	var left, mid, right int
	left = 0
	right = n
	for left <= right {
		mid = (left + right) / 2
		if Pakaian[mid].Id == KeyInt {
			return mid
		} else if Pakaian[mid].Id < KeyInt {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func Mengandung(CekString, KeyString string) bool {
	var j int = 0
	for i := 0; i <= len(CekString)-len(KeyString); i++ {
		for j < len(KeyString) && CekString[i+j] == KeyString[j] {
			j++
		}
		if j == len(KeyString) {
			return true
		}
	}
	return false
}

func SquentialSearchString(Pakaian TabPakaian, n int, KeyString, Mode string) {
	var found bool
	for i := 0; i < n; i++ {
		switch Mode {
		case "Nama":
			if Mengandung(Pakaian[i].Nama, KeyString) {
				found = true
			}
		case "Warna":
			if Mengandung(Pakaian[i].Warna, KeyString) {
				found = true
			}
		}
		if found {
			if i == 0 {
				Header()
			}
			WriteData(Pakaian, i)
		}
	}
	if !found {
		Arlert("DataTidakAda2")
	} else {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	}
}

func SquentialSearcInt(Pakaian TabPakaian, n int, KeyInt int, Mode string) {
	var found bool
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
			fmt.Printf("\n| %-43s%-57s |", " ", "Hasil Pencarian")
		}
		switch Mode {
		case "Kategori":
			if Pakaian[i].Kategori == KeyInt {
				found = true
			}
		case "Formalitas":
			if Pakaian[i].Formalitas == KeyInt {
				found = true
			}
		}
		if found {
			if i == 0 {
				Header()
			}
			WriteData(Pakaian, i)
		}
	}
	if !found {
		Arlert("DataTidakAda2")
	} else {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	}
}

func SoftDelete(Pakaian *TabPakaian, n int, id int) {
	//Memberikan status false pada array yang berarti data telah di nonaktifkan / dihapus
	idx := BinarySearch(*Pakaian, n, id)
	if idx != -1 {
		if (*Pakaian)[idx].Aktif {
			(*Pakaian)[idx].Aktif = false
			fmt.Println("Data berhasil di nonaktifkan")
		} else {
			fmt.Println("Data sudah tidak aktif")
		}
	} else {
		Arlert("DataTidakAda1")
	}
}

func SortByNama(Pakaian *TabPakaian, n int, Ascending bool) {
	var temp DaftarPakaian
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if Ascending {
				if (*Pakaian)[j].Nama > (*Pakaian)[j+1].Nama {
					temp = (*Pakaian)[j]
					(*Pakaian)[j] = (*Pakaian)[j+1]
					(*Pakaian)[j+1] = temp
				}
			} else {
				if (*Pakaian)[j].Nama < (*Pakaian)[j+1].Nama {
					temp = (*Pakaian)[j]
					(*Pakaian)[j] = (*Pakaian)[j+1]
					(*Pakaian)[j+1] = temp
				}
			}
		}
	}
	fmt.Println("Data berhasil diurutkan terkecil ke besar")
}

func main() {
	var Pakaian TabPakaian
	var nPakaian int = -1
	var KeyInt int
	var pilih, KeyString string
	for valid := false; !valid; {
		Welcome("Start")
		Menu("Welcome")
		fmt.Scan(&pilih)
		switch pilih {
		case "1": //Manajemen Pakaian
			for valid1 := false; !valid1; {
				// Welcome("DaftarPakaian")
				Menu("ManajemenPakaian")
				fmt.Scan(&pilih)
				switch pilih {
				case "1": //Daftar Pakaian
					MenuTabPakaian(Pakaian, nPakaian)
				case "2": //Tambah Pakaian
					nPakaian++
					add(&Pakaian, nPakaian)
				case "3": //Edit Pakaian
					fmt.Printf("Masukan Id data yang ingin di edit: ")
					fmt.Scan(&KeyInt)
					SortById(&Pakaian, nPakaian, true) //Melakukan selection sort by Id karena akan menggunakan binary search untuk mencari Id
					edit(&Pakaian, nPakaian, KeyInt)
				case "4": //Hapus Pakaian
					fmt.Printf("Masukan Id data yang ingin di hapus(soft delete): ")
					fmt.Scan(&KeyInt)
					SortById(&Pakaian, nPakaian, true) 
					idx := BinarySearch(Pakaian, nPakaian, KeyInt)
					if idx > -1 && Pakaian[idx].Aktif {
						Pakaian[idx].Aktif = false 
						fmt.Println("Data berhasil di delete (soft delete)")
					} else {
						fmt.Println("Data tidak ditemukan / sudah tidak aktif") // sejauh ini sudah bisa didelete, gua push dulu ada acara gua mya >-<
					}
				case "5": //Cari Pakaian
					Menu("Search")
					fmt.Scan(&pilih)
					switch pilih {
					case "1": //
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyInt)
						SortById(&Pakaian, nPakaian, true)
						if BinarySearch(Pakaian, nPakaian, KeyInt) > -1 {
							Header()
							WriteData(Pakaian, BinarySearch(Pakaian, nPakaian, KeyInt))
							fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
						} else {
							Arlert("DataTidakAda1")
						}
					case "2":
						AturanInput("Nama")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyString)
						SquentialSearchString(Pakaian, nPakaian, KeyString, "Nama")
					case "3":
						AturanInput("Warna")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyString)
						SquentialSearchString(Pakaian, nPakaian, KeyString, "Warna")
					case "4":
						AturanInput("Kategori")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyInt)
						SquentialSearcInt(Pakaian, nPakaian, KeyInt, "Kategori")
					case "5":
						AturanInput("Formalitas")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyInt)
						SquentialSearcInt(Pakaian, nPakaian, KeyInt, "Formalitas")
					case "0":

					}
				case "6":
					Menu("Sort")
					fmt.Scan(&pilih)
					switch pilih {
					case "1":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortById(&Pakaian, nPakaian, true)
						case "2":
							SortById(&Pakaian, nPakaian, false)
						case "0":
						}
					case "2":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByNama(&Pakaian, nPakaian, true)
						case "2":
							SortByNama(&Pakaian, nPakaian, false)
						case "0":
						}
					case "3":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							// SortByNama(&Pakaian, nPakaian, true)
						case "2":
							// SortByNama(&Pakaian, nPakaian, false)
						case "0":
						}
					case "4":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							// SortByNama(&Pakaian, nPakaian, true)
						case "2":
							// SortByNama(&Pakaian, nPakaian, false)
						case "0":
						}
					case "5":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							// SortByNama(&Pakaian, nPakaian, true)
						case "2":
							// SortByNama(&Pakaian, nPakaian, false)
						case "0":
						}
					case "0":
					}
					// case "5":
				// 	sort()
				case "0":
					valid1 = true
				default:
					Arlert("ErrorInput")
				}
			}
		// case "2":
		case "0":
			valid = true
		default:
			Arlert("ErrorInput")
		}
	}
}
