namespace go user

include "base.thrift"

struct CheckUserRequest {
  1: required string username(vt.min_size = "1",vt.max_size = "32")  // 注册用户名，最长32个字符
  2: required string password(vt.min_size = "1",vt.max_size = "32")  // 密码，最长32个字符
}

struct CheckUserResponse {
  1:required base.BaseResp BaseResp
  2:required i64 user_id  // 用户id
}

struct MGetUserRequest {
  1:required list<i64> user_ids  // 用户id
  2:required base.Token token // 用户鉴权token
}

struct MGetUserResponse {
  1:required base.BaseResp BaseResp
  3:required list<User> users  // 用户信息
}

struct CreateUserRequest{
    1: required string username(vt.min_size = "1",vt.max_size = "32")  // 注册用户名，最长32个字符
    2: required string password(vt.min_size = "1",vt.max_size = "32")  // 密码，最长32个字符
}

struct CreateUserResponse{
    1:required base.BaseResp Resp
    2:optional i64 user_id
}

struct User {
  1:required i64 id // 用户id
  2:required string name  // 用户名称
  3:optional i64 follow_count // 关注总数
  4:optional i64 follower_count // 粉丝总数
  5:required bool is_follow  // true-已关注，false-未关注
}

service UserService{
    // 用户登录调用CheckUser  用户注册 调用CreateUser --> CheckUser
    CheckUserResponse CheckUser(1:CheckUserRequest req)
    MGetUserResponse MGetUser(1:MGetUserRequest req)
    CreateUserResponse CreateUser(1:CreateUserRequest req)
}
