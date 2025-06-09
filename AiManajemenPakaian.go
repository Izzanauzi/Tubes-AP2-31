/*
Nama Anggota :
 1. Muhammad Izzan Auzi mulyana Putra (103012400051)
 2. Ghifar Anshari Rabbani (103012400003)

Nomor Topik : 31
Judul Topik : Aplikasi AI Stylist dan Manajemen Pakaian Digital

Deskripsi Program :
Program ini merubakan aplikasi sederhana untuk mengelola data Pakaian dan memberikan rekomendasi pakaian berdasarkan preferensi pengguna, tujuan serta kondisi cuaca, Program ini juga dilengkapi fitur AI sederhana

Tantangan yang  dihadapi :
 1. Membuat program efisien dengan hanya menggunakan package “fmt”.
 2. Membuat UI yang ramah pengguna.
 3. Membuat setiap data memiliki sesuatu yang bersifat unik.
 4. Membuat sistem dapat memberikan rekomendasi pakaian.

Masalah yang dihadapi: Ada beberapa hal yang terbatas karena tidak boleh digunakan, seperti package tambahan, serta masalah eror atau crash yang baru ke detect karena tidak kami sadari sebelumya.

Fitur Utama :

 1. Manajemen Pakaian :
    -Menampilkan daftar data Pakaian
    -Menambah Data Pakaian
    -Mengedit Data pakaian
    -Menghapus Data Pakaian
    -Mencari Data Pakaian berdasarkan ID, Nama, Warna, Kategori, Formalitas
    -Mengsortir Pakaian Berdasarkan ID, Nama, Warna, Kategori, Formalitas

 2. Rekomendasi AI
    -Penggguna memilih tujuan serta kondisi cuaca
    -Sistem Akan Memberi saran set pakaian berdasarkan data pengguna
*/
package main

import (
	"fmt"
)

const NMAX int = 1024 // Konstanta untuk jumlah maksimal array

type DaftarPakaian struct { // Struct untuk menyimpan data pakaian
	id, formalitas, kategori int
	nama, warna              string
	aktif                    bool
}
type TabPakaian [NMAX]DaftarPakaian // Tipe data array untuk menyimpan daftar pakaian

var listKategori = [11]string{ // Daftar kategori pakaian
	"",
	"Kemeja Lengan Panjang",
	"Kemeja Lengan Pendek",
	"Kaos Lengan Panjang",
	"Kaos Lengan Pendek",
	"Celana Panjang",
	"Celana Pendek",
	"Jaket",
	"Luaran",
	"Sendal",
	"Sepatu",
}
var listFormalitas = [4]string{ // Daftar tingkat formalitas pakaian
	"",
	"Santai",
	"Semi Formal",
	"Formal",
}

type Riwayat struct { // Struct untuk menyimpan riwayat rekomendasi pakaian
	day  int
	used [4]int
}
type TabRiwayat [NMAX]Riwayat // Tipe data array untuk menyimpan riwayat rekomendasi pakaian
var nextId int                // Variabel untuk menyimpan ID berikutnya yang akan digunakan
var usingAI int               // Variabel untuk menyimpan indeks terakhir yang digunakan dalam rekomendasi AI
type NeedAI struct {          // Struct untuk menyimpan data yang direkomendasikan AI
	id    int
	count int
}
type TabAI [NMAX]NeedAI // Tipe data array untuk menyimpan data yang direkomendasikan AI

func Welcome(tipe string) { // Fungsi untuk menampilkan pesan di awal program
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
func Menu(tipe string) { // Fungsi untuk menampilkan menu untuk pilih opsi yang tersedia
	fmt.Printf("\n\n+------------------------------------------+")
	fmt.Printf("\n| %-18s%-22s |", " ", "MENU")
	fmt.Printf("\n+------------------------------------------+")
	switch tipe {
	case "Welcome":
		fmt.Printf("\n| %-40s |", "[1] Manajemen Pakaian")
		fmt.Printf("\n| %-40s |", "[2] Rekomendasi AI")
		fmt.Printf("\n| %-40s |", "[0] Exit")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/0)?")
	case "ManajemenPakaian":
		fmt.Printf("\n| %-40s |", "[1] Daftar Pakaian")
		fmt.Printf("\n| %-40s |", "[2] Tambah Pakaian")
		fmt.Printf("\n| %-40s |", "[3] Edit Pakaian")
		fmt.Printf("\n| %-40s |", "[4] Hapus Pakaian")
		fmt.Printf("\n| %-40s |", "[5] Cari Pakaian")
		fmt.Printf("\n| %-40s |", "[6] Sortir Pakaian")
		fmt.Printf("\n| %-40s |", "[0] Home")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Search":
		fmt.Printf("\n| %-40s |", "[1] Search Berdasarkan ID")
		fmt.Printf("\n| %-40s |", "[2] Search Berdasarkan Nama")
		fmt.Printf("\n| %-40s |", "[3] Search Berdasarkan Warna")
		fmt.Printf("\n| %-40s |", "[4] Search Berdasarkan Kategori")
		fmt.Printf("\n| %-40s |", "[5] Search Berdasarkan Formalitas")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Sort":
		fmt.Printf("\n| %-40s |", "[1] Sortir Berdasarkan ID")
		fmt.Printf("\n| %-40s |", "[2] Sortir Berdasarkan Nama")
		fmt.Printf("\n| %-40s |", "[3] Sortir Berdasarkan Warna")
		fmt.Printf("\n| %-40s |", "[4] Sortir Berdasarkan Kategori")
		fmt.Printf("\n| %-40s |", "[5] Sortir Berdasarkan Formalitas")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Ascending":
		fmt.Printf("\n| %-40s |", "[1] Ascending")
		fmt.Printf("\n| %-40s |", "[2] Descending")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/0)?")
	case "delete":
		fmt.Printf("\n| %-40s |", "[1] Soft Delete")
		fmt.Printf("\n| %-40s |", "[2] Hard Delete")
		fmt.Printf("\n| %-40s |", "[3] Restore")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2//3/0)?")
	}
}
func Table() { // Fungsi untuk menampilkan tabel header
	fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	fmt.Printf("\n| %-6s | %-30s | %-15s | %-25s | %-12s |", "Id", "Nama Pakaian", "Warna", "Kategori", "Formalitas")
	fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
}
func MenuTabPakaian(Pakaian TabPakaian, n int) { // Fungsi untuk menampilkan daftar pakaian
	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-43s%-14s%-43s |", " ", "Daftar Pakaian", " ")
	if n == -1 {
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-40s%-20s%-40s |", " ", "Data Pakaian Kosong!", " ")
		fmt.Printf("\n| %-36s%-28s%-36s |", " ", "Silahkan Tambah Data Pakaian", " ")
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	} else {
		Table()
		for i := 0; i <= n; i++ {
			if Pakaian[i].aktif {
				WriteData(Pakaian, i)
			}
		}
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	}
}
func Arlert(Tipe string) { // Fungsi untuk menampilkan pesan peringatan atau informasi
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
func AturanInput(Tipe string) { // Fungsi untuk menampilkan aturan input berdasarkan tipe yang diberikan
	switch Tipe {
	case "ID":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-14s%-26s |", " ", "ID PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan ID Pakaian")
		fmt.Printf("\n| %-40s |", "#2 Masukan berupa angka > 0")
		fmt.Printf("\n| %-40s |", "EX. 1")
		fmt.Printf("\n+------------------------------------------+")
	case "Nama":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-14s%-26s |", " ", "NAMA PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Nama Pakaian")
		fmt.Printf("\n| %-40s |", "#2 Gunakan \"_\" sebagai pengganti spasi")
		fmt.Printf("\n| %-40s |", "EX. Kemeja_Teknik")
		fmt.Printf("\n+------------------------------------------+")
	case "Warna":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-13s%-27s |", " ", "WARNA PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Warna Pakaian")
		fmt.Printf("\n| %-40s |", "EX. Merah")
		fmt.Printf("\n+------------------------------------------+")
	case "Kategori":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-12s%-28s |", " ", "KATEGORI PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Kategori Pakaian")
		fmt.Printf("\n| %-40s |", "#2 Masukan berupa angka")
		fmt.Printf("\n| %-40s |", "EX. 4")
		fmt.Printf("\n+------------------------------------------+")
		for i := 1; i < 11; i++ {
			fmt.Printf("\n| %2d%s%-35s |", i, " : ", listKategori[i])
		}
		fmt.Printf("\n+------------------------------------------+")
	case "Formalitas":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-11s%-29s |", " ", "TINGKAT FORMALITAS")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Tingkat Formalitas")
		fmt.Printf("\n| %-40s |", "#1 Masukan berupa angka")
		fmt.Printf("\n| %-40s |", "EX. 2")
		fmt.Printf("\n+------------------------------------------+")
		for i := 1; i < 4; i++ {
			fmt.Printf("\n| %2d%s%-35s |", i, " : ", listFormalitas[i])
		}
		fmt.Printf("\n+------------------------------------------+")
	case "AITujuan":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-12s%-28s |", " ", "Tujuan")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Tujuan Anda Hari Ini")
		fmt.Printf("\n| %-40s |", "#2 Masukan berupa angka")
		fmt.Printf("\n| %-40s |", "EX. 2")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "1 : Acara Formal")
		fmt.Printf("\n| %-40s |", "2 : Week Day")
		fmt.Printf("\n| %-40s |", "3 : Week End / Holiday")
		fmt.Printf("\n+------------------------------------------+")
	case "AIWeather":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-12s%-28s |", " ", "Cuaca")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Cuaca Hari Ini")
		fmt.Printf("\n| %-40s |", "#2 Masukan berupa angka")
		fmt.Printf("\n| %-40s |", "EX. 2")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "1 : Panas")
		fmt.Printf("\n| %-40s |", "2 : Dingin")
		fmt.Printf("\n+------------------------------------------+")
	case "Edit":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-15s%-25s |", " ", "Edit Berdasarkan")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "1 : ID")
		fmt.Printf("\n| %-40s |", "2 : Nama")
		fmt.Printf("\n+------------------------------------------+")
	case "Acc":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-10s%-30s |", " ", "Apakah rekomendasi ini cocok?")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "[1] Cocok")
		fmt.Printf("\n| %-40s |", "[0] Tidak Cocok")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n Pilih (1/0)?")
	case "Delete":
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-15s%-25s |", " ", "Delete Berdasarkan")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "1 : ID")
		fmt.Printf("\n| %-40s |", "2 : Nama")
		fmt.Printf("\n+------------------------------------------+")
	}
}
func WriteData(Pakaian TabPakaian, i int) { // Fungsi untuk menampilkan data pakaian dalam format tabel
	fmt.Printf("\n| %-6d ", Pakaian[i].id)
	fmt.Printf("| %-30s ", Pakaian[i].nama)
	fmt.Printf("| %-15s ", Pakaian[i].warna)
	fmt.Printf("| %-25s ", listKategori[Pakaian[i].kategori])
	fmt.Printf("| %-12s |", listFormalitas[Pakaian[i].formalitas])
}
func AddData(Pakaian *TabPakaian, n int) { // Fungsi untuk menambahkan data pakaian baru
	AturanInput("Nama")
	fmt.Printf("\n| %-10s%-2s ", "Nama", ":")
	fmt.Scan(&Pakaian[n].nama)
	AturanInput("Warna")
	fmt.Printf("\n| %-10s%-2s ", "Warna", ":")
	fmt.Scan(&Pakaian[n].warna)
	AturanInput("Kategori")
	fmt.Printf("\n| %-10s%-2s ", "Kategori", ":")
	fmt.Scan(&Pakaian[n].kategori)
	AturanInput("Formalitas")
	fmt.Printf("\n| %-10s%-2s ", "Formalitas", ":")
	fmt.Scan(&Pakaian[n].formalitas)
	Pakaian[n].id = nextId + 1
	Pakaian[n].aktif = true
	nextId++
}
func EditData(Pakaian *TabPakaian, n int, KeyInt int) { // Fungsi untuk mengedit data pakaian berdasarkan ID
	var ganti int = BinarySearch(*Pakaian, n, KeyInt)
	var tempEdit, backup DaftarPakaian
	var konfirmasi bool
	if ganti > -1 && Pakaian[ganti].aktif {
		tempEdit.id = (*Pakaian)[ganti].id
		tempEdit.aktif = true
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-15s%-25s |", " ", "EDIT DATA")
		fmt.Printf("\n| %-15s%-3s%6d%15s |", " ", "ID: ", Pakaian[ganti].id, " ")
		fmt.Printf("\n+------------------------------------------+")
		AturanInput("Nama")
		fmt.Printf("\nData Sebelumnya\t: %s", Pakaian[ganti].nama)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&tempEdit.nama)
		AturanInput("Warna")
		fmt.Printf("\nData Sebelumnya\t: %s", Pakaian[ganti].warna)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&tempEdit.warna)
		AturanInput("Kategori")
		fmt.Printf("\nData Sebelumnya\t: %d", Pakaian[ganti].kategori)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&tempEdit.kategori)
		AturanInput("Formalitas")
		fmt.Printf("\nData Sebelumnya\t: %d", Pakaian[ganti].formalitas)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&tempEdit.formalitas)
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-43s%-14s%-43s |", " ", "Perubahan Data", " ")
		Table()
		WriteData(*Pakaian, ganti)
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		backup = (*Pakaian)[ganti]
		(*Pakaian)[ganti] = tempEdit
		fmt.Printf("\n%-51s%-53s", " ", "||")
		fmt.Printf("\n%-50s%-54s", " ", "\\||/")
		fmt.Printf("\n%-51s%-53s", " ", "\\/")
		Table()
		WriteData(*Pakaian, ganti)
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-10s%-30s |", " ", "Konfirmasi Perubahan")
		fmt.Printf("\n| %-40s |", "[1] Accept")
		fmt.Printf("\n| %-40s |", "[0] Reject")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n Pilih (1/0)?")
		fmt.Scan(&konfirmasi)
		if !konfirmasi {
			(*Pakaian)[ganti] = backup
		}
	} else {
		Arlert("DataTidakAda1")
	}
}
func Min(Pakaian TabPakaian, n, i int) int { // Fungsi untuk mencari indeks dengan ID terkecil
	var min = i
	for j := i + 1; j <= n; j++ {
		if Pakaian[min].id > Pakaian[j].id {
			min = j
		}
	}
	return min
}

func Max(Pakaian TabPakaian, n, i int) int { // Fungsi untuk mencari indeks dengan ID terbesar
	var max = i
	for j := i + 1; j <= n; j++ {
		if Pakaian[max].id < Pakaian[j].id {
			max = j
		}
	}
	return max
}

func SelectionSortId(Pakaian *TabPakaian, n int, Ascending bool) { // Fungsi untuk mengurutkan data pakaian berdasarkan ID
	var temp DaftarPakaian
	var Idx int
	for i := 0; i < n; i++ {
		if Ascending {
			Idx = Min(*Pakaian, n, i)
		} else {
			Idx = Max(*Pakaian, n, i)
		}
		temp = (*Pakaian)[i]
		(*Pakaian)[i] = (*Pakaian)[Idx]
		(*Pakaian)[Idx] = temp
	}
}

func BinarySearch(Pakaian TabPakaian, n int, KeyInt int) int { // Fungsi untuk mencari indeks data pakaian berdasarkan ID menggunakan binary search
	var left, mid, right, idx int
	idx = -1
	left = 0
	right = n
	for left <= right {
		mid = (left + right) / 2
		if Pakaian[mid].id == KeyInt {
			idx = mid
			right = -1
		} else if Pakaian[mid].id < KeyInt {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func SquentialSearchNama(Pakaian TabPakaian, n int, KeyString string) int { // Fungsi untuk mencari indeks data pakaian berdasarkan nama menggunakan sequential search
	var idx int = -1
	var found bool = false
	for i := 0; i <= n && !found; i++ {
		if Pakaian[i].nama == KeyString {
			if Pakaian[i].aktif {
				idx = i
			}
			found = true
		}
	}
	return idx
}

func SquentialSearchWarna(Pakaian TabPakaian, n int, KeyString string) { // Fungsi untuk mencari data pakaian berdasarkan warna menggunakan sequential search
	var found bool
	for i := 0; i <= n; i++ {
		if i == 0 {
			fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
			fmt.Printf("\n| %-43s%-57s |", " ", "Hasil Pencarian")
		}
		if Pakaian[i].warna == KeyString && Pakaian[i].aktif {
			if !found {
				Table()
			}
			WriteData(Pakaian, i)
			found = true
		}
	}
	if !found {
		Arlert("DataTidakAda2")
	} else {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	}
}

func SquentialSearcInt(Pakaian TabPakaian, n int, KeyInt int, Mode string) { // Fungsi untuk mencari data pakaian berdasarkan kategori atau formalitas menggunakan sequential search
	var found bool
	for i := 0; i <= n; i++ {
		if i == 0 {
			fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
			fmt.Printf("\n| %-43s%-57s |", " ", "Hasil Pencarian")
		}
		switch Mode {
		case "Kategori":
			if Pakaian[i].kategori == KeyInt && Pakaian[i].aktif {
				if !found {
					Table()
				}
				WriteData(Pakaian, i)
				found = true
			}
		case "Formalitas":
			if Pakaian[i].formalitas == KeyInt && Pakaian[i].aktif {
				if !found {
					Table()
				}
				WriteData(Pakaian, i)
				found = true
			}
		}
	}
	if !found {
		Arlert("DataTidakAda2")
	} else {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	}
}

func Delete(Pakaian *TabPakaian, n int, id int) { // Fungsi untuk menonaktifkan data pakaian berdasarkan ID
	idx := BinarySearch(*Pakaian, n, id)
	if idx > -1 {
		if (*Pakaian)[idx].aktif {
			(*Pakaian)[idx].aktif = false
			fmt.Println("Data berhasil di nonaktifkan")
		} else {
			fmt.Println("Data sudah tidak aktif")
		}
	} else {
		Arlert("DataTidakAda1")
	}
}

func Restore(Pakaian *TabPakaian, n *int, id int) { // Fungsi untuk mengaktifkan kembali data pakaian yang telah dinonaktifkan
	idx := BinarySearch(*Pakaian, *n, id)
	if idx > -1 && !Pakaian[idx].aktif {
		Pakaian[idx].aktif = true
		fmt.Printf("Data Berhasil di aktifkan")
	} else if idx > -1 && Pakaian[idx].aktif {
		fmt.Printf("Data Sudah Aktif")
	} else {
		Arlert("DataTidakAda1")
	}
}

func DeletePermanent(Pakaian *TabPakaian, n *int, id int) { // Fungsi untuk menghapus data pakaian secara permanen berdasarkan ID
	idx := BinarySearch(*Pakaian, *n, id)
	if idx > -1 && !Pakaian[idx].aktif {
		for i := idx; i < *n; i++ {
			Pakaian[i] = Pakaian[i+1]
		}
		*n--
		fmt.Printf("Data Berhasil di hapus")
	} else {
		Arlert("DataTidakAda1")
	}
}

func SortByNama(Pakaian *TabPakaian, n int, Ascending bool) { // Fungsi untuk mengurutkan data pakaian berdasarkan nama
	var temp DaftarPakaian
	var j int
	for i := 0; i < n; i++ {
		temp = (*Pakaian)[i]
		j = i - 1
		if Ascending {
			for j >= 0 && (*Pakaian)[j].nama > temp.nama {
				(*Pakaian)[j+1] = (*Pakaian)[j]
				j--
			}
		} else {
			for j >= 0 && (*Pakaian)[j].nama < temp.nama {
				(*Pakaian)[j+1] = (*Pakaian)[j]
				j--
			}
		}
		(*Pakaian)[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan terkecil ke besar")
}

func SortByWarna(Pakaian *TabPakaian, n int, Ascending bool) { // Fungsi untuk mengurutkan data pakaian berdasarkan warna
	var temp DaftarPakaian
	var j int
	for i := 0; i < n; i++ {
		temp = (*Pakaian)[i]
		j = i - 1
		if Ascending {
			for j >= 0 && (*Pakaian)[j].warna > temp.warna {
				(*Pakaian)[j+1] = (*Pakaian)[j]
				j--
			}
		} else {
			for j >= 0 && (*Pakaian)[j].warna < temp.warna {
				(*Pakaian)[j+1] = (*Pakaian)[j]
				j--
			}
		}
		(*Pakaian)[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan terkecil ke besar")
}

func SortByKategori(Pakaian *TabPakaian, n int, Ascending bool) { // Fungsi untuk mengurutkan data pakaian berdasarkan kategori
	var temp DaftarPakaian
	var j int
	for i := 0; i < n; i++ {
		temp = (*Pakaian)[i]
		j = i - 1
		if Ascending {
			for j >= 0 && (*Pakaian)[j].kategori > temp.kategori {
				(*Pakaian)[j+1] = (*Pakaian)[j]
				j--
			}
		} else {
			for j >= 0 && (*Pakaian)[j].kategori < temp.kategori {
				(*Pakaian)[j+1] = (*Pakaian)[j]
				j--
			}
		}
		(*Pakaian)[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan terkecil ke besar")
}

func SortByFormalitas(Pakaian *TabPakaian, n int, Ascending bool) { // Fungsi untuk mengurutkan data pakaian berdasarkan formalitas
	var temp DaftarPakaian
	var j int
	for i := 0; i < n; i++ {
		temp = (*Pakaian)[i]
		j = i - 1
		if Ascending {
			for j >= 0 && (*Pakaian)[j].formalitas > temp.formalitas {
				(*Pakaian)[j+1] = (*Pakaian)[j]
				j--
			}
		} else {
			for j >= 0 && (*Pakaian)[j].formalitas < temp.formalitas {
				(*Pakaian)[j+1] = (*Pakaian)[j]
				j--
			}
		}
		(*Pakaian)[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan terkecil ke besar")
}

func SortDataAI(AI *TabAI, Trend TabRiwayat, n int) { // Fungsi untuk mengurutkan data AI berdasarkan jumlah kemunculan
	var most int
	var temp NeedAI
	for i := 0; i < n; i++ {
		most = i
		for j := i + 1; j < n; j++ {
			if (*AI)[most].count < (*AI)[j].count {
				most = j
			}
		}
		temp = (*AI)[i]
		(*AI)[i] = (*AI)[most]
		(*AI)[most] = temp
	}
}

func HitungTrend(Pakaian TabPakaian, AI *TabAI, Trend TabRiwayat, n, j int) { // Fungsi untuk menghitung trend pakaian berdasarkan riwayat penggunaan
	for i := 0; i < j; i++ {
		AI[i].count = 0
		for k := usingAI; k > usingAI-4 && k >= 0; k-- {
			if AI[i].id == Trend[k].used[0] || AI[i].id == Trend[k].used[1] || AI[i].id == Trend[k].used[2] || AI[i].id == Trend[k].used[3] {
				AI[i].count++
			}
		}
	}
}
func InputPreferensi(Tujuan, Cuaca *int) { // Fungsi untuk meminta input preferensi pengguna
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-12s%-28s |", " ", "Masukan Preferensi")
	fmt.Printf("\n+------------------------------------------+")
	AturanInput("AITujuan")
	fmt.Printf("\n| %-10s%-2s ", "Tujuan", ":")
	fmt.Scan(Tujuan)
	AturanInput("AIWeather")
	fmt.Printf("\n| %-10s%-2s ", "Cuaca", ":")
	fmt.Scan(Cuaca)
}

func InputDataAI(Pakaian TabPakaian, AI *TabAI, tujuan, cuaca, n int, j *int) { // Fungsi untuk mengisi data AI berdasarkan preferensi pengguna
	var i int
	var a, b bool
	*j = 0
	fmt.Println(tujuan, cuaca)
	for i = 0; i <= n; i++ {
		if tujuan == 1 {
			a = Pakaian[i].formalitas == 3 && Pakaian[i].kategori != 9 && Pakaian[i].kategori != 6
		} else if tujuan == 2 {
			a = Pakaian[i].formalitas == 2 || Pakaian[i].formalitas == 3
		} else {
			a = Pakaian[i].formalitas == 1 || Pakaian[i].formalitas == 2
		}
		if cuaca == 1 {
			b = Pakaian[i].kategori != 1 && Pakaian[i].kategori != 3 && Pakaian[i].kategori != 7 && Pakaian[i].kategori != 8
		} else {
			b = Pakaian[i].kategori != 2 && Pakaian[i].kategori != 4
		}
		if a && b && Pakaian[i].aktif {
			(*AI)[*j].id = Pakaian[i].id
			*j++
		}
	}
}

func PilihPakaian(Pakaian TabPakaian, AI TabAI, Trend *TabRiwayat, n, j int) { // Fungsi untuk memilih pakaian berdasarkan rekomendasi AI
	var pilihan int
	var ada bool = false
	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-47s%-53s |", " ", "Atasan")
	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")

	for i := 0; i < j; i++ {
		if Pakaian[BinarySearch(Pakaian, n, AI[i].id)].kategori >= 1 && Pakaian[BinarySearch(Pakaian, n, AI[i].id)].kategori <= 4 {
			if !ada {
				Table()
			}
			WriteData(Pakaian, BinarySearch(Pakaian, n, AI[i].id))
			ada = true
		}
	}
	if ada {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		AturanInput("ID")
		fmt.Printf("\nMasukan ID Pakaian yang ingin dipilih: ")
		fmt.Scan(&pilihan)
		Trend[usingAI].day = usingAI
		Trend[usingAI].used[0] = pilihan
		ada = false
	} else {
		Arlert("DataTidakAda2")
	}

	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-47s%-53s |", " ", "Bawahan")
	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	ada = false
	for i := 0; i < j; i++ {
		if Pakaian[BinarySearch(Pakaian, n, AI[i].id)].kategori >= 5 && Pakaian[BinarySearch(Pakaian, n, AI[i].id)].kategori <= 6 {
			if !ada {
				Table()
			}
			WriteData(Pakaian, BinarySearch(Pakaian, n, AI[i].id))
			ada = true
		}
	}
	if ada {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		AturanInput("ID")
		fmt.Printf("\nMasukan ID Pakaian yang ingin dipilih: ")
		fmt.Scan(&pilihan)
		Trend[usingAI].used[1] = pilihan
		ada = false
	} else {
		Arlert("DataTidakAda2")
	}

	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-47s%-53s |", " ", "Luaran")
	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	ada = false
	for i := 0; i < j; i++ {
		if Pakaian[BinarySearch(Pakaian, n, AI[i].id)].kategori >= 7 && Pakaian[BinarySearch(Pakaian, n, AI[i].id)].kategori <= 9 {
			if !ada {
				Table()
			}
			WriteData(Pakaian, BinarySearch(Pakaian, n, AI[i].id))
			ada = true
		}
	}
	if ada {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		AturanInput("ID")
		fmt.Printf("\nMasukan ID Pakaian yang ingin dipilih: ")
		fmt.Scan(&pilihan)
		Trend[usingAI].used[2] = pilihan
		ada = false
	} else {
		Arlert("DataTidakAda2")
	}

	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-47s%-53s |", " ", "Alas Kaki")
	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	ada = false
	for i := 0; i < j; i++ {
		if Pakaian[BinarySearch(Pakaian, n, AI[i].id)].kategori >= 10 && Pakaian[BinarySearch(Pakaian, n, AI[i].id)].kategori <= 11 {
			if !ada {
				Table()
			}
			WriteData(Pakaian, BinarySearch(Pakaian, n, AI[i].id))
			ada = true
		}
	}
	if ada {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		AturanInput("ID")
		fmt.Printf("\nMasukan ID Pakaian yang ingin dipilih: ")
		fmt.Scan(&pilihan)
		Trend[usingAI].used[3] = pilihan
	} else {
		Arlert("DataTidakAda2")
	}
	usingAI++
}

func KombinasiPakaian(Pakaian TabPakaian, AI TabAI, Trend *TabRiwayat, n, j int) { // Fungsi untuk menampilkan kombinasi pakaian berdasarkan rekomendasi AI
	var atasanIdx, bawahanIdx, luaranIdx, alasKakiIdx int = -1, -1, -1, -1
	var acc, matching bool = false, false
	for i := 0; i < j && !matching; i++ {
		idx := BinarySearch(Pakaian, n, (AI)[i].id)
		if idx > -1 && Pakaian[idx].aktif && Pakaian[idx].kategori >= 1 && Pakaian[idx].kategori <= 4 {
			atasanIdx = idx
			matching = true
		}
	}
	matching = false
	for i := 0; i < j && !matching; i++ {
		idx := BinarySearch(Pakaian, n, (AI)[i].id)
		if idx > -1 && Pakaian[idx].aktif && Pakaian[idx].kategori >= 5 && Pakaian[idx].kategori <= 6 {
			if atasanIdx != -1 && Pakaian[idx].warna == Pakaian[atasanIdx].warna {
				bawahanIdx = idx
				matching = true
			} else if bawahanIdx == -1 {
				bawahanIdx = idx
			}
		}
	}
	matching = false
	for i := 0; i < j && !matching; i++ {
		idx := BinarySearch(Pakaian, n, (AI)[i].id)
		if idx > -1 && Pakaian[idx].aktif && Pakaian[idx].kategori >= 7 && Pakaian[idx].kategori <= 9 {
			if atasanIdx != -1 && Pakaian[idx].warna == Pakaian[atasanIdx].warna {
				luaranIdx = idx
				matching = true
			} else if luaranIdx == -1 {
				luaranIdx = idx
			}
		}
	}
	matching = false
	for i := 0; i < j && !matching; i++ {
		idx := BinarySearch(Pakaian, n, (AI)[i].id)
		if idx > -1 && Pakaian[idx].aktif && Pakaian[idx].kategori >= 10 && Pakaian[idx].kategori <= 11 {
			if bawahanIdx != -1 && Pakaian[idx].warna == Pakaian[bawahanIdx].warna {
				alasKakiIdx = idx
				matching = true
			} else if alasKakiIdx == -1 {
				alasKakiIdx = idx
			}
		}
	}
	fmt.Printf("\n+-------------------------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-48s%-54s |", " ", "Kombinasi Pakaian")
	fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+------------+")
	if atasanIdx != -1 {
		WriteData(Pakaian, atasanIdx)
		fmt.Printf(" %-10s |", "Atasan")
	}
	if bawahanIdx != -1 {
		WriteData(Pakaian, bawahanIdx)
		fmt.Printf(" %-10s |", "Bawahan")
	}
	if luaranIdx != -1 {
		WriteData(Pakaian, luaranIdx)
		fmt.Printf(" %-10s |", "Luaran")
	}
	if alasKakiIdx != -1 {
		WriteData(Pakaian, alasKakiIdx)
		fmt.Printf(" %-10s |", "Alas Kaki")
	}
	fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+------------+")
	AturanInput("Acc")
	fmt.Scan(&acc)
	if acc {
		Trend[usingAI].day = usingAI
		Trend[usingAI].used[0] = Pakaian[atasanIdx].id
		Trend[usingAI].used[1] = Pakaian[bawahanIdx].id
		if luaranIdx != -1 {
			Trend[usingAI].used[2] = Pakaian[luaranIdx].id
		} else {
			Trend[usingAI].used[2] = 0
		}
		if alasKakiIdx != -1 {
			Trend[usingAI].used[3] = Pakaian[alasKakiIdx].id
		} else {
			Trend[usingAI].used[3] = 0
		}
		usingAI++
	} else {
		PilihPakaian(Pakaian, AI, Trend, n, j)
	}
}

func main() {
	var pakaian TabPakaian
	var nPakaian int = -1
	var keyInt int
	var pilih, keyString string
	var tujuan, cuaca, j int
	var AI TabAI
	var Trend TabRiwayat

	for valid := false; !valid; {
		Welcome("Start")
		Menu("Welcome")
		fmt.Scan(&pilih)

		switch pilih {
		case "1":
			for valid1 := false; !valid1; {
				Menu("ManajemenPakaian")
				fmt.Scan(&pilih)
				switch pilih {
				case "1":
					MenuTabPakaian(pakaian, nPakaian)
				case "2":
					nPakaian++
					AddData(&pakaian, nPakaian)
				case "3":
					AturanInput("Edit")
					fmt.Printf("\nMasukan Pilihan: ")
					fmt.Scan(&pilih)
					switch pilih {
					case "1":
						AturanInput("Id")
						fmt.Printf("\nMasukan Id data yang ingin di edit: ")
						fmt.Scan(&keyInt)
						EditData(&pakaian, nPakaian, keyInt)
					case "2":
						AturanInput("Nama")
						fmt.Printf("\nMasukan Nama data yang ingin di edit: ")
						fmt.Scan(&keyString)
						if SquentialSearchNama(pakaian, nPakaian, keyString) > -1 {
							EditData(&pakaian, nPakaian, pakaian[SquentialSearchNama(pakaian, nPakaian, keyString)].id)
						} else {
							Arlert("DataTidakAda1")
						}
					}
				case "4":
					Menu("delete")
					fmt.Scan(&pilih)
					switch pilih {
					case "1":
						AturanInput("Delete")
						fmt.Printf("\nMasukan Pilihan: ")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							fmt.Printf("Masukan Id data yang ingin di hapus(soft delete): ")
							fmt.Scan(&keyInt)
							Delete(&pakaian, nPakaian, keyInt)
						case "2":
							AturanInput("Nama")
							fmt.Printf("\nMasukan Nama data yang ingin di hapus(soft delete): ")
							fmt.Scan(&keyString)
							idx := SquentialSearchNama(pakaian, nPakaian, keyString)
							if idx > -1 {
								Delete(&pakaian, nPakaian, pakaian[idx].id)
							} else {
								Arlert("DataTidakAda1")
							}
						default:
							Arlert("ErrorInput")
						}
					case "2":
						Table()
						for i := 0; i <= nPakaian; i++ {
							if !pakaian[i].aktif {
								WriteData(pakaian, i)
							}
						}
						fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
						fmt.Printf("\nMasukan Id data yang ingin di hapus(permanen): ")
						fmt.Scan(&keyInt)
						DeletePermanent(&pakaian, &nPakaian, keyInt)
					case "3":
						Table()
						for i := 0; i <= nPakaian; i++ {
							if !pakaian[i].aktif {
								WriteData(pakaian, i)
							}
						}
						fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
						fmt.Printf("\nMasukan Id data yang ingin di aktifkan: ")
						fmt.Scan(&keyInt)
						Restore(&pakaian, &nPakaian, keyInt)
					}
				case "5":
					Menu("Search")
					fmt.Scan(&pilih)
					switch pilih {
					case "1":
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&keyInt)
						SelectionSortId(&pakaian, nPakaian, true)
						if BinarySearch(pakaian, nPakaian, keyInt) > -1 {
							Table()
							WriteData(pakaian, BinarySearch(pakaian, nPakaian, keyInt))
							fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
						} else {
							Arlert("DataTidakAda1")
						}
					case "2":
						AturanInput("Nama")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&keyString)
						if SquentialSearchNama(pakaian, nPakaian, keyString) > -1 {
							Table()
							WriteData(pakaian, SquentialSearchNama(pakaian, nPakaian, keyString))
							fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
						} else {
							Arlert("DataTidakAda2")
						}
					case "3":
						AturanInput("Warna")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&keyString)
						SquentialSearchWarna(pakaian, nPakaian, keyString)
					case "4":
						AturanInput("Kategori")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&keyInt)
						SquentialSearcInt(pakaian, nPakaian, keyInt, "Kategori")
					case "5":
						AturanInput("Formalitas")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&keyInt)
						SquentialSearcInt(pakaian, nPakaian, keyInt, "Formalitas")
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
							MenuTabPakaian(pakaian, nPakaian)
						case "2":
							SelectionSortId(&pakaian, nPakaian, false)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						}
					case "2":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByNama(&pakaian, nPakaian, true)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						case "2":
							SortByNama(&pakaian, nPakaian, false)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						}
					case "3":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByWarna(&pakaian, nPakaian, true)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						case "2":
							SortByWarna(&pakaian, nPakaian, false)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						}
					case "4":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByKategori(&pakaian, nPakaian, true)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						case "2":
							SortByKategori(&pakaian, nPakaian, false)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						}
					case "5":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByFormalitas(&pakaian, nPakaian, true)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						case "2":
							SortByFormalitas(&pakaian, nPakaian, false)
							MenuTabPakaian(pakaian, nPakaian)
							SelectionSortId(&pakaian, nPakaian, true)
						}
					case "0":
						valid1 = true
					default:
						Arlert("ErrorInput")
					}
				case "0":
					valid1 = true
				default:
					Arlert("ErrorInput")
				}
			}
		case "2":
			InputPreferensi(&tujuan, &cuaca)
			InputDataAI(pakaian, &AI, tujuan, cuaca, nPakaian, &j)
			HitungTrend(pakaian, &AI, Trend, nPakaian, j)
			SortDataAI(&AI, Trend, j)
			fmt.Println(usingAI)
			fmt.Printf("\nRekomendasi Pakaian Berdasarkan Preferensi Anda:\n")
			Table()
			for i := 0; i <= nPakaian; i++ {
				if BinarySearch(pakaian, nPakaian, AI[i].id) > -1 {
					WriteData(pakaian, BinarySearch(pakaian, nPakaian, AI[i].id))
				}
			}
			fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
			KombinasiPakaian(pakaian, AI, &Trend, nPakaian, j)
		case "0":
			valid = true
		default:
			Arlert("ErrorInput")
		}
	}
}

/*
Data demo

Untuk mencoba program ini, Anda dapat menggunakan data yang telah disediakan di bawah ini.
Copy dan paste data ini ketika anda sudah berada didalam fitur Manajemen Pakaian.

Berikut adalah data pakaian yang dapat Anda gunakan:
2
Kemeja_Teknik
Biru
1
3
2
Kaos_Tim
Merah
4
2
2
Celana_Kerja
Hitam
5
3
2
Jaket_Motor
Hijau
7
2
2
Luaran_Casual
Coklat
8
2
2
Kaos_Ringan
Abu-Abu
4
1
2
Sepatu_Formal
Hitam
10
3
2
Sendal_Santai
Coklat
9
1
2
Kemeja_Kantor
Putih
1
3
2
Celana_Pendek
Biru
6
1
2
Kaos_Olahraga
Merah
4
1
2
Kemeja_Batik
Coklat
1
3
2
Jaket_Hujan
Hijau
7
2
2
Luaran_Formal
Abu-Abu
8
3
2
Celana_Jeans
Biru
5
2
2
Sendal_Gunung
Hitam
9
1
2
Sepatu_Sport
Putih
10
2
2
Kemeja_Safari
Hijau
2
2
2
Kaos_Distro
Hitam
4
1
2
Kaos_Bayi
Merah
3
1
2
Celana_Kain
Coklat
5
3
2
Jaket_Kulit
Hitam
7
3
2
Kaos_Crop
Putih
4
1
2
Luaran_Jeans
Biru
8
2
2
Sendal_Pantai
Biru
9
1
2
Sepatu_Kets
Merah
10
2
2
Kemeja_Pesta
Merah
1
3
2
Kaos_Muslim
Putih
4
2
2
Celana_Santai
Abu-Abu
6
1
2
Jaket_Parasut
Hijau
7
2
2
Luaran_Batik
Coklat
8
3
2
Sendal_Rumah
Abu-Abu
9
1
2
Sepatu_Boot
Hitam
10
3
2
Kemeja_Hitam
Hitam
1
3
2
Kaos_Tidur
Biru
3
1
2
Celana_Jogger
Hijau
6
1
2
Jaket_Down
Abu-Abu
7
3
2
Kaos_Vintage
Merah
4
2
2
Luaran_Hoodie
Abu-Abu
8
2
2
Sendal_Jepit
Merah
9
1
2
Sepatu_Kerja
Coklat
10
3
2
Kemeja_Slimfit
Biru
1
3
2
Kaos_Polos
Putih
4
1
2
Celana_Chino
Coklat
5
2
2
Jaket_Bomber
Hijau
7
2
2
Luaran_Hitam
Hitam
8
2
2
Sendal_Karet
Hijau
9
1
2
Sepatu_Canvas
Putih
10
2
2
Kemeja_Katun
Putih
1
3
2
Kaos_Stripe
Hitam
4
1
2
Celana_Slimfit
Abu-Abu
5
3
2
Jaket_Parka
Hijau
7
2
2
Luaran_Sport
Merah
8
2
2
Sendal_Sport
Abu-Abu
9
1
2
Sepatu_Lari
Biru
10
2
2
Kemeja_Satin
Merah
1
3
2
Kaos_TieDye
Pelangi
4
1
2
Celana_Linen
Putih
5
2
2
Jaket_Fleece
Hitam
7
3
2
Luaran_Coat
Abu-Abu
8
3
2
Sendal_Balik
Hitam
9
1
2
Sepatu_Hiking
Coklat
10
3
2
Kemeja_Casual
Biru
2
2
2
Kaos_Band
Hitam
4
1
2
Celana_Cargo
Hijau
5
2
2
Jaket_Denim
Biru
7
2
2
Luaran_Cardigan
Coklat
8
2
2
Sendal_Kain
Merah
9
1
2
Sepatu_Tenis
Putih
10
2
2
Kemeja_Tartan
Merah
2
2
2
Kaos_Raglan
Putih
4
1
2
Celana_Training
Biru
6
1
2
Jaket_Kasmir
Coklat
7
3
2
Luaran_Stylish
Hitam
8
2
2
Sendal_Kulit
Coklat
9
1
2
Sepatu_Kanvas
Putih
10
2
2
Kemeja_Flanel
Abu-Abu
1
2
2
Kaos_Graphic
Merah
4
1
2
Celana_Legging
Hitam
6
1
2
Jaket_Tebal
Abu-Abu
7
3
2
Luaran_Elegan
Hitam
8
3
2
Sendal_Tali
Coklat
9
1
2
Sepatu_Casual
Biru
10
2
*/
