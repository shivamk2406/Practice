package llds

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Core Data Structures
type Movie struct {
	ID      string
	Name    string
	Details string
}

type Show struct {
	ID        string
	Movie     *Movie
	ScreenID  string
	StartTime time.Time
}

type Seat struct {
	ID         string
	SeatType   string
	IsOccupied bool
}

type Screen struct {
	ID    string
	Show  *Show
	Seats map[string]*Seat // Map for quick lookup
	mu    sync.Mutex       // Mutex for concurrency control
}

type Theatre struct {
	ID      string
	Screens map[string]*Screen // Map for quick lookup
}

type Ticket struct {
	ID    string
	Seats []*Seat
	Movie *Movie
}

// Managers
type MovieManager struct {
	Movies map[string]*Movie // Map for quick lookup
}

type TheatreManager struct {
	Theatres map[string]*Theatre // Map for quick lookup
}

// Services
type MovieService struct {
	MovieManager   *MovieManager
	TheatreManager *TheatreManager
}

// DTO for Booking
type BookTicketRequestDto struct {
	TheatreID string
	ScreenID  string
	ShowID    string
	SeatIDs   []string
}

// API Interface
type API interface {
	GetAvailableTheatres(movieName string) ([]*Theatre, error)
	BookMovieTicket(request BookTicketRequestDto) (*Ticket, error)
}

// Constructor
func NewMovieService(movieManager *MovieManager, theatreManager *TheatreManager) API {
	return &MovieService{
		MovieManager:   movieManager,
		TheatreManager: theatreManager,
	}
}

// API Implementation
func (ms *MovieService) GetAvailableTheatres(movieName string) ([]*Theatre, error) {
	var result []*Theatre

	for _, theatre := range ms.TheatreManager.Theatres {
		for _, screen := range theatre.Screens {
			if screen.Show != nil && screen.Show.Movie.Name == movieName {
				result = append(result, theatre)
				break
			}
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no theatres found for movie: %s", movieName)
	}

	return result, nil
}

func (ms *MovieService) BookMovieTicket(request BookTicketRequestDto) (*Ticket, error) {
	theatre, ok := ms.TheatreManager.Theatres[request.TheatreID]
	if !ok {
		return nil, errors.New("theatre not found")
	}

	screen, ok := theatre.Screens[request.ScreenID]
	if !ok {
		return nil, errors.New("screen not found")
	}

	screen.mu.Lock()
	defer screen.mu.Unlock()

	// Validate and Book Seats
	ticket := &Ticket{ID: "TICKET-123", Movie: screen.Show.Movie}
	for _, seatID := range request.SeatIDs {
		seat, ok := screen.Seats[seatID]
		if !ok {
			return nil, fmt.Errorf("seat %s not found", seatID)
		}
		if seat.IsOccupied {
			return nil, fmt.Errorf("seat %s is already occupied", seatID)
		}
		seat.IsOccupied = true
		ticket.Seats = append(ticket.Seats, seat)
	}

	return ticket, nil
}

// Driver Function
func Driver3() {
	// Initialize Movies
	movie1 := &Movie{ID: "1", Name: "Bahubali", Details: "Epic"}
	movie2 := &Movie{ID: "2", Name: "Avengers", Details: "Superhero"}

	// Initialize Movie Manager
	movieManager := &MovieManager{
		Movies: map[string]*Movie{
			movie1.ID: movie1,
			movie2.ID: movie2,
		},
	}

	// Initialize Theatre and Screens
	screen1 := &Screen{
		ID: "1",
		Show: &Show{
			ID:        "S1",
			Movie:     movie1,
			StartTime: time.Now(),
		},
		Seats: map[string]*Seat{
			"1": {ID: "1", SeatType: "Gold", IsOccupied: false},
			"2": {ID: "2", SeatType: "Gold", IsOccupied: false},
		},
	}

	theatre1 := &Theatre{
		ID: "T1",
		Screens: map[string]*Screen{
			screen1.ID: screen1,
		},
	}

	// Initialize Theatre Manager
	theatreManager := &TheatreManager{
		Theatres: map[string]*Theatre{
			theatre1.ID: theatre1,
		},
	}

	// Initialize Service
	service := NewMovieService(movieManager, theatreManager)

	// Use Service
	theatres, err := service.GetAvailableTheatres("Bahubali")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Available Theatres:", theatres)
	}

	ticketRequest := BookTicketRequestDto{
		TheatreID: "T1",
		ScreenID:  "1",
		ShowID:    "S1",
		SeatIDs:   []string{"1", "2"},
	}
	ticket, err := service.BookMovieTicket(ticketRequest)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Ticket Booked:", ticket)
	}
}
