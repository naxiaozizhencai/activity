package messages

const (
	AuthCodeErr      = "auth_code_err"       //验证码错误
	DataErr          = "data_err"            //数据错误
	Success          = "success"             //操作成功
	AuthCodeExist    = "auth_code_exist"     //验证码已经发送
	AuthCodeSendFail = "auth_code_send_fail" //发送失败
	ActiveOpenErr    = "active_open_err"     //活动未开启
	AnswerHasSuccess = "active_has_success"  //答题已完成
	AnswerTimesOut   = "answer_times_out"    //问题次数已使用完
	AnswerErr        = "answer_err"          //回答错误
	SaveFail         = "save_err"            //保存失败
	AnswerTimeOut    = "answer_time_out"     //问题答题超时
	Reward_Err1      = "reward_err1"         //未达到领取条件
	Reward_Err2      = "reward_err2"         //奖励已领取，不能重复领取
	SendMailErr      = "send_mail_error"     //邮件发送失败
	ActivityOver     = "activity_over"
)
