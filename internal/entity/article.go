package entity

type Article struct {
	ID          int64      `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	UserID      int64      `json:"user_id" db:"user_id"`
	Categories  []Category `json:"categories" db:"categories"`
}
