namespace go relation

include "base.thrift"

enum ActType{       //action 类型
    Do = 1
    Undo = 2
}

struct ActionRequest {
  1:required base.Token token  // 用户鉴权token
  2:required i64 to_user_id  // 对方用户id
  3:required i32 action_type(vt.in = "1", vt.in = "2")  // 1-关注，2-取消关注
}

struct ActionResponse {
  1:base.BaseResp BaseResp
}

struct FollowListRequest {
  1:required i64 user_id  // 用户id
  2:required base.Token token // 用户鉴权token
}

struct FollowListResponse {
  1:required base.BaseResp BaseResp
  2:list<User> user_list  // 用户信息列表
}

struct User {
  1:required i64 id // 用户id
  2:required string name  // 用户名称
  3:optional i64 follow_count  // 关注总数
  4:optional i64 follower_count// 粉丝总数
  5:required bool is_follow  // true-已关注，false-未关注
}

struct FollowerListRequest {
  1:required i64 user_id  // 用户id
  2:required base.Token token // 用户鉴权token
}

struct FollowerListResponse {
  1:required base.BaseResp BaseResp
  2:list<User> user_list  // 用户列表
}

struct FriendListRequest {
  1:required i64 user_id  // 用户id
  2:required base.Token token  // 用户鉴权token
}

struct FriendListResponse {
  1:required base.BaseResp BaseResp
  2:optional list<FriendUser> user_list  // 用户列表
}

struct FriendUser  {
    1:User user
    2:optional string message // 和该好友的最新聊天消息
    3:required i64 msgType  // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

struct MGetFollowRequest{
    1:required i64 user_id
    2:required list<i64> to_user_id(vt.min_size = "1")
}

struct MGetFollowResponse{
    2:required base.BaseResp BaseResp
    1:optional list<bool> flags
}

service RelationService{
    ActionResponse Action(1:ActionRequest req)
    FollowListResponse FollowList(1:FollowListRequest req)
    FollowerListResponse FollowerList(1:FollowerListRequest req)
    FriendListResponse FriendList(1:FriendListRequest req)
    MGetFollowResponse MgetFollow(1:MGetFollowRequest req)
}