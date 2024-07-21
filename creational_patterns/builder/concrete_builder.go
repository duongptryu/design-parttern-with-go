package builder

// CarBuilder
type CarBuilder struct {
	VehicleProduct
}

func (c *CarBuilder) SetWheels() {
	c.Wheels = 4
}

func (c *CarBuilder) SetSeats() {
	c.Seats = 5
}

func (c *CarBuilder) SetStructure() {
	c.Structure = "Car"
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.VehicleProduct
}

// Bike Builder
type BikeBuilder struct {
	VehicleProduct
}

func (b *BikeBuilder) SetWheels() {
	b.Wheels = 2
}

func (c *BikeBuilder) SetSeats() {
	c.Seats = 2
}

func (c *BikeBuilder) SetStructure() {
	c.Structure = "Bike"
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.VehicleProduct
}

// Bus builder
type BusBuilder struct {
	VehicleProduct
}

func (b *BusBuilder) SetWheels() {
	b.Wheels = 8
}

func (c *BusBuilder) SetSeats() {
	c.Seats = 30
}

func (c *BusBuilder) SetStructure() {
	c.Structure = "Bus"
}

func (c *BusBuilder) GetVehicle() VehicleProduct {
	return c.VehicleProduct
}
