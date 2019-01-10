package authenticate

//作为token
type SimpleAuthorize struct {
	UserId   string
	ClientId string
	UUID     string //用来标记当前token的ID
	Access   bool
	Relation string
}
