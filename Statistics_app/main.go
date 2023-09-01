package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
</form></html></body>`

const error = `<p class="error">%s</p>`

var pageTop = ""
var pageBottom = ""

// Define a root handler for requests to function homePage, and start the webserver combined with error-handling
func main() {

	http.HandleFunc("/", homePage)           // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// Write an HTML header, parse the form, write form to writer and make request for numbers
func homePage(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(writer, form)
	case "POST":
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data with Form */

		numbers, _, _ := processRequest(request)
		fmt.Fprint(writer, formatStats(getStats(numbers)))

	}

}

// Capture the numbers from the request, and format the data and check for errors
func processRequest(request *http.Request) ([]float64, string, bool) {

	var stat statistics

	for _, v := range request.FormValue("numbers") {
		if unicode.IsNumber(v) != true && v != 32 && v != 44 {
			fmt.Println("Please put numbers with spaces and commas only")
			return nil, "Please put numbers with spaces and commas only", false
		}
	}

	for _, v := range strings.FieldsFunc(request.FormValue("numbers"), Split) {
		v, _ := strconv.ParseFloat(v, 64)
		stat.numbers = append(stat.numbers, v)
	}

	fmt.Println(stat.numbers)

	return nil, "Results", true
}

func Split(r rune) bool {
	return r == ' ' || r == ','
}

// sort the values to get mean and median
func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return
}

// seperate function to calculate the sum for mean
func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return
}

// seperate function to calculate the median
func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
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
