package main

import "fmt"

const nmax int = 1024

type AtributePakaian struct {
	Id, Formalitas, Kategori, list int
	Nama, Warna                    string
}
type DaftarPakaian [nmax]AtributePakaian

var pilih int

// type Riwayat struct {
// 	Id   string
// 	Used int
// }
// type TabRiwayat Riwayat

func welcome() {
	//Tampilan Pertama Ketika Memulai Program
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", " ")
	fmt.Printf("\n| %-7s%-33s |", " ", "Selamat Datang Di Aplikasi")
	fmt.Printf("\n| %s |", "AI Stylist dan Manajemen Pakaian Digital")
	fmt.Printf("\n| %-40s |", " ")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n\n+------------------------------------------+")
	fmt.Printf("\n| %-18s%-22s |", " ", "MENU")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", "[1] Daftar Pakaian")
	fmt.Printf("\n| %-40s |", "[2] Rekomendasi AI")
	fmt.Printf("\n| %-40s |", "[0] Exit")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\nPilih (1/2/0)?")
}

func MenuDaftarPakaian(Pakaian DaftarPakaian, n int) {
	//Tampilan Menu Daftar Pakaian
	fmt.Printf("\n+----------------------------------------------------------------------------------------------------+")
	fmt.Printf("\n| %-42s%-14s%-42s |", " ", "Daftar Pakaian", " ")
	if n == -1 { //Jika array blm diisi, maka akan mengeluarkan informasi bahwa data kosong
		fmt.Printf("\n+----------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-39s%-20s%-39s |", " ", "Data Pakaian Kosong!", " ")
		fmt.Printf("\n| %-35s%-28s%-35s |", " ", "Silahkan Tambah Data Pakaian", " ")
		fmt.Printf("\n+----------------------------------------------------------------------------------------------------+")
	} else { //Jika array sudah diisi, maka akan menulis semua data array
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+------------+")
		fmt.Printf("\n| %-6s | %-30s | %-15s | %-25s | %-10s |", "Id", "Nama Pakaian", "Warna", "Kategori", "Formalitas")
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+------------+")
		for i := 0; i <= n; i++ { //Looping untuk menulis semua data array
			WriteData(Pakaian, i)
		}
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+------------+")
	}
	fmt.Printf("\n\n+------------------------------------------+")
	fmt.Printf("\n| %-18s%-22s |", " ", "MENU")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", "[1] Tambah Pakaian")
	fmt.Printf("\n| %-40s |", "[2] Edit Pakaian")
	fmt.Printf("\n| %-40s |", "[3] Hapus Pakaian")
	fmt.Printf("\n| %-40s |", "[4] Cari Pakaian")
	fmt.Printf("\n| %-40s |", "[5] Sortir Pakaian")
	fmt.Printf("\n| %-40s |", "[0] Home")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\nPilih (1/2/3/4/5/6/0)?")
}

func WriteData(Pakaian DaftarPakaian, i int) {
	fmt.Printf("\n| %-6d ", Pakaian[i].Id)
	fmt.Printf("| %-30s ", Pakaian[i].Nama)
	fmt.Printf("| %-15s ", Pakaian[i].Warna)
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
	switch Pakaian[i].Formalitas {
	case 1:
		fmt.Printf("| %-10s |", "Santai")
	case 2:
		fmt.Printf("| %-10s |", "Semi Formal")
	case 3:
		fmt.Printf("| %-10s |", "Formal")
	}
}

func add(Pakaian *DaftarPakaian, n *int) {
	*n++
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-14s%-26s |", " ", "NAMA PAKAIAN")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", "#1 Masukan Nama Pakaian")
	fmt.Printf("\n| %-40s |", "#2 Gunakan \"_\" sebagai pengganti spasi")
	fmt.Printf("\n| %-40s |", "EX. Kemeja_Teknik")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-10s%-2s ", "Nama", ":")
	fmt.Scan(&Pakaian[*n].Nama)
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-13s%-27s |", " ", "WARNA PAKAIAN")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", "#1 Masukan Warna Pakaian")
	fmt.Printf("\n| %-40s |", "#2 Gunakan \",\" tanpa spasi jika terdapat")
	fmt.Printf("\n| %-40s |", "   2 warna atau lebih dalam 1 pakaian")
	fmt.Printf("\n| %-40s |", "EX. Merah,Putih")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-10s%-2s ", "Warna", ":")
	fmt.Scan(&Pakaian[*n].Warna)
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
	fmt.Printf("\n| %-10s%-2s ", "Kategori", ":")
	fmt.Scan(&Pakaian[*n].Kategori)
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
	fmt.Printf("\n| %-10s%-2s ", "Formalitas", ":")
	fmt.Scan(&Pakaian[*n].Formalitas)
	Pakaian[*n].Id = (Pakaian[*n].Formalitas * 100000) + (Pakaian[*n].Kategori * 1000) + (*n + 1)
}

// func edit(A *TabPakaian, n *int) {
// 	fmt.Println("Edit list Pakaian")
// 	// list digunakan untuk memanggil list yang akan di edit
// 	fmt.Print("\nEdit Nama\t: ")
// 	fmt.Scan(&A[*n].Nama)
// 	fmt.Print("\nEdit Warna\t: ")
// 	fmt.Scan(&A[*n].Warna)
// 	fmt.Print("\nEdit Kategori\t: ")
// 	fmt.Scan(&A[*n].Kategori)
// 	fmt.Print("\nEdit Level Formalitas\n\t \nFormalitas\t: ")
// 	fmt.Scan(&A[*n].Formalitas)
// }

func main() {
	var Pakaian DaftarPakaian
	var nPakaian int = -1
	for valid := false; !valid; {
		welcome()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			for valid1 := false; !valid1; {
				MenuDaftarPakaian(Pakaian, nPakaian)
				fmt.Scan(&pilih)
				switch pilih {
				case 1:
					add(&Pakaian, &nPakaian)
				// case 2:
				// 	edit()
				// case 3:
				// 	delete()
				// case 4:
				// 	search()
				// case 5:
				// 	sort()
				case 0:
					valid1 = true
				}
			}
		// case 2:
		case 0:
			valid = true
		}
	}
}
