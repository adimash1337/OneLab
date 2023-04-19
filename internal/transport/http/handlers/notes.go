package handlers

import (
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary CreateNote
// @Security ApiKeyAuth
// @Tags notes
// @Description create note
// @ID create-note
// @Accept  json
// @Produce  json
// @Param input body models.Note true "note info"
// @Success 201 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /note/create [post]
func (h Manager) CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Note
		uid := c.Param("uid")
		bid := c.Param("bid")

		resp, err := h.srv.Note.Create(c.Request.Context(), &req, uid, bid)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

// @Summary DeleteNote
// @Security ApiKeyAuth
// @Tags notes
// @Description delete note
// @ID delete-note
// @Accept  json
// @Produce  json
// @Param id path string true "note id"
// @Success 200 {string} string "message"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /note/delete [get]
func (h Manager) DeleteNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := h.srv.Note.Delete(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, "Note is deleted")
	}
}

// @Summary ListNotes
// @Tags notes
// @Description list notes
// @ID list-notes
// @Accept json
// @Produce json
// @Success 200 {object} []models.Note
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /notes/list_notes [get]
func (h Manager) ListNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h.srv.Note.List(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

// @Summary GetNoteByID
// @Tags notes
// @Description get notes by ID
// @ID get-notes
// @Accept json
// @Produce json
// @Param notes_id path string true "notes id"
// @Success 200 {object} models.Note
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /notes/get_by_id/{notes_id} [get]
func (h Manager) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		resp, err := h.srv.Note.Get(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
