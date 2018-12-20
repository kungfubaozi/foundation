package validate

type Verification struct {
	UserId   string `bson:"user_id"`
	VerId    string `bson:"ver_id"`
	Code     string `bson:"code"`
	Do       int64  `bson:"do"` //做什么
	CreateAt int64  `bson:"create_at"`
	VerifyAt int64  `bson:"verify_at"` //验证时间
}
