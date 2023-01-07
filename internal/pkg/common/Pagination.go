package common

type Pagination struct {
	Size         int         `json:"size" form:"size"`
	Page         int         `json:"page" form:"page"`
	Sort         string      `json:"sort" form:"sort"`
	TotalContent int64       `json:"totalContent" form:"totalContent"`
	Content      interface{} `json:"content" form:"content"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSize() int {
	if p.Size == 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
