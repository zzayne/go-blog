package model

type Article struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (m *Article) QueryList() {
}

func (m *Article) TotalCount() int {
	return getCount("ID>0")
}

func getCount(maps interface{}) (count int) {
	DB.Model(&Article{}).Where(maps).Count(&count)

	return count
}
