package main

import (
	"fmt"
	"math"
)

type Pinjaman struct {
	Nama string
	JumlahPinjaman float64
	Tenor int
	Status string
	Bunga float64
	BayaranPerbulan float64
	sisaPinjam int
	sudahBayar int
	SchemaBunga string
	TotalBunga float64
	TotalBayar float64
	SudahBayar  int
}

var dataPeminjam [1000]Pinjaman
var countPeminjam int = 0

func ClearScreen() {
	//Untuk membersihkan terminal 
	//didapat dari https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go/22892171#22892171
	fmt.Print("\033[H\033[2J")
}

func garis1() {
	fmt.Print("=================================================================================")
}

func garis2() {
	fmt.Println("-------------------------------------------------------------------------------")
}

func DetailPeminjaman(p Pinjaman,idx int) {
	garis2()
	fmt.Printf("  No.             : %d\n", idx+1)
	fmt.Printf("  Nama            : %s\n", p.Nama)
	fmt.Printf("  Jumlah Pinjaman : Rp %.2f\n", p.JumlahPinjaman)
	fmt.Printf("  Tenor           : %d bulan\n", p.Tenor)
	fmt.Printf("  Skema Bunga     : %s\n", p.SchemaBunga)
	fmt.Printf("  Bunga Per Tahun : %.2f%%\n", p.Bunga)
	fmt.Printf("  Bayaran/Bulan   : Rp %.2f\n", p.BayaranPerbulan)
	fmt.Printf("  Total Bunga     : Rp %.2f\n", p.TotalBunga)
	fmt.Printf("  Total Bayar     : Rp %.2f\n", p.TotalBayar)
	fmt.Printf("  Sudah Bayar     : %d dari %d bulan\n", p.SudahBayar, p.Tenor)
	fmt.Printf("  Status          : %s\n", p.Status)
	garis2()
}

func TabelHeader() {
	garis1()
	fmt.Printf("| %-3s | %-15s | %-14s | %-5s | %-8s | %-9s | %-12s |\n",
		"NO","Nama","Pinjaman (Rp)","Tenor","Skema","Status","Bayaran/Bulan")
	garis1()
}

func TabelBaris(p Pinjaman,idx int) {
	fmt.Printf("| %-3d | %-15s | %14.0f | %5d | %-8s | %-9s | %12.0f |\n",
	idx + 1,p.Nama,p.JumlahPinjaman,p.Tenor,p.SchemaBunga,p.Status,p.BayaranPerbulan)
}

func allTable() { 
	if countPeminjam == 0 {
		fmt.Println("Belum ada data peminjam!")
		return
	}

	TabelHeader()
	for i := 0; i < countPeminjam;i++ {
		TabelBaris(dataPeminjam[i],i)	
	}
	garis1()
	fmt.Printf("	Total : %d peminjam\n",countPeminjam)
}

func enter() {
	fmt.Printf("\n Tekan enter untuk kembali ke ...")
	var d string
	fmt.Scan(&d)
}


// UNTUK MENDAFTARKAN ADMIN DENGAN MAKSIMAL ADMIN ADALAH 3
// func regisAdmin(data *[3]admin, i int) {
// 	fmt.Println()
// 	fmt.Printf("=%-10s Buat Akun =%-10s", " ", " ")
// 	fmt.Print("Nama : ")	
// 	fmt.Scan(&data[i].Username)
// 	fmt.Print("Password : ")
// 	fmt.Scan(&data[i].Password)
// 	fmt.Printf("=====================================")
// 	fmt.Print("YEAYYY!DATA KAMU BERHASIL DIBUAT!!")
// }

// CEK APAKAH ADMIN ADA ? 
// func cekDataAdmin(usn string, pw string) bool {
// 	for _,v := range dataAdmin {
// 		if usn == v.Username && pw == v.Password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func cekDataNasabah(usn string,pw string) bool {
// 	for _,v := range dataPeminjam {
// 		if usn == v.Nama && pw == v.Password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func loginAdmin() bool {
// 	var username, password string

// 	fmt.Print("Username : ")
// 	fmt.Scan(&username)
// 	fmt.Print("Password : ")
// 	fmt.Scan(&password)

// 	for _, v := range dataAdmin {
// 		if username == v.Username && password == v.Password {
// 			fmt.Print("Password dan username anda benar!")
// 			Admin(v.Username)
// 			return true
// 		}
// 	}
// 	return false
// }

// func loginPeminjam() bool {
// 	var usn,pwd string

// 	fmt.Print("Username : ")
// 	fmt.Scan(&usn)
// 	fmt.Print("Password : ")
// 	fmt.Scan(&pwd)

// 	for _,v := range akunNasabah {
// 		if v.username == usn && v.password == pwd {
// 			return true
// 		}
// 	}
// 	return false
// }

// func Register() {
// 	ClearScreen()
// 	var pilihan string
// 	garis1()
// 	fmt.Println("--------- DAFTAR AKUN BARU ---------")
// 	fmt.Println("1.Daftar Sebagai admin")
// 	fmt.Println("2.Register sebagai peminjam")
// 	garis1()
// 	fmt.Scan(&pilihan)

// 	var usn string
// 	fmt.Print("Harap masukan username anda : ")
// 	fmt.Scan(&usn)

// 	switch pilihan {
// 	case "1":
// 		found := false
// 		for _, v := range dataAdmin {
// 			if v.Username == usn {
// 				found = true
// 				break
// 			}
// 		}

// 		if found {
// 			fmt.Printf("\nUsername : %s sudah terdaftar!!", usn)
// 			loginAdmin()
// 		} else {
// 			var pw string
// 			fmt.Print("Harap masukan password anda : ")
// 			fmt.Scan(&pw)
// 			dataAdmin[counterAdmin] = admin{Username: usn, Password: pw}
// 			counterAdmin++
// 			fmt.Println("registrasi admin berhasil")
// 		}

// 	case "2":
// 		found := false
// 		for _, v := range dataPeminjam {
// 			if v.Nama == usn {
// 				found = true
// 				break
// 			}
// 		}
// 		if found {
// 			fmt.Printf("\nUsername : %s sudah terdaftar!!", usn)
// 			loginPeminjam()
// 		} else {
// 			nasabahBaru := Pinjaman{
// 				Nama:   usn,
// 				Status: "nggu",
// 				Tenor: 0,
// 				Bunga: 0,
// 				sudahBayar: 0,
// 				BayaranPerbulan: 0,
// 				sisaPinjam: 0,
// 				JumlahPinjaman: 0,
// 			}
// 			dataPeminjam[counterPeminjam] = nasabahBaru
// 			counterPeminjam++
// 			fmt.Println("Registrasi peminjaman berhasil")
// 		}
// 	}
// }

// INI RUMUS UNTUK KALKULASI BUNGA

//Bunga variabel anuitas
// C = p * r(1 + r)^n / ((1 + r)^n - 1)
//didapat dari https://www.bfi.co.id/id/blog/bunga-anuitas

func HitungAnuitas(pokok,bunga float64,tenor int) (bayaran,totalBunga,totalBayar float64) {
	r := (bunga / 100.0) / 12.0
	if r == 0 {
		bayaran = math.Round((pokok/float64(tenor)) * 100) / 100
		return
	}
	rn := math.Pow(1 + r, float64(tenor))
	bayaran = math.Round(pokok * (r * rn/(rn-1))*100)/100
	totalBayar = math.Round(bayaran*float64(tenor)*100) / 100
	totalBunga = math.Round((totalBayar - pokok)*100) / 100
	return
}

//hituhng bunga yang bersifat Flat, artinya bunga tetap dari pokok awal setiap bulan
//r		 = Bunga / (100 * 12)
//TotalBunga = pokok * r * Tenor
//Bayaranperbulan = (pokok + totalBunga) / tenor
//didapat dari : https://www.bfi.co.id/id/blog/bunga-flat-adalah-pengertian-kelebihan-dan-cara-menghitungnya

func bungaFlat(pokok,bunga float64,tenor int) (bayaran,totalBunga,totalBayar float64) {
	r := (bunga / 100.0) / 12.0
	totalBunga = math.Round(pokok*r*float64(tenor)*100) / 100
	totalBayar = math.Round((pokok + totalBunga)*100) / 100
	bayaran = math.Round((totalBayar/float64(tenor))*100) / 100
	return
}

func cetakAmortisasi(p Pinjaman) {
	garis1()
	fmt.Println("Tabel Amortisasi(Jadwal cicilan Anda)")
	garis1()
	fmt.Printf("  %-5s | %-14s | %-14s | %-14s | %-14s\n",
		"Bulan", "Bayaran (Rp)", "Pokok (Rp)", "Bunga (Rp)", "Sisa Pokok (Rp)")
	garis1()
}

func inputSkema() string {
	fmt.Println("Skema Bunga : ")
	fmt.Println("	1.FLAT		-Bunga kamu tetap nih dari pokok awal.")
	fmt.Println("	2.VARIABEL	-Bunga anuitas,menurun tiap bulab")
	fmt.Println("	Pilig(1/2)")
	var s string
	fmt.Scan(&s)
	if s == "1" {
		return "flat"
	}
	return "VARIABEL"
}
//CRUD data peminjam(Create,Read,Update,delete)
func tambahPeminjam() {
	ClearScreen()
	garis1()
	fmt.Println(" Tambah data peminjam")
	garis1()

	if  countPeminjam > 1000 {
		fmt.Println("Data penuh!!, maksimal 1000 peminjam.")
		return
	}

	var p Pinjaman
	
	fmt.Print(" Nama peminjam			:")
	fmt.Scan(&p.Nama)
	fmt.Print("	Jumlah pinjaman			:")
	fmt.Scan(&p.JumlahPinjaman)
	fmt.Print("	Tenor(bulan)			:")
	fmt.Scan(&p.Tenor)
	p.SchemaBunga = inputSkema()
	fmt.Print("	Bunga per tahun(%)		:")
	fmt.Scan(&p.Bunga)

	if p.JumlahPinjaman <= 0 || p.Tenor <= 0 || p.Bunga < 0 {
		fmt.Print("Mohon maaf,Jumlah pinjaman dan tenor harus lebih dari 0.")
		return
	}

	hitungDanSet(&p)

	p.Status = "MENUNGGU"
	p.sudahBayar = 0

	dataPeminjam[countPeminjam] = p
	countPeminjam++

	fmt.Println("\n ===HASIL KALKULASI===")
	DetailPeminjaman(p,countPeminjam - 1)

	var pil string
	fmt.Print("	Apakah anda ingin menampilkan tabel amortisasi anda? (Y/N): ")
	fmt.Scan(&pil)

	if pil == "Y" || pil == "y" {
		cetakAmortisasi(p)
	}

	fmt.Printf("\n Peminjaman \"%s\" berhasil ditambahkan!\n",p.Nama)
}

func ubahPeminjam() {
	ClearScreen()
	garis1()
	fmt.Println("	UBAH DATA PEMINJAM")
	garis1()

	if countPeminjam == 0 {
		fmt.Print("Data peminjam belum ada!")
		return
	}
	allTable()

	fmt.Printf("\n Masukan No. Peminjam yang akan diubah : ")
	var no int
	fmt.Scan(&no)

	if no < 1 || no > countPeminjam {
		fmt.Println("	Nomor tidak valid!")
		return
	}

	idx := no - 1
	p := &dataPeminjam[idx]

	fmt.Printf("\n Data saat ini:\n")
	DetailPeminjaman(*p,idx)

	fmt.Println("\n Yang ingin diubah :")
	fmt.Println("	1.Nama")
	fmt.Println("	2.Jumlah Pinjaman") //sudah hitung ulang secara otomatis bila diubah
	fmt.Println("	3.Tenor")
	fmt.Println("	4.Skema dan Bunga") // sudah hitung ulang secara otomatis
	fmt.Println("	5.Status pembayaran")
	fmt.Println("	6.Jumlah bulan yang sudah dibayar")
	fmt.Println("	Pilihan : ")

	var pil string
	fmt.Scan(&pil)

	switch pil {
	case "1":
		fmt.Print("	Nama baru: ")
		fmt.Scan(&p.Nama)
	case "2":
		fmt.Print("Jumlah pinjaman baru :")
		fmt.Scan(&p.JumlahPinjaman)
		hitungDanSet(p)
	case "3":
		fmt.Print("	Tenor baru (bulan): ")
		fmt.Scan(&p.Tenor)
		hitungDanSet()
	case "4":
		p.SchemaBunga = inputSkema()
		fmt.Print("	Bunga pertahun (%) baru : ")
		fmt.Scan(&p.Bunga)
		hitungDanSet()
	case "5":
		fmt.Print("	Status Pembayaran baru : ")
		fmt.Scan(&p )
	}
}

func hapusPeminjam() {
	ClearScreen()
	garis1()
	fmt.Println("HAPUS DATA PEMINJAM")
	garis1()

	if countPeminjam == 0 {
		fmt.Println("Belum ada data peminjam!")
		return
	}
	allTable()

	fmt.Print("\n Masukan no Peminjam yang ingin dihapus : ")
	var no int
	fmt.Scan(&no)

	if no < 1 || no > countPeminjam {
		fmt.Println("Nomor tidak valid!")
		return
	}
	idx := no -1
	nama := dataPeminjam[idx].Nama

	fmt.Printf( " Apakah anda yakin untuk menghapus \"%s\"? : ",nama)
	var konfirmasi string
	fmt.Scan(&konfirmasi)

	if konfirmasi != "Y" && konfirmasi != "y" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	for i := idx; i < countPeminjam - 1;i++ {
		dataPeminjam[i] = dataPeminjam[i + 1]
	}
	countPeminjam--
	fmt.Printf(" Peminjam \"%s\" berhasil dihapus!\n",nama)
}



func main() {
	fmt.Println("Program Sistem Pinjaman")
}
