package app

import (
	"net/http"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "Pixel tracker API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GenerateLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		uniqueCode := uniuri.New()

		err := s.trackerService.New(uniqueCode, c)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"unique_code": uniqueCode,
			"image_url":   "http://localhost:8000/image/" + uniqueCode + ".png",
		})
	}
}

func (s *Server) HandleImageLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		uniqueCode := c.Param("uniqueCode")
		uniqueCode = strings.ReplaceAll(uniqueCode, ".png", "")

		if uniqueCode == "" {
			c.JSON(http.StatusBadRequest, &gin.H{
				"message": "invalid uri",
			})
		}

		err := s.trackerService.HandleImageLink(uniqueCode, c)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.Writer.Header().Add("Content-Type", "image/png")
		c.File("cbimage.png")
	}
}

func (s *Server) FetchUriInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		uniqueCode := c.Param("uniqueCode")
		uniqueCode = strings.ReplaceAll(uniqueCode, ".png", "")

		if uniqueCode == "" {
			c.JSON(http.StatusBadRequest, &gin.H{
				"message": "invalid uri",
			})
		}

		uriInfo, err := s.trackerService.FetchUriInfo(uniqueCode, c)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, uriInfo)

	}
}
