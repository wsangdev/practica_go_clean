package infra

type UserModel struct {
	ID     int64   `gorm:"column:id;primaryKey;autoIncrement"`
	Nombre string  `gorm:"column:nombre;type:varchar(100);not null"`
	Email  string  `gorm:"column:email;type:varchar(100);not null;unique"`
	Edad   float64 `gorm:"column:edad"`
}

func (UserModel) TableName() string {
	return "users"
}
