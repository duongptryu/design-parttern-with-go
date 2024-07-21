package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	factory := ManufacturingDirector{}

	// Car
	carBuilder := CarBuilder{}

	factory.SetBuilder(&carBuilder)
	factory.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels wanted is 4, got %v", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure wanted is Car, got %v", car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats wanted is 5, got %v", car.Seats)
	}

	// Bike
	bikeBuilder := BikeBuilder{}

	factory.SetBuilder(&bikeBuilder)
	factory.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Wheels wanted is 2, got %v", bike.Wheels)
	}
	if bike.Structure != "Bike" {
		t.Errorf("Structure wanted is bike, got %v", bike.Structure)
	}
	if bike.Seats != 2 {
		t.Errorf("Seats wanted is 2, got %v", bike.Seats)
	}

	// Bus
	// Bike
	busBuilder := BusBuilder{}

	factory.SetBuilder(&busBuilder)
	factory.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != 8 {
		t.Errorf("Wheels wanted is 8, got %v", bus.Wheels)
	}
	if bus.Structure != "Bus" {
		t.Errorf("Structure wanted is bus, got %v", bus.Structure)
	}
	if bus.Seats != 30 {
		t.Errorf("Seats wanted is 30, got %v", bus.Seats)
	}
}
