package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	var ayam string = "adu"

	if length := len(ayam); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}