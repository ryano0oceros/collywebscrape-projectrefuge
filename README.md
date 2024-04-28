# collywebscrape-projectrefuge

## Description
This project is a web scraper that uses the Colly framework to scrape data from Wikipedia. The list of Wikipedia articles are statically set in main.go. The data that is being scraped is the URL of the Wikipedia article and the text sourced from the HTML body with <p>  . The data is then saved to a JSON lines file (output.jl).

## Installation

To install the project, you will need to have Go installed on your machine. You can download Go from the [official website](https://golang.org/).

Once you have Go installed, you can clone the repository and run the following command to install the project dependencies:

```bash
go mod download
```

## Usage

To run the project, you can use the following command:

```bash
go run main.go
```

This will start the web scraper and scrape the data from the website. The data will be saved to a file called `data.csv`.

## Outputs

The latest output of the project can be found in the `output.jl` file. main.go has a helper function that measures execution time and logs the time taken to scrape the data. The latest run took 1.226 seconds to scrape 10 articles.
