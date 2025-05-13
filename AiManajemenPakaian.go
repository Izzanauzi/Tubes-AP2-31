package main

import (
	"fmt"
)

const nmax int = 1024

type DaftarPakaian struct {
	Id, Formalitas, Kategori, list int
	Nama, Warna                    string
	Aktif                          bool // ini untuk soft delete true = aktif dan false = Softdelete
}
type TabPakaian [nmax]DaftarPakaian

//	type Riwayat struct {
//		Id   string
//		Used int
//	}
//
// type TabRiwayat Riwayat
var NextId int

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
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		fmt.Printf("\n| %-6s | %-30s | %-15s | %-25s | %-12s |", "Id", "Nama Pakaian", "Warna", "Kategori", "Formalitas")
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		for i := 0; i <= n; i++ { //Looping untuk menulis semua data array
			WriteData(Pakaian, i)
		}
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
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

func ErrorInput() {
	fmt.Printf("\n ______   _____    _____     ____    _____  ")
	fmt.Printf("\n|  ____| |  __ \\  |  __ \\   / __ \\  |  __ \\ ")
	fmt.Printf("\n| |__    | |__) | | |__) | | |  | | | |__) |")
	fmt.Printf("\n|  __|   |  _  /  |  _  /  | |  | | |  _  / ")
	fmt.Printf("\n| |____  | | \\ \\  | | \\ \\  | |__| | | | \\ \\ ")
	fmt.Printf("\n|______| |_|  \\_\\ |_|  \\_\\  \\____/  |_|  \\_\\")
	fmt.Printf("\n%s", "Masukan Input Yang Sesuai")
}

func WriteData(Pakaian TabPakaian, i int) {
	if Pakaian[i].Aktif { // buat cek ii aktif atau kagak
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
			fmt.Printf("| %-12s |", "Santai")
		case 2:
			fmt.Printf("| %-12s |", "Semi Formal")
		case 3:
			fmt.Printf("| %-12s |", "Formal")
		}
	}
}

func add(Pakaian *TabPakaian, n int, IsEdit bool) {
	//Fitur Tambah Pakaian
	/*
		Menambahkan data pakaian kedalam array dengan ketentuan yang sudah diatur
		Generate Pakaian.Id secara otomatis
	*/
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-14s%-26s |", " ", "NAMA PAKAIAN")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", "#1 Masukan Nama Pakaian")
	fmt.Printf("\n| %-40s |", "#2 Gunakan \"_\" sebagai pengganti spasi")
	fmt.Printf("\n| %-40s |", "EX. Kemeja_Teknik")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-10s%-2s ", "Nama", ":")
	fmt.Scan(&Pakaian[n].Nama)
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-13s%-27s |", " ", "WARNA PAKAIAN")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", "#1 Masukan Warna Pakaian")
	fmt.Printf("\n| %-40s |", "EX. Merah")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-10s%-2s ", "Warna", ":")
	fmt.Scan(&Pakaian[n].Warna)
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
	fmt.Scan(&Pakaian[n].Kategori)
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
	fmt.Scan(&Pakaian[n].Formalitas)
	if !IsEdit {
		Pakaian[n].Id = NextId + 1
	}
	Pakaian[n].Aktif = true
	NextId++
}
func edit(Pakaian *TabPakaian, n *int, Key int) {
	//Fitur edit
	/*
		Data yang diedit dicari berdasarkan 2 angka Id paling akhir
		Menampilkan data sebelum diedit
	*/
	var ganti int = BinarySearch(*Pakaian, *n, Key)
	if ganti > -1 {
		fmt.Printf("\n+------------------------------------------------------------------------------------------------------+")
		fmt.Printf("\n| %-43s%-14s%-43s |", " ", "Edit Data", " ")
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		fmt.Printf("\n| %-6s | %-30s | %-15s | %-25s | %-12s |", "Id", "Nama Pakaian", "Warna", "Kategori", "Formalitas")
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		WriteData(*Pakaian, ganti)
		fmt.Printf("\n+--------+--------------------------------+-----------------+---------------------------+--------------+")
		add(Pakaian, ganti, true)
	} else {
		fmt.Printf("Data tidak ditemukan")
	}
}

func min(Pakaian TabPakaian, n, i int) int {
	//Mencari nilai minimum dalam rentang tertentu
	var min = i
	for j := i + 1; j < n; j++ {
		if Pakaian[i].Id > Pakaian[j].Id {
			min = j
		}
	}
	return min
}

func SortById(Pakaian *TabPakaian, n int) {
	//Mengurutkan Index array berdasarkan Id
	var temp DaftarPakaian
	var minIdx int
	for i := 0; i < n; i++ {
		minIdx = min(*Pakaian, n, i)
		temp = (*Pakaian)[i]
		(*Pakaian)[i] = (*Pakaian)[minIdx]
		(*Pakaian)[minIdx] = temp
	}
}

func BinarySearch(Pakaian TabPakaian, n int, key int) int {
	//Pencarian Id menggunakan binary search
	var left, mid, right int
	left = 0
	right = n
	for left <= right {
		mid = (left + right) / 2
		if Pakaian[mid].Id == key {
			return mid
		} else if Pakaian[mid].Id < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// func SquentialSearch(Pakaian TabPakaian, n int, key int) {

// }

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
		fmt.Println("Data Tidak ada")
	}
}

func sortPakaianKeBesar(Pakaian *TabPakaian, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*Pakaian)[j].Aktif && (*Pakaian)[j+1].Aktif {
				if (*Pakaian)[j].Nama > (*Pakaian)[j+1].Nama {
					(*Pakaian)[j], (*Pakaian)[j+1] = (*Pakaian)[j+1], (*Pakaian)[j]
				}
			}
		}
	}
	fmt.Println("Data berhasil diurutkan terkecil ke besar")
}

func sortPakaianKeKecil(Pakaian *TabPakaian, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*Pakaian)[j].Aktif && (*Pakaian)[j+1].Aktif {
				if (*Pakaian)[j].Nama < (*Pakaian)[j+1].Nama {
					(*Pakaian)[j], (*Pakaian)[j+1] = (*Pakaian)[j+1], (*Pakaian)[j]
				}
			}
		}
	}
	fmt.Println("Data berhasil diurutkan Terbesar ke kecil")
}

func main() {
	var Pakaian TabPakaian
	var nPakaian int = -1
	var Key int
	var pilih string
	var urut int
	for valid := false; !valid; {
		welcome()
		fmt.Scan(&pilih)
		switch pilih {
		case "1":
			for valid1 := false; !valid1; {
				MenuTabPakaian(Pakaian, nPakaian)
				fmt.Scan(&pilih)
				switch pilih {
				case "1":
					nPakaian++
					add(&Pakaian, nPakaian, false)
				case "2":
					fmt.Printf("Masukan Id data yang ingin di edit: ")
					fmt.Scan(&Key)
					SortById(&Pakaian, nPakaian) //Melakukan selection sort by Id karena akan menggunakan binary search untuk mencari Id
					edit(&Pakaian, &nPakaian, Key)
				case "3":
					fmt.Printf("Masukan Id data yang ingin di hapus(soft delete): ")
					fmt.Scan(&Key)
					SortById(&Pakaian, nPakaian)
					idx := BinarySearch(Pakaian, nPakaian, Key)
					if idx > -1 && Pakaian[idx].Aktif {
						Pakaian[idx].Aktif = false
						fmt.Println("Data berhasil di delete (soft delete)")
					} else {
						fmt.Println("Data tidak ditemukan / sudah tidak aktif") // sejauh ini sudah bisa didelete, gua push dulu ada acara gua mya >-<
					}
				case "4":
					fmt.Scan(&pilih)
					switch pilih {
					case "1":
						BinarySearch(Pakaian, nPakaian, Key)
					case "2":
						// SquentialSearch()
					case "0":

					}
				case "5":
					fmt.Println("1.mengurut keatas, 2.mengurut kebawah")
					fmt.Scan(&urut)
					switch urut {
						case 1 :
							sortPakaianKeBesar(&Pakaian , nPakaian)
						case 2 :
							sortPakaianKeKecil(&Pakaian , nPakaian)
						}
					// case "5":
				// 	sort()
				case "0":
					valid1 = true
				default:
					ErrorInput()
				}
			}
		// case "2":
		case "0":
			valid = true
		default:
			ErrorInput()
		}
	}
}

//   ______   _____    _____     ____    _____
//  |  ____| |  __ \  |  __ \   / __ \  |  __ \
//  | |__    | |__) | | |__) | | |  | | | |__) |
//  |  __|   |  _  /  |  _  /  | |  | | |  _  /
//  | |____  | | \ \  | | \ \  | |__| | | | \ \
//  |______| |_|  \_\ |_|  \_\  \____/  |_|  \_\

/*

ERRORERRORERRORERRORERRORERRORERRORERRORERROR
RRORERROR+-------------------------+RORERRORE
RORERRORE|          ERROR          |ORERRORER
ORERRORER|MASUKAN INPUT YANG SESUAI|RERRORERR
RERRORERR+-------------------------+ERRORERRO
ERRORERRORERRORERRORERRORERRORERRORERRORERROR

ERRORERRORERRORERRORERRORERRORERRORERRORERROR
RRORERRORERRORERRORERRORERRORERRORERRORERRORE
RORERRORERRORERRORERRORERRORERRORERRORERRORER
ORERRORERRORERRORERRORERRORERRORERRORERRORERR
RERRORERRORERRORERRORERRORERRORERRORERRORERRO
*/
