package controllers

import (
	"base-api/app/formatters"
	"base-api/app/helper"
	"base-api/app/inputs"
	"base-api/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type participantController struct {
	participantService services.IParticipantService
}

func ParticipantController(participantService services.IParticipantService) *participantController {
	return &participantController{participantService}
}

// GetAllParticipant godoc
// @Summary Get all Participant.
// @Description Get a list of Participant.
// @Tags Participant
// @Produce json
// @Success 200 {object} []formatters.ParticipantFormatter
// @Router /api/v1/participant [get]
func (h *participantController) GetParticipants(c *gin.Context) {
	result, err := h.participantService.GetAll()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Error to get participants", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of participants", http.StatusOK, "success", formatters.FormatParticipants(result))
	c.JSON(http.StatusOK, response)
}

// GetParticipantById godoc
// @Summary Get Participant.
// @Description Get a Participant by id.
// @Tags Participant
// @Produce json
// @Param id path string true "participant id"
// @Success 200 {object} formatters.ParticipantFormatter
// @Router /api/v1/participant/{id} [get]
func (h *participantController) GetParticipant(c *gin.Context) {
	var input inputs.GetIDParticipantInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("please add ID", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := h.participantService.GetByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get detail of participant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if result.ID == uuid.Nil {
		response := helper.APIResponse("Participant data", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("Participant data", http.StatusOK, "success", formatters.FormatParticipant(result))
	c.JSON(http.StatusOK, response)
}

// CreateParticipant godoc
// @Summary Create New Participant.
// @Description Creating a new Participant.
// @Tags Participant
// @Param Body body inputs.CreateParticipantInput true "the body to create a new participant"
// @Produce json
// @Success 200 {object} helper.Response
// @Router /api/v1/participant [post]
func (h *participantController) CreateParticipant(c *gin.Context) {
	var input inputs.CreateParticipantInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create participant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.participantService.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to create participant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create participant", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

// UpdateParticipant godoc
// @Summary Update Participant.
// @Description Update participant by id.
// @Tags Participant
// @Produce json
// @Param id path string true "participant id"
// @Param Body body inputs.UpdateParticipantInput true "the body to update an participant"
// @Success 200 {object} helper.Response
// @Router /api/v1/participant/{id} [put]
func (h *participantController) UpdateParticipant(c *gin.Context) {
	var inputID inputs.GetIDParticipantInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("please add ID", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.UpdateParticipantInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update participant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.participantService.Update(inputID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to update participant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update participant", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

// DeleteParticipant godoc
// @Summary Delete one Participant.
// @Description Delete a participant by id.
// @Tags Participant
// @Produce json
// @Param id path string true "participant id"
// @Success 200 {object} helper.Response
// @Router /api/v1/participant/{id} [delete]
func (h *participantController) DeleteParticipant(c *gin.Context) {
	var input inputs.GetIDParticipantInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("please add ID", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.participantService.Delete(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Participant deleted", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
