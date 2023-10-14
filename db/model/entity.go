package model

import (
	"time"
)

type Node struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	NodeId     int       `gorm:"column:node_id;default:NULL"`
	VersionId  int       `gorm:"column:version_id;default:NULL"`
	Name       string    `gorm:"column:name;default:NULL"`
	Content    string    `gorm:"column:content;default:NULL"`
	Comment    string    `gorm:"column:comment;default:NULL"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"column:update_time;default:NULL"`
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

type Edge struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	SrcId       int       `gorm:"column:src_id;NOT NULL"`
	TarId       int       `gorm:"column:tar_id;NOT NULL"`
	Description string    `gorm:"column:description;default:NULL"`
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

func (e *Edge) TableName() string {
	return "edge"
}
