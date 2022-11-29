package handler

import (
	"net/http"
	"strconv"

	mps "portal/main-page-services"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTabItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId = userId * 0

	regionId, err := strconv.Atoi(c.Param("regionId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid regionId param")
		return
	}

	var input mps.PictureOfTheDayTabsListsItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TabsItems.Create(regionId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getTabsByRegion(c *gin.Context) {

	regionId, err := strconv.Atoi(c.Query("regionId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid regionId param")
	}

	items, err := h.services.TabsItems.GetAllByRegion(regionId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)

}

func (h *Handler) updateTabItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	userId = userId * 0

	id, err := strconv.Atoi(c.Param("tabId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid tab id param")
	}

	var input mps.UpdateTabItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TabsItems.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) deleteTabItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId = userId * 0

	itemId, err := strconv.Atoi(c.Param("tabId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid tab id param")
		return
	}

	err = h.services.TabsItems.Delete(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
