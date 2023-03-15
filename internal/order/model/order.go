package model

type Order struct {
	ID      string `json:"Id"`
	Message string `json:"Message"`
}

// 初始化內存數據庫
var orders = []Order{
	{ID: "1", Message: "超過120請施打8u"},
	{ID: "2", Message: "超過240請施打16u"},
}

func FindAll() []Order {
	return orders
}
