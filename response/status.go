package response

const Status_ignore int = -1 // -1 忽略
const (
	Status_fail            int = iota // 0 失败
	Status_success                    // 1 成功
	Status_invalid_user               // 2 无效用户
	Status_inactive_user              // 3 未激活用户
	Status_invalid_token              // 4 token超时
	Status_abandoned_token            // 5 token变更
	Status_no_permission              // 6 无权限
	Status_no_data                    // 7 无数据
	Status_disuse                     // 8 无用
	Status_duplication                // 9 重复
	Status_fail_server                // 10 服务器错误
	Status_fail_sql                   // 11 数据错误
	Status_fail_captcha               // 12 验证码错误
	Status_fail_request               // 13 请求错误
	Status_fail_meta                  // 14 meta错误
	Status_fail_query                 // 15 query错误
	Status_fail_order                 // 16 order错误
	Status_fail_limit                 // 17 limit错误
	Status_fail_illegal               // 18 非用户拥有
	Status_fail_frequently            // 19 频繁请求
	Status_fail_arg0                  // 20 第1个参数错误
	Status_fail_arg1                  // 21 第2个参数错误
	Status_fail_arg2                  // 22 第3个参数错误
	Status_fail_arg3                  // 23 第4个参数错误
	Status_fail_arg4                  // 24 第5个参数错误
	Status_fail_arg5                  // 25 第6个参数错误
	Status_fail_arg6                  // 26 第7个参数错误
	Status_fail_arg7                  // 27 第8个参数错误
	Status_fail_arg8                  // 28 第9个参数错误
	Status_fail_arg9                  // 29 第10个参数错误
	Status_fail_arg                   // 30 参数错误
	Status_fail_arg11                 // 31 第11个参数错误
	Status_fail_arg12                 // 32 第12个参数错误
	Status_fail_arg13                 // 33 第13个参数错误
	Status_fail_arg14                 // 34 第14个参数错误
	Status_fail_arg15                 // 35 第15个参数错误
	Status_fail_arg16                 // 36 第16个参数错误
	Status_fail_arg17                 // 37 第17个参数错误
	Status_fail_arg18                 // 38 第18个参数错误
	Status_fail_arg19                 // 39 第19个参数错误
	Status_fail_arg20                 // 40 第20个参数错误
	Status_fail_arg21                 // 41 第21个参数错误
	Status_fail_arg22                 // 42 第22个参数错误
	Status_fail_arg23                 // 43 第23个参数错误
	Status_fail_arg24                 // 44 第24个参数错误
	Status_fail_arg25                 // 45 第25个参数错误
	Status_fail_arg26                 // 46 第26个参数错误
	Status_fail_arg27                 // 47 第27个参数错误
	Status_fail_arg28                 // 48 第28个参数错误
	Status_fail_arg29                 // 49 第29个参数错误
	Status_fail_arg30                 // 50 第30个参数错误
	Status_fail_arg31                 // 51 第31个参数错误
	Status_fail_arg32                 // 52 第32个参数错误
	Status_fail_arg33                 // 53 第33个参数错误
	Status_fail_arg34                 // 54 第34个参数错误
	Status_fail_arg35                 // 55 第35个参数错误
	Status_fail_arg36                 // 56 第36个参数错误
	Status_fail_arg37                 // 57 第37个参数错误
	Status_fail_arg38                 // 58 第38个参数错误
	Status_fail_arg39                 // 59 第39个参数错误
)

const (
	Status_forbidden = 403 // 403 服务拒绝
)

const (
	Status_service_close int = 500 // 500 服务器关闭
)

const (
	Status_out_workday int = 601 // 601 非工作时间
)
