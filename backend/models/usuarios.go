package models

type Usuario struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Email string `gorm:"unique;not null" json:"email"`
	Rol   string `gorm:"not null" json:"rol"`
}

// Tabla personalizada en plural
func (Usuario) TableName() string {
	return "usuarios"
}
