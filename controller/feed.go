package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed video list for every request
func Feed(c *gin.Context) {
	//token := c.Query("token")
	lstTime := time.Now().Format("2006-01-02 15:04:05")
	lstTimeStr := c.Query("latest_time")
	if lstTimeStr != "" {
		lstTimeNum, _ := strconv.ParseInt(lstTimeStr, 10, 64)
		//除1000，
		//1692110538 当前时间
		//1692110528343 参数
		lstTime = time.Unix(lstTimeNum/1000, 0).Format("2006-01-02 15:04:05")
	}
	videoList := service.VideoGet(lstTime)
	var nextTime int64
	var feedVideoList []Video = make([]Video, len(videoList))
	for i, v := range videoList {
		feedVideoList[i] = Video{
			Id: v.VideoId,
			//Author:
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			//IsFavorite:
			//title:
		}
	}
	if len(feedVideoList) > 0 {
		nextTime = videoList[0].CreateAt.Unix() * 1000
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: feedVideoList,
			NextTime:  nextTime,
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 0},
			NextTime: 0,
		})
	}
	fmt.Println(nextTime)
}
