package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		fmt.Println("\n==== OOTD Planner ====")
		fmt.Println("1. Tampilkan semua pakaian")
		fmt.Println("2. Tambah pakaian")
		fmt.Println("3. Edit pakaian")
		fmt.Println("4. Hapus pakaian")
		fmt.Println("5. Cari pakaian (Binary Search)")
		fmt.Println("6. Urutkan berdasarkan tanggal terakhir dipakai (Insertion Sort)")
		fmt.Println("7. Lihat kombinasi OOTD")
		fmt.Println("8. Rekomendasi berdasarkan cuaca")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
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
			cariBinary(scanner)
		case "6":
			urutkanTanggal()
		case "7":
			tampilkanOutfit()
		case "8":
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
		{"Sweater Abu", "Casual", "Abu-abu", 7, parseTanggal("2025-03-05", layout)},
		{"Mantel Hitam", "Formal", "Hitam", 9, parseTanggal("2025-01-15", layout)},
		{"Hoodie Merah", "Sporty", "Merah", 4, parseTanggal("2025-04-10", layout)},
		{"Kaos Lengan Panjang", "Casual", "Putih", 5, parseTanggal("2025-04-18", layout)},
		{"Jaket Waterproof", "Sporty", "Biru", 6, parseTanggal("2025-03-27", layout)},
		{"Rompi Rajut", "Casual", "Coklat", 6, parseTanggal("2025-03-15", layout)},
		{"Blazer Abu", "Formal", "Abu-abu", 8, parseTanggal("2025-02-22", layout)},
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
		{"Kencan Dingin", "cuaca dingin", []string{"Sweater Abu", "Rompi Rajut", "Celana Jeans"}},
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
	fmt.Println("\nDaftar Pakaian:")

	for i := 0; i < len(daftarPakaian); i++ {
		p := daftarPakaian[i]
		fmt.Printf("%d. %s | %s | %s | Formalitas: %d | Terakhir dipakai: %s\n",
			i+1, p.Nama, p.Kategori, p.Warna, p.Formalitas, p.TerakhirDipakai.Format("2006-01-02"))
	}
}

func tambahPakaian(scanner *bufio.Scanner) {
	var p Pakaian
	layout := "2006-01-02"

	fmt.Print("Nama: ")
	scanner.Scan()
	p.Nama = scanner.Text()

	fmt.Print("Kategori: ")
	scanner.Scan()
	p.Kategori = scanner.Text()

	fmt.Print("Warna: ")
	scanner.Scan()
	p.Warna = scanner.Text()

	fmt.Print("Formalitas (1-10): ")
	fmt.Scanln(&p.Formalitas)

	fmt.Print("Tanggal terakhir dipakai (YYYY-MM-DD): ")
	scanner.Scan()
	p.TerakhirDipakai = parseTanggal(scanner.Text(), layout)

	daftarPakaian = append(daftarPakaian, p)
	fmt.Println("Pakaian berhasil ditambahkan.")
}

func editPakaian(scanner *bufio.Scanner) {
	tampilkanPakaian()
	fmt.Print("Masukkan nomor pakaian yang ingin diedit: ")
	var index int
	fmt.Scanln(&index)
	index--

	if index < 0 || index >= len(daftarPakaian) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	layout := "2006-01-02"

	fmt.Print("Nama baru: ")
	scanner.Scan()
	daftarPakaian[index].Nama = scanner.Text()

	fmt.Print("Kategori baru: ")
	scanner.Scan()
	daftarPakaian[index].Kategori = scanner.Text()

	fmt.Print("Warna baru: ")
	scanner.Scan()
	daftarPakaian[index].Warna = scanner.Text()

	fmt.Print("Formalitas baru (1-10): ")
	fmt.Scanln(&daftarPakaian[index].Formalitas)

	fmt.Print("Tanggal terakhir dipakai baru (YYYY-MM-DD): ")
	scanner.Scan()
	daftarPakaian[index].TerakhirDipakai = parseTanggal(scanner.Text(), layout)

	fmt.Println("Data pakaian berhasil diubah.")
}

func hapusPakaian(scanner *bufio.Scanner) {
	tampilkanPakaian()
	fmt.Print("Masukkan nomor pakaian yang ingin dihapus: ")
	var index int
	fmt.Scanln(&index)
	index--

	if index < 0 || index >= len(daftarPakaian) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	daftarPakaian = append(daftarPakaian[:index], daftarPakaian[index+1:]...)
	fmt.Println("Data pakaian berhasil dihapus.")
}

func cariBinary(scanner *bufio.Scanner) {
	sort.Slice(daftarPakaian, func(i, j int) bool {
		return strings.ToLower(daftarPakaian[i].Nama) < strings.ToLower(daftarPakaian[j].Nama)
	})

	fmt.Print("Masukkan nama pakaian yang dicari: ")
	scanner.Scan()
	cari := strings.ToLower(scanner.Text())

	low := 0
	high := len(daftarPakaian) - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		namaMid := strings.ToLower(daftarPakaian[mid].Nama)

		if namaMid == cari {
			fmt.Printf("Ditemukan: %s | %s | %s | %d | %s\n",
				daftarPakaian[mid].Nama, daftarPakaian[mid].Kategori,
				daftarPakaian[mid].Warna, daftarPakaian[mid].Formalitas,
				daftarPakaian[mid].TerakhirDipakai.Format("2006-01-02"))
			found = true
			break
		} else if namaMid < cari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Pakaian tidak ditemukan.")
	}
}

func urutkanTanggal() {
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
}

func tampilkanOutfit() {
	fmt.Println("\nKombinasi OOTD:")
	for i := 0; i < len(daftarOutfit); i++ {
		o := daftarOutfit[i]
		fmt.Printf("%d. %s | Acara: %s | Pakaian: %s\n",
			i+1, o.Nama, o.KategoriAcara, strings.Join(o.DaftarPakaian, ", "))
	}
}

func rekomendasiCuaca(scanner *bufio.Scanner) {
	fmt.Println("\nMasukkan kondisi cuaca (cerah / hujan / dingin):")
	scanner.Scan()
	kondisi := strings.ToLower(scanner.Text())

	fmt.Println("Rekomendasi Outfit:")
	found := false
	for i := 0; i < len(daftarOutfit); i++ {
		o := daftarOutfit[i]
		if strings.Contains(strings.ToLower(o.KategoriAcara), kondisi) {
			fmt.Printf("- %s | Pakaian: %s\n", o.Nama, strings.Join(o.DaftarPakaian, ", "))
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada outfit yang sesuai dengan cuaca tersebut.")
	}
}