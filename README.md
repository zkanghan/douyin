划分为几个服务合适呢
避免环形依赖

?? 不同服务是共享数据库好还是不共享好   ans:不共享，微服务是自治的

只有视频流接口不用登录态

user    table user
    CheckUser()          //检查用户名和密码
    MGetUser(userID)  //用户的所有信息  依赖follow
   
relation
    Action() //关注操作
    FollowList()  //关注列表
    FollowerList() //粉丝列表
    FriendList() //互关列表   依赖message
    MGetFollow(userID, []toUserID) []bool 

video   table: video favorite comment
    Feed() //视频流  依赖user relation
    Publish() //发布视频
    PublishList() //发布列表 依赖user relation
    CommentAction() //发布评论
    CommentList()   //评论列表
    FavoriteAction()    //视频点赞
    FavoriteList()  //点赞列表  依赖user relation

message
    Action() //发送消息
    List() //消息列表
   