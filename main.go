package main //Golang is open source, mixture of C and python

import (
	"fmt"
	"os"

	// "html/template"
	// "log"
	"net/http"
	"strings"
	// "github.com/gin-gonic/gin"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = []string{}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "<html><body>")
		fmt.Fprintf(w, "<h1>Welcome to %v booking application</h1>", conferenceName)
		fmt.Fprintf(w, "<p>Number of tickets available are %v and the remaining ones are %v</p>", conferenceTickets, remainingTickets)
		fmt.Fprint(w, "<hr>")
		fmt.Fprint(w, "<h2>Get your tickets here to attend the conference</h2>")
		fmt.Fprint(w, "<form method='POST'>")
		http.ServeFile(w, r, "index.html")
	case "POST":
		r.ParseForm()
		firstName := r.FormValue("first-name")
		lastName := r.FormValue("last-name")
		email := r.FormValue("email")
		ticketNumber := r.FormValue("ticket-number")

		fileName := "bookings.txt"
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Write data to the file
		data := fmt.Sprintf("First Name: %s\nLast Name: %s\nEmail: %s\nNumber of Tickets: %s\n\n\n", firstName, lastName, email, ticketNumber)
		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println("Error writing data to file:", err)
			return
		}

		fmt.Fprint(w, "Booking successful")

	}
}

func main() {
	greetUsers()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)

	// r := gin.Default()

	// // Route to render HTML page
	// r.GET("/", func(c *gin.Context) {
	// 	tmpl, err := template.ParseFiles("index.html")
	// 	if err != nil {
	// 		c.AbortWithError(http.StatusInternalServerError, err)
	// 		return
	// 	}
	// 	data := gin.H{
	// 		"Bookings": bookings,
	// 	}
	// 	err = tmpl.Execute(c.Writer, data)
	// 	if err != nil {
	// 		c.AbortWithError(http.StatusInternalServerError, err)
	// 		return
	// 	}

	// })

	// go func() {
	// 	if err := r.Run(":8000"); err != nil {
	// 		log.Fatalf("Failed to start server: %v", err)
	// 	}
	// }()

	// // Start the server
	// r.Run()

	go func() {

		for {
			firstName, lastName, email, userTicket := getUserInput()

			isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remainingTickets)

			if isValidName && isValidEmail && isValidTicketNumber {

				bookTicket(remainingTickets, userTicket, &bookings, firstName, lastName, email, conferenceName)
				remainingTickets -= userTicket
				// if userTicket > remainingTickets{
				// 	fmt.Printf("We only have %v tickets available, so you can't book %v tickets\n ", remainingTickets, userTicket)
				// 	break // loop will break here
				// 	// continue  **This will not break and jump to next iteration in case user wants to correct number of tickets**
				// }

				// fmt.Printf("The whole booking array: %v\n", bookings)
				// fmt.Printf("The first value booking array: %v \n", bookings[0])
				// fmt.Printf("Array type: %T \n", bookings)
				// fmt.Printf("Array length: %v \n", len(bookings))
				// fmt.Printf("Slice type: %T \n", bookings)
				// fmt.Printf("Slice length: %v \n", len(bookings))

				// fmt.Printf("These all are booking: %v\n", bookings)

				firstNames := getFirstName(bookings)
				fmt.Printf("The first names of bookings are: %v\n", firstNames)

				if remainingTickets == 0 {
					fmt.Println("Our Conference is booked out. Come back next year")
					break

				}
			} else {
				if !isValidName {
					fmt.Println("You entered wrong name, please check")

				}
				if !isValidEmail {
					fmt.Println("You entered wrong mail id, please check")
				}
				if !isValidTicketNumber {
					fmt.Println("You entered wrong Ticket Number, please check")
				}

				fmt.Println("Your input data is invalid, try again")
			}
		}

	}()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Number of tickets available are %v and the remaining ones are %v \n", conferenceTickets, remainingTickets)

	fmt.Print("*** Get your tickets here to attend the conference *** \n")

}

func getFirstName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	println("Enter your first name")
	fmt.Scan(&firstName) // & it is used for pointer in Golang

	println("Enter your last name")
	fmt.Scan(&lastName)

	println("Enter your email")
	fmt.Scan(&email)

	println("Enter number of tickets you want to book")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(remainingTickets uint, userTicket uint, bookings *[]string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTicket

	fullname := firstName + " " + lastName

	//**for appending the array when we dnt knw the size of array** i.e., dynamic list
	*bookings = append(*bookings, fullname)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will get confirmation at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
