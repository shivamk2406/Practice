package llds

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// User flow
// 1. Users city is selected.
// 2. Stores are shown
// 3. available vehicles are shown
// 4. users reserves the vehicle

type User struct {
	userId    string
	userName  string
	licenseNo string
}

type vehicle struct {
	vehicleNo   string
	vehicleType string
	kmsDriven   float64
	isAvaialble bool
}

type Stores struct {
	id         string
	address    string
	name       string
	vehicleMap map[string]map[string]*vehicle
}

type Reservation struct {
	userInfo  *User
	vehicle   *vehicle
	ticketId  string
	startDate time.Time
	mutex     sync.Mutex
}

type StoreManager struct {
	storeCityMap map[string][]*Stores
	storeIdMap   map[string]*Stores
	reservations []*Reservation
}

type ReserveVehicleRequestDto struct {
	storeId   string
	userInfo  User
	vehicle   vehicle
	startDate time.Time
}

type AvailableVehicleResponseDto struct {
	vehicleStoreMap map[string][]*vehicle
}

type AvailableVehicleRequestDto struct {
	city            string
	vehicleCategory string
}

func (s *StoreManager) ReserveVehicle(reservationRequest ReserveVehicleRequestDto) (*Reservation, error) {
	store, ok := s.storeIdMap[reservationRequest.storeId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no such store exists with id %s", reservationRequest.storeId))
	}

	vehicles, ok := store.vehicleMap[reservationRequest.vehicle.vehicleType]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no vehicle available for this type", reservationRequest.vehicle.vehicleType))
	}

	reservation := &Reservation{
		userInfo:  &reservationRequest.userInfo,
		ticketId:  uuid.NewString(),
		startDate: time.Now(),
	}

	reservation.mutex.Lock()
	defer reservation.mutex.Unlock()

	if vehicles[reservationRequest.vehicle.vehicleNo].isAvaialble {
		vehicles[reservationRequest.vehicle.vehicleNo].isAvaialble = false
		reservation.vehicle = vehicles[reservationRequest.vehicle.vehicleNo]
		s.reservations = append(s.reservations, reservation)
		return reservation, nil
	} else {
		return nil, errors.New("requested vehicle is not available for reservation")
	}

}

func (s *StoreManager) GetAllAvailableVehicles(availableVehicleRequest AvailableVehicleRequestDto) (*AvailableVehicleResponseDto, error) {
	stores, ok := s.storeCityMap[availableVehicleRequest.city]
	if !ok {
		return nil, errors.New("no stores are available in requested city")
	}

	response := &AvailableVehicleResponseDto{
		vehicleStoreMap: make(map[string][]*vehicle),
	}

	for _, store := range stores {
		vehicles, ok := store.vehicleMap[availableVehicleRequest.vehicleCategory]
		if ok {
			for _, vehicle := range vehicles {
				if vehicle.isAvaialble {
					response.vehicleStoreMap[store.name] = append(response.vehicleStoreMap[store.name], vehicle)
				}
			}
		}
	}
	return response, nil
}


func Driver5(){
	storeManager := &StoreManager{
		storeCityMap: map[string][]*Stores{},
		storeIdMap: map[string]*Stores{},
	}

	// Create Vehicles
	vehicles := map[string]*vehicle{
		"V101": {vehicleNo: "V101", vehicleType: "Car", kmsDriven: 15000, isAvaialble: true},
		"V102": {vehicleNo: "V102", vehicleType: "Car", kmsDriven: 20000, isAvaialble: true},
		"V201": {vehicleNo: "V201", vehicleType: "Bike", kmsDriven: 5000, isAvaialble: true},
	}

	store := &Stores{
		id:         "S101",
		address:    "123 Main Street",
		name:       "Downtown Rentals",
		vehicleMap: map[string]map[string]*vehicle{"Car": {"V101": vehicles["V101"], "V102": vehicles["V102"]}, "Bike": {"V201": vehicles["V201"]}},
	}

	storeManager.storeCityMap["New York"] = []*Stores{store}
	storeManager.storeIdMap[store.id] = store

	user := User{
		userId:    "U001",
		userName:  "John Doe",
		licenseNo: "LIC1234",
	}

	// Reserve a Vehicle
	reserveRequest := ReserveVehicleRequestDto{
		storeId:     "S101",
		userInfo:    user,
		vehicle: vehicle{
			vehicleNo: "V101",
			vehicleType: "Car",
		},
		startDate:   time.Now(),
	}

	reservation, err := storeManager.ReserveVehicle(reserveRequest)
	if err != nil {
		fmt.Println("Error while reserving vehicle:", err)
	} else {
		fmt.Printf("Reservation successful! Ticket ID: %s\n", reservation.ticketId)
		fmt.Printf("Vehicle Reserved: %s (%s)\n", reservation.vehicle.vehicleNo, reservation.vehicle.vehicleType)
	}

	// Check Available Vehicles in New York
	availableRequest := AvailableVehicleRequestDto{
		city:            "New York",
		vehicleCategory: "Car",
	}

	availableResponse, err := storeManager.GetAllAvailableVehicles(availableRequest)
	if err != nil {
		fmt.Println("Error fetching available vehicles:", err)
	} else {
		fmt.Println("\nAvailable Vehicles:")
		for storeName, vehicles := range availableResponse.vehicleStoreMap {
			fmt.Printf("Store: %s\n", storeName)
			for _, vehicle := range vehicles {
				fmt.Printf("- %s (%s), Kms Driven: %.2f\n", vehicle.vehicleNo, vehicle.vehicleType, vehicle.kmsDriven)
			}
		}
	}

}
