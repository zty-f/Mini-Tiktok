# Mini-Tiktok
bytedance camp project


### 2022.6.2
1.实体类和vo对象的区别和协同，刚开始有点模糊
2.使用ffmpeg实现视频帧图片获取
3.学习使用gorm查询语句
4.在数据对象操作时要记得加 &  ，否则会出现很多错误，比如db.Find(videos),应该是db.Find(&videos)
5.注意前端需要的字段格式，需要严格匹配 ！定义返回对象结构体时严格声明 `json:"video_list,omitempty"`

### 2022.6.3
方法的注意事项:
1. 结构体是值类型,在方法调用中遵守值类型的传递机制;
2. 如果需要修改结构体变量的值,类似上述例子,通过结构体指针的方式处理;
3. Go 中的方法作用在指定的数据类型的方法上(与指定的数据类型绑定);
4. **方法的控制范围的规则与函数一样,方法名首字母小写,只能在本包内访问,方法首字母大写,可以在本包和其它包访问;**
5. 如果一个结构体实现了String() 方法,那么fmt.Println() 默认会调用String 方法进行输出;
6. 对于方法接收者是值类型时,可以直接用指针类型的变量调用,反过来同样可以;
7. 不管接收者是值类型还是引用类型都可以直接使用实例变量调用方法; 
其他注意事项：
8. 写清楚前后端的POST和GET方法匹配，不然就会404
9. gorm查询数据库得到结果传给变量需要加&，不然会报错。

### 2022.6.5
1. 前端发布视频数量显示一直为0，应该是前端还没设置好
2. 所有点赞或关注的标志位都应该用登录用户去做判断！！！！！！

#### import cycle not allowed 循环导包问题待解决，可以多加一层或者吧循环调用放在下面的层！！！！！！

### 2022.6.6
1. 解决循环导包问题
2. ？获取用户信息/douyin/user/接口需不需要进行登录效验？（未登录能不能调用此接口？？？？）
   答：目前来说应该是不行吧，传入参数有id和token，查id用户信息，用token判断是否登录，未登录不能查询用户信息
3. 项目所有接口基本完成，待优化~~~~~~

### 2022.6.8
1. sql注入问题gorm框架已经防止了，通过占位符？的方式来解决
2. 评论删除功能还存在权限问题，应该只能登录才能删除评论，并且只能删除自己的评论。
3. 其他应该拦截的接口已经做了接口统一拦截，能够保证基本的安全问题
4. 因为网络是http明文传输，所以密码应该进行加密处理然后保存在数据库！
5. 登录功能应该先判断是否包含此用户名，然后再进行密码效验。