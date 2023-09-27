package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	// Menyimpan referensi ke node kedua
	second := l.head

	// Mengganti node pertama dengan node baru
	l.head = n

	// Menghubungkan node baru ke node kedua
	l.head.next = second

	// Menambah panjang linked list
	l.length++
}

// printListData mencetak semua data dalam linked list ke layar.
func (l linkedList) printListData() {
	// Inisialisasi variabel toPrint dengan node pertama (head) dari linked list
	toPrint := l.head

	// Melakukan perulangan selama panjang linked list tidak sama dengan 0
	for l.length != 0 {
		// Mencetak data dari node saat ini
		fmt.Printf("%d ", toPrint.data)

		// Melanjutkan ke node berikutnya dalam linked list
		toPrint = toPrint.next

		// Dekrement panjang linked list
		l.length--
	}

	// Mencetak baris kosong sebagai pemisah setelah mencetak seluruh linked list
	fmt.Print("\n")
}

// deleteLinkedList menghapus node dengan nilai tertentu dari linked list.
// Jika node dengan nilai yang sesuai ditemukan, itu akan dihapus dari linked list.
// Jika linked list kosong atau tidak ada node dengan nilai yang sesuai, tidak akan ada tindakan yang diambil.
func (l *linkedList) deleteLinkedList(value int) {
	// Simpan referensi ke node yang akan dihapus
	previusToDelete := l.head

	// Jika linked list kosong, tidak ada yang perlu dihapus
	if l.length == 0 {
		return
	}

	// Jika node yang akan dihapus adalah node pertama
	if l.head.data == value {
		l.head = l.head.next // Mengganti head dengan node berikutnya
		l.length--           // Mengurangi panjang linked list
		return
	}

	// Iterasi melalui linked list untuk menemukan node sebelum node yang akan dihapus
	for previusToDelete.next.data != value {

		// Jika kita mencapai akhir linked list tanpa menemukan node yang sesuai, keluar
		if previusToDelete.next.next == nil {
			return
		}

		previusToDelete = previusToDelete.next
	}

	// Menghapus node dengan mengganti referensi ke node sebelumnya
	previusToDelete.next = previusToDelete.next.next

	// Mengurangi panjang linked list
	l.length--
}

func main() {

	myList := linkedList{}
	node1 := &node{data: 30}
	node2 := &node{data: 41}
	node3 := &node{data: 70}
	node4 := &node{data: 50}
	node5 := &node{data: 10}
	node6 := &node{data: 90}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.printListData()
	myList.deleteLinkedList(10)
	myList.printListData()
}
