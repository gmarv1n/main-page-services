package handler

import (
	"portal/main-page-services/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	v1 := router.Group("/v1")
	{
		// themes := v1.Group("/v1/theme")
		// {
		// 	themes.POST("/", h.createTheme)
		// 	themes.GET("/", h.getThemes)
		// 	themes.GET("/:id", h.getTheme)
		// 	themes.PUT("/:id", h.updateTheme)
		// 	themes.DELETE("/:id", h.deleteTheme)
		// }

		quotes := v1.Group("/quote")
		{
			quotes.POST("/", h.createQuote)
			quotes.GET("/quotes", h.getQuotes)
			quotes.GET("/:regionId", h.getQuote)
			quotes.PUT("/:quoteId", h.updateQuote)
			quotes.DELETE("/:quoteId", h.deleteQuote)
		}

		pictureOfTheDay := v1.Group("/picture-of-the-day")
		{
			tabs := pictureOfTheDay.Group("/tabs")
			{
				tabs.POST("/", h.createTabsList)
				tabs.GET("/tabs", h.getTabsByRegion)
				tabs.DELETE("/:regionId", h.deleteTabsList)

				tabsItems := tabs.Group("/items")
				{
					tabsItems.POST("/regionId", h.createTabItem)
					tabsItems.PUT("/:tabId", h.updateTabItem)
					tabsItems.DELETE("/:id", h.deleteTabItem)
				}
			}
		}
	}

	return router
}
