package response

type Menus struct {
	Id              string `json:"id"`
	MenuTypeName    string `json:"menuTypeName"`
	MenuName        string `json:"menuName"`
	MenuDescription string `json:"menuDescription"`
	MenuPrice       int    `json:"menuPrice"`
}
