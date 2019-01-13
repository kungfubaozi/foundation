package strategydef

import (
	"time"
	"zskparker.com/foundation/base/pb"
)

func GetOnVerificationDefault() *fs_base.OnVerification {
	return &fs_base.OnVerification{
		EffectiveTime:   10, //10分钟
		CombinationMode: 1,  //数字验证
		VoucherDuration: 60, //60秒
	}
}

func DefStrategy(projectId, creator string) *fs_base.ProjectStrategy {
	return &fs_base.ProjectStrategy{
		ProjectId: projectId,
		CreateAt:  time.Now().UnixNano(),
		Creator:   creator,
		Version:   1,
		Configuration: &fs_base.Configuration{
			OpenTime: "0-24",
		},
		Events: &fs_base.Events{
			OnRegister: &fs_base.OnRegister{
				OpenReview:                   1, //不开启审核
				Mode:                         1, //手机注册
				AnIPRegistrationInterval:     5,
				AnDeviceRegistrationInterval: 5,
				Submitlal:                    1, //不提交经纬度信息
				AllowNewRegistrations:        2,
			},
			OnLogin: &fs_base.OnLogin{
				AllowLogin:                   2,
				AllowOtherProjectUserToLogin: 2,
				Mode: []int64{
					1, 2, 3, 4, 5, 6,
				},
				MaxCountOfOnline: &fs_base.MaxCountOfOnline{
					Android: 1,
					IOS:     1,
					Windows: 1,
					MacOS:   1,
					Web:     0, //无限制
				},
				Verification: 1,
				MaxCountOfErrorPassword: []*fs_base.MaxCountOfErrorPassword{
					{
						Count:  3,
						Action: 4,
					},
					{
						Count:  5,
						Action: 5,
					},
					{
						Count:       8,
						Action:      3,
						ExpiredTime: 10 * 60, //10分钟
					},
				},
				MaxCountOfInvalidAccount: []*fs_base.MaxCountOfInvalidAccount{
					{
						Count:  3,
						Action: 4,
					},
					{
						Count:  5,
						Action: 5,
					},
					{
						Count:       8,
						Action:      3,
						ExpiredTime: 10 * 60, //10分钟
					},
				},
				Submitlal: 1, //不用提交经纬度
			},
			OnVerification: GetOnVerificationDefault(),
			OnQRLogin: &fs_base.OnQRLogin{
				RefreshDuration: 60, //单位秒
			},
			OnFaceLogin: &fs_base.OnFaceLogin{
				Degree: 80,
			},
			OnCommonEquipmentChanges: &fs_base.OnCommonEquipmentChanges{
				SendMessageToUser: 2,
			},
			OnRequestFrozen: &fs_base.OnRequestFrozen{
				Verification: 2,
			},
			OnCancelFrozen: &fs_base.OnCancelFrozen{
				Verification: 2,
			},
			OnChangePhoneNumber: &fs_base.OnChangePhoneNumber{
				Verification: 2,
			},
			OnChangeEmail: &fs_base.OnChangeEmail{
				Verification: 2,
			},
			OnChangeFace: &fs_base.OnChangeFace{
				Verification: 2,
			},
			OnChangeOAuth: &fs_base.OnChangeOAuth{
				Verification: 1,
			},
			OnResetPassword: &fs_base.OnResetPassword{ //两种方式都可重置密码
				Phone: 2,
				Email: 2,
			},
			OnElsewhereLogin: &fs_base.OnElsewhereLogin{
				SendMessageToUser: 2,
				Verification:      2,
			},
			OnSubmitReview: &fs_base.OnSubmitReview{},
			OnInviteUser: &fs_base.OnInviteUser{
				ExpireTime: 48, //48小时
				Review:     2,
			},
		}}
}
