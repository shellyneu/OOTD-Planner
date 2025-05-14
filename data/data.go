package data

import (
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

func DataDummy() {
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

func parseTanggal(tanggal string, layout string) time.Time {
	t, _ := time.Parse(layout, tanggal)
	return t
}

func ShowData() ([]Pakaian, []Outfit) {
	return daftarPakaian, daftarOutfit
}
