package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrandrey24/url-shortening-server/internal/domain"
	"github.com/mrandrey24/url-shortening-server/internal/service"
)

type URLController struct {
	service *service.URLService
}

func NewURLController(s *service.URLService) *URLController {
	return &URLController{
		service: s,
	}
}

type Url struct {
	URL string `json:"url"`
}

func (a *URLController) CreateShortUrl(c *gin.Context) {
	var req Url

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL invalid"})
		return
	}

	url, err := a.service.CreateShortUrl(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, url)
}

func (a *URLController) UpdateShortUrl(c *gin.Context) {
	code := c.Param("code")
	var req Url

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The URL is not valid"})
	}

	urlData := &domain.URL{
		URL: req.URL,
	}

	updatedURL, err := a.service.UpdateShortUrl(code, urlData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := domainURLToResponseUpdate(updatedURL)

	c.JSON(http.StatusOK, response)
}

func (a *URLController) GetShortUrlByCode(c *gin.Context) {
	code := c.Param("code")

	url, err := a.service.RetrieveByShortCode(code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short code not found"})
	}

	response := domainURLToResponseCode(url)

	c.JSON(http.StatusOK, response)
}

func (a *URLController) GetShortUrlStats(c *gin.Context) {
	code := c.Param("code")

	url, err := a.service.RetrieveByShortCode(code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short code not found"})
	}

	response := domainURLToResponseStats(url)

	c.JSON(http.StatusOK, response)
}

func (a *URLController) DeleteShortUrl(c *gin.Context) {
	code := c.Param("code")

	if err := a.service.DeleteSHortUrl(code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL deleted "})
}
