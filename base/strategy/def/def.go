package strategydef

import "zskparker.com/foundation/base/pb"

func GetOnVerificationDefault() *fs_base.OnVerification {
	return &fs_base.OnVerification{
		EffectiveTime:   60 * 10, //10分钟
		CombinationMode: 1,       //数字验证
	}
}
