package service

import "github.com/RaymondCode/simple-demo/models"

var videoDao = models.NewVideoDaoInstance()
var limVideoNum = 30

//获取lastTime之前的视频，最新视频的位于最前面
func VideoGet(lastTime string) []*models.Video {
	return videoDao.QueryVideo(&lastTime, limVideoNum)
}
