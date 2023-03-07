package youtube

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetData ...
func (d *dependencies) GetData(c echo.Context) (err error) {
	// Gets page numer from request
	currentPage, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || currentPage == 0 {
		log.Println("GetData.BadRequest", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"err": "bad request", "msg": "invalid page number"})
	}
	log.Println("GetData.Controller.Request.Received: ", currentPage)

	// Gets data
	res, err := d.dalServices.GetAllData(int64(currentPage))
	if err != nil {
		log.Println("GetData.Dal.Call.Failed", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"err": "internal error", "msg": "failed to get data"})
	}
	log.Println("GetData.Controller.Response.Received: ", len(res))

	// Returns
	return c.JSON(http.StatusOK, map[string]interface{}{"paging": map[string]interface{}{"nextPage": currentPage + 1}, "data": res})
}
