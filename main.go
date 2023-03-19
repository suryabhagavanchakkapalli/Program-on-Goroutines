package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type trainTicket struct {
	trainName    string
	departure    time.Time
	departureStr string
	destination  string
}

type ticketRequest struct {
	passengerName string
	ticket        trainTicket
}

const (
	numBookingAgents = 3
)

func main() {

	fmt.Println("Booking Tickets between Visakhapatnam and Secunderabad")

	// Create a wait group to track when all booking agents have finished.
	var wg sync.WaitGroup
	wg.Add(numBookingAgents)

	// Create a channel to hold the ticket requests to be processed.
	requestCh := make(chan ticketRequest)

	// Start the booking agent goroutines.
	for i := 0; i < numBookingAgents; i++ {
		go bookingAgent(i, &wg, requestCh)
	}

	// Define a map to keep track of the number of tickets booked to each train.
	ticketsBooked := make(map[string]int)

	ticketsBookedInEachCity := make(map[string]int)

	// Send the ticket requests to the booking agent goroutines.
	for i := 0; i < 50; i++ {
		// Simulate some delay between ticket requests.
		time.Sleep(500 * time.Millisecond)

		// Generate random train name and city name based on a random number.
		var trainName string
		switch rand.Intn(5) {
		case 0:
			trainName = "Vandemataram"
		case 1:
			trainName = "Janmabhomi"
		case 2:
			trainName = "LTT"
		case 3:
			trainName = "Garib Rath"
		case 4:
			trainName = "Godavari"
		}

		var cityName string
		switch rand.Intn(5) {
		case 0:
			cityName = "Visakhapatnam"
		case 1:
			cityName = "Annavaram"
		case 2:
			cityName = "Rajahmundry"
		case 3:
			cityName = "Vijayawada"
		case 4:
			cityName = "Secunderabad"
		}

		departure := time.Now().Add(time.Duration(rand.Int63n(60)) * time.Minute)
		departureStr := departure.Format("2006-01-02 15:04")

		ticket := trainTicket{
			trainName:    trainName,
			departure:    departure,
			departureStr: departureStr,
			destination:  cityName,
		}

		request := ticketRequest{
			passengerName: fmt.Sprintf("PASSENGER%d", i),
			ticket:        ticket,
		}

		requestCh <- request

		// Increment the count for the corresponding train.
		ticketsBooked[trainName]++
		ticketsBookedInEachCity[ticket.destination]++
	}

	// Close the request channel to signal to the booking agents that there are no more requests.
	close(requestCh)

	// Wait for all booking agents to finish.
	wg.Wait()

	fmt.Println("All tickets booked.")

	// fmt.Println(ticketsBooked)
	for trainName, numTickets := range ticketsBooked {
		fmt.Printf("Tickets booked for train %s: %d\n", trainName, numTickets)
	}

	for destination, numTickets := range ticketsBookedInEachCity {
		fmt.Printf("Tickets booked for city %s: %d\n", destination, numTickets)
	}
}

func bookingAgent(id int, wg *sync.WaitGroup, requestCh <-chan ticketRequest) {
	// Seed the random number generator with the booking agent ID.
	rand.Seed(time.Now().UnixNano() + int64(id))

	// Create a slice to hold the waiting list requests.
	var waitingList []ticketRequest

	// Keep track of the number of tickets booked by each agent.
	var numTicketsBooked int

	// Loop over the ticket requests until the request channel is closed.
	for request := range requestCh {
		// Simulate some delay in processing the ticket request.
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		if numTicketsBooked < 10 {
			// Book the ticket.
			fmt.Printf("Booking agent %d processed ticket request for passenger %s on train %s departing at %s to %s.\n",
				id, request.passengerName, request.ticket.trainName, request.ticket.departureStr, request.ticket.destination)
			numTicketsBooked++
		} else {
			// Put the request in the waiting list.
			waitingList = append(waitingList, request)
		}

		if len(waitingList) > 0 {
			fmt.Printf("Booking agent %d has %d requests in the waiting list.\n", id, len(waitingList))
			for _, request := range waitingList {
				// Book the ticket.
				fmt.Printf("Booking agent %d processed waiting list ticket request for passenger %s on train %s departing at %s to %s.\n",
					id, request.passengerName, request.ticket.trainName, request.ticket.departureStr, request.ticket.destination)
			}
		}
	}

	// Notify the wait group that this booking agent has finished.
	wg.Done()
}
