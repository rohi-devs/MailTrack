package Model

import "github.com/google/uuid"

type TrackEvent struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CampId   string    `gorm:"type:varchar(255);not null;index:idx_camp_id"`
	UserMail string    `gorm:"type:varchar(255);not null;index:idx_user_mail"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;index:idx_user_id"`
	User     User      `gorm:"foreignKey:UserID"`
	IpAddr   string    `gorm:"type:varchar(255);not null;index:idx_ip_addr"`
}
