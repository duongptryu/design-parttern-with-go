package builder

type BuildProcess interface {
	SetWheels()
	SetSeats()
	SetStructure()
	GetVehicle() VehicleProduct
}
