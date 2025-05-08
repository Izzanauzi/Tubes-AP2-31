package main

import "fmt"

const nmax int = 1024

type Pakaian struct {
	Id, Formalitas, list  int // ku tambahkan list
	Nama, Warna, Kategori string
}
type TabPakaian Pakaian
type Riwayat struct {
	Id   string
	Used int
}
type TabRiwayat Riwayat

func menu() {

}

func Add(A *TabPakaian, n *int) {
	fmt.Printf("Tambah Pakaian")
	fmt.Printf("\nNama\t: ")
	fmt.Scan(&A[*n].Nama)
	fmt.Printf("\nWarna\t: ")
	fmt.Scan(&A[*n].Warna)
	fmt.Printf("\nKategori\t: ")
	fmt.Scan(&A[*n].Kategori)
	fmt.Printf("\nLevel Formalitas\n\t \nFormalitas\t: ")
	fmt.Scan(&A[*n].Formalitas)
	A[*n].list = *n + 1
	*n = *n + 1
}

func edit(A *TabPakaian, n *int) {
	fmt.Println("Edit list Pakaian")
	// list digunakan untuk memanggil list yang akan di edit
	fmt.Print("\nEdit Nama\t: ")
	fmt.Scan(&A[*n].Nama)
	fmt.Print("\nEdit Warna\t: ")
	fmt.Scan(&A[*n].Warna)
	fmt.Print("\nEdit Kategori\t: ")
	fmt.Scan(&A[*n].Kategori)
	fmt.Print("\nEdit Level Formalitas\n\t \nFormalitas\t: ")
	fmt.Scan(&A[*n].Formalitas)
}

func main() {
	var DaftarPakaian TabPakaian
	var menu int
	var nEdit int
	// var Tren TabRiwayat
	var nPakaian int = 0
	fmt.Print("ADD\n EDIT\n MENU\n EXIT\n")
	fmt.Scan(&menu)
	switch menu {
	case 1:
		Add(&DaftarPakaian, &nPakaian)
	case 2:
		fmt.Scan(&nEdit)
		edit(&DaftarPakaian, &nEdit)
	}
}
