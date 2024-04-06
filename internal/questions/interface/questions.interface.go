package inter

type Resp map[string][]Question

type Question struct {
	Number string `json:"number"`
	Desc   string `json:"desc"`
}

type BaseQuestions struct {
	FormsName  string `gorm:"column:Forms_name;not null" json:"forms_name"`
	Categories string `gorm:"column:Categories;not null" json:"categories"`
	Question   string `gorm:"column:question;not null" json:"question"`
	Age        int    `gorm:"column:Age;not null" json:"age"`
}
