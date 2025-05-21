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
		fmt.Println("│ 10. Rekomendasi Berdasarkan Cuaca  │")
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
			urutkanOutfitFormalitas()
		case "9":
			tampilkanOutfit()
		case "10":
			rekomendasiCuaca(scanner)
		case "0":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func parseTanggal(tanggal string, layout string) time.Time {
	t, _ := time.Parse(layout, tanggal)
	return t
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

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func tampilkanPakaian() {
	clearScreen()
	fmt.Println("╔════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         DAFTAR PAKAIAN ANDA                            ║")
	fmt.Println("╠═════╦══════════════════╦═══════════╦════════╦═══════════╦══════════════╣")
	fmt.Println("║ No. ║ Nama             ║ Kategori  ║ Warna  ║ Formalitas║ Terakhir     ║")
	fmt.Println("╠═════╬══════════════════╬═══════════╬════════╬═══════════╬══════════════╣")

	for i := 0; i < len(daftarPakaian); i++ {
		p := daftarPakaian[i]
		fmt.Printf("║ %-3d ║ %-16s ║ %-9s ║ %-6s ║ %-9d ║ %-12s ║\n",
			i+1, batasiPanjang(p.Nama, 16), p.Kategori, p.Warna, p.Formalitas, p.TerakhirDipakai.Format("02/01/06"))
		if i < len(daftarPakaian)-1 {
			fmt.Println("╟─────╫──────────────────╫───────────╫────────╫───────────╫──────────────╢")
		}
	}

	fmt.Println("╚═════╩══════════════════╩═══════════╩════════╩═══════════╩══════════════╝")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	bufio.NewScanner(os.Stdin).Scan()
}

func batasiPanjang(text string, panjang int) string {
	if len(text) > panjang {
		return text[:panjang-3] + "..."
	}
	return text
}

func tambahPakaian(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("╔════════════════════════════════╗")
	fmt.Println("║       TAMBAH PAKAIAN BARU      ║")
	fmt.Println("╚════════════════════════════════╝")

	var p Pakaian
	layout := "2006-01-02"

	fmt.Print("➤ Nama: ")
	scanner.Scan()
	p.Nama = scanner.Text()

	fmt.Print("➤ Kategori: ")
	scanner.Scan()
	p.Kategori = scanner.Text()

	fmt.Print("➤ Warna: ")
	scanner.Scan()
	p.Warna = scanner.Text()

	fmt.Print("➤ Formalitas (1-10): ")
	fmt.Scanln(&p.Formalitas)

	fmt.Print("➤ Tanggal terakhir dipakai (YYYY-MM-DD): ")
	scanner.Scan()
	p.TerakhirDipakai = parseTanggal(scanner.Text(), layout)

	daftarPakaian = append(daftarPakaian, p)

	fmt.Println("\n┌────────────────────────────────┐")
	fmt.Println("│  Pakaian berhasil ditambahkan! │")
	fmt.Println("└────────────────────────────────┘")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	bufio.NewScanner(os.Stdin).Scan()
}

func editPakaian(scanner *bufio.Scanner) {
	tampilkanPakaian()
	fmt.Print("Masukkan nama pakaian yang ingin diedit: ")
	scanner.Scan()
	namaCari := scanner.Text()

	// Sequential search
	idx := -1
	for i := 0; i < len(daftarPakaian); i++ {
		if strings.EqualFold(daftarPakaian[i].Nama, namaCari) {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("Pakaian tidak ditemukan.")
		fmt.Println("\nTekan enter untuk selanjutnya...")
		bufio.NewScanner(os.Stdin).Scan()
		return
	}

	layout := "2006-01-02"

	fmt.Print("➤ Nama Baru: ")
	scanner.Scan()
	daftarPakaian[idx].Nama = scanner.Text()

	fmt.Print("➤ Kategori baru: ")
	scanner.Scan()
	daftarPakaian[idx].Kategori = scanner.Text()

	fmt.Print("➤ Warna baru: ")
	scanner.Scan()
	daftarPakaian[idx].Warna = scanner.Text()

	fmt.Print("➤ Formalitas baru (1-10): ")
	fmt.Scanln(&daftarPakaian[idx].Formalitas)

	fmt.Print("➤ Tanggal terakhir dipakai (YYYY-MM-DD): ")
	scanner.Scan()
	daftarPakaian[idx].TerakhirDipakai = parseTanggal(scanner.Text(), layout)

	fmt.Println("\n┌────────────────────────────────┐")
	fmt.Println("│    Pakaian berhasil diedit!    │")
	fmt.Println("└────────────────────────────────┘")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	bufio.NewScanner(os.Stdin).Scan()
}

func hapusPakaian(scanner *bufio.Scanner) {
	tampilkanPakaian()
	fmt.Print("Masukkan nama pakaian yang ingin dihapus: ")
	scanner.Scan()
	namaCari := strings.ToLower(scanner.Text())

	// Insertion sort
	for i := 1; i < len(daftarPakaian); i++ {
		key := daftarPakaian[i]
		j := i - 1
		for j >= 0 && strings.ToLower(daftarPakaian[j].Nama) > strings.ToLower(key.Nama) {
			daftarPakaian[j+1] = daftarPakaian[j]
			j--
		}
		daftarPakaian[j+1] = key
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
		bufio.NewScanner(os.Stdin).Scan()
		return
	}

	daftarPakaian = append(daftarPakaian[:idx], daftarPakaian[idx+1:]...)
	fmt.Println("Data pakaian berhasil dihapus.")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	bufio.NewScanner(os.Stdin).Scan()
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
		bufio.NewScanner(os.Stdin).Scan()
		return
	}

	outfit.DaftarPakaian = daftar
	daftarOutfit = append(daftarOutfit, outfit)

	fmt.Println("\n┌────────────────────────────────┐")
	fmt.Println("│   Outfit berhasil ditambah!    │")
	fmt.Println("└────────────────────────────────┘")
	fmt.Println("\nTekan enter untuk selanjutnya...")
	bufio.NewScanner(os.Stdin).Scan()
}

func cekNamaPakaian(nama string) bool {
	//Sequential Search
	for i := 0; i < len(daftarPakaian); i++ {
		if strings.EqualFold(daftarPakaian[i].Nama, nama) {
			return true
		}
	}
	return false
}

func cariPakaian(scanner *bufio.Scanner) {
	//selection sort
	n := len(daftarPakaian)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(daftarPakaian[j].Warna) < strings.ToLower(daftarPakaian[minIdx].Warna) {
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
			bufio.NewScanner(os.Stdin).Scan()
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
		bufio.NewScanner(os.Stdin).Scan()
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
func tampilkanOutfit() {
	clearScreen()
	fmt.Println("╔════════╦════════════════════════════════════════════╦════════════════════════════╦════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║   No.  ║ Nama Kombinasi                             ║ Kategori Acara             ║ Daftar Pakaian                                                                     ║")
	fmt.Println("╠════════╬════════════════════════════════════════════╬════════════════════════════╬════════════════════════════════════════════════════════════════════════════════════╣")

	for i := 0; i < len(daftarOutfit); i++ {
		o := daftarOutfit[i]
		fmt.Printf("║ %-6d ║ %-42s ║ %-26s ║ %-82s ║\n",
			i+1,
			batasiPanjang(o.Nama, 42),
			batasiPanjang(o.KategoriAcara, 26),
			batasiPanjang(strings.Join(o.DaftarPakaian, ", "), 82),
		)

		if i < len(daftarOutfit)-1 {
			fmt.Println("╟────────╫────────────────────────────────────────────╫────────────────────────────╫────────────────────────────────────────────────────────────────────────────────────╢")
		}
	}

	fmt.Println("╚════════╩════════════════════════════════════════════╩════════════════════════════╩════════════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println("\nTekan enter untuk kembali...")
	bufio.NewScanner(os.Stdin).Scan()
}

func rekomendasiCuaca(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║        REKOMENDASI BERDASARKAN         ║")
	fmt.Println("║               CUACA                    ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Print("\n➤ Masukkan kondisi cuaca (cerah/hujan/dingin): ")
	scanner.Scan()
	kondisi := strings.ToLower(scanner.Text())

	fmt.Println("┌───────────────────────────────┬──────────────────────────────────────────────────┐")
	fmt.Println("│        Nama Kombinasi         │                 Daftar Pakaian                   │")
	fmt.Println("├───────────────────────────────┼──────────────────────────────────────────────────┤")

	found := false
	for i := 0; i < len(daftarOutfit); i++ {
		o := daftarOutfit[i]

		if strings.Contains(strings.ToLower(o.KategoriAcara), kondisi) {
			nama := o.Nama
			if len(nama) > 30 {
				nama = nama[:29] + "..."
			}

			pakaian := strings.Join(o.DaftarPakaian, ", ")
			if len(pakaian) > 52 {
				pakaian = pakaian[:48] + "..."
			}

			fmt.Printf("│ %-29s │ %-48s │\n", nama, pakaian)
			found = true
		}
	}

	if !found {
		fmt.Println("│ Tidak ada outfit yang sesuai dengan cuaca tersebut                             │")
	}

	fmt.Println("└───────────────────────────────┴──────────────────────────────────────────────────┘")
	fmt.Println("\nTekan enter untuk kembali...")
	bufio.NewScanner(os.Stdin).Scan()
}

func urutkanOutfitFormalitas() {
	clearScreen()
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║         URUTAN OUTFIT BERDASARKAN FORMALITAS (SELECTION SORT)              ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════╝")

	n := len(daftarOutfit)
	formalitasOutfit := make([]float64, n)
	for i, o := range daftarOutfit {
		total := 0
		count := 0
		for _, nama := range o.DaftarPakaian {
			for _, p := range daftarPakaian {
				if strings.EqualFold(p.Nama, nama) {
					total += p.Formalitas
					count++
					break
				}
			}
		}
		if count > 0 {
			formalitasOutfit[i] = float64(total) / float64(count)
		} else {
			formalitasOutfit[i] = 0
		}
	}
	// Selection sort dari formalitas tertinggi ke terendah
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if formalitasOutfit[j] > formalitasOutfit[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			daftarOutfit[i], daftarOutfit[maxIdx] = daftarOutfit[maxIdx], daftarOutfit[i]
			formalitasOutfit[i], formalitasOutfit[maxIdx] = formalitasOutfit[maxIdx], formalitasOutfit[i]
		}
	}
	fmt.Println("Daftar outfit berhasil diurutkan berdasarkan formalitas (tinggi ke rendah).")
	tampilkanOutfit()
}
