package service 



type Book struct {
	ID uint `gorm:"primaryKey"`
	Title string 
	Author string 
	UserID int 
	
}

