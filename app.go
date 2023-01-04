package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorings = 3
const delay = 5

func main() {
	sayIntro()
	for {
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
}

func sayIntro() {
	username := "hana"
	fmt.Println("hello", username)
}

func showMenu() {
	fmt.Println("")
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
	sites := readSites()

	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	response, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
	}
	nameSite := site[8:]
	if response.StatusCode == 200 {
		fmt.Println(nameSite, "is up!")
	} else {
		fmt.Println(nameSite, "is down! status code:", response.StatusCode)
	}
}

func showLogs() {
	fmt.Println("Showing logs...")
}

func readSites() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		sites = append(sites, line)
		fmt.Println(line)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error:", err)
		}
	}

	file.Close()

	return sites
}
