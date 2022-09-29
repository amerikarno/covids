package app_test

import (
	"context"
	"testing"

	"github.com/amerikarno/covids/app"
	"github.com/amerikarno/covids/app/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type AppUsecaseTestSuite struct {
	suite.Suite
	ctx     context.Context
	ctrl    *gomock.Controller
	repo    *mock.MockIRepository
	usecase *app.Usecase
}

func (s *AppUsecaseTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.ctrl = gomock.NewController(s.T())

	s.repo = mock.NewMockIRepository(s.ctrl)
	s.usecase = app.NewUsecase(s.repo)
}

func (s *AppUsecaseTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestRun(t *testing.T) {
	suite.Run(t, new(AppUsecaseTestSuite))
}

func (s *AppUsecaseTestSuite) Test_case_success() {
	expect := ""

	s.repo.EXPECT().DecodeJSONUrl().Return(map[string]interface{}{
		"Provice": map[string]int{
			"Bangkok": 1,
		},
	}, nil).Times(1)

	actual, err := s.usecase.GetProvince()
	s.NoError(err)
	s.Equal(expect, actual)
}
