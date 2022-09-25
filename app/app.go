package app

type IRepository interface {
	DecodeJSONUrl() (Data, error) 
	GetJSONUrl() ([]byte,error)
	QueryInput() error
	GetInput() Data
}