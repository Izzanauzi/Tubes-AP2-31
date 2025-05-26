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

const NMAX int = 1024 // Maksimun jumlah data pakaian

type DaftarPakaian struct {
	id, formalitas, kategori, list int    // Atribut numerik pakaian
	nama, warna                    string // Atribut string pakaian
	aktif                          bool   // Flag status untuk data
}

type TabPakaian [NMAX]DaftarPakaian

var listKategori = [11]string{
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

var listFormalitas = [4]string{
	"",
	"Santai",
	"Semi Formal",
	"Formal",
}

type Riwayat struct {
	id   string
	used [5]int
}

type TabRiwayat Riwayat

var nextId int //Variabel untuk generate Id Otomatis

type NeedAI struct {
	id    int
	count int
}

type TabAI [NMAX]NeedAI

func Welcome(tipe string) {
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
	case "ManajemenPakaian": //Tampilan opsi setelah memilih Manajemen Pakaian
		fmt.Printf("\n| %-40s |", "[1] Daftar Pakaian")
		fmt.Printf("\n| %-40s |", "[2] Tambah Pakaian")
		fmt.Printf("\n| %-40s |", "[3] Edit Pakaian")
		fmt.Printf("\n| %-40s |", "[4] Hapus Pakaian")
		fmt.Printf("\n| %-40s |", "[5] Cari Pakaian")
		fmt.Printf("\n| %-40s |", "[6] Sortir Pakaian")
		fmt.Printf("\n| %-40s |", "[0] Home")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Search": //Tampilan saat memilih Cari Pakaian
		fmt.Printf("\n| %-40s |", "[1] Search Berdasarkan ID")
		fmt.Printf("\n| %-40s |", "[2] Search Berdasarkan Nama")
		fmt.Printf("\n| %-40s |", "[3] Search Berdasarkan Warna")
		fmt.Printf("\n| %-40s |", "[4] Search Berdasarkan Kategori")
		fmt.Printf("\n| %-40s |", "[5] Search Berdasarkan Formalitas")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Sort": //Tampilan saat memilih sotir Pakaian
		fmt.Printf("\n| %-40s |", "[1] Sortir Berdasarkan ID")
		fmt.Printf("\n| %-40s |", "[2] Sortir Berdasarkan Nama")
		fmt.Printf("\n| %-40s |", "[3] Sortir Berdasarkan Warna")
		fmt.Printf("\n| %-40s |", "[4] Sortir Berdasarkan Kategori")
		fmt.Printf("\n| %-40s |", "[5] Sortir Berdasarkan Formalitas")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/3/4/5/0)?")
	case "Ascending": //Tampilan menu sorting ascending/descending
		fmt.Printf("\n| %-40s |", "[1] Ascending")
		fmt.Printf("\n| %-40s |", "[2] Descending")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/0)?")
	case "delete": //Tampilan menu delete
		fmt.Printf("\n| %-40s |", "[1] Soft Delete")
		fmt.Printf("\n| %-40s |", "[2] Hard Delete")
		fmt.Printf("\n| %-40s |", "[0] Back")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\nPilih (1/2/0)?")
	}
}

func Table() {
	//digunakan sebagai bagian atas tabel daftar pakaian
	fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	fmt.Printf("\n| %-6s | %-30s | %-15s | %-25s | %-12s |", "Id", "Nama Pakaian", "Warna", "Kategori", "Formalitas")
	fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
}

func MenuTabPakaian(Pakaian TabPakaian, n int) {
	fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-43s%-14s%-43s |", " ", "Daftar Pakaian", " ")
	if n == -1 { //Jika array blm diisi, maka akan mengeluarkan informasi bahwa data kosong
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-40s%-20s%-40s |", " ", "Data Pakaian Kosong!", " ")
		fmt.Printf("\n| %-36s%-28s%-36s |", " ", "Silahkan Tambah Data Pakaian", " ")
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	} else { //Jika array sudah diisi, maka akan menulis semua data array
		Table()
		for i := 0; i <= n; i++ {
			if Pakaian[i].aktif {
				WriteData(Pakaian, i)
			}
		}
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
	}
}

func Arlert(Tipe string) {
	switch Tipe {
	case "ErrorInput": //Tampilan salah membrikan masukan
		fmt.Printf("\n ______   _____    _____     ____    _____  ")
		fmt.Printf("\n|  ____| |  __ \\  |  __ \\   / __ \\  |  __ \\ ")
		fmt.Printf("\n| |__    | |__) | | |__) | | |  | | | |__) |")
		fmt.Printf("\n|  __|   |  _  /  |  _  /  | |  | | |  _  / ")
		fmt.Printf("\n| |____  | | \\ \\  | | \\ \\  | |__| | | | \\ \\ ")
		fmt.Printf("\n|______| |_|  \\_\\ |_|  \\_\\  \\____/  |_|  \\_\\")
		fmt.Printf("\n%-9s%s", " ", "Masukan Input Yang Sesuai!")
	case "DataTidakAda1": //Tampilan data tidak ditemukan
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-10s%-30s |", " ", "Data Tidak Ditemukan")
		fmt.Printf("\n+------------------------------------------+")
	case "DataTidakAda2": //Tampilan data tidak ditemukan
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-40s%-60s |", " ", "Data Tidak Ditemukan")
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
	}
}

func AturanInput(Tipe string) {
	switch Tipe {
	case "Nama": //tampilan menu Nama saat akan memasukan data
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-14s%-26s |", " ", "NAMA PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Nama Pakaian")
		fmt.Printf("\n| %-40s |", "#2 Gunakan \"_\" sebagai pengganti spasi")
		fmt.Printf("\n| %-40s |", "EX. Kemeja_Teknik")
		fmt.Printf("\n+------------------------------------------+")
	case "Warna": //tampilan menu warna saat akan memasukan data
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-13s%-27s |", " ", "WARNA PAKAIAN")
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-40s |", "#1 Masukan Warna Pakaian")
		fmt.Printf("\n| %-40s |", "EX. Merah")
		fmt.Printf("\n+------------------------------------------+")
	case "Kategori": //tampilan menu kategori saat akan memasukan data
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
	case "Formalitas": //tampilan menu formalitas saat akan memasukan data
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
	case "AITujuan": //tampilan menu AI saat akan memasukan data
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
	case "AIWeather": //tampilan menu AI saat akan memasukan data
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
	}
}

func WriteData(Pakaian TabPakaian, i int) {
	/*
		Initial State	: Terdefinisi tipe data buatan TabPakaian Pakaian yang berisi array data pakaian, dan integer i yang merupakan indeks array yang akan diakses
		Process			:
			1. Menampilkan informasi pakaian (ID, Nama, Warna, Kategori, dan Formalitas) sesuai format tabel
		Final State		: Informasi pakaian pada indeks ke-i telah ditampilkan ke layar dalam format tabel
	*/
	fmt.Printf("\n| %-6d ", Pakaian[i].id)
	fmt.Printf("| %-30s ", Pakaian[i].nama)
	fmt.Printf("| %-15s ", Pakaian[i].warna)
	fmt.Printf("| %-25s ", listKategori[Pakaian[i].kategori])
	fmt.Printf("| %-12s |", listFormalitas[Pakaian[i].formalitas])
}

func AddData(Pakaian *TabPakaian, n int) {
	//Fitur Tambah Pakaian
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TabPakaian Pakaian, integer n sebagai indeks untuk data baru yang akan ditambahkan
		Process			:
			1. Meminta input dari pengguna untuk mengisi data Nama, Warna, Kategori, dan Formalitas pakaian
			2. Menetapkan ID otomatis dari variabel global nextId dan mengaktifkan data
		Final State		: Data pakaian pada indeks ke-n telah terisi lengkap dan aktif dan variabel global nextId bertambah 1 untuk ID selanjutnya
	*/
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
	nextId++ //Update nextId untuk ID berikutnya
}

func EditData(Pakaian *TabPakaian, n int, KeyInt int) {
	//Fitur edit
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TabPakaian, integer n sebagai jumlah data, dan KeyInt sebagai ID pakaian yang akan diedit
		Process			:
			1. Melakukan pencarian indeks data berdasarkan ID menggunakan Binary Search
			2. Jika data ditemukan:
				1. Menyimpan ID dan status aktif ke dalam variabel sementara
				2. Meminta input pengguna untuk data baru: Nama, Warna, Kategori, dan Formalitas
				3. Menampilkan perbandingan data sebelum dan sesudah perubahan
				4. Meminta konfirmasi perubahan dari pengguna
				5. Jika konfirmasi ditolak, data dikembalikan ke kondisi semula
			3. Jika data tidak ditemukan, tampilkan pesan error
		Final State		: Jika konfirmasi diterima, data pada indeks hasil pencarian telah diperbarui. Jika ditolak, data tetap seperti semula
	*/
	var ganti int = BinarySearch(*Pakaian, n, KeyInt) // Cari index data berdasarkan ID
	var TempEdit, backup DaftarPakaian
	var konfirmasi bool
	if ganti > -1 && Pakaian[ganti].aktif {
		TempEdit.id = (*Pakaian)[ganti].id
		TempEdit.aktif = true
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-15s%-25s |", " ", "EDIT DATA")
		fmt.Printf("\n| %-15s%-3s%6d%15s |", " ", "ID: ", Pakaian[ganti].id, " ")
		fmt.Printf("\n+------------------------------------------+")
		// Edit nama
		AturanInput("Nama")
		fmt.Printf("\nData Sebelumnya\t: %s", Pakaian[ganti].nama)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&TempEdit.nama)
		// Edit warna
		AturanInput("Warna")
		fmt.Printf("\nData Sebelumnya\t: %s", Pakaian[ganti].warna)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&TempEdit.warna)
		// Edit kategori
		AturanInput("Kategori")
		fmt.Printf("\nData Sebelumnya\t: %d", Pakaian[ganti].kategori)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&TempEdit.kategori)
		// Edit formalitas
		AturanInput("Formalitas")
		fmt.Printf("\nData Sebelumnya\t: %d", Pakaian[ganti].formalitas)
		fmt.Printf("\nData Baru\t: ")
		fmt.Scan(&TempEdit.formalitas)
		// Menampilkan perbandingan sebelum dan sesudah
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-43s%-14s%-43s |", " ", "Perubahan Data", " ")
		Table()
		WriteData(*Pakaian, ganti)
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		backup = (*Pakaian)[ganti]
		(*Pakaian)[ganti] = TempEdit
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
			(*Pakaian)[ganti] = backup // Batalkan perubahan jika ditolak
		}
	} else {
		Arlert("DataTidakAda1") // Tampilkan pesan jika ID tidak ditemukan
	}
}

func Min(Pakaian TabPakaian, n, i int) int {
	/*
		Mengembalikan nilai minimum dari array Pakaian dari index array integer i sampai jumlah sebanyak integer n
	*/
	var min = i
	for j := i + 1; j <= n; j++ {
		if Pakaian[min].id > Pakaian[j].id {
			min = j // Perbarui indeks minimum jika ditemukan nilai lebih kecil
		}
	}
	return min // kembalikan indeks minimum
}

func Max(Pakaian TabPakaian, n, i int) int {
	/*
		Mengembalikan nilai maksimum array Pakaian dari index array integer i sampai jumlah sebanyak integer n
	*/
	var max = i
	for j := i + 1; j <= n; j++ {
		if Pakaian[max].id < Pakaian[j].id {
			max = j // Perbarui indeks maksimum jika ditemukan nilai lebih besar
		}
	}
	return max // Kembalikan indeks maksimum
}

func SelectionSortId(Pakaian *TabPakaian, n int, Ascending bool) {
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TabPakaian, integer n sebagai jumlah data array, dan boolean Ascending untuk menentukan urutan (true = naik, false = turun)
		Process			:
			1. Melakukan iterasi dari indeks 0 hingga n
			2. Pada setiap iterasi:
				1. Menentukan indeks dengan ID terkecil (jika Ascending) atau terbesar (jika Descending)
					1. Jika Ascending maka akan mencari nilai minimum dari array dengan cara memanggil fungsi min dan menempatkan hasilnya ke Idx
					2. Jika Descending maka akan mencari nilai maksimum dari array dengan cara memanggil fungsi max dan menempatkan hasilnya ke Idx
				2. Menukar posisi data index ke i dengan data index ke Idx dengan cara menggunakan temp sebagai tempat penyimpanan sementara
		Final State		: Data pakaian dalam array telah terurut berdasarkan ID sesuai urutan yang diinginkan (naik atau turun)
	*/
	var temp DaftarPakaian
	var Idx int
	for i := 0; i < n; i++ {
		if Ascending {
			Idx = Min(*Pakaian, n, i) // Cari indeks terkecil
		} else {
			Idx = Max(*Pakaian, n, i) // Cari indeks terbesar
		}
		temp = (*Pakaian)[i]            // Simpan sementara
		(*Pakaian)[i] = (*Pakaian)[Idx] // Tukar posisi
		(*Pakaian)[Idx] = temp          // Selesaikan pertukaran
	}
}

func BinarySearch(Pakaian TabPakaian, n int, KeyInt int) int {
	/*
		Mengembalikan index array Pakaian yang memiliki Id yang sama seperti KeyInt, jika data sebanyak integer n tidak memiliki Id yang sesuai dengan KeyInt kembalikan integer -1
	*/
	var left, mid, right int
	left = 0
	right = n
	for left <= right {
		mid = (left + right) / 2
		if Pakaian[mid].id == KeyInt {
			return mid // Ditemukan
		} else if Pakaian[mid].id < KeyInt {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // Tidak ditemukan
}

func SquentialSearchString(Pakaian TabPakaian, n int, KeyString, Mode string) {
	/*
		Initial State	: Terdefinisi array Pakaian bertipe TabPakaian, integer n sebagai jumlah array, string KeyString sebagai kata kunci pencarian, dan string Mode sebagai jenis data yang dicari ("Nama" atau "Warna")
		Process			:
			1. Melakukan iterasi dari indeks 0 hingga n-1
			2. Untuk setiap elemen:
				- Jika Mode adalah "Nama", periksa apakah Nama mengandung KeyString dan memiliki status Aktif true
				- Jika Mode adalah "Warna", periksa apakah Warna mengandung KeyString dan memiliki status Aktif true
				- Jika ditemukan, tampilkan Table (hanya sekali) ketika iterasi pertama dan tampilkan data pakaian pada indeks tersebut
			- Jika tidak ditemukan, tampilkan pesan bahwa data tidak ditemukan
		Final State		: Data pakaian yang mengandung KeyString di field yang sesuai telah ditampilkan, atau ditampilkan pesan bahwa data tidak ditemukan
	*/
	var found bool
	for i := 0; i <= n; i++ {
		switch Mode {
		case "Nama":
			if Pakaian[i].nama == KeyString && Pakaian[i].aktif {
				if i == 0 {
					Table()
				}
				WriteData(Pakaian, i)
				found = true
			}
		case "Warna":
			if Pakaian[i].warna == KeyString && Pakaian[i].aktif {
				if i == 0 {
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

func SquentialSearcInt(Pakaian TabPakaian, n int, KeyInt int, Mode string) {
	/*
		Initial State	: Terdefinisi array Pakaian bertipe TabPakaian, integer n sebagai jumlah array, integer KeyInt sebagai kata kunci pencarian, dan string Mode sebagai jenis data yang dicari ("Kategori" atau "Formalitas")
		Process			:
			1. Melakukan iterasi dari indeks 0 hingga n-1
			2. Untuk setiap elemen:
				- Jika Mode adalah "Kategori", periksa apakah Kategori mengandung KeyInt dan memiliki status Aktif true
				- Jika Mode adalah "Formalitas", periksa apakah Formalitas mengandung KeyInt dan memiliki status Aktif true
				- Jika ditemukan, tampilkan Table (hanya sekali) ketika iterasi pertama dan tampilkan data pakaian pada indeks tersebut
			- Jika tidak ditemukan, tampilkan pesan bahwa data tidak ditemukan
		Final State		: Data pakaian yang mengandung KeyInt di field yang sesuai telah ditampilkan, atau ditampilkan pesan bahwa data tidak ditemukan
	*/
	var found bool
	for i := 0; i <= n; i++ {
		if i == 0 {
			fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
			fmt.Printf("\n| %-43s%-57s |", " ", "Hasil Pencarian")
		}
		switch Mode {
		case "Kategori":
			if Pakaian[i].kategori == KeyInt && Pakaian[i].aktif {
				if i == 0 {
					Table()
				}
				WriteData(Pakaian, i)
				found = true
			}
		case "Formalitas":
			if Pakaian[i].formalitas == KeyInt && Pakaian[i].aktif {
				if i == 0 {
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

func Delete(Pakaian *TabPakaian, n int, id int) {
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TanPakaian, integer n sebagai jumlah data array, dan integer id array yang akan di nonaktifkan
		Process			:
		Final State		:
	*/
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

func DeletePermanent(Pakaian *TabPakaian, n *int, id int) {
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TanPakaian, integer n sebagai jumlah data array, dan integer id array yang akan di delete permanen
		Process			:
		Final State		:
	*/

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

func SortByNama(Pakaian *TabPakaian, n int, Ascending bool) {
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TabPakaian, integer n sebagai jumlah data array, dan boolean Ascending untuk menentukan urutan (true = naik, false = turun)
		Process			:
			1. Melakukan iterasi dari indeks 0 hingga n
			2. Pada setiap iterasi:
				1. Menyimpan data pada indeks ke-i ke dalam variabel sementara temp
				2. Menginisialisasi variabel j dengan i - 1
				3. Jika Ascending, lakukan perbandingan nama pakaian dari indeks j hingga 0
					- Jika nama pakaian pada indeks j lebih besar dari nama pakaian pada temp, geser data ke kanan
					- Jika tidak, keluar dari loop
				4. Jika Descending, lakukan perbandingan nama pakaian dari indeks j hingga 0
					- Jika nama pakaian pada indeks j lebih kecil dari nama pakaian pada temp, geser data ke kanan
					- Jika tidak, keluar dari loop
			5. Setelah loop selesai, simpan data temp ke dalam indeks j + 1
		Final State		: Data pakaian dalam array telah terurut berdasarkan nama sesuai urutan yang diinginkan (naik atau turun)
	*/
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

func SortByWarna(Pakaian *TabPakaian, n int, Ascending bool) {
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TabPakaian, integer n sebagai jumlah data array, dan boolean Ascending untuk menentukan urutan (true = naik, false = turun)
		Process			:
			1. Melakukan iterasi dari indeks 0 hingga n
			2. Pada setiap iterasi:
				1. Menyimpan data pada indeks ke-i ke dalam variabel sementara temp
				2. Menginisialisasi variabel j dengan i - 1
				3. Jika Ascending, lakukan perbandingan warna pakaian dari indeks j hingga 0
					- Jika warna pakaian pada indeks j lebih besar dari warna pakaian pada temp, geser data ke kanan
					- Jika tidak, keluar dari loop
				4. Jika Descending, lakukan perbandingan warna pakaian dari indeks j hingga 0
					- Jika warna pakaian pada indeks j lebih kecil dari warna pakaian pada temp, geser data ke kanan
					- Jika tidak, keluar dari loop
			5. Setelah loop selesai, simpan data temp ke dalam indeks j + 1
		Final State		: Data pakaian dalam array telah terurut berdasarkan warna sesuai urutan yang diinginkan (naik atau turun)
	*/
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

func SortByKategori(Pakaian *TabPakaian, n int, Ascending bool) {
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TabPakaian, integer n sebagai jumlah data array, dan boolean Ascending untuk menentukan urutan (true = naik, false = turun)
		Process			:
			1. Melakukan iterasi dari indeks 0 hingga n
			2. Pada setiap iterasi:
				1. Menyimpan data pada indeks ke-i ke dalam variabel sementara temp
				2. Menginisialisasi variabel j dengan i - 1
				3. Jika Ascending, lakukan perbandingan kategori pakaian dari indeks j hingga 0
					- Jika kategori pakaian pada indeks j lebih besar dari kategori pakaian pada temp, geser data ke kanan
					- Jika tidak, keluar dari loop
				4. Jika Descending, lakukan perbandingan kategori pakaian dari indeks j hingga 0
					- Jika kategori pakaian pada indeks j lebih kecil dari kategori pakaian pada temp, geser data ke kanan
					- Jika tidak, keluar dari loop
			5. Setelah loop selesai, simpan data temp ke dalam indeks j + 1
		Final State		: Data pakaian dalam array telah terurut berdasarkan kategori sesuai urutan yang diinginkan (naik atau turun)
	*/
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

func SortByFormalitas(Pakaian *TabPakaian, n int, Ascending bool) {
	/*
		Initial State	: Terdefinisi pointer ke array Pakaian bertipe TabPakaian, integer n sebagai jumlah data array, dan boolean Ascending untuk menentukan urutan (true = naik, false = turun)
		Process			:
			1. Melakukan iterasi dari indeks 0 hingga n
			2. Pada setiap iterasi:
				1. Menyimpan data pada indeks ke-i ke dalam variabel sementara temp
				2. Menginisialisasi variabel j dengan i - 1
				3. Jika Ascending, lakukan perbandingan formalitas pakaian dari indeks j hingga 0
					- Jika formalitas pakaian pada indeks j lebih besar dari formalitas pakaian pada temp, geser data ke kanan
					- Jika tidak, keluar dari loop
				4. Jika Descending, lakukan perbandingan formalitas pakaian dari indeks j hingga 0
					- Jika formalitas pakaian pada indeks j lebih kecil dari formalitas pakaian pada temp, geser data ke kanan
					- Jika tidak, keluar dari loop
			5. Setelah loop selesai, simpan data temp ke dalam indeks j + 1
		Final State		: Data pakaian dalam array telah terurut berdasarkan formalitas sesuai urutan yang diinginkan (naik atau turun)
	*/
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

func SortDataAI(AI *TabAI, n int) {
	/*
		Initial State	: Terdefinisi pointer ke array AI bertipe TabAI, integer n sebagai jumlah data array
		Process			:
		1. Melakukan iterasi dari indeks 0 hingga n
		2. Pada setiap iterasi:
			1. Menyimpan indeks dengan nilai count terbesar ke dalam variabel most
			2. Melakukan iterasi dari indeks i+1 hingga n
				- Jika count pada indeks most lebih besar dari count pada indeks j, perbarui most
			3. Menukar posisi data pada indeks i dengan data pada indeks most menggunakan variabel sementara temp
		Final State		: Data AI dalam array telah terurut berdasarkan count dari yang terbesar ke terkecil
	*/
	var most int
	var temp NeedAI
	for i := 0; i < n; i++ {
		most = i
		for j := i + 1; j < n; j++ {
			if (*AI)[most].count > (*AI)[j].count {
				most = j
			}
		}
		temp = (*AI)[i]
		(*AI)[i] = (*AI)[most]
		(*AI)[most] = temp
	}
}

// Fitur AI

func InputPreferensi(Tujuan, Cuaca *int) {
	/*
		Initial State	: Terdefinisi pointer integer Tujuan dan Cuaca sebagai input preferensi pengguna
		Process			:
			1. Pengguna memasukkan preferensi tujuan dan cuaca sesuai format yang ditentukan
		Final State		: Preferensi pengguna telah didapatkan untuk digunakan dalam fitur AI
	*/
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

// Next Function AI, mengisi data array AI sesuai preferensi pengguna
func InputDataAI(Pakaian TabPakaian, AI *TabAI, tujuan, cuaca, n, j int) {
	/*
		Initial State	: Terdefinisi array Pakaian bertipe TabPakaian, pointer AI bertipe TabAI untuk menyimpan data AI, integer tujuan dan cuaca sebagai preferensi pengguna, dan integer n sebagai jumlah data pakaian
		Process			:
			1. Menginisialisasi variabel j sebagai indeks AI
			2. Menginisialisasi variabel a dan b sebagai boolean untuk menyimpan kondisi formalitas dan cuaca
			3. Menggunakan kondisi if untuk menentukan formalitas berdasarkan tujuan
			4. Menggunakan kondisi if untuk menentukan cuaca berdasarkan kategori pakaian
			5. Melakukan iterasi dari indeks 0 hingga n
			6. Jika kondisi formalitas dan cuaca terpenuhi, simpan data pakaian ke dalam AI
		Final State		: Data AI telah terisi sesuai preferensi pengguna
	*/
	var i int
	var a, b bool
	j = 0
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
			(*AI)[j].id = Pakaian[i].id
			j++
		}
	}
}

func main() {
	var Pakaian TabPakaian      //array utama penyimpan pakaian
	var nPakaian int = -1       // indeks akhir array
	var KeyInt int              // input ID atau nilai numerik lainnya
	var pilih, KeyString string //pilihan menu dan kata kunci pencarian
	var Tujuan, Cuaca, j int
	var AI TabAI //array untuk menyimpan data AI

	for valid := false; !valid; {
		Welcome("Start") //menampilkan halaman awal
		Menu("Welcome")  // menampilkan menu utama
		fmt.Scan(&pilih)

		switch pilih {
		case "1": //Menu Manajemen Pakaian
			for valid1 := false; !valid1; {
				// Welcome("DaftarPakaian")
				Menu("ManajemenPakaian")
				fmt.Scan(&pilih)
				switch pilih {
				case "1": //Daftar Pakaian
					MenuTabPakaian(Pakaian, nPakaian)
				case "2": //Tambah Pakaian
					nPakaian++
					AddData(&Pakaian, nPakaian)
				case "3": //Edit Pakaian
					fmt.Printf("Masukan Id data yang ingin di edit: ")
					fmt.Scan(&KeyInt)
					SelectionSortId(&Pakaian, nPakaian, true) //Melakukan selection sort by Id karena akan menggunakan binary search untuk mencari Id
					EditData(&Pakaian, nPakaian, KeyInt)
				case "4": //Hapus Pakaian
					Menu("delete")
					fmt.Scan(&pilih)
					switch pilih {
					case "1": //Hapus Soft
						fmt.Printf("Masukan Id data yang ingin di hapus(soft delete): ")
						fmt.Scan(&KeyInt)
						SelectionSortId(&Pakaian, nPakaian, true)
						Delete(&Pakaian, nPakaian, KeyInt)
					case "2": //Hapus Permanen
						Table()
						for i := 0; i <= nPakaian; i++ {
							if !Pakaian[i].aktif {
								WriteData(Pakaian, i)
							}
						}
						fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
						fmt.Printf("\nMasukan Id data yang ingin di hapus(permanen): ")
						fmt.Scan(&KeyInt)
						SelectionSortId(&Pakaian, nPakaian, true)
						DeletePermanent(&Pakaian, &nPakaian, KeyInt)
					}
				case "5": //Cari Pakaian
					Menu("Search")
					fmt.Scan(&pilih)
					switch pilih {
					case "1": //Berdasarkan ID
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyInt)
						SelectionSortId(&Pakaian, nPakaian, true)
						if BinarySearch(Pakaian, nPakaian, KeyInt) > -1 {
							Table()
							WriteData(Pakaian, BinarySearch(Pakaian, nPakaian, KeyInt))
							fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
						} else {
							Arlert("DataTidakAda1")
						}
					case "2": //Berdasarkan Nama
						AturanInput("Nama")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyString)
						SquentialSearchString(Pakaian, nPakaian, KeyString, "Nama")

					case "3": //Berdasarkan warna
						AturanInput("Warna")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyString)
						SquentialSearchString(Pakaian, nPakaian, KeyString, "Warna")

					case "4": //Berdasarkan kategori
						AturanInput("Kategori")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyInt)
						SquentialSearcInt(Pakaian, nPakaian, KeyInt, "Kategori")

					case "5": //berdasarkan Formalitas
						AturanInput("Formalitas")
						fmt.Printf("\nMasukan Kata Kunci\t: ")
						fmt.Scan(&KeyInt)
						SquentialSearcInt(Pakaian, nPakaian, KeyInt, "Formalitas")
					case "0":
						//Kembali
					}
				case "6": // Menu sortir
					Menu("Sort")
					fmt.Scan(&pilih)
					switch pilih {
					case "1": // Berdasarkan ID
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SelectionSortId(&Pakaian, nPakaian, true)
						case "2":
							SelectionSortId(&Pakaian, nPakaian, false)
						case "0":
						}
					case "2": // Berdasarkan Nama
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByNama(&Pakaian, nPakaian, true)
						case "2":
							SortByNama(&Pakaian, nPakaian, false)
						case "0":
							//Kembali
						}
					case "3":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByWarna(&Pakaian, nPakaian, true)
						case "2":
							SortByWarna(&Pakaian, nPakaian, false)
						case "0":
						}
					case "4":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByKategori(&Pakaian, nPakaian, true)
						case "2":
							SortByKategori(&Pakaian, nPakaian, false)
						case "0":
						}
					case "5":
						Menu("Ascending")
						fmt.Scan(&pilih)
						switch pilih {
						case "1":
							SortByFormalitas(&Pakaian, nPakaian, true)
						case "2":
							SortByFormalitas(&Pakaian, nPakaian, false)
						case "0":
						}
					case "0":
						// kembali
					}
					// case "5":
				// 	sort()
				case "0":
					valid1 = true // kembali ke menu utama
				default:
					Arlert("ErrorInput")
				}
			}
		case "2": //Menu AI
			InputPreferensi(&Tujuan, &Cuaca)
			InputDataAI(Pakaian, &AI, Tujuan, Cuaca, nPakaian, j)
			SortDataAI(&AI, j) // Mengurutkan data AI berdasarkan Trend
			fmt.Printf("\nRekomendasi Pakaian Berdasarkan Preferensi Anda:\n")
			Table()
			SelectionSortId(&Pakaian, nPakaian, true)
			for i := 0; i <= nPakaian; i++ {
				BinarySearch(Pakaian, nPakaian, AI[i].id)
				if BinarySearch(Pakaian, nPakaian, AI[i].id) > -1 {
					WriteData(Pakaian, BinarySearch(Pakaian, nPakaian, AI[i].id))
				}
			}
			fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		case "0":
			valid = true // keluar dari program
		default:
			Arlert("ErrorInput")
		}
	}
}
