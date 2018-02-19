package models

// UserUserRelationship Model
type UserUserRelationship struct {
	ID          int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID      int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	LikedUserID int `xorm:"liked_user_id INT NOTNULL INDEX(liked_user_id_idx)"`
}

func NewUserUserRelationship(user_id int, liked_user_id int) {
	return &UserUserRelationship{UserID:user_id, LikedUserID:liked_user_id}
}