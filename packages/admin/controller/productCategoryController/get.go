package productCategoryController

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"strconv"
)

func GetList(c *gin.Context) {
	api := adminApp.ApiInit(c)
	var err error
	idString := c.DefaultQuery("id", "")
	id, err := strconv.Atoi(idString)

	if funk.IsEmpty(id) {
		err = errors.New(e.GetMsg(e.NotFoundId))
	}

	// 查询逻辑 todo 暂时这样可以保障不报错
	fmt.Println(err)
	fmt.Println(api)
}
