package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stroiman/go-automapper"
	"gopkg.in/jeevatkm/go-model.v1"
	"net/http"
	"possant-admin/internal/domain/user/entity"
	"possant-admin/internal/domain/user/request"
	"possant-admin/internal/domain/user/response"
	"possant-admin/internal/domain/user/service"
)

type UserHandler struct {
	service service.UserService
}

func UserHandlerProvider(service service.UserService) UserHandler {
	return UserHandler{service: service}
}

func (h *UserHandler) FindAll(c *gin.Context) {
	users := h.service.FindAll()
	userResponse := make([]response.UserList, 0)
	automapper.MapLoose(users, &userResponse)
	c.JSON(http.StatusOK, gin.H{"users": userResponse})
}

func (h *UserHandler) Save(c *gin.Context) {
	var userRequest request.UserCreate
	err := c.BindJSON(&userRequest)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	user := entity.User{}
	errs := model.Copy(&user, userRequest)
	if errs != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	h.service.Save(user)
	c.Status(http.StatusOK)
}
