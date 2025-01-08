# Web Scraper CLI Tool

This is a flexible web scraper CLI tool built with Node.js. It allows you to extract specific information from web pages and save it in a JSON format.

## Features

- Scrape titles, prices, and image URLs from web pages
- Flexible configuration with custom headers and cookies
- Save scraped data to a JSON file
- Command-line interface for easy use

## Prerequisites

- Node.js (v12 or later)

## Installation

1. Make sure you have Node.js installed on your system.
2. Clone this repository or download the source code.
3. Navigate to the project directory and run:

```
npm install
```

## Usage

```
node index.js -u <URL> [-o <output_file>] [-H <header>...] [-C <cookie>...]
```

Arguments:
- `-u, --url`: URL to scrape (required)
- `-o, --output`: Output JSON file (default: "output.json")
- `-H, --header`: Custom headers (can be used multiple times)
- `-C, --cookie`: Custom cookies (can be used multiple times)

Example:
```
node index.js -u https://example.com -o results.json -H "User-Agent: MyBot/1.0" -C "session=abc123"
```

## Output

The scraped data will be saved in the specified JSON file with the following structure:

```json
{
  "titles": ["Product 1", "Product 2", ...],
  "prices": ["$10.99", "$24.99", ...],
  "images": ["https://example.com/image1.jpg", "https://example.com/image2.jpg", ...]
}
```

## Customization

You can customize the scraping behavior by modifying the cheerio selectors in the `index.js` file. Adjust the selectors to match the structure of the websites you want to scrape.

## Error Handling

The scraper includes basic error handling for connection and scraping errors. These errors will be logged to the console.


