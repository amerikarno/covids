package app

// Data struct
type Data struct{
	Data []JSONData `json:"Data"`
}

//JSONData struct create for url
type JSONData struct{
	ConfirmDate string `json:"ConfirmDate"`
	No string `json:"No"`
	Age int `json:"Age"`
	Gender string `json:"Gender"`
	GenderEn string `json:"GenderEn"`
	Nation string `json:"Nation"`
	NationEn string `json:"NationEn"`
	Province string `json:"Province"`
	ProvinceID int `json:"ProvinceId"`
	District string `json:"District"`
	ProvinceEn string `json:"ProvinceEn"`
	StatQuarantine int `json:"StatQuarantine"`
}

//Province struct
type Province struct{ 
	ProvinceName string `json:"ProvincName"`
	Total int `json:"Total"`
}

//Ages struct
type Ages struct{
	Below30 int `json:"0-30"`
	Above30lt60 int `json:"30-60"`
	Above60 int `json:"60+"`
	NotAvailable int `json:"N/A"`
}