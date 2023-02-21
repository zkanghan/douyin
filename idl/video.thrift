namespace go video

include "base.thrift"

enum ActType{
    Do = 1
    Undo = 2
}

struct Video {
  1:required i64 id // 视频唯一标识
  2:required User author  // 视频作者信息
  3:required string play_url  // 视频播放地址
  4:required string cover_url  // 视频封面地址
  5:required i64 favorite_count  // 视频的点赞总数
  6:required i64 comment_count  // 视频的评论总数
  7:required bool is_favorite  // true-已点赞，false-未点赞
  8:required string title  // 视频标题
}

struct User {
  1: required i64 id  // 用户id
  2: required string name  // 用户名称
  3:optional i64 follow_count  // 关注总数
  4:optional i64 follower_count  // 粉丝总数
  5:required bool is_follow  // true-已关注，false-未关注
}

struct Comment {
  1:required i64 id  // 视频评论id
  2:required User user  // 评论用户信息
  3:required string content // 评论内容
  4:required string create_date  // 评论发布日期，格式 mm-dd
}

struct FeedRequest{
    1:optional i64 latest_time // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2:optional base.Token token // 可选参数，登录用户设置
}

struct FeedResponse{
    1: base.BaseResp BaseResp
    2: list<Video> video_list// 视频列表
    3: optional i64 next_time// 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct PublishRequest {
  1:required base.Token token // 用户鉴权token
  2:required binary data (vt.min_size = "1")  // 视频数据
  3:required string title(vt.min_size = '1')  // 视频标题
}

struct PublishResponse {
  1:base.BaseResp BaseResp
}

struct PublishListRequest {
  1:required i64 user_id  // 用户id
  2:required base.Token token  // 用户鉴权token
}

struct PublishListResponse {
  1:base.BaseResp BaseResp
  2:list<Video>  video_list  // 用户发布的视频列表
}

struct CommentActionRequest {
  1:required base.Token token  // 用户鉴权token
  2:required i64 video_id  // 视频id
  3:required i32 action_type  // 1-发布评论，2-删除评论
  4:optional string comment_text  // 用户填写的评论内容，在action_type=1的时候使用
  5:optional i64 comment_id  // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
  1:required base.BaseResp BaseResp
  2:optional Comment comment  // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct CommentListRequest {
  1:required base.Token token  // 用户鉴权token
  2:required i64 video_id // 视频id
}

struct CommentListResponse {
  1:required base.BaseResp BaseResp
  3:list<Comment> comment_list  // 评论列表
}


struct FavoriteActionRequest {
  1:required base.Token token  // 用户鉴权token
  2:required i64 video_id  // 视频id
  3:required i32 action_type  // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
  1:required base.BaseResp BaseResp
}

struct FavoriteListRequest {
  1:required i64 user_id  // 用户id
  2:required base.Token token  // 用户鉴权token
}

struct FavoriteListResponse {
  1:required base.BaseResp BaseResp
  2:list<Video> video_list  // 用户点赞视频列表
}

service VideoService{
    FeedResponse Feed(1:FeedRequest req)
    PublishListResponse Publish(1:PublishListRequest req)
    PublishListResponse PublishList(1:PublishListRequest req)
    CommentActionResponse CommentAction(1:CommentActionRequest req)
    CommentListResponse CommentList(1:CommentListRequest req)
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1:FavoriteListRequest req)
}