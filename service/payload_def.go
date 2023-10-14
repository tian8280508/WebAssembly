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

type GetNodeByIdPayload struct {
	Id int `json:"id"`
}

type GetAllNodePayload struct {
}

type AddEdgePayload struct {
	SrcId int    `json:"src_id"`
	TarId int    `json:"tar_id"`
	Desc  string `json:"desc"`
}

type GetNodeVersionListPayload struct {
	NodeId int `json:"node_id"`
}
