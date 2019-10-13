package structs

// UserInfo 用户信息，这个结构体用于管理控制台使用
type UserInfo struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Mobile   string `json:"mobile"`
	Mail     string `json:"mail"`
	WX       string `json:"wx"`
}
