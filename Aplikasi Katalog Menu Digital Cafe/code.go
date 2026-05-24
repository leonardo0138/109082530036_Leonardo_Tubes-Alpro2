package main

import "fmt"

const NMAX int = 100

type menu struct {
	id        int
	nama      string
	kategori  string
	harga     int
	komposisi string
	tersedia  bool
}

type tabMenu [NMAX]menu

// ================= TAMBAH MENU =================
func tambahMenu(T *tabMenu, n *int) {

	fmt.Println("\n   ~Tambah Menu~   ")

	fmt.Print("ID: ")
	jumlah, err := fmt.Scan(&T[*n].id)

	if err != nil || jumlah != 1 {
		fmt.Print("ID Tidak Valid")
		return
	}

	fmt.Print("Nama Menu: ")
	fmt.Scan(&T[*n].nama)

	fmt.Print("Kategori (coffee/non-coffee): ")
	fmt.Scan(&T[*n].kategori)

	fmt.Print("Harga: ")
	fmt.Scan(&T[*n].harga)

	fmt.Print("Komposisi: ")
	fmt.Scan(&T[*n].komposisi)

	fmt.Print("Tersedia (true/false): ")
	fmt.Scan(&T[*n].tersedia)

	*n = *n + 1

	fmt.Println("Menu berhasil ditambahkan.")
}

// ================= TAMPIL MENU =================
func tampilMenu(T tabMenu, n int) {
	var i int

	fmt.Println("\n  ~Data Menu~  ")

	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	for i = 0; i < n; i++ {
		fmt.Println("ID         :", T[i].id)
		fmt.Println("Nama       :", T[i].nama)
		fmt.Println("Kategori   :", T[i].kategori)
		fmt.Println("Harga      :", T[i].harga)
		fmt.Println("Komposisi  :", T[i].komposisi)
		fmt.Println("Tersedia   :", T[i].tersedia)
		fmt.Println("---------------------")
	}
}

// ================= CARI ID =================
func cariID(T tabMenu, n, id int) int {
	var i int

	for i = 0; i < n; i++ {
		if T[i].id == id {
			return i
		}
	}

	return -1
}

// ================= UBAH MENU =================
func ubahMenu(T *tabMenu, n int) {
	var id, idx int
	var pilih int

	fmt.Print("Masukkan ID yang diubah: ")
	fmt.Scan(&id)

	idx = cariID(*T, n, id)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Println("/nData yang ditemukan")
	fmt.Println("1. Nama")
	fmt.Println("2. Kategori")
	fmt.Println("3. Harga")
	fmt.Println("4. Komposisi ")
	fmt.Println("5. Status Tersedia")
	fmt.Print("Pilih Data Yang Ingin Di Ubah: ")
	fmt.Scan(&pilih)

	switch pilih{
	case 1:
	fmt.Print("Nama baru:")
	fmt.Scan(&T[idx].nama)

	case 2: 
	fmt.Print("Kategori baru: ")
	fmt.Scan(&T[idx].kategori)

	case 3:
	fmt.Print("Harga baru: ")
	fmt.Scan(&T[idx].harga)

	case 4:
	fmt.Print("Komposisi baru: ")
	fmt.Scan(&T[idx].komposisi)

	case 5:
	fmt.Print("Tersedia (True/False): ")
	fmt.Scan(&T[idx].tersedia)

	default:
		fmt.Print("Pilihan tidak tersedia")
		return

	}
	fmt.Println("Data berhasil diubah")

}

// ================= HAPUS MENU =================
func hapusMenu(T *tabMenu, n *int) {
	var id, idx, i int

	fmt.Print("Masukkan ID yang dihapus: ")
	fmt.Scan(&id)

	idx = cariID(*T, *n, id)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	for i = idx; i < *n-1; i++ {
		T[i] = T[i+1]
	}

	*n = *n - 1

	fmt.Println("Data berhasil dihapus.")
}

// ================= SEQUENTIAL SEARCH =================
func sequentialSearch(T tabMenu, n int, kategori string) {
	var i int
	found := false

	fmt.Println("\n   ~Hasil Sequential Search~    ")

	for i = 0; i < n; i++ {
		if T[i].kategori == kategori {
			fmt.Println(T[i].nama, "-", T[i].harga)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

// ================= SELECTION SORT =================
func selectionSort(T *tabMenu, n int) {
	var i, j, min int
	var temp menu

	for i = 0; i < n-1; i++ {
		min = i

		for j = i + 1; j < n; j++ {
			if T[j].harga < T[min].harga {
				min = j
			}
		}

		temp = T[i]
		T[i] = T[min]
		T[min] = temp
	}
}

// ================= INSERTION SORT =================
func insertionSort(T *tabMenu, n int) {
	var i, j int
	var temp menu

	for i = 1; i < n; i++ {
		temp = T[i]
		j = i - 1

		for j >= 0 && T[j].harga > temp.harga {
			T[j+1] = T[j]
			j--
		}

		T[j+1] = temp
	}
}

// ================= BINARY SEARCH =================
func binarySearch(T tabMenu, n int, kategori string) {
	var left, right, mid int
	found := false

	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2

		if T[mid].kategori == kategori {
			found = true
			break
		} else if T[mid].kategori < kategori {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found {
		fmt.Println("Kategori ditemukan pada menu:", T[mid].nama)
	} else {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

// ================= STATISTIK =================
func statistik(T tabMenu, n int) {
	var i int
	var total int
	var coffee, noncoffee int
	var rata float64

	for i = 0; i < n; i++ {
		total += T[i].harga

		if T[i].kategori == "coffee" {
			coffee++
		} else {
			noncoffee++
		}
	}

	if n > 0 {
		rata = float64(total) / float64(n)
	}

	fmt.Println("\n   ~Statistik~    ")
	fmt.Println("Jumlah coffee     :", coffee)
	fmt.Println("Jumlah non-coffee :", noncoffee)
	fmt.Println("Rata-rata harga   :", rata)
}

// ================= MAIN =================
func main() {

	var T tabMenu
	var n int
	var pilih int
	var kategori string

	for {

		fmt.Println("\n  ~CAFE MENU~   ")
		fmt.Println("1. Tambah Menu")
		fmt.Println("2. Tampil Menu")
		fmt.Println("3. Ubah Menu")
		fmt.Println("4. Hapus Menu")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Binary Search")
		fmt.Println("7. Selection Sort")
		fmt.Println("8. Insertion Sort")
		fmt.Println("9. Statistik")
		fmt.Println("0. Exit")

		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			tambahMenu(&T, &n)

		case 2:
			tampilMenu(T, n)

		case 3:
			ubahMenu(&T, n)

		case 4:
			hapusMenu(&T, &n)

		case 5:
			fmt.Print("Kategori dicari: ")
			fmt.Scan(&kategori)
			sequentialSearch(T, n, kategori)

		case 6:
			fmt.Print("Kategori dicari: ")
			fmt.Scan(&kategori)
			binarySearch(T, n, kategori)

		case 7:
			selectionSort(&T, n)
			fmt.Println("Data berhasil diurutkan dengan Selection Sort.")

		case 8:
			insertionSort(&T, n)
			fmt.Println("Data berhasil diurutkan dengan Insertion Sort.")

		case 9:
			statistik(T, n)

		case 0:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}