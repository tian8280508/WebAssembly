package db

import (
	"WebAssembly/db/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	DB_User     = "DB_User"
	DB_Password = "DB_Password"
	DB_Host     = "124.220.1.170"
	DB_Port     = "3307"
	DB_database = "webassembly"
)

var db *gorm.DB

func init() {
	dsn := initDSN()
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func initDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv(DB_User), os.Getenv(DB_Password), DB_Host, DB_Port, DB_database)
}

func CreateNode(parentId int, name, content string) error {
	if err := db.Create(&model.Node{
		ParentId: parentId,
		Name:     name,
		Content:  content,
		Comment:  "",
	}).Error; err != nil {
		log.Println("CreateNode error:%v", err)
		return err
	}
	return db.Error
}

func UpdateNode() error {
	return nil

}

func GetOneNode(id int) (*model.Node, error) {
	res := &model.Node{}
	if err := db.First(&res, id).Error; err != nil {
		log.Printf("GetOneNode error:%v", err)
		return nil, err
	}
	return res, nil

}

// GetNodeBatch 后续要优化，目前走全量
func GetNodeBatch() ([]*model.Node, error) {
	res := make([]*model.Node, 0, 10)
	if err := db.Find(&res).Error; err != nil {
		log.Println("GetNodeBatch error:%v", err)
		return nil, err
	}
	return res, nil
}

func GetCommentByNodeId(nodeId int) ([]*model.Comment, error) {
	//res := make([]*model.Comment, 0, 10)
	//db.Model(&res, "node")
	return nil, nil
}

func AddComment(nodeId int, content, avatarUrl, name string) error {
	db.Create(&model.Comment{
		NodeId:    nodeId,
		Content:   content,
		AvatarUrl: avatarUrl,
		Name:      name,
	})
	return db.Error
}
