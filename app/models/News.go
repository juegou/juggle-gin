package models

type News struct {
	Id   int `xorm:"pk autoincr" uri:"id" binding:"required,gt=0"`
	NewsTitle string `xorm:"varchar(255) 'news_title' comment('新闻标题')"`
}

func NewNewsModel() *News {
	return &News{}
}

func (this *News) String() string {
	return "news"
}
