package notas

import "time"

type Nota struct {
	ID        string    `gorm:"primarykey" json:"id"`
	Titulo    string    `gorm:"size:200;not null" json:"titulo"`
	Contenido string    `gorm:"type:text" json:"contenido"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
