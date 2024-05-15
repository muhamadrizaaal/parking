package domain

type Car struct {
    RegistrationNumber string
    Color              string
}

type Slot struct {
    Number int
    Car    *Car
}

type ParkingLot struct {
    Capacity int
    Slots    []*Slot
}
