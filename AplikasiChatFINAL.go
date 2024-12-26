package main

import (
	"fmt"
)

const NMAX int = 100

type dataUser struct {
	usn, pass string
	personal  [NMAX]dataPersonal
	group     [NMAX]dataGroup
	nPl, nG   int
}

type dataPersonal struct {
	namaPersonal string
	personalChat [NMAX]string
	nPC          int
}

type dataGroup struct {
	namaGroup string
	anggota   [NMAX]string
	groupChat [NMAX]string
	nGC       int
	nA        int
}

type tabUser [NMAX]dataUser

func main() {
	var registrasi, dataPengguna tabUser
	var nRegistrasi, nPengguna int
	menuWelcome()
	menuOption(&registrasi, &dataPengguna, &nRegistrasi, &nPengguna)
}

func menuWelcome() { //ini untuk menu home ketika aplikasi baru dibuka / dijalankan
	fmt.Println("_________________________")
	fmt.Println("======== WELCOME ========")
	fmt.Println("|      ~Chat Wave~      |")
	fmt.Println("|   Aplikasi Chatting   |")
	fmt.Println("|   2024 / (IF-47-09)   |")
	fmt.Println("_________________________")
}

func menuOption(R, dP *tabUser, nR, nP *int) { //ini menu untuk memilih apakah pengguna ini menjadi user / admin
	var pilih int
	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println("|      OPTION MENU      |")
	fmt.Println("-------------------------")
	fmt.Println("1. User")
	fmt.Println("2. Admin")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------")
	fmt.Println("Pilih 1/2/3?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 && pilih != 3 { //loop agar pilihan tetap sesuai dengan yang diinginkan
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 3 {
		keluar()
	}
	if pilih == 1 {
		menuUser(R, dP, *nR, *nP)
	} else if pilih == 2 {
		menuLoginAdmin(R, dP, *nR, *nP)
	}
}

func keluar() { //func untuk back / keluar dari aplikasi
	fmt.Println("Anda telah keluar dari aplikasi.")
}

func menuUser(R, dP *tabUser, nR, nP int) { //menu user
	var pilih int
	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println("|       USER MENU       |")
	fmt.Println("-------------------------")
	fmt.Println("1. Registration")
	fmt.Println("2. Login")
	fmt.Println("3. Back")
	fmt.Println("-------------------------")
	fmt.Println("Pilih 1/2/3?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 && pilih != 3 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		menuRegistration(R, dP, &nR, &nP)
	} else if pilih == 2 {
		menuLogin(R, dP, nR, nP)
	} else {
		menuOption(R, dP, &nR, &nP)
	}
}

func menuRegistration(R, dP *tabUser, nR, nP *int) { //ini untuk registrasi akun sebagai user
	fmt.Println()
	fmt.Println("*Silahkan lengkapi data Anda untuk melakukan proses registrasi.")
	fmt.Println("*Isi '-' pada Username atau Password untuk kembali ke USER MENU")
	fmt.Println("-------------------------")
	fmt.Println("|   REGISTRATION MENU   |")
	fmt.Println("-------------------------")
	fmt.Print("Username: ")
	fmt.Scan(&R[*nR].usn)
	if R[*nR].usn == "-" {
		fmt.Println("*Registrasi Anda tidak berhasil.")
		menuUser(R, dP, *nR, *nP)
	}
	fmt.Print("Password: ")
	fmt.Scan(&R[*nR].pass)
	if R[*nR].pass == "-" {
		fmt.Println("*Registrasi Anda tidak berhasil.")
		menuUser(R, dP, *nR, *nP)
	}
	fmt.Println("-------------------------")
	fmt.Println("*Registrasi Anda sudah berhasil, mohon untuk menunggu verifikasi data dari admin untuk dapat Login.")
	fmt.Println("*Jika anda tidak dapat Login dalam waktu 24 jam, maka verifikasi anda ditolak.")
	*nR++
	menuUser(R, dP, *nR, *nP)
}

func menuLogin(R, dP *tabUser, nR, nP int) { //ini menu login jika akun sudah di acc oleh admin
	var namaUser, passUser string
	var ketemu bool = false
	var idxUser1 int
	fmt.Println()
	fmt.Println("*Isi '-' pada Username atau Password untuk kembali ke USER MENU")
	fmt.Println("--------------------------")
	fmt.Println("|       USER LOGIN       |")
	fmt.Println("--------------------------")
	fmt.Print("Username: ")
	fmt.Scan(&namaUser)
	if namaUser == "-" {
		fmt.Println("*Kembali ke USER MENU")
		menuUser(R, dP, nR, nP)
	}
	fmt.Print("Password: ")
	fmt.Scan(&passUser)
	if passUser == "-" {
		fmt.Println("*Kembali ke USER MENU")
		menuUser(R, dP, nR, nP)
	}
	fmt.Println("--------------------------")
	i := 0
	for i < nP {
		if namaUser == dP[i].usn && passUser == dP[i].pass {
			ketemu = true
			idxUser1 = i
		}
		i++
	}
	if ketemu {
		fmt.Print("*Login berhasil.")
		fmt.Scanln()
		fmt.Scanln()
		menuMain(R, dP, nR, nP, idxUser1)
	} else {
		fmt.Print("*Login gagal, username tidak ditemukan atau password tidak cocok.")
		fmt.Scanln()
		fmt.Scanln()
		menuLogin(R, dP, nR, nP)
	}
}

func menuMain(R, dP *tabUser, nR, nP, idxUser1 int) { //ini menu ketika user berhasil login
	var pilih int
	fmt.Println()
	fmt.Println("-------------------------------------")
	fmt.Printf("Anda berhasil Login sebagai, %s! \n", dP[idxUser1].usn)
	fmt.Println("-------------------------------------")
	fmt.Println("1. Personal Chat")
	fmt.Println("2. Group Chat")
	fmt.Println("3. Ubah Password")
	fmt.Println("4. Logout")
	fmt.Println("-------------------------------------")
	fmt.Println("Pilih 1/2/3/4?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 && pilih != 3 && pilih != 4 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		menuPersonalChat(R, dP, nR, nP, idxUser1)
	} else if pilih == 2 {
		menuGroupChat(R, dP, nR, nP, idxUser1)
	} else if pilih == 3 {
		menuEditPass(R, dP, nR, nP, idxUser1)
	} else {
		menuUser(R, dP, nR, nP)
	}
}

func menuEditPass(R, dP *tabUser, nR, nP, idxUser1 int) { //menu untuk ubah password
	var password string
	var pilih int
	fmt.Println()
	fmt.Println("*Ketik '-' untuk kembali ke MAIN MENU")
	fmt.Println("-------------------------")
	fmt.Println("|       EDIT DATA       |")
	fmt.Println("-------------------------")
	fmt.Print("Password baru: ")
	fmt.Scan(&password)
	if password == "-" {
		menuMain(R, dP, nR, nP, idxUser1)
	}
	fmt.Println("-------------------------")
	fmt.Println("Simpan Password baru?")
	fmt.Println("1. Ya  2.Tidak")
	fmt.Println("Pilih 1/2?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 && pilih != 111 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		dP[idxUser1].pass = password
		fmt.Println("Password berhasil diubah.")
		menuMain(R, dP, nR, nP, idxUser1)
	} else if pilih == 2 {
		fmt.Println("Password gagal diubah.")
		menuMain(R, dP, nR, nP, idxUser1)
	}
}

func menuPersonalChat(R, dP *tabUser, nR, nP, idxUser1 int) { //menu opsi chat personal (kontak teman)
	var pilih, p1 int
	fmt.Println()
	fmt.Println("*Ketik '111' untuk kembali ke MAIN MENU")
	fmt.Println("*Ketik '0' untuk menambah Personal Chat")
	fmt.Println("-------------------------")
	fmt.Println("|     PERSONAL CHAT     |")
	fmt.Println("-------------------------")
	if dP[idxUser1].nPl == 0 {
		fmt.Println("          empty          ")
	}
	urutPersonal(dP, idxUser1)
	for i := 0; i < dP[idxUser1].nPl; i++ {
		fmt.Printf("%d. %s \n", i+1, dP[idxUser1].personal[i].namaPersonal)
	}
	fmt.Println("-------------------------")
	fmt.Print("Pilih 0")
	for i := 0; i < dP[idxUser1].nPl; i++ {
		fmt.Printf("/%d", i+1)
	}
	fmt.Print("?")
	fmt.Println()
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 0 && pilih != 111 && pilih > dP[idxUser1].nPl {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 111 {
		menuMain(R, dP, nR, nP, idxUser1)
	} else if pilih == 0 {
		addPC(R, dP, nR, nP, idxUser1)
	} else if pilih <= dP[idxUser1].nPl {
		p1 = pilih - 1
		personalConversation(R, dP, nR, nP, idxUser1, p1)
	}
}

func urutPersonal(dP *tabUser, idxUser1 int) { //ini untuk mengurutkan nama kontak dari abjad terkecil (ascending)
	var temp dataPersonal
	var pass, idx, i int
	pass = 1
	for pass <= dP[idxUser1].nPl-1 {
		idx = pass - 1
		i = pass
		for i < dP[idxUser1].nPl {
			if dP[idxUser1].personal[idx].namaPersonal > dP[idxUser1].personal[i].namaPersonal {
				idx = i
			}
			i++
		}
		temp = dP[idxUser1].personal[pass-1]
		dP[idxUser1].personal[pass-1] = dP[idxUser1].personal[idx]
		dP[idxUser1].personal[idx] = temp
		pass++
	}
}

func addPC(R, dP *tabUser, nR, nP, idxUser1 int) { //untuk add teman (personal chat)
	var cariUsn string
	var indeks int
	var nP1, nP2 int
	fmt.Println()
	fmt.Println("*Ketik username yang ingin ditambahkan.")
	fmt.Println("*Ketik '-' untuk kembali ke PERSONAL CHAT")
	fmt.Println("-------------------------")
	fmt.Println("|   ADD PERSONAL CHAT   |")
	fmt.Println("-------------------------")
	fmt.Print("Username: ")
	fmt.Scan(&cariUsn)
	if cariUsn == "-" {
		menuPersonalChat(R, dP, nR, nP, idxUser1)
	}
	indeks = searchUsn(*dP, nP, cariUsn)
	if indeks == -1 {
		fmt.Println("-------------------------")
		fmt.Printf("*Username '%s' tidak ada. \n", cariUsn)
		addPC(R, dP, nR, nP, idxUser1)
	} else {
		var double bool = false
		if cariUsn == dP[idxUser1].usn {
			double = true
		}
		for j := 0; j < dP[idxUser1].nPl; j++ {
			if dP[idxUser1].personal[j].namaPersonal == cariUsn {
				double = true
			}
		}
		if double {
			fmt.Println("-------------------------")
			fmt.Printf("*Username '%s' gagal ditambahkan \n", cariUsn)
			addPC(R, dP, nR, nP, idxUser1)
		} else if !double {
			var pilih int
			fmt.Println("-------------------------")
			fmt.Printf("Username '%s' ditemukan. \n", cariUsn)
			fmt.Println("Tambahkan personal?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Println("Pilih 1/2?")
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&pilih)
			for pilih != 1 && pilih != 2 {
				fmt.Println()
				fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
				fmt.Print("Pilihan Anda: ")
				fmt.Scan(&pilih)
			}
			if pilih == 1 {
				nP1 = dP[idxUser1].nPl
				dP[idxUser1].personal[nP1].namaPersonal = dP[indeks].usn
				dP[idxUser1].nPl++
				nP2 = dP[indeks].nPl
				dP[indeks].personal[nP2].namaPersonal = dP[idxUser1].usn
				dP[indeks].nPl++
				fmt.Println("-------------------------")
				fmt.Printf("*Username '%s' berhasil ditambahkan \n", cariUsn)
			}
			menuPersonalChat(R, dP, nR, nP, idxUser1)
		}
	}
}

func searchUsn(dP tabUser, nP int, cariUsn string) int { //untuk mencari username yg terdapat pada array data pengguna (akun yang telah diacc)
	var idx int
	idx = -1
	for i := 0; i < nP; i++ {
		if dP[i].usn == cariUsn {
			idx = i
		}
	}
	return idx
}

func personalConversation(R, dP *tabUser, nR, nP, idxUser1, p1 int) { //room chat personal chat
	var pesan string
	var nPC1, nPC2, p2 int
	var indeks int
	var arrPesan [NMAX]string
	fmt.Println()
	fmt.Println("*Ketik '-' untuk kembali ke PERSONAL MENU")
	fmt.Println("---------------------------------------")
	fmt.Printf("𖠋 %s \n", dP[idxUser1].personal[p1].namaPersonal)
	fmt.Println("---------------------------------------")
	if dP[idxUser1].personal[p1].nPC == 0 {
		fmt.Println("                 empty                 ")
	} else {
		for i := 0; i < dP[idxUser1].personal[p1].nPC; i++ {
			fmt.Println(dP[idxUser1].personal[p1].personalChat[i])
		}
	}
	fmt.Println("---------------------------------------")
	fmt.Print("Send: ")
	var p bool = true

	for i := 0; i < NMAX && p; i++ {
		fmt.Scan(&arrPesan[i])
		if arrPesan[i] == "-" {
			p = false
		}
	}
	p = true
	for i := 0; i < NMAX && p; i++ {
		if arrPesan[i] == "-" {
			p = false
		} else {
			pesan += arrPesan[i] + " "
		}
	}
	indeks = searchUsn(*dP, nP, dP[idxUser1].personal[p1].namaPersonal)
	if pesan != "" {
		spasi := ""
		panjangPesan := len(pesan)
		banyakSpasi := 40 - panjangPesan
		for i := 0; i < banyakSpasi; i++ {
			spasi += " "
		}
		nPC1 = dP[idxUser1].personal[p1].nPC
		dP[idxUser1].personal[p1].personalChat[nPC1] = spasi + pesan
		dP[idxUser1].personal[p1].nPC++
		p2 = searchP(*dP, nP, idxUser1, indeks)
		nPC2 = dP[indeks].personal[p2].nPC
		dP[indeks].personal[p2].personalChat[nPC2] = pesan
		dP[indeks].personal[p2].nPC++
		personalConversation(R, dP, nR, nP, idxUser1, p1)
	} else {
		menuPersonalChat(R, dP, nR, nP, idxUser1)
	}
}

func searchP(dP tabUser, nP, idxUser1, indeks int) int { //buat search nama pengguna, sehingga chat tersimpan di array orang yang dichat
	var idx int
	idx = -1
	for i := 0; i < nP; i++ {
		if dP[idxUser1].usn == dP[indeks].personal[i].namaPersonal {
			idx = i
		}
	}
	return idx
}

func menuGroupChat(R, dP *tabUser, nR, nP, idxUser1 int) { //menu group chat
	var pilih int
	var g1 int
	fmt.Println()
	fmt.Println("*Ketik '111' untuk kembali ke USER MENU")
	fmt.Println("*Ketik '0' untuk menambah Group Chat")
	fmt.Println("-------------------------")
	fmt.Println("|       GROUP CHAT      |")
	fmt.Println("-------------------------")
	if dP[idxUser1].nG == 0 {
		fmt.Println("          empty          ")
	}
	urutGroup(dP, idxUser1)
	for i := 0; i < dP[idxUser1].nG; i++ {
		fmt.Printf("%d. %s \n", i+1, dP[idxUser1].group[i].namaGroup)
	}
	fmt.Println("-------------------------")
	fmt.Print("Pilih 0")
	for i := 0; i < dP[idxUser1].nG; i++ {
		fmt.Printf("/%d", i+1)
	}
	fmt.Print("?")
	fmt.Println()
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 0 && pilih != 111 && pilih > dP[idxUser1].nG {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 111 {
		menuMain(R, dP, nR, nP, idxUser1)
	} else if pilih == 0 {
		addGC(R, dP, nR, nP, idxUser1)
	} else if pilih <= dP[idxUser1].nG {
		g1 = pilih - 1
		groupConversation(R, dP, nR, nP, idxUser1, g1)
	}
}

func urutGroup(dP *tabUser, idxUser1 int) { //ini sama kaya func urutPersonal, bedanya func ini ngurutin nama grup (ascending)
	var temp dataGroup
	var pass, idx, i int
	pass = 1
	for pass <= dP[idxUser1].nG-1 {
		idx = pass - 1
		i = pass
		for i < dP[idxUser1].nG {
			if dP[idxUser1].group[idx].namaGroup > dP[idxUser1].group[i].namaGroup {
				idx = i
			}
			i++
		}
		temp = dP[idxUser1].group[pass-1]
		dP[idxUser1].group[pass-1] = dP[idxUser1].group[idx]
		dP[idxUser1].group[idx] = temp
		pass++
	}
}

func addGC(R, dP *tabUser, nR, nP, idxUser1 int) { //untuk nambahin group
	var namagroup string
	fmt.Println()
	fmt.Println("*Ketik nama group yang ingin dibuat.")
	fmt.Println("*Ketik '-' untuk kembali ke GROUP CHAT")
	fmt.Println("--------------------------")
	fmt.Println("|      CREATE GROUP      |")
	fmt.Println("--------------------------")
	fmt.Print("Nama Group: ")
	fmt.Scan(&namagroup)
	if namagroup == "-" {
		menuGroupChat(R, dP, nR, nP, idxUser1)
	}
	fmt.Println("--------------------------")
	anggotaGroup(R, dP, nR, nP, idxUser1, namagroup)
}

func anggotaGroup(R, dP *tabUser, nR, nP, idxUser1 int, namagroup string) { //nambahin anggota group
	var tambah bool = true
	var cariUsn string
	var temp [NMAX]string
	var indeks, idx, nGs int
	ntemp := 0
	temp[0] = dP[idxUser1].usn
	ntemp++
	i := 0
	fmt.Println()
	fmt.Println("*Ketik '-' untuk berhenti menambahkan anggota")
	fmt.Println("-------------------------")
	fmt.Println("|       ADD MEMBER      |")
	fmt.Println("-------------------------")
	for tambah && i < NMAX {
		fmt.Print("Username: ")
		fmt.Scan(&cariUsn)
		if cariUsn != "-" {
			tambah = true
			idx = searchUsn(*dP, nP, cariUsn)
			if idx == -1 {
				fmt.Printf("*Username '%s' tidak ditemukan. \n", cariUsn)
				fmt.Println()
			} else {
				var double bool = false
				for j := 0; j < ntemp; j++ {
					if temp[j] == cariUsn {
						double = true
					}
				}
				if double {
					fmt.Printf("*Username '%s' SUDAH ditambahkan ke dalam group. \n", cariUsn)
					fmt.Println()
				} else if !double {
					temp[ntemp] = cariUsn
					ntemp++
					fmt.Printf("*Username '%s' ditambahkan ke dalam group. \n", cariUsn)
					fmt.Println()
				}
			}
		} else {
			tambah = false
		}
		i++
	}
	fmt.Println("-------------------------")
	var pilih int
	fmt.Println("Tambahkan group?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Println("Pilih 1/2?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		for i := 0; i < ntemp; i++ {
			indeks = searchUsn(*dP, nP, temp[i])
			nGs = dP[indeks].nG
			dP[indeks].group[nGs].namaGroup = namagroup
			dP[indeks].nG++
			for j := 0; j < ntemp; j++ {
				dP[indeks].group[nGs].anggota[j] = temp[j]
				dP[indeks].group[nGs].nA++
			}
		}
	}
	menuGroupChat(R, dP, nR, nP, idxUser1)
}

func groupConversation(R, dP *tabUser, nR, nP, idxUser1, g1 int) { //room chat group
	var pesan, chat string
	var nGC1, nGCs, g2 int
	var indeks int
	var arrPesan [NMAX]string
	fmt.Println()
	fmt.Println("*Ketik ';' untuk melihat anggota.")
	fmt.Println("*Ketik '-' untuk kembali ke GROUP MENU")
	fmt.Println("---------------------------------------")
	fmt.Printf("✧ %s \n", dP[idxUser1].group[g1].namaGroup)
	fmt.Println("---------------------------------------")
	if dP[idxUser1].group[g1].nGC == 0 {
		fmt.Println("                 empty                 ")
	} else {
		for i := 0; i < dP[idxUser1].group[g1].nGC; i++ {
			fmt.Println(dP[idxUser1].group[g1].groupChat[i])
		}
	}
	fmt.Println("---------------------------------------")
	fmt.Print("Send: ")
	var p bool = true

	for i := 0; i < NMAX && p; i++ {
		fmt.Scan(&chat)
		if chat == ";" {
			liatAnggota(R, dP, nR, nP, idxUser1, g1)
		}
		arrPesan[i] = chat
		if arrPesan[i] == "-" {
			p = false
		}
	}

	p = true
	for i := 0; i < NMAX && p; i++ {
		if arrPesan[i] == "-" {
			p = false
		} else {
			pesan += arrPesan[i] + " "
		}
	}

	if pesan != "" {
		spasi := ""
		panjangPesan := len(pesan)
		banyakSpasi := 40 - panjangPesan
		for i := 0; i < banyakSpasi; i++ {
			spasi += " "
		}
		nGC1 = dP[idxUser1].group[g1].nGC
		dP[idxUser1].group[g1].groupChat[nGC1] = spasi + pesan
		dP[idxUser1].group[g1].nGC++
		i := 0
		for i < dP[idxUser1].group[g1].nA {
			if dP[idxUser1].group[g1].anggota[i] != dP[idxUser1].usn {
				cari := dP[idxUser1].group[g1].anggota[i]
				indeks = searchUsn(*dP, nP, cari)
				g2 = searchG(*dP, nP, idxUser1, indeks, g1)
				nGCs = dP[indeks].group[g2].nGC
				dP[indeks].group[g2].groupChat[nGCs] = dP[idxUser1].usn + ": " + pesan
				dP[indeks].group[g2].nGC++
			}
			i++
		}
		groupConversation(R, dP, nR, nP, idxUser1, g1)
	} else {
		menuGroupChat(R, dP, nR, nP, idxUser1)
	}
}

func searchG(dP tabUser, nP, idxUser1, indeks, g1 int) int { //search group di akun orang lain, sehingga chat tersimpan di array orang lain
	var idx int
	idx = -1
	for i := 0; i < nP; i++ {
		if dP[idxUser1].group[g1].namaGroup == dP[indeks].group[i].namaGroup {
			idx = i
		}
	}
	return idx
}

func liatAnggota(R, dP *tabUser, nR, nP, idxUser1, g1 int) { //untuk melihat nama anggota di dalam group
	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println("|      LIST MEMBERS     |")
	fmt.Println("-------------------------")
	for i := 0; i < dP[idxUser1].group[g1].nA; i++ {
		fmt.Println(dP[idxUser1].group[g1].anggota[i])
	}
	fmt.Print("-------------------------")
	fmt.Scanln()
	fmt.Scanln()
	groupConversation(R, dP, nR, nP, idxUser1, g1)
}

func menuLoginAdmin(R, dP *tabUser, nR, nP int) { //login sebagai admin
	var namaAdmin, passAdmin string
	fmt.Println()
	fmt.Println("*Isi '-' pada Username Admin atau Password Admin untuk kembali ke OPTION MENU")
	fmt.Println("--------------------------")
	fmt.Println("|       ADMIN LOGIN      |")
	fmt.Println("--------------------------")
	fmt.Print("Username Admin: ")
	fmt.Scan(&namaAdmin)
	if namaAdmin == "-" {
		fmt.Println("*Kembali ke OPTION MENU")
		menuOption(R, dP, &nR, &nP)
	}
	fmt.Print("Password Admin: ")
	fmt.Scan(&passAdmin)
	if passAdmin == "-" {
		fmt.Println("*Kembali ke OPTION MENU")
		menuOption(R, dP, &nR, &nP)
	}
	fmt.Println("--------------------------")
	if namaAdmin == "adminadmin" && passAdmin == "admin123" {
		menuAdmin(R, dP, nR, nP)
	} else {
		fmt.Println("*Username Admin atau Password Admin salah, silahkan coba lagi.")
		menuLoginAdmin(R, dP, nR, nP)
	}
}

func menuAdmin(R, dP *tabUser, nR, nP int) { //menu admin
	var pilih int
	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println("|       ADMIN MENU      |")
	fmt.Println("-------------------------")
	fmt.Println("1. Account Verification")
	fmt.Println("2. User List")
	fmt.Println("3. Delete Account")
	fmt.Println("4. Logout")
	fmt.Println("-------------------------")
	fmt.Println("Pilih 1/2/3/4?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 && pilih != 3 && pilih != 4 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		menuVerification(R, dP, &nR, &nP)
	} else if pilih == 2 {
		menuList(R, dP, nR, nP)
	} else if pilih == 3 {
		menuDelete(R, dP, &nR, &nP)
	} else {
		menuOption(R, dP, &nR, &nP)
	}
}

func menuDelete(R, dP *tabUser, nR, nP *int) { //ini untuk menghapus usn & password / akun
	var username string
	var pilih int
	var indeks int
	fmt.Println()
	fmt.Println("-------------------------")
	fmt.Println("|        USERNAMES      |")
	fmt.Println("-------------------------")
	PrintDataUsn(*dP, *nP)
	fmt.Println("-------------------------")
	fmt.Println()
	fmt.Println("*Ketik '-' untuk kembali ke ADMIN MENU")
	fmt.Println("--------------------------------------------")
	fmt.Println("|              DELETE ACCOUNT              |")
	fmt.Println("--------------------------------------------")
	fmt.Print("Username akun yang akan dihapus: ")
	fmt.Scan(&username)
	fmt.Println("--------------------------------------------")
	if username == "-" {
		menuAdmin(R, dP, *nR, *nP)
	}
	indeks = searchUsn(*dP, *nP, username)
	if indeks == -1 {
		fmt.Printf("*Akun dengan username '%s' tidak ditemukan. \n", username)
		menuAdmin(R, dP, *nR, *nP)
	} else {
		fmt.Println("Hapus akun?")
		fmt.Println("1. Ya  2.Tidak")
		fmt.Println("Pilih 1/2?")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
		for pilih != 1 && pilih != 2 && pilih != 111 {
			fmt.Println()
			fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&pilih)
		}
		if pilih == 1 {
			for indeks < *nP {
				dP[indeks] = dP[indeks+1]
				indeks++
			}
			for i := 0; i < *nP; i++ {
				for j := 0; j < dP[i].nPl; j++ {
					if dP[i].personal[j].namaPersonal == username {
						dP[i].personal[j] = dP[i].personal[j+1]
						dP[i].nPl--
					}
					for k := 0; k < dP[i].group[k].nA; k++ {
						if dP[i].group[j].anggota[k] == username {
							dP[i].group[j].anggota[k] = dP[i].group[j].anggota[k+1]
							dP[i].group[j].nA--
						}
					}
				}
			}
			*nP--
			fmt.Printf("*Akun dengan username '%s' berhasil dihapus. \n", username)
			menuAdmin(R, dP, *nR, *nP)
		} else {
			fmt.Printf("*Akun dengan username '%s' berhasil dihapus. \n", username)
			menuAdmin(R, dP, *nR, *nP)
		}
	}

}

func menuVerification(R, dP *tabUser, nR, nP *int) { //func verifikasi admin untuk acc user yg udah registrasi
	var cari string
	var pilih int
	var ketemu bool = false
	if *nR == 0 {
		fmt.Println()
		fmt.Print("*Tidak ada akun yang harus diverifikasi.")
		fmt.Scanln()
		fmt.Scanln()
		menuAdmin(R, dP, *nR, *nP)
	}
	fmt.Println("--------------------------")
	for i := 0; i < *nR; i++ {
		fmt.Println()
		fmt.Println("Data Pengguna yang diajukan:")
		fmt.Printf("Username: %s \n", R[i].usn)
		fmt.Printf("Password: %s \n", R[i].pass)
		fmt.Println("--------------------------")
		cari = R[i].usn
		for j := 0; j < *nP; j++ {
			if cari == dP[j].usn {
				ketemu = true
			}
		}
		if ketemu {
			fmt.Printf("*Akun dengan username '%s' SUDAH terpakai. \n", R[i].usn)
		} else {
			fmt.Printf("*Akun dengan username '%s' BELUM terpakai. \n", R[i].usn)
		}
		fmt.Println("1. Terima")
		fmt.Println("2. Tolak")
		fmt.Println("-------------------------")
		fmt.Println("Pilih 1/2?")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
		for pilih != 1 && pilih != 2 {
			fmt.Println()
			fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&pilih)
		}
		if pilih == 1 {
			dP[*nP] = R[i]
			*nP++
			fmt.Printf("*Akun dengan username '%s' diTERIMA oleh admin. \n", dP[i].usn)
		} else {
			fmt.Printf("*Akun dengan username '%s' diTOLAK oleh admin. \n", R[i].usn)
		}
		ketemu = false
	}
	*nR = 0
	menuAdmin(R, dP, *nR, *nP)
}

func menuList(R, dP *tabUser, nR, nP int) { //menu list admin
	var pilih int
	for {
		fmt.Println()
		fmt.Println("---------------------------------------------")
		fmt.Println("|                 LIST MENU                 |")
		fmt.Println("---------------------------------------------")
		fmt.Println("1. Username & Password")
		fmt.Println("2. Sort usernames (Ascending)")
		fmt.Println("3. Sort usernames (Descending)")
		fmt.Println("4. Back")
		fmt.Println("---------------------------------------------")
		fmt.Println("Pilih 1/2/3/4?")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
		for pilih != 1 && pilih != 2 && pilih != 3 && pilih != 4 {
			fmt.Println()
			fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&pilih)
		}
		if pilih == 1 {
			fmt.Println()
			fmt.Println("-------------------------")
			fmt.Println("| USERNAMES & PASSWORDS |")
			fmt.Println("-------------------------")
			PrintDataAll(*dP, nP)
			fmt.Print("-------------------------")
			fmt.Scanln()
			fmt.Scanln()
		} else if pilih == 2 {
			ascending(R, dP, nR, nP)
		} else if pilih == 3 {
			descending(R, dP, nR, nP)
		} else {
			menuAdmin(R, dP, nR, nP)
		}
	}
}

func ascending(R, dP *tabUser, nR, nP int) { //buat ngurutin yg udah daftar di menu admin dari yg terkecil
	var pilih int
	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println("[1] Ascending by Alphabet")
	fmt.Println("[2] Ascending by Length")
	fmt.Println("--------------------------")
	fmt.Println("Pilih 1/2?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	fmt.Println()
	if pilih == 1 {
		ascendingSelection1(dP, nP)
		fmt.Println("Ascending by Alphabet")
	} else {
		ascendingSelection2(dP, nP)
		fmt.Println("Ascending by Length")
	}
	fmt.Println("-------------------------")
	fmt.Println("|     LIST USERNAMES    |")
	fmt.Println("-------------------------")
	PrintDataUsn(*dP, nP)
	fmt.Println("-------------------------")
	fmt.Println("Cari Username?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Println("Pilih 1/2?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		cariUsernameAscending(R, dP, nR, nP)
	} else {
		return
	}
}

func descending(R, dP *tabUser, nR, nP int) { //buat ngurutin yg udah daftar di menu admin dari yg terbesar
	var pilih int
	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println("[1] Descending by Alphabet")
	fmt.Println("[2] Descending by Length")
	fmt.Println("---------------------------")
	fmt.Println("Pilih 1/2?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	fmt.Println()
	if pilih == 1 {
		descendingInsertion1(dP, nP)
		fmt.Println("Descending by Alphabet")
	} else {
		descendingInsertion2(dP, nP)
		fmt.Println("Descending by Length")
	}
	fmt.Println("-------------------------")
	fmt.Println("|     LIST USERNAMES    |")
	fmt.Println("-------------------------")
	PrintDataUsn(*dP, nP)
	fmt.Println("-------------------------")
	fmt.Println("Cari Username?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Println("Pilih 1/2?")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 {
		fmt.Println()
		fmt.Printf("*Tidak terdapat pilihan %d, silahkan coba lagi. \n", pilih)
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		cariUsernameDescending(R, dP, nR, nP)
	} else {
		return
	}
}

func cariUsernameAscending(R, dP *tabUser, nR, nP int) { //untuk search username yg udah di acc dari terlecil
	var cari string
	var ketemu int
	for {
		fmt.Println()
		fmt.Println("*Ketik '-' untuk kembali ke LIST MENU")
		fmt.Println("-------------------------")
		fmt.Println("|    SEARCH USERNAME    |")
		fmt.Println("-------------------------")
		fmt.Print("Cari username: ")
		fmt.Scan(&cari)
		if cari == "-" {
			return
		}
		fmt.Println("-------------------------")
		ketemu = binaryAscending(*dP, nP, cari)
		if ketemu != -1 {
			fmt.Printf("*Username '%s' ditemukan pada indeks %d. \n", dP[ketemu].usn, ketemu)
			cariUsernameAscending(R, dP, nR, nP)
		} else if ketemu == -1 {
			fmt.Println("*Username tidak ditemukan.")
		}
	}
}

func cariUsernameDescending(R, dP *tabUser, nR, nP int) { //untuk search username yg udah di acc dari terbesar
	var cari string
	var ketemu int
	for {
		fmt.Println()
		fmt.Println("*Ketik '-' untuk kembali ke LIST MENU")
		fmt.Println("-------------------------")
		fmt.Println("|    SEARCH USERNAME    |")
		fmt.Println("-------------------------")
		fmt.Print("Cari username: ")
		fmt.Scan(&cari)
		if cari == "-" {
			return
		}
		fmt.Println("-------------------------")
		ketemu = binaryDescending(*dP, nP, cari)
		if ketemu != -1 {
			fmt.Printf("*Username '%s' ditemukan pada indeks %d. \n", dP[ketemu].usn, ketemu)
			cariUsernameDescending(R, dP, nR, nP)
		} else if ketemu == -1 {
			fmt.Println("*Username tidak ditemukan.")
		}
	}
}

func ascendingSelection1(dP *tabUser, nP int) {
	var temp dataUser
	var pass, idx, i int
	pass = 1
	for pass <= nP-1 {
		idx = pass - 1
		i = pass
		for i < nP {
			if dP[idx].usn > dP[i].usn {
				idx = i
			}
			i++
		}
		temp = dP[pass-1]
		dP[pass-1] = dP[idx]
		dP[idx] = temp
		pass = pass + 1
	}
}

func ascendingSelection2(dP *tabUser, nP int) {
	var temp dataUser
	var pass, idx, i int
	pass = 1
	for pass <= nP-1 {
		idx = pass - 1
		i = pass
		for i < nP {
			if len(dP[idx].usn) > len(dP[i].usn) {
				idx = i
			}
			i++
		}
		temp = dP[pass-1]
		dP[pass-1] = dP[idx]
		dP[idx] = temp
		pass = pass + 1
	}
}

func descendingInsertion1(dP *tabUser, nP int) {
	var i int
	var pass int
	var temp dataUser
	pass = 1
	for pass <= nP-1 {
		i = pass
		temp = dP[pass]
		for i > 0 && temp.usn > dP[i-1].usn {
			dP[i] = dP[i-1]
			i = i - 1
		}
		dP[i] = temp
		pass = pass + 1
	}
}

func descendingInsertion2(dP *tabUser, nP int) {
	var i int
	var pass int
	var temp dataUser
	pass = 1
	for pass <= nP-1 {
		i = pass
		temp = dP[pass]
		for i > 0 && len(temp.usn) > len(dP[i-1].usn) {
			dP[i] = dP[i-1]
			i = i - 1
		}
		dP[i] = temp
		pass = pass + 1
	}
}

func binaryAscending(dP tabUser, nP int, usn string) int {
	var kanan, kiri, tengah int
	var ketemu int
	kiri = 0
	kanan = nP - 1
	ketemu = -1
	for kiri <= kanan && ketemu == -1 {
		tengah = (kanan + kiri) / 2
		if usn < dP[tengah].usn {
			kanan = tengah - 1
		} else if usn > dP[tengah].usn {
			kiri = tengah + 1
		} else {
			ketemu = tengah
		}
	}
	return ketemu
}

func binaryDescending(dP tabUser, nP int, usn string) int {
	var kanan, kiri, tengah int
	var ketemu int
	kiri = 0
	kanan = nP - 1
	ketemu = -1
	for kiri <= kanan && ketemu == -1 {
		tengah = (kanan + kiri) / 2
		if usn > dP[tengah].usn {
			kanan = tengah - 1
		} else if usn < dP[tengah].usn {
			kiri = tengah + 1
		} else {
			ketemu = tengah
		}
	}
	return ketemu
}

func PrintDataAll(dP tabUser, nP int) {
	for i := 0; i < nP; i++ {
		fmt.Print(dP[i].usn, " & ")
		fmt.Println(dP[i].pass)
	}
}

func PrintDataUsn(dP tabUser, nP int) {
	for i := 0; i < nP; i++ {
		fmt.Println(dP[i].usn)
	}
}
