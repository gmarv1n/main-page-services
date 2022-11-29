package handler

import (
	"net/http"
	"strconv"

	mps "portal/main-page-services"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTabsList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userId = userId * 0

	var input mps.PictureOfTheDayTabsList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TabsLists.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllTabsListsResponse struct {
	Data []mps.PictureOfTheDayTabsList `json:"data"`
}

func (h *Handler) getTabsLists(c *gin.Context) {

	lists, err := h.services.TabsLists.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTabsListsResponse{
		Data: lists,
	})
}

func (h *Handler) deleteTabsList(c *gin.Context) {

	regionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	err = h.services.TabsLists.Delete(regionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
