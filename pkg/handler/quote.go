package handler

import (
	"net/http"
	"strconv"

	mps "portal/main-page-services"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createQuote(c *gin.Context) {

	var input mps.Quote
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	quoteId, err := h.services.Quote.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"quoteId": quoteId,
	})
}

func (h *Handler) getQuotes(c *gin.Context) {
	items, err := h.services.Quote.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getQuote(c *gin.Context) {
	regionId, err := strconv.Atoi(c.Param("regionId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid regionId param")
		return
	}

	item, err := h.services.Quote.GetByRegionId(regionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

func (h *Handler) updateQuote(c *gin.Context) {

	regionId, err := strconv.Atoi(c.Param("quoteId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	var input mps.UpdateQuoteInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Quote.Update(regionId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) deleteQuote(c *gin.Context) {

	quoteId, err := strconv.Atoi(c.Param("quoteId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quoteId param")
		return
	}

	err = h.services.Quote.Delete(quoteId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
