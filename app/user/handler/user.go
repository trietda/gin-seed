package handler

import (
	"fmt"
	"gin-seed/app/user/model"
	"gin-seed/app/user/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CredentialDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDto struct {
	CredentialDto
}

func Register(c *gin.Context) {
	var body RegisterDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guest := model.NewGuest(c.ClientIP())
	credential, _ := model.NewCredential(body.Username, body.Password)
	user := guest.Register(credential)

	if err := repository.SaveUser(*user); err != nil {
		if err.Code == repository.UserExistedError {
			c.JSON(http.StatusConflict, gin.H{})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func Login(c *gin.Context) {
	var body CredentialDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := repository.GetByUsername(body.Username)

	if user == nil || !user.Credential.IsValidPassword(body.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	user.Guest = model.Guest{
		Ip: c.ClientIP(),
	}
	session := user.Login()
	repository.SaveSession(*session)

	if jwt, err := session.GenerateAccessToken(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"accessToken": jwt})
		return
	}
}

func Test(c *gin.Context) {
	claim := c.MustGet("userClaim")
	c.JSON(http.StatusOK, claim)
}

type RefreshTokenDto struct {
	RefreshToken string `json:"refreshToken"`
}

func RefreshToken(c *gin.Context) {
	var body RefreshTokenDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	session := repository.GetSession(body.RefreshToken)

	if session == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	accessToken, err := session.GenerateAccessToken()

	if err != nil {
		log.Panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"accessToken": accessToken})
}
