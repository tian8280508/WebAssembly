package service

type AddNodePayload struct {
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
}

type UpdateNodePayload struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
}

type GetNodeById struct {
	Id int `json:"id"`
}

type GetAllNode struct {
}
