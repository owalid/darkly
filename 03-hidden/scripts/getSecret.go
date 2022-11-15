package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func getUrlsFromPage(url string) (urls []string) {
	page, err := getPage(url)
	if err != nil {
		log.Fatal(err)
	}

	// Parse html and get all pages
	return parse(url, page)
}

func work(urls []string) bool {
	var newUrls []string
	// Get the page
	for _, url := range urls {
		urlsFormPage := getUrlsFromPage(url)
		for _, urlFromPage := range urlsFormPage {
			newUrls = append(newUrls, urlFromPage)
		}
		// If the page contains README, check if it contains looseStrings
		readMe, err := getPage(url + "README")

		// If it contains looseStrings, call work() with the url of the pages.
		if err != nil {
			log.Fatal(err)
		}

		if !strings.Contains(readMe, "Tu veux de l'aide ? Moi aussi !") && !strings.Contains(readMe, "Demande à ton voisin") && !strings.Contains(readMe, "Non ce n'est toujours pas") && !strings.Contains(readMe, "Toujours pas tu vas") {
			log.Printf("Found the secret: %s", readMe)
			return true
		}
		log.Printf("[+] Page: %s \n", url)
	}
	return work(newUrls)
}

func parse(baseUrl string, text string) (data []string) {
	tkn := html.NewTokenizer(strings.NewReader(text))

	var vals []string
	var isA bool

	for {
		tt := tkn.Next()
		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()
			isA = t.Data == "a"

		case tt == html.TextToken:

			t := tkn.Token()

			if isA && t.Data != "../" && t.Data != "README" {
				vals = append(vals, baseUrl+t.Data)
			}

			isA = false
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a host")
	}
	var baseUrl = ""

	// check if host have http or https
	if os.Args[1][:4] != "http" && os.Args[1][:5] != "https" {
		baseUrl = "http://" + os.Args[1]
	} else {
		baseUrl = os.Args[1]
	}

	var end bool
	for {
		urlsStart := getUrlsFromPage(baseUrl + "/.hidden/")
		end = work(urlsStart)
		if end {
			break
		}
	}
}
