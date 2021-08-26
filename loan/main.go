package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/health/live", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "up")
	})

	http.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "yes")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()                           // Parses the request body
		loanPostParameter := r.Form.Get("loan") // x will be "" if parameter is not set

		loanAmount, err := strconv.Atoi(loanPostParameter)
		if err != nil {
			loanAmount = 0
		}
		quote := offerQuote(loanAmount, 12)

		fmt.Fprintf(w, `<html>
		<form method="post">
		Enter your loan amount to see the interest. $<input name="loan" type="number" value="%d">
		<br/>
		<input type="submit">
		</form>
		
		<br/>
		%s
		
		</html>
		
		
		`, loanAmount, quote)
	})

	fmt.Println("Listening now at port 8080")
	http.ListenAndServe(":8080", nil)
}

func offerQuote(loan int, interest int) string {
	if loan <= 0 {
		return ""
	}

	total := loan * interest / 100
	return fmt.Sprintf("With rate %d%% you will pay  %d extra interest", interest, total)

}
