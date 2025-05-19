package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// ↓リンク記載
	baseURL := "https://hoge/index.php?p=hoge&s=16&b="
	re := regexp.MustCompile(`T(\d+).*?\((\w)\)`)

	log.Printf("start")

	// ↓paramのbをincrementしてページごとに処理。適宜変更
	for b := 27; b <= 40; b++ {
		log.Printf("start page %d", b)

		url := baseURL + strconv.Itoa(b)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Error fetching URL %s: %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Printf("Non-200 response from %s: %d\n", url, resp.StatusCode)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Printf("Error parsing HTML from %s: %v\n", url, err)
			continue
		}

		// ↓htmlのclassに対して、valueを取得して適宜実装している。適宜変更。実装に関してはgithub.com/PuerkitoBio/goqueryを参照する
		doc.Find(".list-box").Each(func(i int, s *goquery.Selection) {
			size := s.Find(".size").Text()
			ageStr := s.Find(".age").Text()
			age, err := strconv.Atoi(ageStr[1:3])
			if err != nil {
				return
			}
			match := re.FindStringSubmatch(size)
			if match == nil {
				return
			}
			if len(match) < 3 {
				log.Printf("unexpected size text %s \n", size)
				return
			}

			name := s.Find(".name").Text()
			tall, err := strconv.Atoi(match[1])
			if err != nil {
				log.Printf("Error converting %s to int: %v\n", match[1], err)
				return
			}

			if (tall >= 157 && tall < 160) && (age > 20 && age < 23) {
				link, _ := s.Find("a").Attr("href")
				fmt.Printf("名前: %s\nサイズ: %s\nリンク: %s\n\n", name, size, link)
			}
		})
	}
}
