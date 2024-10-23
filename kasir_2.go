//ini adalah tubes Ahmad Raja Fadhil
package main

// komentar
import (
	"fmt"
	"time"
)

// Maksimum jumlah barang dan transaksi
const maxItems = 100
const maxTransactions = 100

// Struct untuk menyimpan data barang
type Item struct {
	Code  string
	Name  string
	Price int
}

// Struct untuk menyimpan data transaksi
type Transaction struct {
	Date       string
	Items      []Item
	Quantities []int // Menyimpan jumlah barang
	Total      int
}

// Array statis untuk menyimpan barang dan transaksi
var items [maxItems]Item
var transactions [maxTransactions]Transaction
var itemCount, transactionCount int

func main() {
	for {
		displayMainMenu()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handleItemManagement()
		case 2:
			handleTransaction()
		case 3:
			displayTransactionsAndOmzet()
		case 4:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

// displayMainMenu menampilkan menu utama kepada pengguna
func displayMainMenu() {
	fmt.Println("Menu Utama:")
	fmt.Println("1. Kelola Barang")
	fmt.Println("2. Catat Transaksi")
	fmt.Println("3. Lihat Transaksi dan Omzet Harian")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih menu: ")
}

// handleItemManagement menangani submenu pengelolaan barang
func handleItemManagement() {
	fmt.Println("1. Tambah Barang")
	fmt.Println("2. Ubah Barang")
	fmt.Println("3. Hapus Barang")
	fmt.Print("Pilih submenu: ")
	var subChoice int
	fmt.Scan(&subChoice)

	switch subChoice {
	case 1:
		addNewItem()
	case 2:
		updateExistingItem()
	case 3:
		deleteExistingItem()
	default:
		fmt.Println("Submenu tidak valid.")
	}
}

// addNewItem menambah barang baru ke dalam daftar barang
func addNewItem() {
	var code, name string
	var price int
	fmt.Print("Input Kode Barang: ")
	fmt.Scan(&code)
	fmt.Print("Input Nama Barang: ")
	fmt.Scan(&name)
	fmt.Print("Input Harga Barang: ")
	fmt.Scan(&price)

	addItem(code, name, price)
	displayItems()
}

// updateExistingItem memperbarui data barang yang ada
func updateExistingItem() {
	var code, name string
	var price int
	fmt.Print("Input Kode Barang yang ingin diubah: ")
	fmt.Scan(&code)
	index := binarySearch(code)
	if index != -1 {
		fmt.Print("Input Nama Barang Baru: ")
		fmt.Scan(&name)
		fmt.Print("Input Harga Barang Baru: ")
		fmt.Scan(&price)

		updateItem(index, code, name, price)
		displayItems()
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

// deleteExistingItem menghapus barang dari daftar barang
func deleteExistingItem() {
	var code string
	fmt.Print("Input Kode Barang yang ingin dihapus: ")
	fmt.Scan(&code)
	index := binarySearch(code)
	if index != -1 {
		deleteItem(index)
		displayItems()
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

// handleTransaction mencatat transaksi baru
func handleTransaction() {
	var itemsInTransaction []Item
	var quantities []int
	total := 0
	for {
		var code string
		fmt.Print("Input Kode Barang (ketik 'selesai' untuk menyelesaikan): ")
		fmt.Scan(&code)
		if code == "selesai" {
			break
		}
		index := binarySearch(code)
		if index != -1 {
			var quantity int
			fmt.Print("Input Jumlah Barang: ")
			fmt.Scan(&quantity)
			item := items[index]
			itemsInTransaction = append(itemsInTransaction, item)
			quantities = append(quantities, quantity)
			total += item.Price * quantity
		} else {
			fmt.Println("Barang tidak ditemukan.")
		}
	}
	addTransaction(itemsInTransaction, quantities, total)
	fmt.Println("Transaksi berhasil dicatat.")
}

// displayItems menampilkan daftar barang yang ada
func displayItems() {
	fmt.Println("Daftar Barang:")
	for i := 0; i < itemCount; i++ {
		fmt.Printf("%d. %s - %s: %d\n", i+1, items[i].Code, items[i].Name, items[i].Price)
	}
}

// displayTransactionsAndOmzet menampilkan daftar transaksi dan omzet harian
func displayTransactionsAndOmzet() {
	displayTransactions()
	omzet := 0
	for i := 0; i < transactionCount; i++ {
		omzet += transactions[i].Total
	}
	fmt.Println("Omzet Harian:", omzet)
}

// displayTransactions menampilkan daftar transaksi yang ada
func displayTransactions() {
	fmt.Println("Daftar Transaksi:")
	for i := 0; i < transactionCount; i++ {
		fmt.Printf("%d. %s\n", i+1, transactions[i].Date)
		for j, item := range transactions[i].Items {
			fmt.Printf("  %s - %s: %d x %d = %d\n", item.Code, item.Name, transactions[i].Quantities[j], item.Price, transactions[i].Quantities[j]*item.Price)
		}
		fmt.Printf("Total: %d\n", transactions[i].Total)
	}
}

// addItem menambahkan barang baru ke dalam array items
func addItem(code, name string, price int) {
	items[itemCount] = Item{code, name, price}
	itemCount++
	selectionSort(items[:itemCount], true)
}

// updateItem memperbarui data barang yang ada di dalam array items
func updateItem(index int, code, name string, price int) {
	items[index] = Item{code, name, price}
	selectionSort(items[:itemCount], true)
}

// deleteItem menghapus barang dari dalam array items
func deleteItem(index int) {
	for i := index; i < itemCount-1; i++ {
		items[i] = items[i+1]
	}
	itemCount--
	selectionSort(items[:itemCount], true)
}

// sequentialSearch mencari barang secara sequential berdasarkan kode
func sequentialSearch(code string) int {
	for i := 0; i < itemCount; i++ {
		if items[i].Code == code {
			return i
		}
	}
	return -1
}

// binarySearch mencari barang secara binary berdasarkan kode
func binarySearch(code string) int {
	low, high := 0, itemCount-1
	for low <= high {
		mid := (low + high) / 2
		if items[mid].Code == code {
			return mid
		} else if items[mid].Code < code {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// addTransaction menambahkan transaksi baru ke dalam array transactions
func addTransaction(items []Item, quantities []int, total int) {
	transactions[transactionCount] = Transaction{
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		Items:      items,
		Quantities: quantities,
		Total:      total,
	}
	transactionCount++
}

// selectionSort mengurutkan array items berdasarkan kode secara ascending atau descending
func selectionSort(data []Item, ascending bool) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minMaxIdx := i
		for j := i + 1; j < n; j++ {
			if (ascending && data[j].Code < data[minMaxIdx].Code) || (!ascending && data[j].Code > data[minMaxIdx].Code) {
				minMaxIdx = j
			}
		}
		data[i], data[minMaxIdx] = data[minMaxIdx], data[i]
	}
}

// insertionSort mengurutkan array transactions berdasarkan tanggal secara ascending atau descending
func insertionSort(data []Transaction, ascending bool) {
	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && ((ascending && data[j].Date > key.Date) || (!ascending && data[j].Date < key.Date)) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
