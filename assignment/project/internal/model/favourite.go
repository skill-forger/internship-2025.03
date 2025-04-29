package model

// FollowUser represents follow_user table from the database
type FollowUser struct {
	UserID       int `gorm:"primaryKey;column:user_id"`
	FollowUserID int `gorm:"primaryKey;column:follow_user_id"`
}

// TableName specifies the table name for FollowUser
func (FollowUser) TableName() string {
	return "follow_user"
}