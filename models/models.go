package models

type Tasks struct {
	ID          int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title       string `json:"title" gorm:"type:text;not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	DueDate     string `json:"due_date" gorm:"type:text;not null"`
	CreatedAt   string `json:"created_at" gorm:"type:text;not null"`
	UpdatedAt   string `json:"updated_at" gorm:"type:text;not null"`
}
