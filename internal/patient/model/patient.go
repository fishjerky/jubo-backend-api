package model

type Patient struct {
	ID      string `json:"Id"`
	Name    string `json:"Name"`
	OrderId string `json:"OrderId"`
}

// 初始化內存數據庫
var patients = []Patient{
	{ID: "1", Name: "Patient1", OrderId: "1"},
	{ID: "2", Name: "Patient2", OrderId: "2"},
}

func FindAll() []Patient {
	return patients
}

func Insert(newPatient Patient) {
	patients = append(patients, newPatient)

}
