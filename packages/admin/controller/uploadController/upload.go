package uploadController

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"hd-mall-ed/packages/common/config"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"net/http"
	"time"
)

func Upload(c *gin.Context) {
	api := adminApp.ApiInit(c)

	fileHeader, _ := c.FormFile("file")
	fileRealName := fileHeader.Filename
	file, _ := fileHeader.Open()

	dist := make([]byte, fileHeader.Size) //开辟存储空间
	n, _ := file.Read(dist)
	sourceString := base64.StdEncoding.EncodeToString(dist[:n])

	fileName := fmt.Sprintf("%d-%s", time.Now().UnixNano()/1e6, fileHeader.Filename)
	urlPath := fmt.Sprintf(
		"https://gitee.com/api/v5/repos/%s/%s/contents/%s/",
		config.AppConfig.Owner,
		config.AppConfig.Repository,
		config.AppConfig.Path,
	)

	// 请求
	requestUrl := fmt.Sprintf(
		"%s%s",
		urlPath,
		fileName,
	)
	//	https://gitee.com/api/v5/repos/yanleweb/static/contents/hd_client/demo-2.jpg

	res, err := grequests.Post(requestUrl, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Length": "multipart/form-data; boundary=<calculated when request is sent>",
		},
		Data: map[string]string{
			"access_token": config.AppConfig.AccessToken,
			"message":      fileName,
			"content":      sourceString,
		},
	})

	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"download_url": gjson.Get(res.String(), "content.download_url").String(),
		"file_name": fileRealName,
	})
}
