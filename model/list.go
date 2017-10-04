package model

//List is hogehoge
type List struct {
	StartMap []BusStop `json:"start_list"`
	EndMap   []BusStop `json:"end_list"`
}

//BusStop is hoge
type BusStop struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
