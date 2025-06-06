package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Pakaian struct {
	Nama            string
	Kategori        string
	Warna           string
	Formalitas      int
	TerakhirDipakai time.Time
}

type Outfit struct {
	Nama          string
	KategoriAcara string
	DaftarPakaian []string
}

var daftarPakaian []Pakaian
var daftarOutfit []Outfit

func main() {
	dataDummy()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearScreen()
		fmt.Println("╔════════════════════════════════╗")
		fmt.Println("║          OOTD PLANNER          ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Println("┌────────────────────────────────────┐")
		fmt.Println("│ 1. Tampilkan Semua Pakaian         │")
		fmt.Println("│ 2. Tambah Pakaian Baru             │")
		fmt.Println("│ 3. Edit Data Pakaian               │")
		fmt.Println("│ 4. Hapus Pakaian                   │")
		fmt.Println("│ 5. Tambah kombinasi OOTD           │")
		fmt.Println("│ 6. Cari Pakaian                    │")
		fmt.Println("│ 7. Urutkan Berdasarkan Tanggal     │")
		fmt.Println("│ 8. Urutkan Berdasarkan Formalitas  │")
		fmt.Println("│ 9. Lihat Kombinasi OOTD            │")
		fmt.Println("│ 10. Rekomendasi Acara / Cuaca      │")
		fmt.Println("│ 0. Keluar                          │")
		fmt.Println("└────────────────────────────────────┘")
		fmt.Print("➤ Pilih menu: ")
		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			tampilkanPakaian()
		case "2":
			tambahPakaian(scanner)
		case "3":
			editPakaian(scanner)
		case "4":
			hapusPakaian(scanner)
		case "5":
			tambahOutfit(scanner)
		case "6":
			cariPakaian(scanner)
		case "7":
			urutkanTanggal()
		case "8":
			urutkanFormalitas()
		case "9":
			tampilkanOutfit()
		case "10":
			rekomendasi(scanner)
		case "0":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func dataDummy() {
	layout := "2006-01-02"
	daftarPakaian = []Pakaian{
		{"Kemeja Putih", "Formal", "Putih", 9, parseTanggal("2025-04-01", layout)},
		{"Kaos Hitam", "Casual", "Hitam", 3, parseTanggal("2025-04-25", layout)},
		{"Jaket Jeans", "Casual", "Biru", 5, parseTanggal("2025-03-20", layout)},
		{"Setelan Jas", "Formal", "Hitam", 10, parseTanggal("2025-02-10", layout)},
		{"Celana Chino", "Casual", "Beige", 6, parseTanggal("2025-03-30", layout)},
		{"Celana Jeans", "Casual", "Biru", 4, parseTanggal("2025-04-20", layout)},
		{"Sweater Abu", "Casual", "Abu", 7, parseTanggal("2025-03-05", layout)},
		{"Mantel Hitam", "Formal", "Hitam", 9, parseTanggal("2025-01-15", layout)},
		{"Hoodie Merah", "Sporty", "Merah", 4, parseTanggal("2025-04-10", layout)},
		{"Kaos Lengan Panjang", "Casual", "Putih", 5, parseTanggal("2025-04-18", layout)},
		{"Jaket Waterproof", "Sporty", "Biru", 6, parseTanggal("2025-03-27", layout)},
		{"Rompi Rajut", "Casual", "Coklat", 6, parseTanggal("2025-03-15", layout)},
		{"Blazer Abu", "Formal", "Abu", 8, parseTanggal("2025-02-22", layout)},
	}

	daftarOutfit = []Outfit{
		{"Meeting Formal", "meeting formal", []string{"Kemeja Putih", "Setelan Jas"}},
		{"Casual Hangout", "casual weekend", []string{"Kaos Hitam", "Celana Jeans", "Jaket Jeans"}},
		{"Cuaca Hujan 1", "cuaca hujan", []string{"Jaket Waterproof", "Kaos Lengan Panjang"}},
		{"Cuaca Dingin 1", "cuaca dingin", []string{"Sweater Abu", "Celana Chino"}},
		{"Olahraga Outdoor", "cuaca cerah", []string{"Hoodie Merah", "Celana Jeans"}},
		{"Work From Cafe", "cuaca cerah", []string{"Rompi Rajut", "Kaos Lengan Panjang", "Celana Chino"}},
		{"Kantor Santai", "meeting formal", []string{"Blazer Abu", "Kaos Lengan Panjang", "Celana Chino"}},
		{"Jalan Saat Hujan", "cuaca hujan", []string{"Mantel Hitam", "Kaos Hitam", "Celana Jeans"}},
		{"Kencan Dingin", "cuaca ding`in", []string{"Sweater Abu", "Rompi Rajut", "Celana Jeans"}},
		{"Santai Cerah", "cuaca cerah", []string{"Kaos Hitam", "Celana Chino"}},
		{"Weekend Sporty", "cuaca cerah", []string{"Hoodie Merah", "Celana Chino"}},
		{"Rainy Day Cozy", "cuaca hujan", []string{"Sweater Abu", "Jaket Waterproof"}},
		{"Windy Walk", "cuaca dingin", []string{"Mantel Hitam", "Rompi Rajut", "Celana Jeans"}},
		{"After Work Casual", "casual weekend", []string{"Kaos Lengan Panjang", "Celana Chino"}},
		{"Formal Presentation", "meeting formal", []string{"Blazer Abu", "Kemeja Putih", "Celana Chino"}},
		{"Coffee Date Cerah", "cuaca cerah", []string{"Rompi Rajut", "Kaos Lengan Panjang", "Celana Jeans"}},
		{"Hujan dan Kantor", "cuaca hujan", []string{"Blazer Abu", "Jaket Waterproof", "Celana Chino"}},
		{"Liburan Dingin", "cuaca dingin", []string{"Mantel Hitam", "Sweater Abu", "Celana Chino"}},
		{"Senin Pagi Formal", "meeting formal", []string{"Setelan Jas", "Kemeja Putih"}},
		{"Outfit Harian Cerah", "cuaca cerah", []string{"Kaos Hitam", "Celana Jeans"}},
	}
}

func tampilkanPakaian() {
	clearScreen()
	fmt.Println("╔════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         DAFTAR PAKAIAN ANDA                            ║")
	fmt.Println("╠═════╦══════════════════╦═══════════╦════════╦═══════════╦══════════════╣")
	fmt.Println("║ No. ║ Nama             ║ Kategori  ║ Warna  ║ Formalitas║ Terakhir     ║")
	fmt.Println("╠═════╬══════════════════╬═══════════╬════════╬═══════════╬══════════════╣")

	for i := 0; i < len(daftarPakaian); i++ {
		pakaian := daftarPakaian[i]
		fmt.Printf("║ %-3d ║ %-16s ║ %-9s ║ %-6s ║ %-9d ║ %-12s ║\n",
			i+1, batasiPanjang(pakaian.Nama, 16), pakaian.Kategori, pakaian.Warna, 
			pakaian.Formalitas, pakaian.TerakhirDipakai.Format("02/01/06"))
		if i < len(daftarPakaian)-1 {
			fmt.Println("╟─────╫──────────────────╫───────────╫────────╫───────────╫──────────────╢")
		}
	}

	fmt.Println("╚═════╩══════════════════╩═══════════╩════════╩═══════════╩══════════════╝")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	fmt.Scanln()
}

func tambahPakaian(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("╔════════════════════════════════╗")
	fmt.Println("║       TAMBAH PAKAIAN BARU      ║")
	fmt.Println("╚════════════════════════════════╝")

	var pakaian Pakaian
	layout := "2006-01-02"

	fmt.Print("➤ Nama: ")
	scanner.Scan()
	pakaian.Nama = scanner.Text()

	for {
		fmt.Print("➤ Kategori: ")
		scanner.Scan()
		kategori := scanner.Text()
		isAngka := true
		for i := 0; i < len(kategori); i++ {
			if kategori[i] < '0' || kategori[i] > '9' {
				isAngka = false
				break
			}
		}
		if isAngka || len(kategori) == 0 {
			fmt.Println("Kategori tidak boleh berupa angka atau kosong.")
			continue
		}
		pakaian.Kategori = kategori
		break
	}

	for {
		fmt.Print("➤ Warna: ")
		scanner.Scan()
		warna := scanner.Text()
		isAngka := true
		for i := 0; i < len(warna); i++ {
			if warna[i] < '0' || warna[i] > '9' {
				isAngka = false
				break
			}
		}
		if isAngka || len(warna) == 0 {
			fmt.Println("Warna tidak boleh berupa angka atau kosong.")
			continue
		}
		pakaian.Warna = warna
		break
	}

	for {
		fmt.Print("➤ Formalitas (1-10): ")
		var formalitas int
		fmt.Scanln(&formalitas)
		if formalitas < 1 || formalitas > 10 {
			fmt.Println("Nilai formalitas harus antara 1 sampai 10.")
			continue
		}
		pakaian.Formalitas = formalitas
		break
	}

	fmt.Print("➤ Tanggal terakhir dipakai (YYYY-MM-DD): ")
	scanner.Scan()
	pakaian.TerakhirDipakai = parseTanggal(scanner.Text(), layout)

	daftarPakaian = append(daftarPakaian, pakaian)

	fmt.Println("\n┌────────────────────────────────┐")
	fmt.Println("│  Pakaian berhasil ditambahkan! │")
	fmt.Println("└────────────────────────────────┘")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	fmt.Scanln()
}

func editPakaian(scanner *bufio.Scanner) {
	tampilkanPakaian()
	fmt.Print("Masukkan nama pakaian yang ingin diedit: ")
	scanner.Scan()
	namaCari := scanner.Text()

	// Sequential search
	idx := -1
	for i := 0; i < len(daftarPakaian); i++ {
		if daftarPakaian[i].Nama == namaCari {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Pakaian tidak ditemukan.")
		fmt.Println("\nTekan enter untuk selanjutnya...")
		fmt.Scanln()
		return
	}

	layout := "2006-01-02"

	fmt.Print("➤ Nama Baru: ")
	scanner.Scan()
	daftarPakaian[idx].Nama = scanner.Text()

	for {
		fmt.Print("➤ Kategori: ")
		scanner.Scan()
		input := scanner.Text()
		isAngka := true
		for i := 0; i < len(input); i++ {
			if input[i] < '0' || input[i] > '9' {
				isAngka = false
				break
			}
		}
		if isAngka || len(input) == 0 {
			fmt.Println("Kategori tidak boleh berupa angka atau kosong.")
			continue
		}
		daftarPakaian[idx].Kategori = input
		break
	}

	for {
		fmt.Print("➤ Warna: ")
		scanner.Scan()
		input := scanner.Text()
		isAngka := true
		for i := 0; i < len(input); i++ {
			if input[i] < '0' || input[i] > '9' {
				isAngka = false
				break
			}
		}
		if isAngka || len(input) == 0 {
			fmt.Println("Warna tidak boleh berupa angka atau kosong.")
			continue
		}
		daftarPakaian[idx].Warna = input
		break
	}

	for {
		fmt.Print("➤ Formalitas baru (1-10): ")
		var formalitas int
		fmt.Scanln(&formalitas)
		if formalitas < 1 || formalitas > 10 {
			fmt.Println("Nilai formalitas harus antara 1 sampai 10.")
			scanner.Scan()
			continue
		}
		daftarPakaian[idx].Formalitas = formalitas
		break
	}

	fmt.Print("➤ Tanggal terakhir dipakai (YYYY-MM-DD): ")
	scanner.Scan()
	daftarPakaian[idx].TerakhirDipakai = parseTanggal(scanner.Text(), layout)

	fmt.Println("\n┌────────────────────────────────┐")
	fmt.Println("│    Pakaian berhasil diedit!    │")
	fmt.Println("└────────────────────────────────┘")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	fmt.Scanln()
}

func hapusPakaian(scanner *bufio.Scanner) {
	tampilkanPakaian()
	fmt.Print("Masukkan nama pakaian yang ingin dihapus: ")
	scanner.Scan()
	namaCari := strings.ToLower(scanner.Text())

	// Insertion sort
	for i := 1; i < len(daftarPakaian); i++ {
		pakaian := daftarPakaian[i]
		j := i - 1
		for j >= 0 && strings.ToLower(daftarPakaian[j].Nama) > strings.ToLower(pakaian.Nama) {
			daftarPakaian[j+1] = daftarPakaian[j]
			j--
		}
		daftarPakaian[j+1] = pakaian
	}

	// Binary search
	low, high := 0, len(daftarPakaian)-1
	idx := -1
	for low <= high {
		mid := (low + high) / 2
		namaMid := strings.ToLower(daftarPakaian[mid].Nama)
		if namaMid == namaCari {
			idx = mid
			break
		} else if namaMid < namaCari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if idx == -1 {
		fmt.Println("Pakaian tidak ditemukan.")
		fmt.Println("\nTekan enter untuk selanjutnya...")
		fmt.Scanln()
		return
	}

	daftarPakaian = append(daftarPakaian[:idx], daftarPakaian[idx+1:]...)
	fmt.Println("Data pakaian berhasil dihapus.")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	fmt.Scanln()
}

func tambahOutfit(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("╔════════════════════════════════╗")
	fmt.Println("║      TAMBAH KOMBINASI OOTD     ║")
	fmt.Println("╚════════════════════════════════╝")

	var outfit Outfit

	fmt.Print("➤ Nama kombinasi: ")
	scanner.Scan()
	outfit.Nama = scanner.Text()

	fmt.Print("➤ Kategori acara: ")
	scanner.Scan()
	outfit.KategoriAcara = scanner.Text()

	fmt.Println("Masukkan nama-nama pakaian yang ingin ditambahkan ke kombinasi (pisahkan dengan koma):")
	scanner.Scan()
	masukan := scanner.Text()
	namaPakaian := strings.Split(masukan, ",")

	valid := true
	var daftar []string
	for i := 0; i < len(namaPakaian); i++ {
		nama := strings.TrimSpace(namaPakaian[i])
		if cekNamaPakaian(nama) {
			daftar = append(daftar, nama)
		} else {
			fmt.Printf("Pakaian dengan nama \"%s\" tidak ditemukan.\n", nama)
			valid = false
		}
	}

	if !valid {
		fmt.Println("Gagal menambahkan kombinasi karena ada nama pakaian yang tidak valid.")
		fmt.Println("\nTekan enter untuk kembali...")
		fmt.Scanln()
		return
	}

	outfit.DaftarPakaian = daftar
	daftarOutfit = append(daftarOutfit, outfit)

	fmt.Println("\n┌────────────────────────────────┐")
	fmt.Println("│   Outfit berhasil ditambah!    │")
	fmt.Println("└────────────────────────────────┘")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	fmt.Scanln()
}

func cariPakaian(scanner *bufio.Scanner) {
	// Selection sort
	n := len(daftarPakaian)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(daftarPakaian[j].Nama) < strings.ToLower(daftarPakaian[minIdx].Nama) {
				minIdx = j
			}
		}
		if minIdx != i {
			daftarPakaian[i], daftarPakaian[minIdx] = daftarPakaian[minIdx], daftarPakaian[i]
		}
	}

	fmt.Print("Masukkan nama pakaian yang dicari: ")
	scanner.Scan()
	cari := strings.ToLower(scanner.Text())

	//Binary Search
	low := 0
	high := len(daftarPakaian) - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		namaMid := strings.ToLower(daftarPakaian[mid].Nama)

		if namaMid == cari {
			fmt.Println("Pakaian ditemukan:")
			fmt.Println("╔═════╦══════════════════╦═══════════╦════════╦═══════════╦══════════════╗")
			fmt.Printf("║ %-3d ║ %-16s ║ %-9s ║ %-6s ║ %-9d ║ %-12s ║\n",
				mid+1, batasiPanjang(daftarPakaian[mid].Nama, 16),
				daftarPakaian[mid].Kategori, daftarPakaian[mid].Warna,
				daftarPakaian[mid].Formalitas, daftarPakaian[mid].TerakhirDipakai.Format("2006-01-02"))
			fmt.Println("╚═════╩══════════════════╩═══════════╩════════╩═══════════╩══════════════╝")
			found = true
			fmt.Println("\nTekan enter untuk selanjutnya...")
			fmt.Scanln()
			break
		} else if namaMid < cari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Pakaian tidak ditemukan.")
		fmt.Println("\nTekan enter untuk selanjutnya...")
		fmt.Scanln()
	}
}

func urutkanTanggal() {
	//Insertion Sort
	for i := 1; i < len(daftarPakaian); i++ {
		key := daftarPakaian[i]
		j := i - 1
		for j >= 0 && daftarPakaian[j].TerakhirDipakai.After(key.TerakhirDipakai) {
			daftarPakaian[j+1] = daftarPakaian[j]
			j--
		}
		daftarPakaian[j+1] = key
	}
	fmt.Println("Data diurutkan berdasarkan tanggal terakhir dipakai (dari lama ke terbaru).")
	tampilkanPakaian()
}

func urutkanFormalitas() {
	//Selection Sort
	n := len(daftarPakaian)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if daftarPakaian[j].Formalitas > daftarPakaian[maxIdx].Formalitas {
				maxIdx = j
			}
		}
		// Tukar nilai formalitas
		daftarPakaian[i].Formalitas, daftarPakaian[maxIdx].Formalitas = 
		daftarPakaian[maxIdx].Formalitas, daftarPakaian[i].Formalitas
	}
	fmt.Println("Data diurutkan berdasarkan formalitas (dari tetringgi ke terendah).")
	tampilkanPakaian()
}

func tampilkanOutfit() {
	clearScreen()
	fmt.Println("╔════════╦════════════════════════════════════════════╦════════════════════════════╦════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║   No.  ║ Nama Kombinasi                             ║ Kategori Acara             ║ Daftar Pakaian                                                                     ║")
	fmt.Println("╠════════╬════════════════════════════════════════════╬════════════════════════════╬════════════════════════════════════════════════════════════════════════════════════╣")

	for i := 0; i < len(daftarOutfit); i++ {
		outfit := daftarOutfit[i]
		fmt.Printf("║ %-6d ║ %-42s ║ %-26s ║ %-82s ║\n",
			i+1,
			batasiPanjang(outfit.Nama, 42),
			batasiPanjang(outfit.KategoriAcara, 26),
			batasiPanjang(strings.Join(outfit.DaftarPakaian, ", "), 82),
		)

		if i < len(daftarOutfit)-1 {
			fmt.Println("╟────────╫────────────────────────────────────────────╫────────────────────────────╫────────────────────────────────────────────────────────────────────────────────────╢")
		}
	}

	fmt.Println("╚════════╩════════════════════════════════════════════╩════════════════════════════╩════════════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println("\nTekan enter untuk kembali...")
	fmt.Scanln()
}

func rekomendasi(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║        REKOMENDASI BERDASARKAN         ║")
	fmt.Println("║            CUACA / ACARA               ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Print("\n➤ Masukkan kondisi cuaca/acara: ")
	scanner.Scan()
	kondisi := strings.ToLower(scanner.Text())

	fmt.Println("┌───────────────────────────────┬──────────────────────────────────────────────────┐")
	fmt.Println("│        Nama Kombinasi         │                 Daftar Pakaian                   │")
	fmt.Println("├───────────────────────────────┼──────────────────────────────────────────────────┤")

	idx := -1
	for i := 0; i < len(daftarOutfit); i++ {
		outfit := daftarOutfit[i]

		if strings.Contains(strings.ToLower(outfit.KategoriAcara), kondisi) {
			nama := outfit.Nama
			if len(nama) > 30 {
				nama = nama[:29] + "..."
			}

			pakaian := strings.Join(outfit.DaftarPakaian, ", ")
			if len(pakaian) > 52 {
				pakaian = pakaian[:48] + "..."
			}

			fmt.Printf("│ %-29s │ %-48s │\n", nama, pakaian)
			idx = 1
		}
	}

	if idx == -1 {
		fmt.Println("│ Tidak ada outfit yang sesuai dengan acara/cuaca tersebut                               │")
	}

	fmt.Println("└───────────────────────────────┴──────────────────────────────────────────────────┘")
	fmt.Println("\nTekan enter untuk kembali...")
	fmt.Scanln()
}

func parseTanggal(tanggal string, layout string) time.Time {
	t, _ := time.Parse(layout, tanggal)
	return t
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func cekNamaPakaian(nama string) bool {
	nama = strings.ToLower(nama)
	for i := 0; i < len(daftarPakaian); i++ {
		if strings.ToLower(daftarPakaian[i].Nama) == nama {
			return true
		}
	}
	return false
}

func batasiPanjang(text string, panjang int) string {
	if len(text) > panjang {
		return text[:panjang-3] + "..."
	}
	return text
}
