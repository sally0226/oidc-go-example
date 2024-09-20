package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sally0226/oidc-go-example/model"
	"github.com/sally0226/oidc-go-example/service"
	"github.com/sally0226/oidc-go-example/types"
	"log"
	"net/http"
)

type handler struct {
	oauthService map[types.Provider]service.IOAuthService
	userService  service.IUserService
}

func (h *handler) OAuthLogin(c *gin.Context) {
	p := c.Param("provider")
	provider, err := types.ValidateProvider(p)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	authURL := h.oauthService[provider].AuthURL()

	// TODO : state 적용
	c.Redirect(http.StatusFound, authURL)
}

func (h *handler) OAuthCallback(c *gin.Context) {
	//p := c.Param("provider")
	//provider, err := types.ValidateProvider(p)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithError(http.StatusNotFound, err)
	//	return
	//}
	provider := types.ProviderGoogle
	code := c.Query("code")

	idToken, err := h.oauthService[provider].ExchangeToken(code)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	oauthUser, err := h.oauthService[provider].ParseUser(idToken)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	user, err := h.userService.GetUserByProvider(provider, oauthUser.ID)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	if user == nil {
		user, err = h.userService.CreateUser(&model.User{
			Email:          oauthUser.Email,
			Name:           oauthUser.Name,
			Picture:        oauthUser.Picture,
			Provider:       provider,
			ProviderUserID: oauthUser.ID,
		})
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "login successful"})
	return
}

func NewHandler(oauthService map[types.Provider]service.IOAuthService, userService service.IUserService) *handler {
	return &handler{oauthService: oauthService, userService: userService}
}
