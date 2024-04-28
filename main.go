package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Page struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

func main() {
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	measureExecutionTime(func() { scrapePages(urls) }, "Scraping pages")

}

func scrapePages(urls []string) {
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		page := Page{
			URL:  e.Request.URL.String(),
			Text: e.ChildText("p"),
		}

		file, err := os.OpenFile("output.jl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(page); err != nil {
			fmt.Println("Error encoding JSON:", err)
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	for _, url := range urls {
		c.Visit(url)
	}
}

func measureExecutionTime(f func(), label string) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", label, elapsed)
}
