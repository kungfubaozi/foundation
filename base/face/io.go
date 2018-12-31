package face

type faceset struct {
	UserId    string `bson:"user_id"`
	FaceToken string `bson:"face_token"`
	CreateAt  int64  `bson:"create_at"`
	ModifyAt  int64  `bson:"modify_at"`
}
