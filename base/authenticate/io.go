package authenticate

//作为token
type SimpleAuthorize struct {
	UserId   string
	ClientId string
	UUID     string
	Access   bool
	Relation string
}
