package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// UsersFollowUsers Model
type UsersFollowUsers struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID         int       `xorm:"user_id INT NOTNULL INDEX(user_id_idx)" json:"user_id"`
	FollowedUserID int       `xorm:"followed_user_id INT NOTNULL INDEX(liked_user_id_idx)" json:"followed_user_id"`
	Timestamp      time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// UsersFollowUsersFull Model
type UsersFollowUsersFull struct {
	ID           int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	User         User      `xorm:"extends" json:"user"`
	FollowedUser User      `xorm:"extends" json:"followed_user"`
	Timestamp    time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// UsersFollowUsersDataAccessObject provides database access for
// Model UsersFollowUsers
type UsersFollowUsersDataAccessObject struct{}

// UsersFollowUsersDAO instance of UsersFollowUsersDataAccessObject
var UsersFollowUsersDAO *UsersFollowUsersDataAccessObject

// NewUsersFollowUsers creates a new user-user relationship
func NewUsersFollowUsers(userID int, focusedUserID int) *UsersFollowUsers {
	return &UsersFollowUsers{
		UserID:         userID,
		FollowedUserID: focusedUserID,
		Timestamp:      time.Now(),
	}
}

// TableName returns table name
func (*UsersFollowUsersDataAccessObject) TableName() string {
	return "users_follow_users"
}

// FindFullByUserID finds all full models by a user ID
func (*UsersFollowUsersDataAccessObject) FindFullByUserID(
	id int) []UsersFollowUsersFull {
	l := make([]UsersFollowUsersFull, 0)
	err := orm.Table(UsersFollowUsersDAO.TableName()).
		Join("INNER", UserDAO.TableName(),
			UsersFollowUsersDAO.TableName()+".user_id="+
				UserDAO.TableName()+".id").
		Join("INNER", OrganizationDAO.TableName(),
			UsersFollowUsersDAO.TableName()+".organization_id="+
				OrganizationDAO.TableName()+".id").
		Find(&l)
	logger.LogIfError(err)
	return l
}
