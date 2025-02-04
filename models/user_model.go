package models

type User struct {
	ID string `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserBase
}

type UserBase struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

// TableName is Database TableName of this model
func (e *User) TableName() string {
	return "tbUser"
}
