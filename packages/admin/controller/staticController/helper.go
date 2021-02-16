package staticController

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"strconv"
)

func handleQueryParamsHelper(c *gin.Context) *GetListByQueryParams {
	queryParams := &GetListByQueryParams{}

	//params
	var id int
	var productId int
	var typeParam int
	var fileName string
	var url string
	var link string

	id, _ = strconv.Atoi(c.DefaultQuery("id", ""))
	productId, _ = strconv.Atoi(c.DefaultQuery("product_id", ""))
	typeParam, _ = strconv.Atoi(c.DefaultQuery("type", ""))
	fileName = c.DefaultQuery("file_name", "")
	url = c.DefaultQuery("url", "")
	link = c.DefaultQuery("link", "")

	if !funk.IsEmpty(id) {
		queryParams.ID = uint(id)
	}

	if !funk.IsEmpty(productId) {
		queryParams.ProductId = uint(productId)
	}

	if !funk.IsEmpty(typeParam) {
		queryParams.Type = typeParam
	}

	if !funk.IsEmpty(fileName) {
		queryParams.FileName = fileName
	}

	if !funk.IsEmpty(url) {
		queryParams.Url = url
	}

	if !funk.IsEmpty(link) {
		queryParams.Link = link
	}

	return queryParams
}
