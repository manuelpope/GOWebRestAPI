package main

//Metric model
type Metric struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Temp     float64 `gorm:"not null" json:"temp" binding:"required"`
	Moisture float64 `json:"moisture" binding:"required"`
}

//List list to delete model
type List struct {
	ListID []int `json:"listID" binding:"required"`
}
