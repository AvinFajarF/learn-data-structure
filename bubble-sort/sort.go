package main

import "fmt"

func bubbleSort(arr []int) {
	//  langkah pertama kita hitung panjang array nya
	n := len(arr)
	// langkah kedua deklarasi nilai boolean untuk for nya
	swap := true

	for swap {
		swap = false

		for i := 1; i < n; i++ {
			// pengecekan
			if arr[i-1] > arr[i] {
				// menukar urutan array nya
				arr[i-1], arr[i] = arr[i], arr[i-1]
				// karena condisi nya true maka, swap akan bernilai true juga
				swap = true
			}
		}
		n--
	}

}

func main() {
	// Data yang akan diurutkan
	data := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Data sebelum diurutkan:", data)
	bubbleSort(data)
	fmt.Println("Data setelah diurutkan:", data)
}
