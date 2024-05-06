# Web Crawling Assignment

This assignment involves using Colly to extract paragraph text from webpages and write the extracted data into a JSON line file.

## Output
The output consists of text content collected from the provided URLs. Along with the text content, the output file also includes the following fields: "url", "title", and "text" (which is the main content).

## Issues
There are several issues in this simple application. As Colly runs concurrently, the variable `data` may cause data discrepancy.

## Performance
In comparison to a Python crawler, experiments run purely on the terminal with "time python run-articles-spider.py" and "time go run main.go" showed the following results:
- Python: 5.53s user 1.79s system 31% cpu 23.178 total
- Go: 1.01s user 0.97s system 65% cpu 3.009 total

The issue with this experiment is that the Go code only extracts text content, which is an inaccurate measurement compared to Python. However, it still shows a significant difference in the performance between the two languages.