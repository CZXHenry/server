package models

// PhotoComment Model
type PhotoComment struct {
	ID      int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Title   string `xorm:"title VARCHAR(45) NOTNULL"`
	Content string `xorm:"content VARCHAR(200) NOTNULL"`
	UserID  int    `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	PhotoID int    `xorm:"photo_id INT NOTNULL INDEX(photo_id_idx)"`
}
