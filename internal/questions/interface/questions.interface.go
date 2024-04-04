package inter

type Question struct {
	Number string `json:"number"`
	Desc   string `json:"desc"`
}

type Category struct {
	Age      string     `json:"age"`
	Question []Question `json:"question"`
}

type Resp map[string][]Category
