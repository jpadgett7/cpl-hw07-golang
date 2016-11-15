package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// A brief description of program usage
const usage = `
This program accepts one or more URLs as positional arguments and
outputs the number of times the specified target word was found
on each page.

`

// Parses options passed on the command line.
//
// Returns the word to search for, the number of workers to use, and a
// slice of links to look at. If the CLI options do not pass a sanity
// check, an error message is shown followed by the program usage.
func parseCLI() (word string, numWorkers uint, links []string) {
	// Set the usage message for the cli parser
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] URL1 [URL2 [URL3 ...]]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, usage)
		flag.PrintDefaults()
	}

	// Setup the flags we're looking for
	flag.StringVar(&word, "word", "", "The word to search for.")
	flag.UintVar(&numWorkers, "workers", 1, "The number of workers to use.")

	// Parse the flags
	flag.Parse()

	if word == "" {
		fmt.Fprintf(os.Stderr, "Need a word to process.\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if numWorkers < 1 {
		fmt.Fprintf(os.Stderr, "Number of workers must be greater than 0.\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Need links to process.\n\n")
		flag.Usage()
		os.Exit(1)
	}

	links = flag.Args()
	return
}

// A Result represents the outcome from counting the occurrence of a
// word on a web page
type Result struct {
	link  string // the path to the page that was inspected
	count uint   // the number of occurrences of the word
	err   error  // an encountered error, if there was one (otherwise nil)
}

// Formats a Result as a string.
//
// String returns the string representation of the received Result
func (r Result) String() string {
	return fmt.Sprintf("%s\n\tcount: %d\n\terror: %v", r.link, r.count, r.err)
}

// countOccurrences counts the number of occurrences of `word` in
// `s`. It splits `s` into words and compares each word to `word`,
// counting the number of matches.
//
// It returns the number of occurences found while scanning and the
// first error encountered, if any.
func countOccurrences(word string, s io.Reader) (uint, error) {
	// Initialize Scanner and return variables
	scanner := bufio.NewScanner(s)
	scanner.Split(bufio.ScanWords)
	var counter uint = 0
	// Scan file word-by-word
	for scanner.Scan() {
		// If word matches a word found in file, increment counter
		if scanner.Text() == word {
			counter++
		}
	}
	return counter, scanner.Err()
}

// wordsOnPage reads links from the `links` channel searching for
// occurrences of `word` and sending Results over the `results`
// channel.
//
// wordsOnPage continually receives values from the `links`
// channel. When it receives a value, it fetches the page data, counts
// occurrences of `word` on the page, and sends results over
// `results`. If any error is encountered, it is packed into a Result
// and sent over the `results` channel. When there are no more links
// to read, the function returns.
func wordsOnPage(word string, links chan string, results chan Result) {
	// Loop over every entry in buffered channel links
	for link := range links {
		// Issue a GET to the url, store data in resp, error in err
		resp, err := http.Get(link)
		if err != nil {
			// Get() returns error, URL not open
			results <- Result{link, 0, err}
		} else if resp.Status != "200 OK" {
			// Something went wrong opening URL
			resErr := errors.New("Did not receive 200 OK")
			results <- Result{link, 0, resErr}
		} else {
			// No errors detected
			wordCount, resErr := countOccurrences(word, resp.Body)
			results <- Result{link, wordCount, resErr}
		}
	}
	return
}

//Parses CLI args. Spins up workers (goroutines) as specified by the
//user and sets up channels for communication. Sends links over
//channel for processing, closes the channel, then reads all results
//from goroutines.
func main() {
	searchTerm, numWorkers, links := parseCLI()
	c_links := make(chan string, len(links))
	c_results := make(chan Result, len(links))
	// Create our workers (goroutines) for processing urls
	for i := uint(0); i < numWorkers; i++ {
		go wordsOnPage(searchTerm, c_links, c_results)
	}
	// Add urls to links channel
	for j := 0; j < len(links); j++ {
		c_links <- links[j]
	}
	close(c_links) // Leave this alone now
	// Print whatever the workers bring back
	for k := 0; k < len(links); k++ {
		fmt.Println(<-c_results)
	}
	close(c_results) // All done, closing up shop
	return
}
