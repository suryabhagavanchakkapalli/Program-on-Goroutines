# Train Ticket Booking System

This is a train ticket booking system between Visakhapatnam and Secunderabad, which allows customers to book train tickets through booking agents. The system creates ticketRequest objects with passenger information and train details, which are sent to the booking agent goroutines via a channel. Each booking agent has a limit of booking ten tickets at a time. If the agent receives more than ten requests, the remaining requests are stored in a waiting list.

## Technologies Used

```
Go Programming Language
```

## Getting Started

To run the program, clone the repository and run the following command in the terminal:

```go
go run main.go
```

## Program Flow

1. A wait group is created to track when all booking agents have finished, and a channel is created to hold the ticket requests.
2. Booking agent goroutines are started, and a map is defined to keep track of the number of tickets booked to each train.
3. Fifty ticketRequest objects are created and sent to the booking agent goroutines via the channel.
4. Each booking agent processes the ticket requests by booking tickets or storing them in a waiting list if they have already booked ten tickets.
5. The program waits for all booking agents to finish and then prints the number of tickets booked for each train and city.