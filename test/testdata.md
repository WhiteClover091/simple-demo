# 插入数据

## test1

url改成自己的ip

insert into videos
(
    video_id,user_id,play_url,cover_url,favorite_count,comment_count,create_at
)
values
(
    1,1,"http://192.168.2.128:8080/static/bear.mp4",
    "http://192.168.2.128:8080/static/bear.jpg",
    100,
    200,
    Current_timestamp()
)

## test2

insert into videos
(
    video_id,user_id,play_url,cover_url,favorite_count,comment_count,create_at
)
values
(
    2,2,"http://192.168.2.128:8080/static/arknight.mp4",
    "http://192.168.2.128:8080/static/arknight.jpg",
    200,
    300,
    Current_timestamp()
)

## test3

insert into videos
(
    video_id,user_id,play_url,cover_url,favorite_count,comment_count,create_at
)
values
(
    3,3,"http://192.168.2.128:8080/static/wanglei.mp4",
    "http://192.168.2.128:8080/static/wanglei.jpg",
    300,
    400,
    Current_timestamp()
)