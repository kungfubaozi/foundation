package usersync

type repository interface {
	Add()
}

type model struct {
	UserId    string `bson:"user_id"`
	Synced    bool   `bson:"synced"`
	ProjectId string `bson:"project_id"`
}
