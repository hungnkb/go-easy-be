package database

type ProviderEnum string

const (
	PROVIDER_PASSWORD ProviderEnum = "password"
	PROVIDER_GOOGLE   ProviderEnum = "google"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
}

type Credential struct {
	ID       uint         `json:"id"`
	UserID   uint         `json:"userId"`
	Provider ProviderEnum `json:"provider"`
	Key      string       `json:"key"`
	Password string       `json:"password"`
	User     User         `gorm:"foreignKey:UserID" json:"user"`
}

func (conn *Conn) Login() {

}
