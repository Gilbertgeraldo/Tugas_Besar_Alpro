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

func hitungDanSet(p *Pinjaman) {
	var bayaran,totalBunga,totalBayar float64
	if p.SchemaBunga == "FLAT" {
		bayaran,totalBunga,totalBayar = bungaFlat(p.JumlahPinjaman,p.Bunga,p.Tenor)
	}else {
		bayaran,totalBunga,totalBayar = HitungAnuitas(p.JumlahPinjaman,p.Bunga,p.Tenor)
	}
	p.BayaranPerbulan = bayaran
	p.TotalBunga = totalBunga
	p.TotalBayar = totalBayar
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
		hitungDanSet(p)
	case "4":
		p.SchemaBunga = inputSkema()
		fmt.Print("	Bunga pertahun (%) baru : ")
		fmt.Scan(&p.Bunga)
		if p.Bunga < 0 {
			fmt.Println(" Bunga tidak boleh negatif.")
		}
		hitungDanSet(p)
	case "5":
		fmt.Print("	Status Pembayaran baru : ")
		fmt.Println("Pilih status baru : ")
		fmt.Println("	1.MENUNGGU")
		fmt.Println("	2.AKTIF")
		fmt.Println("	3.LUNAS")
		fmt.Println("	4.MACET")
		fmt.Println("Pilihan :")
		var s string
		fmt.Scan(&s)
		switch s {
		case "1":
			p.Status = "MENUNGGU"
		case "2":
			p.Status = "AKTIF"
		case "3":
			p.Status = "LUNAS"
		case "4":
			p.Status = "MACET"
		default:
			enter()
			return
		}
	case "6":
		fmt.Printf("Kamu udah bayar berapa bulan nih : (max : %d)",p.Tenor)
		fmt.Scan(&p.sudahBayar)
		if p.sudahBayar < 0 || p.sudahBayar > p.Tenor {
			fmt.Printf("Kamu gabisa input nilai yang negatif ya untuk tenor...")
			enter()
			return
		}
		if p.sudahBayar == p.Tenor {
			p.Status = "LUNAS"
		}
	default:
		fmt.Println("Pilihan yang kamu masukan tidak valid nih,masukan pilihan yang benar ya..")
		enter()
		return
	}
	fmt.Printf("\n Data peminjam \"%s\" Berhasil diubah!\n",p.Nama)
	DetailPeminjaman(*p,idx)
	enter()
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

//PENCARIAN : Sequential & Binary Searc  

//ALGORITMA SORTING :

//INSERTION

func InsertionSortPinjaman(arr *[1000]Pinjaman,n int) {
	for i := 0;i < n;i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].JumlahPinjaman > key.JumlahPinjaman {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = key
	}
}

func InsertionSortTenor(arr *[1000]Pinjaman,n int) {
	for i := 1;i < n;i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Tenor > key.Tenor {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = key
	}
}

func InsertionSortNama(arr *[1000]Pinjaman,n int) {
	for i := 1;i < n;i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Nama > key.Nama {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = key
	}
}

//SELECTION
func SelectionSortPinjaman(arr *[1000]Pinjaman,n int) {
	for i := 1;i < n-1;i++ {
		mIdx := i
		for j := i + 1;j < n;j++ {
			if arr[j].JumlahPinjaman < arr[mIdx].JumlahPinjaman {
				mIdx = j
			}
		}
		arr[i],arr[mIdx] = arr[mIdx],arr[i]
	}
}

func SelectionSortTenor(arr *[1000]Pinjaman,n int) {
	for i := 1;i < n - 1;i++ {
		minIdx := i
		for j := i + 1;j < n; j++ {
			if arr[j].Tenor < arr[minIdx].Tenor {
				minIdx = j
			}
		}
		arr[i],arr[minIdx] = arr[minIdx],arr[i]
	}
}

func MenuSorting() {
	ClearScreen()
	garis1()
	fmt.Print(" URUTKAN DATA PEMINJAM")
	garis1()

	if countPeminjam == 0 {
		fmt.Println(" Belum ada data peminjam!")
		enter()
		return
	}

	fmt.Println(" Urutkan berdasarkan : ")
	fmt.Println("	1.Jumlah Pinjaman(Selection Sort)")
	fmt.Println("	2.Tenor (Selection Sort)")
	fmt.Println("	3.Jumlah Pinjaman(Insertion Sort)")
	fmt.Println("	4.Tenor (Selection Sort)")

	var pil string
	fmt.Scan(&pil)

	switch pil {
	case "1":
		SelectionSortPinjaman(&dataPeminjam,countPeminjam)
		fmt.Println("Data diurutkan berdasarkan Jumlah Pinjaaman (Selection Sort)")
	case "2":
		SelectionSortTenor(&dataPeminjam,countPeminjam)
		fmt.Println("Data diurutkan berdasarkan Tenor (Selection Sort)")
	case "3":
		InsertionSortPinjaman(&dataPeminjam,countPeminjam)
		fmt.Println("Data diurutkan berdasarkan Jumlah Pinjaman(Insertion Sort)")
	case "4":
		InsertionSortTenor(&dataPeminjam,countPeminjam)
		fmt.Println("Data diurutkan berdasarkan Tenor (Insertion Sort)")
	default:
		fmt.Println("Error!!,Pilihan tidak valid.")
		enter()
		return
	}
	allTable()
	enter()
}

func Laporan() {
	ClearScreen()
	garis1()
	fmt.Println(" LAPORAN SISTEM PEMINJAMAN")
	garis1()

	if countPeminjam == 0 {
		fmt.Println("Belum ada data peminjam")
		enter()
		return
	}
	var totalPokok,totalBayar,totalBunga float64
	var cMen,CAk,cLun,CMcet int

	for i := 0;i < countPeminjam;i++{
		p := dataPeminjam[i]
		totalPokok += p.JumlahPinjaman
		totalBayar += p.TotalBayar
		totalBunga += p.TotalBunga

		switch p.Status {
		case "MENUNGGU":
			cMen++
		case "AKTIF":
			CAk++
		case "LUNAS":
			cLun++
		case "MACET":
			CMcet++
		}
	}
	fmt.Printf("	Total peminjam				: %d orang\n",countPeminjam)
	fmt.Printf("	Total pokok pinjaman		: Rp %.2f\n",totalPokok)
	fmt.Printf("	Total Bunga					: Rp %.2f\n",totalBunga)
	fmt.Printf("	Total Nilai Bayar			: Rp %.2f\n",totalBayar)
	garis2()	
	fmt.Printf("	Total peminjam				: %d orang\n",countPeminjam)
	fmt.Printf("	Total peminjam				: %d orang\n",countPeminjam)
	fmt.Printf("	Total peminjam				: %d orang\n",countPeminjam)
	fmt.Printf("	Total peminjam				: %d orang\n",countPeminjam)
	fmt.Printf("	Total peminjam				: %d orang\n",countPeminjam)

}
func main() {
	fmt.Println("Program Sistem Pinjaman")
}
