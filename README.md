# Mini-Tiktok
bytedance camp project
# Mini-Tiktok

1.实体类和vo对象的区别和协同，刚开始有点模糊
2.使用ffmpeg实现视频帧图片获取
3.学习使用gorm查询语句
4.在数据对象操作时要记得加 &  ，否则会出现很多错误，比如db.Find(videos),应该是db.Find(&videos)
5.注意前端需要的字段格式，需要严格匹配 ！定义返回对象结构体时严格声明 `json:"video_list,omitempty"`