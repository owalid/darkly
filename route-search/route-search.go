package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func fileNotExists(fileName string) bool {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	return os.IsNotExist(error)
}

func main() {
	// get host from args
	if len(os.Args) < 2 {
		log.Fatal("Please provide a host")
	}
	host := os.Args[1]
	var baseUrl = ""

	// check if host have http or https
	if host[:4] != "http" && host[:5] != "https" {
		baseUrl = "http://" + host
	} else {
		baseUrl = host
	}

	log.Printf("[+] Scanning %s", host)
	file := "dir-list-lowercase.txt"

	if fileNotExists(file) {
		log.Printf("[*] Dictinary %s does not exist", file)
		log.Printf("[+] Download dictonnary.....")
		cmd := exec.Command("/bin/sh", "download_dir_dict.sh")
		_, err := cmd.Output()

		if err != nil {
			log.Println(err.Error())
			return
		}

		// Print the output
		log.Printf("[+] Dictonnary downloaded")
	}

	log.Printf("[+] Open dictonnary")
	dirDict, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(dirDict)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		route := fileScanner.Text()

		var url = baseUrl + "/" + route
		ret, err := http.Get(url)

		// get html from ret
		bodyBytes, _ := io.ReadAll(ret.Body)
		bodyString := string(bodyBytes)

		// if not have 404 in bodyString
		if err == nil && ret.StatusCode != 404 && !strings.Contains(bodyString, "404") {
			log.Println("[ " + strconv.Itoa(ret.StatusCode) + " ] /" + route)
		}
	}
}
