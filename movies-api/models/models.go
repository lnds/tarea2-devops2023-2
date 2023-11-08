package models

type Director struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Movie struct {
	ID          uint     `json:"id,string" gorm:"primaryKey"`
	Title       string   `json:"title"`
	Year        int      `json:"year,string"`
	Description string   `json:"description" gorm:"text"`
	DirectorID  int      `json:"director_id,string"`
	Director    Director `json:"director"`
}
