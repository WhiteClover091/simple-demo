#### 简介

我的工作为发布视频, 将这个任务分解为几个小任务, 先看哪些任务能够自己独立完成,先完成这些任务, 之后再根据小组成员的接口, 实现完整的功能



环境

- 需要提前安装`ffmpeg`





发布视频的流程

- 鉴权 ----- 使用别人的接口
- 存储视频   ---- (  demo中 有 )
- 抽取封面并存储  
- 视频 和 封面 上传到 `cdn` (  没有, 因此没有做 )
- 将收到的视频信息( 视频名称, 用户名, 播放地址, 封面的地址 )存储到数据库中   ---- 使用数据库的接口
- 返回消息 ( 发布成功)



因此最先做的工作就是

- 确定资源**存储的位置**, 通过查找**路由**得知, 有这样一个映射`r.Static("/static", "./public")`, 我们访问静态资源的时候, 网址为`xxx/static/fileName.txt`, 通过` c.Request.Host`字段, 来获得域名 : `2f781ee3592dd7a9ff0bbd0007fe40ce-app.1024paas.com`, 加头(协议)加尾(路径 + 文件名), 进行字符串拼接
- 使用`ffmpeg`来抽取某一帧来做封面, 需要设置 其 **环境**





```go
// controller/ publish.go
// 没有其他合作者的提供的接口
func Publish(c *gin.Context) {
    // token, 鉴权
 	
    // 存储
    err := c.SaveUploadedFile(data, saveFile);
    
    // 抽取并存储图片
    Vedio2Jpeg(saveFile, 6)
    
    // 拼接网址
    vedio_url := "https://"+ domain + "/static/"+ finalName
    
    // 放入数据库
    db.creat().....
    
}



```



抽取图片 demo

```go
//  将视频抽取一帧, 转化为流
package examples

import (
	"bytes"
	"fmt"
	"io"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
```



#### 队友的接口

```go
// 鉴权接口
// 参数: token, 返回: username
func ParseToken(token string) (string, error)


// 数据库接口---用户
func (*UserDao) FindUserByName(username string) (*User, error)
func (d *UserDao) FindUserById(id int64) (*User, error) 
_, err := models.NewUserDaoInstance().FindUserByName("qong");

// 数据库接口--- 视频
func (*VideoDao) CreateVideo(video *Video) (*Video, error) 

```





#### 使用了接口之后的伪代码

```go
func Publish(c *gin.Context) {
    // token, 鉴权
 	token := c.PostForm("token")
    user_name, err := ParseToken(token);
    user, err := models.NewUserDaoInstance().FindUserByName(user_name)
    var user_id = user.UserId
    
    // 存储
    err := c.SaveUploadedFile(data, saveFile);
    
    // 抽取并存储图片
    Vedio2Jpeg(saveFile, 6)
    
    // 拼接网址
    vedio_url := "https://"+ domain + "/static/"+ finalName
    
    // 放入数据库
    video1 := models.Video{UserId : user_id, PlayUrl : vedio_url, CoverUrl : img_url , }
    _, err = models.NewVideoDaoInstance().CreateVideo(&video1)
    
}
```







