package model

// TodoModel 备忘录模型
type TodoModel struct {
	Id     int32  `gorm:"column:id" json:"id"`
	Title  string `gorm:"column:title" json:"title"`
	Status bool   `gorm:"column:status" json:"status"`
}
