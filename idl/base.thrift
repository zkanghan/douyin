namespace go base

struct BaseResp {
  1:required i32 status_code  // 状态码，0-成功，其他值-失败
  2:optional string status_msg  // 返回状态描述
}

struct Token{
    1:optional i64 user_id      //token解析出的user_id
}