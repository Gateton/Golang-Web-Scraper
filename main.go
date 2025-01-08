package main

import (
  "encoding/json"
  "fmt"
  "log"
  "os"
  "strings"

  "github.com/gocolly/colly/v2"
  "github.com/spf13/cobra"
)

type ScrapedData struct {
  Titles  []string `json:"titles"`
  Prices  []string `json:"prices"`
  Images  []string `json:"images"`
}

var (
  url         string
  outputFile  string
  headers     []string
  cookies     []string
)

var rootCmd = &cobra.Command{
  Use:   "webscraper",
  Short: "A flexible web scraper CLI tool",
  Run:   runScraper,
}

func init() {
  rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL to scrape")
  rootCmd.Flags().StringVarP(&outputFile, "output", "o", "output.json", "Output JSON file")
  rootCmd.Flags().StringSliceVarP(&headers, "header", "H", []string{}, "Custom headers")
  rootCmd.Flags().StringSliceVarP(&cookies, "cookie", "C", []string{}, "Custom cookies")
  rootCmd.MarkFlagRequired("url")
}

func runScraper(cmd *cobra.Command, args []string) {
  c := colly.NewCollector()

  data := ScrapedData{}

  // Set custom headers
  for _, header := range headers {
    parts := strings.SplitN(header, ":", 2)
    if len(parts) == 2 {
      c.OnRequest(func(r *colly.Request) {
        r.Headers.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
      })
    }
  }

  // Set custom cookies
  for _, cookie := range cookies {
    parts := strings.SplitN(cookie, "=", 2)
    if len(parts) == 2 {
      c.SetCookies(url, []*http.Cookie{{
        Name:  strings.TrimSpace(parts[0]),
        Value: strings.TrimSpace(parts[1]),
      }})
    }
  }

  // Scrape titles
  c.OnHTML("h2.product-title", func(e *colly.HTMLElement) {
    data.Titles = append(data.Titles, e.Text)
  })

  // Scrape prices
  c.OnHTML("span.price", func(e *colly.HTMLElement) {
    data.Prices = append(data.Prices, e.Text)
  })

  // Scrape image URLs
  c.OnHTML("img.product-image", func(e *colly.HTMLElement) {
    data.Images = append(data.Images, e.Attr("src"))
  })

  c.OnError(func(r *colly.Response, err error) {
    log.Printf("Error scraping '%s': %s", r.Request.URL, err)
  })

  err := c.Visit(url)
  if err != nil {
    log.Fatalf("Error visiting URL: %s", err)
  }

  // Save data to JSON file
  jsonData, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    log.Fatalf("Error marshaling JSON: %s", err)
  }

  err = os.WriteFile(outputFile, jsonData, 0644)
  if err != nil {
    log.Fatalf("Error writing to file: %s", err)
  }

  fmt.Printf("Scraping completed. Data saved to %s\n", outputFile)
}

func main() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
