package main
import (
	"fmt"
	"math"
)

type Pinjaman struct {
	Nama            string
	JumlahPinjaman  float64
	Tenor           int
	Status          string
	Bunga           float64
	BayaranPerbulan float64
	sisaPinjam      int
	sudahBayar      int
	SchemaBunga     string
	TotalBunga      float64
	TotalBayar      float64
	SudahBayar      int
}

var dataPeminjam [1000]Pinjaman
var countPeminjam int = 0

func ClearScreen() {
	//Untuk membersihkan terminal
	//didapat dari https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go/22892171#22892171
	fmt.Print("\033[H\033[2J")
}

func garis1() {
	fmt.Println("=================================================================================")
}

func garis2() {
	fmt.Println("-------------------------------------------------------------------------------")
}

func DetailPeminjaman(p Pinjaman, idx int) {
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
		"NO", "Nama", "Pinjaman (Rp)", "Tenor", "Skema", "Status", "Bayaran/Bulan")
	garis1()
}

func TabelBaris(p Pinjaman, idx int) {
	fmt.Printf("| %-3d | %-15s | %14.0f | %5d | %-8s | %-9s | %12.0f |\n",
		idx+1, p.Nama, p.JumlahPinjaman, p.Tenor, p.SchemaBunga, p.Status, p.BayaranPerbulan)
}

func allTable() {
	if countPeminjam == 0 {
		fmt.Println("  Belum ada data peminjam!")
		return
	}
	TabelHeader()
	for i := 0; i < countPeminjam; i++ {
		TabelBaris(dataPeminjam[i], i)
	}
	garis1()
	fmt.Printf("  Total : %d peminjam\n", countPeminjam)
}

func enter() {
	fmt.Println("\nTekan ENTER untuk kembali ke Menu Utama...")
	fmt.Scanln()
	fmt.Scanln()
}

// INI RUMUS UNTUK KALKULASI BUNGA
// Bunga variabel anuitas
// C = p * r(1 + r)^n / ((1 + r)^n - 1)
// didapat dari https://www.bfi.co.id/id/blog/bunga-anuitas

func HitungAnuitas(pokok, bunga float64, tenor int) (bayaran, totalBunga, totalBayar float64) {
	r := (bunga / 100.0) / 12.0
	if r == 0 {
		bayaran = math.Round((pokok/float64(tenor))*100) / 100
		return
	}
	rn := math.Pow(1+r, float64(tenor))
	bayaran = math.Round(pokok*(r*rn/(rn-1))*100) / 100
	totalBayar = math.Round(bayaran*float64(tenor)*100) / 100
	totalBunga = math.Round((totalBayar-pokok)*100) / 100
	return
}

// Hitung bunga yang bersifat Flat, artinya bunga tetap dari pokok awal setiap bulan
// r           = Bunga / (100 * 12)
// TotalBunga  = pokok * r * Tenor
// Bayaran/bln = (pokok + totalBunga) / tenor
// didapat dari : https://www.bfi.co.id/id/blog/bunga-flat-adalah-pengertian-kelebihan-dan-cara-menghitungnya

func bungaFlat(pokok, bunga float64, tenor int) (bayaran, totalBunga, totalBayar float64) {
	r := (bunga / 100.0) / 12.0
	totalBunga = math.Round(pokok*r*float64(tenor)*100) / 100
	totalBayar = math.Round((pokok+totalBunga)*100) / 100
	bayaran = math.Round((totalBayar/float64(tenor))*100) / 100
	return
}

func hitungDanSet(p *Pinjaman) {
	var bayaran, totalBunga, totalBayar float64
	if p.SchemaBunga == "FLAT" {
		bayaran, totalBunga, totalBayar = bungaFlat(p.JumlahPinjaman, p.Bunga, p.Tenor)
	} else {
		bayaran, totalBunga, totalBayar = HitungAnuitas(p.JumlahPinjaman, p.Bunga, p.Tenor)
	}
	p.BayaranPerbulan = bayaran
	p.TotalBunga = totalBunga
	p.TotalBayar = totalBayar
}

func cetakAmortisasi(p Pinjaman) {
	garis1()
	fmt.Println("  Tabel Amortisasi (Jadwal cicilan Anda)")
	garis1()
	fmt.Printf("  %-5s | %-14s | %-14s | %-14s | %-14s\n",
		"Bulan", "Bayaran (Rp)", "Pokok (Rp)", "Bunga (Rp)", "Sisa Pokok (Rp)")
	garis1()
}

func inputSkema() string {
	fmt.Println("  Skema Bunga :")
	fmt.Println("    1. FLAT     - Bunga kamu tetap dari pokok awal.")
	fmt.Println("    2. VARIABEL - Bunga anuitas, menurun tiap bulan.")
	fmt.Print("  Pilih (1/2) : ")
	var s string
	fmt.Scan(&s)
	if s == "1" {
		return "FLAT"
	}
	return "VARIABEL"
}

// =====================================================================
// CRUD DATA PEMINJAM
// =====================================================================

func tambahPeminjam() {
	ClearScreen()
	garis1()
	fmt.Println(" Tambah Data Peminjam")
	garis1()

	if countPeminjam > 1000 {
		fmt.Println("  Data penuh! Maksimal 1000 peminjam.")
		return
	}

	var p Pinjaman

	fmt.Print("  Nama peminjam        : ")
	fmt.Scan(&p.Nama)
	fmt.Print("  Jumlah pinjaman (Rp) : ")
	fmt.Scan(&p.JumlahPinjaman)
	fmt.Print("  Tenor (bulan)        : ")
	fmt.Scan(&p.Tenor)
	p.SchemaBunga = inputSkema()
	fmt.Print("  Bunga per tahun (%)  : ")
	fmt.Scan(&p.Bunga)

	if p.JumlahPinjaman <= 0 || p.Tenor <= 0 || p.Bunga < 0 {
		fmt.Println("  Mohon maaf, jumlah pinjaman dan tenor harus lebih dari 0.")
		return
	}

	hitungDanSet(&p)
	p.Status = "MENUNGGU"
	p.sudahBayar = 0

	dataPeminjam[countPeminjam] = p
	countPeminjam++

	fmt.Println("\n  === HASIL KALKULASI ===")
	DetailPeminjaman(p, countPeminjam-1)

	var pil string
	fmt.Print("  Apakah anda ingin menampilkan tabel amortisasi? (Y/N) : ")
	fmt.Scan(&pil)
	if pil == "Y" || pil == "y" {
		cetakAmortisasi(p)
	}

	fmt.Printf("\n  Peminjaman \"%s\" berhasil ditambahkan!\n", p.Nama)
}

func ubahPeminjam() {
	ClearScreen()
	garis1()
	fmt.Println("  UBAH DATA PEMINJAM")
	garis1()

	if countPeminjam == 0 {
		fmt.Println("  Data peminjam belum ada!")
		return
	}
	allTable()

	fmt.Print("\n  Masukkan No. peminjam yang akan diubah : ")
	var no int
	fmt.Scan(&no)

	if no < 1 || no > countPeminjam {
		fmt.Println("  Nomor tidak valid!")
		return
	}

	idx := no - 1
	p := &dataPeminjam[idx]

	fmt.Println("\n  Data saat ini:")
	DetailPeminjaman(*p, idx)

	fmt.Println("\n  Yang ingin diubah :")
	fmt.Println("    1. Nama")
	fmt.Println("    2. Jumlah Pinjaman") // sudah hitung ulang secara otomatis bila diubah
	fmt.Println("    3. Tenor")
	fmt.Println("    4. Skema dan Bunga") // sudah hitung ulang secara otomatis
	fmt.Println("    5. Status Pembayaran")
	fmt.Println("    6. Jumlah bulan yang sudah dibayar")
	fmt.Print("  Pilihan : ")

	var pil string
	fmt.Scan(&pil)

	switch pil {
	case "1":
		fmt.Print("  Nama baru : ")
		fmt.Scan(&p.Nama)
	case "2":
		fmt.Print("  Jumlah pinjaman baru (Rp) : ")
		fmt.Scan(&p.JumlahPinjaman)
		hitungDanSet(p)
	case "3":
		fmt.Print("  Tenor baru (bulan) : ")
		fmt.Scan(&p.Tenor)
		hitungDanSet(p)
	case "4":
		p.SchemaBunga = inputSkema()
		fmt.Print("  Bunga per tahun (%) baru : ")
		fmt.Scan(&p.Bunga)
		if p.Bunga < 0 {
			fmt.Println("  Bunga tidak boleh negatif.")
		}
		hitungDanSet(p)
	case "5":
		fmt.Print("  Status Pembayaran baru : ")
		fmt.Scan(&p)
	}

	fmt.Printf("\n  Data peminjam \"%s\" berhasil diubah!\n", p.Nama)
	DetailPeminjaman(*p, idx)
	enter()
}

func hapusPeminjam() {
	ClearScreen()
	garis1()
	fmt.Println("  HAPUS DATA PEMINJAM")
	garis1()

	if countPeminjam == 0 {
		fmt.Println("  Belum ada data peminjam!")
		return
	}
	allTable()

	fmt.Print("\n  Masukkan No. peminjam yang ingin dihapus : ")
	var no int
	fmt.Scan(&no)

	if no < 1 || no > countPeminjam {
		fmt.Println("  Nomor tidak valid!")
		return
	}

	idx := no - 1
	nama := dataPeminjam[idx].Nama

	fmt.Printf("  Apakah anda yakin untuk menghapus \"%s\"? (Y/N) : ", nama)
	var konfirmasi string
	fmt.Scan(&konfirmasi)

	if konfirmasi != "Y" && konfirmasi != "y" {
		fmt.Println("  Penghapusan dibatalkan.")
		return
	}

	for i := idx; i < countPeminjam-1; i++ {
		dataPeminjam[i] = dataPeminjam[i+1]
	}
	countPeminjam--
	fmt.Printf("  Peminjam \"%s\" berhasil dihapus!\n", nama)
}

// =====================================================================
// PENCARIAN
// =====================================================================

func SequentialSearch() {
	if countPeminjam == 0 {
		fmt.Println("  Belum ada data peminjam!")
		return
	}

	var keyword string
	fmt.Print("  Masukkan nama peminjam yang dicari : ")
	fmt.Scan(&keyword)

	keyLower := ""
	for i := 0; i < len(keyword); i++ {
		c := keyword[i]
		if c >= 'A' && c <= 'Z' {
			c = c + 32
		}
		keyLower += string(c)
	}

	found := false
	fmt.Println()
	garis1()
	fmt.Println("  HASIL SEQUENTIAL SEARCH")
	garis1()

	for i := 0; i < countPeminjam; i++ {
		namaLower := ""
		for j := 0; j < len(dataPeminjam[i].Nama); j++ {
			c := dataPeminjam[i].Nama[j]
			if c >= 'A' && c <= 'Z' {
				c = c + 32
			}
			namaLower += string(c)
		}

		match := false
		if len(namaLower) == len(keyLower) {
			match = true
			for k := 0; k < len(namaLower); k++ {
				if namaLower[k] != keyLower[k] {
					match = false
					break
				}
			}
		}

		if match {
			DetailPeminjaman(dataPeminjam[i], i)
			found = true
		}
	}

	if !found {
		fmt.Printf("  Peminjam dengan nama \"%s\" tidak ditemukan.\n", keyword)
	}
}

func BinarySearch() {
	if countPeminjam == 0 {
		fmt.Println("  Belum ada data peminjam!")
		return
	}

	var temp [1000]Pinjaman
	for i := 0; i < countPeminjam; i++ {
		temp[i] = dataPeminjam[i]
	}

	for i := 1; i < countPeminjam; i++ {
		key := temp[i]

		keyNamaLower := ""
		for x := 0; x < len(key.Nama); x++ {
			c := key.Nama[x]
			if c >= 'A' && c <= 'Z' {
				c = c + 32
			}
			keyNamaLower += string(c)
		}

		j := i - 1
		for j >= 0 {
			namaJLower := ""
			for x := 0; x < len(temp[j].Nama); x++ {
				c := temp[j].Nama[x]
				if c >= 'A' && c <= 'Z' {
					c = c + 32
				}
				namaJLower += string(c)
			}

			isGreater := false
			minLen := len(namaJLower)
			if len(keyNamaLower) < minLen {
				minLen = len(keyNamaLower)
			}
			for k := 0; k < minLen; k++ {
				if namaJLower[k] > keyNamaLower[k] {
					isGreater = true
					break
				} else if namaJLower[k] < keyNamaLower[k] {
					break
				}
			}
			if !isGreater && len(namaJLower) > len(keyNamaLower) {
				isGreater = true
			}

			if isGreater {
				temp[j+1] = temp[j]
				j--
			} else {
				break
			}
		}
		temp[j+1] = key
	}

	var keyword string
	fmt.Print("  Masukkan nama peminjam yang dicari : ")
	fmt.Scan(&keyword)

	keyLower := ""
	for i := 0; i < len(keyword); i++ {
		c := keyword[i]
		if c >= 'A' && c <= 'Z' {
			c = c + 32
		}
		keyLower += string(c)
	}

	low := 0
	high := countPeminjam - 1
	result := -1

	for low <= high {
		mid := int(math.Floor(float64(low+high) / 2))

		midNamaLower := ""
		for x := 0; x < len(temp[mid].Nama); x++ {
			c := temp[mid].Nama[x]
			if c >= 'A' && c <= 'Z' {
				c = c + 32
			}
			midNamaLower += string(c)
		}

		cmp := 0
		minLen := len(midNamaLower)
		if len(keyLower) < minLen {
			minLen = len(keyLower)
		}
		for k := 0; k < minLen; k++ {
			if midNamaLower[k] < keyLower[k] {
				cmp = -1
				break
			} else if midNamaLower[k] > keyLower[k] {
				cmp = 1
				break
			}
		}
		if cmp == 0 {
			if len(midNamaLower) < len(keyLower) {
				cmp = -1
			} else if len(midNamaLower) > len(keyLower) {
				cmp = 1
			}
		}

		if cmp == 0 {
			result = mid
			break
		} else if cmp < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println()
	garis1()
	fmt.Println("  HASIL BINARY SEARCH")
	garis1()

	if result != -1 {
		DetailPeminjaman(temp[result], result)
	} else {
		fmt.Printf("  Peminjam dengan nama \"%s\" tidak ditemukan.\n", keyword)
	}
}

// =====================================================================
// PENGURUTAN
// =====================================================================

func InsertionSortPinjaman(arr *[1000]Pinjaman, n int) {
	for i := 0; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].JumlahPinjaman > key.JumlahPinjaman {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func InsertionSortTenor(arr *[1000]Pinjaman, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Tenor > key.Tenor {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func InsertionSortNama(arr *[1000]Pinjaman, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Nama > key.Nama {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func SelectionSortPinjaman(arr *[1000]Pinjaman, n int) {
	for i := 1; i < n-1; i++ {
		mIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].JumlahPinjaman < arr[mIdx].JumlahPinjaman {
				mIdx = j
			}
		}
		arr[i], arr[mIdx] = arr[mIdx], arr[i]
	}
}

func SelectionSortTenor(arr *[1000]Pinjaman, n int) {
	for i := 1; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].Tenor < arr[minIdx].Tenor {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// =====================================================================
// MENU SEARCHING & SORTING
// =====================================================================

func MenuSearching() {
	ClearScreen()
	garis1()
	fmt.Println("  CARI DATA PEMINJAM")
	garis1()

	if countPeminjam == 0 {
		fmt.Println("  Belum ada data peminjam!")
		enter()
		return
	}

	fmt.Println("  Pilih metode pencarian :")
	fmt.Println("    1. Sequential Search")
	fmt.Println("    2. Binary Search")
	fmt.Print("  Pilihan : ")

	var pil string
	fmt.Scan(&pil)

	switch pil {
	case "1":
		SequentialSearch()
	case "2":
		BinarySearch()
	default:
		fmt.Println("  Error! Pilihan tidak valid.")
	}
	enter()
}

func MenuSorting() {
	ClearScreen()
	garis1()
	fmt.Println("  URUTKAN DATA PEMINJAM")
	garis1()

	if countPeminjam == 0 {
		fmt.Println("  Belum ada data peminjam!")
		enter()
		return
	}

	fmt.Println("  Urutkan berdasarkan :")
	fmt.Println("    1. Jumlah Pinjaman (Selection Sort)")
	fmt.Println("    2. Tenor           (Selection Sort)")
	fmt.Println("    3. Jumlah Pinjaman (Insertion Sort)")
	fmt.Println("    4. Tenor           (Insertion Sort)")
	fmt.Print("  Pilihan : ")

	var pil string
	fmt.Scan(&pil)

	switch pil {
	case "1":
		SelectionSortPinjaman(&dataPeminjam, countPeminjam)
		fmt.Println("  Data diurutkan berdasarkan Jumlah Pinjaman (Selection Sort).")
	case "2":
		SelectionSortTenor(&dataPeminjam, countPeminjam)
		fmt.Println("  Data diurutkan berdasarkan Tenor (Selection Sort).")
	case "3":
		InsertionSortPinjaman(&dataPeminjam, countPeminjam)
		fmt.Println("  Data diurutkan berdasarkan Jumlah Pinjaman (Insertion Sort).")
	case "4":
		InsertionSortTenor(&dataPeminjam, countPeminjam)
		fmt.Println("  Data diurutkan berdasarkan Tenor (Insertion Sort).")
	default:
		fmt.Println("  Error! Pilihan tidak valid.")
		enter()
		return
	}
	allTable()
	enter()
}

// =====================================================================
// LAPORAN
// =====================================================================

func Laporan() {
	ClearScreen()
	garis1()
	fmt.Println("  LAPORAN SISTEM PEMINJAMAN")
	garis1()

	if countPeminjam == 0 {
		fmt.Println("  Belum ada data peminjam.")
		enter()
		return
	}

	var totalPokok, totalBayar, totalBunga float64
	var cMen, cAk, cLun, cMcet int

	for i := 0; i < countPeminjam; i++ {
		p := dataPeminjam[i]
		totalPokok += p.JumlahPinjaman
		totalBayar += p.TotalBayar
		totalBunga += p.TotalBunga

		switch p.Status {
		case "MENUNGGU":
			cMen++
		case "AKTIF":
			cAk++
		case "LUNAS":
			cLun++
		case "MACET":
			cMcet++
		}
	}

	fmt.Printf("  Total Peminjam       : %d orang\n", countPeminjam)
	fmt.Printf("  Total Pokok Pinjaman : Rp %.2f\n", totalPokok)
	fmt.Printf("  Total Bunga          : Rp %.2f\n", totalBunga)
	fmt.Printf("  Total Nilai Bayar    : Rp %.2f\n", totalBayar)
	garis2()
	fmt.Println("  Status Pembayaran :")
	fmt.Printf("    MENUNGGU : %d peminjam\n", cMen)
	fmt.Printf("    AKTIF    : %d peminjam\n", cAk)
	fmt.Printf("    LUNAS    : %d peminjam\n", cLun)
	fmt.Printf("    MACET    : %d peminjam\n", cMcet)
	garis2()
	fmt.Println("\n  Daftar Semua Peminjam :")
	allTable()
	enter()
}

// =====================================================================
// JUDUL
// =====================================================================

func cetakJudul() {
	judul := []string{
		`________          __                                   ______   __                               `,
		`/        |        /  |                                 /      \ /  |                              `,
		`$$$$$$$$/__    __ $$ |____    ______    _______       /$$$$$$  |$$ |  ______    ______    ______  `,
		`   $$ | /  |  /  |$$      \  /      \  /       |      $$ |__$$ |$$ | /      \  /      \  /      \`,
		`   $$ | $$ |  $$ |$$$$$$$  |/$$$$$$  |/$$$$$$$/       $$    $$ |$$ |/$$$$$$  |/$$$$$$  |/$$$$$$  |`,
		`   $$ | $$ |  $$ |$$ |  $$ |$$    $$ |$$      \       $$$$$$$$ |$$ |$$ |  $$ |$$ |  $$/ $$ |  $$ |`,
		`   $$ | $$ \__$$ |$$ |__$$ |$$$$$$$$/  $$$$$$  |      $$ |  $$ |$$ |$$ |__$$ |$$ |      $$ \__$$ |`,
		`   $$ | $$    $$/ $$    $$/ $$       |/     $$/       $$ |  $$ |$$ |$$    $$/ $$ |      $$    $$/  `,
		`   $$/   $$$$$$/  $$$$$$$/   $$$$$$$/ $$$$$$$/        $$/   $$/ $$/ $$$$$$$/  $$/        $$$$$$/   `,
		`                                                                    $$ |                            `,
		`                                                                    $$ |                            `,
		`                                                                    $$/                             `,
	}

	warna := []string{
		"\033[31m", // Merah
		"\033[33m", // Kuning
		"\033[32m", // Hijau
		"\033[36m", // Cyan
		"\033[34m", // Biru
		"\033[35m", // Ungu
		"\033[91m", // Merah terang
		"\033[92m", // Hijau terang
		"\033[93m", // Kuning terang
		"\033[94m", // Biru terang
		"\033[95m", // Ungu terang
		"\033[96m", // Cyan terang
	}

	for i, baris := range judul {
		fmt.Println(warna[i%len(warna)] + baris + "\033[0m")
	}
}

// =====================================================================
// MAIN
// =====================================================================

func main() {
	for {
		ClearScreen()
		cetakJudul()
		garis2()
		fmt.Println("  Anggota :")
		fmt.Println("  1. Gilbert Geraldo (103052500054)")
		fmt.Println("  2. Jafar Shiddiq   (103052500002)")
		garis2()
		fmt.Println("  1. Tambah Peminjam")
		fmt.Println("  2. Ubah Data Peminjam")
		fmt.Println("  3. Hapus Peminjam")
		fmt.Println("  4. Lihat Semua Pinjaman")
		fmt.Println("  5. Cari Peminjam")
		fmt.Println("  6. Urutkan Peminjaman")
		fmt.Println("  7. Laporan")
		fmt.Println("  0. Keluar")
		garis1()
		fmt.Print("  Silahkan masukkan pilihan anda : ")

		var p string
		fmt.Scan(&p)

		switch p {
		case "1":
			tambahPeminjam()
		case "2":
			ubahPeminjam()
		case "3":
			hapusPeminjam()
		case "4":
			ClearScreen()
			garis1()
			fmt.Println("  DAFTAR SEMUA PEMINJAM")
			garis1()
			allTable()
			enter()
		case "5":
			MenuSearching()
		case "6":
			MenuSorting()
		case "7":
			Laporan()
		case "0":
			fmt.Println("  Terima kasih... Sampai Jumpa!")
			return
		default:
			fmt.Println("  Pilihan tidak valid.")
			enter()
		}
	}
}