package userservice 




type User struct {
	ID uint `gorm:"primaryKey"`
	Name string 
	Username string
	Password string 

}
