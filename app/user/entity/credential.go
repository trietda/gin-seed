package entity

type Credential struct {
	UserId   string
	Username string
	Password []byte `gorm:"type:string"`
}
