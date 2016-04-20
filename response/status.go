package goresponse

const Status_ignore int = -1 // -1 忽略
const (
	Status_fail          int = iota // 0 失败
	Status_success                  // 1 成功
	Status_invalid_id               // 2 无效id
	Status_invalid_user             // 3 无效用户
	Status_invalid_token            // 4 无效token
	Status_token_timeout            // 5 token超时
	Status_no_permission            // 6 无权限
	Status_no_data                  // 7 无数据
	Status_disuse                   // 8 无用
	Status_duplication              // 9 重复
	Status_fail_server              // 10 服务器错误
	Status_fail_sql                 // 11 数据错误
	Status_fail_captcha             // 12 验证码错误
	Status_fail_request             // 13 请求错误
	status_fail_none0               // 14 (预留)
	status_fail_none1               // 15 (预留)
	status_fail_none2               // 16 (预留)
	status_fail_none3               // 17 (预留)
	status_fail_none4               // 18 (预留)
	status_fail_none5               // 19 (预留)
	Status_fail_arg0                // 20 第1个参数错误
	Status_fail_arg1                // 21 第2个参数错误
	Status_fail_arg2                // 22 第3个参数错误
	Status_fail_arg3                // 23 第4个参数错误
	Status_fail_arg4                // 24 第5个参数错误
	Status_fail_arg5                // 25 第6个参数错误
	Status_fail_arg6                // 26 第7个参数错误
	Status_fail_arg7                // 27 第8个参数错误
	Status_fail_arg8                // 28 第9个参数错误
	Status_fail_arg9                // 29 第10个参数错误
	Status_fail_arg                 // 30 参数错误
)