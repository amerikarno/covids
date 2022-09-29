package app

//go:generate mockgen -source=app.go -destination=mock/mock_app.go -package=mock
type IUsecase interface {
	GetProvince() (map[string]interface{}, error)
	GetAges() (*Ages, error)
}

type IRepository interface {
	DecodeJSONUrl() (Data, error)
	GetJSONUrl() ([]byte, error)
	QueryInput() error
	GetInput() Data
}
