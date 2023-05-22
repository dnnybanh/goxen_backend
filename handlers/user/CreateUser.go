package user

import (
	"net/http"

	"github.com/dnnybanh/goxen_backend/models"
	"github.com/dnnybanh/goxen_backend/utils/services"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad Request Error": err.Error()})
		return
	}	

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Internal Error": err.Error()})
		return
	}

	mailchimp := services.Mailchimp{APIKey: "0fc79de674d516f36f108025a27f6373-us21", ListID: "10401", BaseURL: "https://us21.admin.mailchimp.com"}

	if err := mailchimp.SubscribeUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Internal Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}