package main

import "fmt"

/*
Dear Ghifar
Udah sampe fitur add
Tapi masih error pas nampilin datanya, mungkin nanti gw coba buat pake fungsi lagi
anjay banyak beut fungsinya
udh jalan sih, lu coba aja
oiya fitur fitur yang blm jalan gw jadiin comment
*/

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
	if n == -1 {
		fmt.Printf("\n+------------------------------------------+")
		fmt.Printf("\n| %-10s%-20s%-10s |", " ", "Data Pakaian Kosong", " ")
		fmt.Printf("\n+------------------------------------------+")
	} else {
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+------------+")
		fmt.Printf("\n| %-6s | %-30s | %-15s | %-25s | %-10s |", "Id", "Nama Pakaian", "Warna", "Kategori", "Formalitas")
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+------------+")
		for i := 0; i <= n; i++ {
			//Write Data Array Pakaian
			fmt.Printf("\n| %-6d | %-30s | %-15s | %-25d | %-10d |", Pakaian[i].Id, Pakaian[i].Nama, Pakaian[i].Warna, Pakaian[i].Kategori, Pakaian[i].Formalitas)
		}
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+------------+")
	}
	fmt.Printf("\n+------------------------------------------+")
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

func add(Pakaian *DaftarPakaian, n *int) {
	*n++
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-13s%-27s |", " ", "TAMBAH PAKAIAN")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-10s%-2s ", "Nama", ":")
	fmt.Scan(&Pakaian[*n].Nama)
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-10s%-2s ", "Warna", ":")
	fmt.Scan(&Pakaian[*n].Warna)
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-16s%-24s |", " ", "KATEGORI")
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
	fmt.Printf("\n| %-15s%-25s |", " ", "FORMALITAS")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", "1 : Santai")
	fmt.Printf("\n| %-40s |", "2 : Semi Formal")
	fmt.Printf("\n| %-40s |", "3 : Formal")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-10s%-2s ", "Formalitas", ":")
	fmt.Scan(&Pakaian[*n].Formalitas)
	fmt.Printf("\n+------------------------------------------+")
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
			}
		// case 2:
		case 0:
			valid = true
		}
	}
}
