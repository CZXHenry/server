package models

import "time"

// UserLoginLog Model
type UserLoginLog struct {
	ID            int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime     time.Time `xorm:"login_time DATETIME NOTNULL"`
	LoginLocation string    `xorm:"login_location VARCHAR(50) NOTNULL"`
	LoginDevice   string    `xorm:"login_device VARCHAR(20) NOTNULL"`
	UserID        int       `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
}

// NewUserLoginLog creates a new user login log
func NewUserLoginLog(loginTime time.Time, loginLocation string,
	loginDevice string, userID int) *UserLoginLog {
	return &UserLoginLog{
		LoginTime:     loginTime,
		LoginLocation: loginLocation,
		LoginDevice:   loginDevice,
		UserID:        userID,
	}
}
