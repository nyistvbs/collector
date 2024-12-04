package model

type TourismDB struct {
	ID               uint   `gorm:"primaryKey"`
	HotelName        string `gorm:"size:100;not null"`
	Star             string `gorm:"size:100;not null"`
	Price            string `gorm:"size:100;not null"`
	PriceBeforeTaxes string `gorm:"size:100;not null"`
	CheckInDate      string `gorm:"size:100;not null"`
	CheckOutDate     string `gorm:"size:100;not null"`
	Guests           string `gorm:"size:100;not null"`
}
