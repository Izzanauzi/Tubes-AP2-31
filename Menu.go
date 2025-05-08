package main

import "fmt"

func menu() {
	fmt.Printf("+------------------------------------------+")
	fmt.Printf("\n| %-40s |", " ")
	fmt.Printf("\n| %-7s%-33s |", " ", "Selamat Datang Di Aplikasi")
	fmt.Printf("\n| %s |", "AI Stylist dan Manajemen Pakaian Digital")
	fmt.Printf("\n| %-40s |", " ")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-18s%-22s |", " ", "MENU")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\n| %-40s |", "[1]")
	fmt.Printf("\n| %-40s |", "[2]")
	fmt.Printf("\n| %-40s |", "[3]")
	fmt.Printf("\n| %-40s |", "[4]")
	fmt.Printf("\n+------------------------------------------+")
	fmt.Printf("\nPilih (1,2,3,4)?")
}
func main() {
	menu()
}
