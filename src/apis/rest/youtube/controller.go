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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"err": "internal error", "msg": "failed to get data"})
	}
	log.Println("GetData.Controller.Response.Received: ", len(res))

	// Returns
	return c.JSON(http.StatusOK, map[string]interface{}{"paging": map[string]interface{}{"nextPage": currentPage + 1}, "data": res})
}

// Search ...
func (d *dependencies) Search(c echo.Context) (err error) {
	// Get search text
	searchText := c.Param("searchtext")

	// Checks
	if searchText == "" {
		log.Println("Search.BadRequest", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"err": "bad request", "msg": "empty search key"})
	}
	log.Println("Search.Controller.Request.Received: ", searchText)

	// Search data
	res, err := d.dalServices.SearchData(searchText)
	if err != nil {
		log.Println("SearchData.Dal.Call.Failed", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"err": "internal error", "msg": "failed to get search data"})
	}
	log.Println("Search.Controller.Response.Received: ", len(res))

	// Returns
	return c.JSON(http.StatusOK, map[string]interface{}{"data": res})
}
