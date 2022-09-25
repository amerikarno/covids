package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase
}

type Groups struct{
	AgeGroups *Ages `json:"AgeGroups"`
}

func NewHandler(u Usecase) *Handler {
	return &Handler{u}
}

func (h *Handler) GetCovidSum(gc *gin.Context){
	provinces, _ := h.Usecase.GetProvince()
	ages, _ := h.Usecase.GetAges()
	age := Groups{
		AgeGroups: ages,
	}
	gc.JSON(http.StatusOK,provinces)
	gc.JSON(http.StatusOK,age)
}