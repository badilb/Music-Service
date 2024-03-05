package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`
	Name     string `gorm:"size:35" json:"name"`
	Surname  string `gorm:"size:35" json:"surname"`
	Username string `gorm:"type:varchar(35);unique_index" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Roles    []Role `gorm:"many2many:user_roles"`
}

type Role struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:35"`
	Users []User `gorm:"many2many:user_roles"`
}

type Tracks struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Title     string     `gorm:"type:varchar(255)" json:"title"`
	Artist    string     `gorm:"type:varchar(255)" json:"artist"`
	AudioData []byte     `gorm:"type:bytea;not null" json:"audioData"`
	Playlist  []Playlist `gorm:"many2many:playlist_music"`
}

type Playlist struct {
	gorm.Model
	ID     uint     `gorm:"primaryKey;autoIncrement"`
	Title  string   `gorm:"type:varchar(255)" json:"title"`
	Tracks []Tracks `gorm:"many2many:playlist_music"`
}
