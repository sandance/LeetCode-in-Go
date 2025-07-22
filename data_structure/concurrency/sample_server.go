package main

import (
	"fmt"
	"net/http"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

const form = `<html><body><form action="/" method="POST">
<h1>Statistics</h1>
<h5>Compute base statistics for a given list of numbers</h5>
<label for="numbers">Numbers (comma or space-separated):</label><br>
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form></body></html>`

const error = `<p class="error">%s</p>`

var pageTop = ""
var pageBottom = ""

// Define a root handler for requests to function homePage, and start the webserver combined with error-handling
func main() {

}

// Write an HTML header, parse the form, write form to writer and make request for numbers
func homePage(writer http.ResponseWriter, request *http.Request) {
	// write your code here'
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprint(writer, error, err)
	} else {
		
	}

}

// Capture the numbers from the request, and format the data and check for errors
func processRequest(request *http.Request) ([]float64, string, bool) {
	// write your code here
	return nil, "", false
}

// sort the values to get mean and median
func getStats(numbers []float64) (stats statistics) {
	// write your code here
	return stats
}

// seperate function to calculate the sum for mean
func sum(numbers []float64) (total float64) {
	// write your code here
	return 0
}

// seperate function to calculate the median
func median(numbers []float64) float64 {
	// write your code here
	return 0
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}
