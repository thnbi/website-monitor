package main

import (
	"fmt"
	"os"
)

func main() {
	sayIntro()
	showMenu()
	option := readOption()

	switch option {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
		os.Exit(-1)
	}
}

func sayIntro() {
	username := "hana"
	fmt.Println("hello", username)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)
	return option
}

func startMonitoring() {
	fmt.Println("Monitoring...")

}

func showLogs() {
	fmt.Println("Showing logs...")
}