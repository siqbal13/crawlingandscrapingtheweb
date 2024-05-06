package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PageContent struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var Urls = []string{
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

func main() {
	//output file
	file, err := os.Create("output1.jl")
	if err != nil {
		fmt.Println("Cannot create output file")
	}
	defer file.Close()
	//only access to wikipedia domain
	c := colly.NewCollector(colly.AllowedDomains(("en.wikipedia.org")))

	data := PageContent{}

	//err handling
	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while crawling:%s\n", e.Error())
	})

	//get header <h1>
	c.OnHTML("h1", func(a *colly.HTMLElement) {
		data.Title = a.Text
	})

	c.OnHTML("p", func(h *colly.HTMLElement) {
		data.Text += h.Text + "\n"

	})
	
	c.OnScraped(func(r *colly.Response) {
		data.URL = r.Request.URL.String()
		err := CreateJson(file, data)
		if err != nil {
			log.Println("Error writing to file:", err)
		}
	})

	for _, url := range Urls {
		c.Visit(url)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func CreateJson(file *os.File, data PageContent) error {
	//encoder and write data into it
	encoder := json.NewEncoder(file)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println("error in convert json line file")
	}
	return nil
}