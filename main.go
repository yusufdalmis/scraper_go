package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	url_pointer := flag.String("url", "", "Hedef Web Sitesi")
	flag.Parse()

	if *url_pointer == "" {
		log.Fatal("boş olamamlı => go run main.go -url https://www.google.com")
	}

	targetURL := *url_pointer
	

	fmt.Printf("Hedef: %s için işlem başlatılıyor\n", targetURL)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.WindowSize(1920, 1080)	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()


	var html_content string
	var buf []byte

	error := chromedp.Run(ctx,
		chromedp.Navigate(targetURL),
		chromedp.Sleep(2*time.Second), 
		chromedp.OuterHTML("html", &html_content),
		chromedp.FullScreenshot(&buf, 90)	)

	if error != nil {
		log.Fatalf("işlem tamamlanamadı: %v", error)
	}

	

	filenametxt := safeFileName(targetURL,"", ".txt")
	filenamepng := safeFileName(targetURL,"", ".png")


	error = os.WriteFile(filenametxt, []byte(html_content), 0644)
	if error != nil {
		log.Fatalf("HTML kaydedilemedi: %v", error)
	}
	fmt.Printf("HTML içeriği '%s' kaydedildi\n",filenametxt)

	
	error = os.WriteFile(filenamepng, buf, 0644)
	if error != nil {
		log.Fatalf("Ekran görüntüsü kaydedilemedi: %v", error)
	}
	fmt.Printf(" Ekran görüntüsü '%s' kaydedildi\n",filenamepng)

	//extra
	filenamelinktxt := safeFileName(targetURL,"_link", ".txt")

	links := extractLinks(html_content)
	linksText := strings.Join(links, "\n")

	error = os.WriteFile(filenamelinktxt,  []byte(linksText), 0644)
	if error != nil {
	log.Fatalf("Linkler kaydedilemedi: %v", error)
	}
	fmt.Printf(" Linkler '%s' kaydedildi\n",filenametxt)

}


func extractLinks(html string) []string{

	re := regexp.MustCompile(`href=["'](http[^"']+)["']`)
	matches := re.FindAllStringSubmatch(html, -1)

	var links []string

	
	for idx, match := range matches { 
		if len(match) > 1 {
			links = append(links, fmt.Sprintf("%d\t%s", idx, match[1]))
		}
		
	}
	
	return links
}

func safeFileName(url string, extra string,  suffix string ) string {
	name := strings.TrimPrefix(url, "https://www.")
	name = strings.TrimPrefix(name, "http://www.")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, ":", "")
	name = strings.ReplaceAll(name, ".", "_")
	return name+ extra+ suffix
}