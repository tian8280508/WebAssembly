package db

import (
	"WebAssembly/db/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	DB_User     = "DB_User"
	DB_Password = "DB_Password"
	DB_InRemote = "124.220.1.170"
	DB_InLocal  = "0.0.0.0"
	DB_Port     = "3307"
	DB_database = "webassembly"
	Hostname    = "Hostname"
)

var db *gorm.DB

func init() {
	dsn := initDSN()
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	rand.Seed(time.Now().UnixNano())
}

func initDSN() string {
	var DB_Host string
	if os.Getenv(Hostname) == "local" {
		DB_Host = DB_InRemote
	} else {
		DB_Host = DB_InLocal
	}
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv(DB_User), os.Getenv(DB_Password), DB_Host, DB_Port, DB_database)
}

func GetNodeVersionList(nodeId int) ([]*model.Node, error) {
	res := make([]*model.Node, 0, 10)
	if err := db.Where("node_id = ?", nodeId).Find(&res).Error; err != nil {
		log.Printf("GetNodeVersionList error:%v", err)
		return nil, err
	}
	return res, nil
}

func CreateNode(name, content string) error {
	if err := db.Create(&model.Node{
		NodeId:    int(generateRandomInt32()),
		VersionId: int(generateRandomInt32()),
		Name:      name,
		Content:   content,
		Comment:   "",
	}).Error; err != nil {
		log.Printf("CreateNode error:%v", err)
		return err
	}
	return nil
}

func CreateEdge(srcId, tarId int, desc string) error {
	if err := db.Create(&model.Edge{
		SrcId:       srcId,
		TarId:       tarId,
		Description: desc,
	}).Error; err != nil {
		log.Printf("CreateEdge error:%v", err)
		return err
	}
	return nil

}

func UpdateNode(id int, name, content string) error {
	res := &model.Node{Id: id}
	if err := db.Model(res).Updates(&model.Node{Name: name, Content: content}).Error; err != nil {
		log.Printf("UpdateNode error:%v", err)
		return err
	}

	return nil
}

func GetNodeOneVersion(nodeId, versionId int) (*model.Node, error) {
	res := &model.Node{}
	if err := db.Where("node_id = ? and version_id = ?", nodeId, versionId).First(&res).Error; err != nil {
		log.Printf("GetNodeOneVersion error:%v", err)
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

// GetEdgeBatch 后续要优化，目前走全量
func GetEdgeBatch() ([]*model.Edge, error) {
	res := make([]*model.Edge, 0, 10)
	if err := db.Find(&res).Error; err != nil {
		log.Println("GetEdgeBatch error:%v", err)
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

func generateRandomInt32() int32 {
	return rand.Int31()
}
