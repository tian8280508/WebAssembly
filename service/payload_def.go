package service

type AddNodePayload struct {
	SrcId   int    `json:"src_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type UpdateNodePayload struct {
	Id      int    `json:"id"`
	SrcId   int    `json:"src_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type GetNodeVersion struct {
	NodeId    int `json:"node_id"`
	VersionId int `json:"version_id"`
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
