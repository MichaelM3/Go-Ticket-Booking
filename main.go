package main

import (
    "fmt"
    "strings"
)

func main() {
    const conferenceName = "Go Conference"
    const conferenceTickets = 50
    var remainingTickets uint = 50
    bookings := []string{}

    // fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")
   
    for remainingTickets > 0 {
        var firstName string
        var lastName string
        var email string
        var userTickets uint
        // ask user for their name
        fmt.Println("Enter your first name: ")
        fmt.Scan(&firstName)

        fmt.Println("Enter your last name: ")
        fmt.Scan(&lastName)

        fmt.Println("Enter your email: ")
        fmt.Scan(&email)

        fmt.Println("Enter number of tickets: ")
        fmt.Scan(&userTickets)

        isValidName := len(firstName) >= 2 && len(lastName) >= 2

        if userTickets <= remainingTickets {
            remainingTickets = remainingTickets - userTickets
            bookings = append(bookings, firstName + " " + lastName)
            // bookings[0] = firstName + " " + lastName

            // fmt.Printf("The first value: %v\n", bookings[0])
            // fmt.Printf("Slice type: %T\n", bookings)
            // fmt.Printf("Slice length %v\n", len(bookings))
            // fmt.Printf("The whole slice: %v\n", bookings)

            fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
            fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

            firstNames := []string{}
            for _, booking := range bookings {
                var names = strings.Fields(booking)        
                firstNames = append(firstNames, names[0])
            }

            fmt.Printf("These are all our bookings: %v\n", firstNames)

            // if remainingTickets == 0 {
            //     // End the program
            //     fmt.Println("Our conference is booked out. Come back next year.")
            //     break
            // }
        } else {
            fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
            // continue
        }
    }
}
