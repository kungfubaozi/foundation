package functionsvc

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/invite"
	"zskparker.com/foundation/base/project"
	"zskparker.com/foundation/base/refresh"
	"zskparker.com/foundation/base/strategy"
	"zskparker.com/foundation/base/usersync"
	"zskparker.com/foundation/entry/login"
	"zskparker.com/foundation/entry/register"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/safety/blacklist"
	"zskparker.com/foundation/safety/froze"
	"zskparker.com/foundation/safety/unblock"
	"zskparker.com/foundation/safety/update"
	"zskparker.com/foundation/safety/verification"
)

func StartService() {

}

func upsert(c *mgo.Collection, function *fs_pkg_model.APIFunction) {
	c.Upsert(bson.M{"api": function.Function.Api}, function)
}

func insertDef(session *mgo.Session) {
	c := session.DB("foundation").C("functions")

	//login functions
	upsert(c, login.GetEntryByFaceFunc())
	upsert(c, login.GetEntryByValidateCodeFunc())
	upsert(c, login.GetEntryByAPFunc())
	upsert(c, login.GetEntryByOAuthFunc())
	upsert(c, login.GetEntryByQRCodeFunc())

	//safety verification functions
	upsert(c, verification.GetNewFunc())

	//register functions
	upsert(c, register.GetFromAPFunc())
	upsert(c, register.GetFromOAuthFunc())

	//safety update functions
	upsert(c, update.GetUpdateEmailFunc())
	upsert(c, update.GetUpdateEnterpriseFunc())
	upsert(c, update.GetUpdatePasswordFunc())
	upsert(c, update.GetUpdatePhoneFunc())

	//unblock
	upsert(c, unblock.GetUnlockFunc())

	//blacklist
	upsert(c, blacklist.GetAddBlacklistFunc())
	upsert(c, blacklist.GetRemoveBlacklistFunc())

	//function
	upsert(c, function.GetAddFunc())
	upsert(c, function.GetAllFunc())
	upsert(c, function.GetFindFunc())
	upsert(c, function.GetRemoveFunc())
	upsert(c, function.GetUpdateFunc())

	//authorization token refresh functions
	upsert(c, refresh.GetRefreshFunc())

	//project functions
	upsert(c, project.GetCreateProject())
	upsert(c, project.GetRemoveProject())
	upsert(c, project.GetUpdateProject())

	//usersync functions
	upsert(c, usersync.GetAddUserSyncHookFunc())
	upsert(c, usersync.GetRemoveUserSyncHookFunc())
	upsert(c, usersync.GetUpdateUserSyncHookFunc())

	//strategy functions
	upsert(c, strategy.GetUpdateProjectStrategyFunc())

	//review functions

	//invite functions
	upsert(c, invite.GetInviteUserFunc())

	//froze
	upsert(c, froze.GetRequestFrozeFunc())

}
