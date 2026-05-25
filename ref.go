// package 

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"runtime"
// 	"strings"
// 	"time"
// )

// type admin struct {
// 	Username string
// 	Password string
// }

// type Pinjaman struct {
// 	Username string
// 	Password string
// 	Nama string
// 	JumlahPinjaman int
// 	Tenor int
// 	Status string
// 	Bunga float64
// 	BayaranPerbulan float64
// 	sisaPinjam int
// 	sudahBayar bool
// }

// var dataAdmin [3]admin
// var dataPeminjam [1000]Pinjaman

// func ClearScreen() {
// 	//clear screen ini digunakan untuk membersihkan output yang keluar di CLI,dengan clear screen kita bisa dapat lebih nyaman melihat CLI yang langsung bersih dan tidak ada sisa-sisa dari hasil execute code sebelumnya
// 	var cmd *exec.Cmd
// 	if runtime.GOOS == "windows" {
// 		cmd = exec.Command("cls")
// 	}else {
// 		cmd = exec.Command("clear")
// 	}
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

// func menuAdmin(user string) {
// 	now := time.Now() //fungsinya itu untuk menampilkan waktu real time pada CLI nanti...
// 	a := strings.Repeat("-",15) //untuk header
// 	fmt.Printf("TANGGAL: %s \n",now.Format("02-01-2006"))
// 	fmt.Printf("JAM: %s \n",now.Format("15:04:05"))
// 	fmt.Print(a)
// }

// func menu(user string) {
// 	now := time.Now()
// 	sama := strings.Repeat("-",85) //untuk membuat sama dengan sebanyak 30 kali( --------------- )
// 	fmt.Printf("TANGGAL: %s \n",now.Format("02-01-2006"))
// 	fmt.Printf("JAM: %s \n",now.Format("15:04:05"))
// 	fmt.Print(sama)
// 	fmt.Printf("| Selamat datang, %-60s |\n",user)
// 	fmt.Printf("| %-76s |\n","Apa yang bisa kami bantu hari ini ? ")
// 	fmt.Print(sama)
// }

// //UNTUK MENDAFTARKAN ADMIN DENGAN MAKSIMAL ADMIN ADALAH 3
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

// //CEK APAKAH ADMIN ADA ? 
// // func cekDataAdmin(usn string, pw string) bool {
// // 	for _,v := range dataAdmin {
// // 		if usn == v.Username && pw == v.Password {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // func cekDataNasabah(usn string,pw string) bool {
// // 	for _,v := range dataPeminjam {
// // 		if usn == v.Nama && pw == v.Password {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// func loginAdmin()(string,bool) {
// 	var username,password string

// 	fmt.Print("Username : ")
// 	fmt.Scan(&username)
// 	fmt.Print("Password : ")
// 	fmt.Scan(&password)

// 	for _,v := range dataAdmin {
// 		if username == v.Username && password == v.Password {
// 			fmt.Print("Password dan username anda benar!")
// 			menu(v.Username)

// 			return v.Username,true
// 		}
// 	}
// 	return "username atau password kamu salah!",false
// }

// func loginPeminjam() (string,bool) {
// 	var username,password string

// 	fmt.Print("Username : ")
// 	fmt.Scan(&username)
// 	fmt.Print("Password : ")
// 	fmt.Scan(&password)

// 	for _,v := range dataPeminjam {
// 		if username == v.Username && password == v.Password {
// 			fmt.Print("Password anda benar!")
// 			menu(v.Username)
// 			return v.Username,true
// 		}
// 	}
// 	return "username atau password anda salah!",false
// }
// func Register() {
// 	ClearScreen()
// 	var pilihan string
// 	fmt.Println(strings.Repeat("-",5),"DAFTAR AKUN BARU",strings.Repeat("-",5))
// 	fmt.Println("1.Daftar Sebagai admin")
// 	fmt.Println("2.Register sebagai peminjam")
// 	fmt.Scan(&pilihan)

// 	var usn string
// 	fmt.Print("Harap masukan username anda : ")
// 	fmt.Scan(&usn)
				
// 	switch pilihan {
// 	case "1":
// 		found := false
// 		for _,v := range dataAdmin {
// 			if v.Username == usn {
// 				found = true 
// 				break
// 			}
// 		}

// 		if found {
// 			fmt.Printf("\nUsername : %s sudah terdaftar!!",usn)
// 			time.Sleep(2 * time.Second)
// 			loginAdmin() //ini tuh langsung mengalihkan ke login admin ketika admin ternyata sudah permah mendaftar sebelumnya
// 		}else {
// 			var pw string
// 			fmt.Print("Harap masukan password anda : ")
// 			fmt.Scan(&pw)
// 			dataAdmin = append(dataAdmin,admin{Username: usn,Password: pw})
// 			fmt.Println("registrasi admin berhasil")
// 		}

// 	case "2":
// 		found := false
// 		for _,v := range dataPeminjam {
// 			if v.Nama == usn {
// 				found = true
// 				break
// 			}
// 		}
// 		if found {
// 			fmt.Print("\nUsername : %s sudah terdaftar!!",usn)
// 			time.Sleep(2 * time.Second)
// 			loginPeminjam()
// 		}else {
// 			nasabahBaru := Pinjaman {
// 				Nama : usn,
// 				Status : "Calon Peminjam",
// 				//Masih mikir wkwkkwk
// 			}
// 			dataAdmin = append(dataAdmin,nasabahBaru)
// 			fmt.Println("Registrasi peminjaman berhasil")
// 		}
// 	}
// }


// func HitungAnuitas(pokok,bunga float64,tenor int) (bayaran,totalBunga,totalBayar float64) {
// 	r := (bunga / 100.0) / 12.0
// 	if r == 0 {
// 		bayaran = math.Round((pokok/float64(tenor)) * 100) / 100
// 		return
// 	}
// 	rn := math.Pow(1 + r, float64(tenor))
// 	bayaran = math.Round(pokok * (r * rn/(rn-1))*100)/100
// 	totalBayar = math.Round()
// 	totalBunga = math.Round()
// }