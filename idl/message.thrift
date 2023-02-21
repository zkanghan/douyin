namespace go message

include "base.thrift"

// 聊天记录
struct RecordRequest {
  1:required base.Token token  // 用户鉴权token
  2:required i64 to_user_id  // 对方用户id
}

struct RecordResponse {
  1:required base.BaseResp BaseResp
  2:list<Message> message_list  // 消息列表
}

struct Message {
  1:required i64 id  // 消息id
  2:required i64 to_user_id  // 该消息接收者的id
  3:required i64 from_user_id  // 该消息发送者的id
  4:required string content  // 消息内容
  5:optional string create_time  // 消息创建时间
}

struct ActionRequest {
  1:required base.Token token  // 用户鉴权token
  2:required i64 to_user_id  // 对方用户id
  3:required i32 action_type  // 1-发送消息
  4:required string content  // 消息内容
}

struct ActionResponse {
  1:required base.BaseResp BaseResp
}

service MessageSerivce{
    ActionResponse Send(1:ActionRequest req)
    RecordResponse Record(1:RecordRequest req)
}