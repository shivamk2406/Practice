package llds

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type VehicleType int

const(
	TWO_WHEELER VehicleType=iota
	THREE_WHEELER
	FOUR_WHEELER
)

type OwnerDetails struct{
	Name string
	Address string
}

type Vehicle struct{
	VehicleNo string
	VehicleCatgeory VehicleType 
	OwnerInfo OwnerDetails
}

type ParkingSpot struct {
	Available bool
	VehicleDetails Vehicle
	VehicleCategory VehicleType
	Id string
}

type ParkingSpotManager struct {
	parking_spots []*ParkingSpot
}

func(p ParkingSpotManager) ParkVehicle(vehicle Vehicle) (*ParkingSpot,error) {
	vehicleType:=vehicle.VehicleCatgeory

	for _, val := range p.parking_spots{
		if val.VehicleCategory==vehicleType && val.Available{
			val.Available=false
			val.VehicleDetails=vehicle
			val.Id=uuid.NewString()
			fmt.Println("Vehicle Parked sucessfully with id: ", val.Id)
			return val, nil
		}
	}
	return nil,  errors.New("error while parking vehicle no spot available")
} 

func(p ParkingSpotManager) RemoveVehicle(spot ParkingSpot) error {
	for _, val := range p.parking_spots{
		if val.Id==spot.Id{
			val.Available=true
			val.VehicleDetails=Vehicle{}
			fmt.Println("Vehicle un Parked sucessfully with id: ", spot.Id)
			return nil
		}
	}
	return errors.New("error while unparking vehicle")
}

func Driver(){
	parkingSpots:=ParkingSpotManager{
		[]*ParkingSpot{
			{
				Available: true,
				VehicleCategory: TWO_WHEELER,
			},
			{
				Available: true,
				VehicleCategory: TWO_WHEELER,
			},
			{
				Available: true,
				VehicleCategory: THREE_WHEELER,
			},
			{
				Available: true,
				VehicleCategory: TWO_WHEELER,
			},
			{
				Available: true,
				VehicleCategory: TWO_WHEELER,
			},
		},
	}


	vehicle1:=Vehicle{
		VehicleNo: "BR-01A-1763",
		VehicleCatgeory: TWO_WHEELER,
		OwnerInfo: OwnerDetails{
			Name: "SHIVAM",
			Address: "BIHAR",
		},
	}

	vehicle3:=Vehicle{
		VehicleNo: "BR-01A-1745",
		VehicleCatgeory: THREE_WHEELER,
		OwnerInfo: OwnerDetails{
			Name: "SATYAM",
			Address: "BIHAR",
		},
	}

	vehicle2:=Vehicle{
		VehicleNo: "BR-01A-1789",
		VehicleCatgeory: FOUR_WHEELER,
		OwnerInfo: OwnerDetails{
			Name: "SUNDARAM",
			Address: "BIHAR",
		},
	}

	_,err:=parkingSpots.ParkVehicle(vehicle1)
	if err!=nil{
		log.Println(err)
	}

	_,err=parkingSpots.ParkVehicle(vehicle2)
	if err!=nil{
		log.Println(err)
	}

	_,err=parkingSpots.ParkVehicle(vehicle3)
	if err!=nil{
		log.Println(err)
	}

	fmt.Println(parkingSpots.parking_spots)
}

/*
1. Park Vehicle
2. Remove Vehicle
*/