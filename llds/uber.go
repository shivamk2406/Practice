package llds

import "errors"

// Scope of Work
// User Management:

// Riders and drivers can sign up, log in, and manage their profiles.
// Ride Booking:

// Riders can request a ride by providing pickup and drop-off locations.
// Drivers receive ride requests and can accept or reject them.
// Driver Matching:

// Match the nearest available driver to the rider.
// Basic Ride Flow:

// Update ride status: "Requested," "Accepted," "In Progress," and "Completed."
// Payment (Optional for Scope):

// Assume payment is handled outside the app.

type uberUser struct {
	Id     string
	name   string
	city   string
	lat    string
	long   string
	mobile string
}

type Location struct {
	latitude  string
	longitude string
}

type uberVehicle struct {
	modelName   string
	vehicleNo   string
	vehicleType string
}

type uberDriver struct {
	Id     string
	name   string
	mobile string
}

type uberRide struct {
	vehicle     uberVehicle
	driver      uberDriver
	status      string
	startPoint  Location
	destination Location
}

type UberService struct {
	users   map[string]*uberUser
	drivers map[string]*uberDriver
	rides   map[string][]*uberRide
}

type SignupRequestDto struct {
	isDriver   bool
	userInfo   uberUser
	driverInfo uberDriver
}

func (u *UberService) SignUp(request SignupRequestDto) error {
	if request.isDriver {
		_, ok := u.drivers[request.driverInfo.mobile]
		if ok {
			return errors.New("mobile no already registered")
		}
		u.drivers[request.driverInfo.mobile] = &request.driverInfo
	} else {
		_, ok := u.users[request.userInfo.mobile]
		if ok {
			return errors.New("mobile no already registered")
		}
		u.users[request.userInfo.mobile] = &request.userInfo
	}

	return nil
}

func (u *UberService) CreateRide(request SignupRequestDto) error {
	if request.isDriver {
		_, ok := u.drivers[request.driverInfo.mobile]
		if ok {
			return errors.New("mobile no already registered")
		}
		u.drivers[request.driverInfo.mobile] = &request.driverInfo
	} else {
		_, ok := u.users[request.userInfo.mobile]
		if ok {
			return errors.New("mobile no already registered")
		}
		u.users[request.userInfo.mobile] = &request.userInfo
	}

	return nil
}
