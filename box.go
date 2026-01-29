package pet3

type BoxList struct {
	Id          int    `json:"id"`
	Title       string `json:"titlr"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type BoxItem struct {
	Id          int    `json:"id"`
	Title       string `json:"titlr"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	UserId int
	ListId int
}
