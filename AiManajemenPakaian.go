package main

import "fmt"

const nmax int = 1024

type Pakaian struct {
	Id, Formalitas        string
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
}
func main() {
	var DaftarPakaian TabPakaian
	// var Tren TabRiwayat
	var nPakaian int = 0
	Add(&DaftarPakaian, &nPakaian)
}
