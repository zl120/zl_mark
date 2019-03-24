package vo

type (
	UserAccount struct {
		Name string `json:"name"`
		Pwd string `json:"pwd"`
		Role int `json:"role"`
	}

	UserInfo struct {
		Name        string
		Age         int
		IdCard      string
		Gender      string
		Status      int
		Tel         string
	}
)


