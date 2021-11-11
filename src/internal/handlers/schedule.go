package handlers

import (
	"net/http"
	"strconv"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/ports"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	horarioService ports.HorarioService
}

func NewHTTPHandler(horarioService ports.HorarioService) *HTTPHandler {
	return &HTTPHandler{
		horarioService: horarioService,
	}
}

//GetAvailableHours is the handler for getting available hours endpoint
//@Sumary Get available hours
//@Description List all the hours remaining for creaiting an entrie on the schedule
//@Descriptionby type of hour (lessons, lab or problems)
//@Tag Scheduler
//@Produce json
//@Param titulacion query string true "titulacion de las horas a obtener"
//@Param curso query int true "curso de las horas a obtener"
//@Param grupo query int true "grupo de las horas a obtener"
//@Success 200 {array} domain.AvailableHours
// @Failure 400,404 {object} ErrorHttp
//@Router /availableHours/ [get]
func (hdl *HTTPHandler) GetAvailableHours(c *gin.Context) {

	titulacion := c.Query("titulacion")
	curso, _ := strconv.Atoi(c.Query("year"))
	grupo, _ := strconv.Atoi(c.Query("group"))
	terna := domain.Terna{
		Curso:      curso,
		Titulacion: titulacion,
		Grupo:      grupo,
	}
	availableHours, err := hdl.horarioService.GetAvailableHours(terna)
	if err == apperrors.ErrInvalidInput {

		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorHttp{Message: "Parámetros incorrectos"})

	} else if err == apperrors.ErrNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorHttp{Message: "La terna no existe"})

	} else if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	} else {
		c.JSON(http.StatusOK, availableHours)
	}

}

//PostNewEntry is the handler for creating a new schedluer entry
//@Sumary Post new scheduler entry
//@Description Requesting a new entry for the scheduler. The entry will be definied by the initial hour
//@Description and the ending hour, adintional info must be indicated depending of the kind of hours
//@Description the kinds of subject hours are:
//@Description  - Theorical = 1
//@Description  - Practices = 2
//@Description  - Exercises = 3
//@Tag Scheduler
//@Param entry body  EntryDTO true "Entry to create"
//@Produce text/plain
//@Success 200 "Receive the date of the latests entry modification with format dd/mm/aaaa"
//@Router /newEntry/ [post]
func (hdl *HTTPHandler) PostNewEntry(c *gin.Context) {
	//Read the body request
	body := EntryDTO{}
	c.BindJSON(&body)
	//Execute service
	lastMod, err := hdl.horarioService.CreateNewEntry(body.ToEntry())
	if err == nil {
		c.String(http.StatusOK, lastMod)

	}

}

//ListDegrees is the handler for getting the list of all degrees' descriptions avaiable
//@Sumary Get degrees description
//@Description List all degrees' descriptions avaiable, it do not require any parameter
//@Tag Scheduler
//@Produce json
//@Success 200 {array} handlers.ListDegreesDTO
// @Failure 500 {object} ErrorHttp
//@Router /listDegrees/ [get]
func (hdl *HTTPHandler) ListDegrees(c *gin.Context) {
	list, err := hdl.horarioService.ListAllDegrees()
	if err == nil {

		c.JSON(http.StatusOK, NewListDegrees(list))
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorHttp{Message: "unkown"})
	}
}
