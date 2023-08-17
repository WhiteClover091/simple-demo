package controller

import (
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	// "io"
	"os"

   "github.com/RaymondCode/simple-demo/models"
  // "time"
	"github.com/disintegration/imaging"

	"github.com/gin-gonic/gin"
	"strings"

	"net/http"
  // "net"

	"path/filepath"
)

type VideoListResponse struct {
	Response

	VideoList []Video `json:"video_list"`
}

// ----------------------------TO DO------------------
// 抽取某一帧作为图片
func Vedio2Jpeg(inFileName string, frameNum int) {
	buf := bytes.NewBuffer(nil)

	// 使用 ffmpeg 命令行工具提取视频的指定帧作为 JPEG 图像
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		panic(err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return
	}

	index := strings.LastIndex(inFileName, ".")
	// var outFileName strings
	if index != -1 {
		outFileName := strings.Join([]string{inFileName[:index+1], "jpeg"}, "")
    fmt.Print(outFileName)
		err = imaging.Save(img, outFileName)
		if err != nil {
			return 
		} else {
			return
		}
	}
}

// 获得 域名 
func getDomain(c *gin.Context) (string, error) {
	domain := c.Request.Host
  fmt.Println("My domain:", domain)
	return domain, nil
}

// -------------------------------------------------------------------
// Publish check token then save upload file to public directory

func Publish(c *gin.Context) {

  // mysql -h ${MYSQL_HOST} -P ${MYSQL_PORT} -u ${MYSQL_USER} -p
	token := c.PostForm("token")

	// 鉴权, var usersLoginInfo = map[string]User
  user_name, err := ParseToken(token);
	// if  err != nil {

	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})

	// 	return

	// }

  // var user_id int64;
  user, err := models.NewUserDaoInstance().FindUserByName(user_name); 
   if err != nil {
        panic(err)
    }
  var user_id = user.UserId

  
  

	//
	data, err := c.FormFile("data")

	if err != nil {

		c.JSON(http.StatusOK, Response{

			StatusCode: 1,

			StatusMsg: err.Error(),
		})

		return

	}

	// 本地存储文件
	filename := filepath.Base(data.Filename)

	// user := usersLoginInfo[token]

	finalName := fmt.Sprintf("%d_%s", user_id, filename)

	saveFile := filepath.Join("./public/", finalName)

  // 获得域名
  domain :=  c.Request.Host
  // My domain: 2f781ee3592dd7a9ff0bbd0007fe40ce-app.1024paas.com
  // fmt.Println("My domain:", domain)

  vedio_url := "https://"+ domain + "/static/"+ finalName
  // index := "https://" + strings.LastIndex(finalName, ".")
  
  img_url := domain + "/static/" + strings.TrimSuffix(finalName, filepath.Ext(finalName)) + ".jpeg"
  // fmt.Println("vedio:", vedio_url, "     ", img_url)
  // vedio: https://2f781ee3592dd7a9ff0bbd0007fe40ce-app.1024paas.com/static/0_20230329_133339.mp4       2f781ee3592dd7a9ff0bbd0007fe40ce-app.1024paas.com/static/0_20230329_133339.jpeg
  
  

  

	// func (c *gin.Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
    
    
		c.JSON(http.StatusOK, Response{

			StatusCode: 1,

			StatusMsg: err.Error(),
		})

		return

	}
  Vedio2Jpeg(saveFile, 6)
  

	video1 := models.Video{
		UserId : user_id,
		PlayUrl : vedio_url,
    CoverUrl : img_url , 
	}

  // put it to database
  _, err = models.NewVideoDaoInstance().CreateVideo(&video1)


  
	c.JSON(http.StatusOK, Response{

		StatusCode: 0,

		StatusMsg: finalName + " uploaded successfully",
	})

}

// PublishList all users have same publish video list

func PublishList(c *gin.Context) {

	c.JSON(http.StatusOK, VideoListResponse{

		Response: Response{

			StatusCode: 0,
		},

		VideoList: DemoVideos,
	})

}
