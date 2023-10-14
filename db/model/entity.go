package model

import (
	"time"
)

type Node struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	ParentId   int       `gorm:"column:parent_id;type:int(11)" json:"parent_id"`
	Name       string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Content    string    `gorm:"column:content;type:varchar(255)" json:"content"`
	Comment    string    `gorm:"column:comment;type:varchar(255)" json:"comment"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"update_time"`
}

func (m *Node) TableName() string {
	return "node"
}

type Comment struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	NodeId     int       `gorm:"column:node_id;type:int(11)" json:"node_id"`
	Content    string    `gorm:"column:content;type:varchar(255)" json:"content"`
	AvatarUrl  string    `gorm:"column:avatar_url;type:varchar(255)" json:"avatar_url"`
	Name       string    `gorm:"column:name;type:varchar(255)" json:"name"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP" json:"update_time"`
}

func (m *Comment) TableName() string {
	return "comment"
}
