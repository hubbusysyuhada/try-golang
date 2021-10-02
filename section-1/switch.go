package main

import "fmt"

func main() {
	color := "grey"
	switch color {
	case "green":
		fmt.Println("lampu hijau, boleh jalan")
	case "yellow":
		fmt.Println("lampu kuning, siap-siap")
	case "red":
		fmt.Println("lampu merah, berhenti")
	default:
		fmt.Println("wrong input")
	}

	// no condition switch
	switch {
	case color == "green":
		fmt.Println("lampu hijau, boleh jalan")
	case color == "yellow":
		fmt.Println("lampu kuning, siap-siap")
	case color == "red":
		fmt.Println("lampu merah, berhenti")
	default:
		fmt.Println("wrong input")
	}

}
