package functionsvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net"
	"os"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/invite"
	"zskparker.com/foundation/base/project"
	"zskparker.com/foundation/base/refresh"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/strategy"
	"zskparker.com/foundation/base/usersync"
	"zskparker.com/foundation/entry/login"
	"zskparker.com/foundation/entry/register"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/registration"
	"zskparker.com/foundation/pkg/serv"
	"zskparker.com/foundation/safety/blacklist"
	"zskparker.com/foundation/safety/froze"
	"zskparker.com/foundation/safety/unblock"
	"zskparker.com/foundation/safety/update"
	"zskparker.com/foundation/safety/verification"
)

func StartService() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var otTracer stdopentracing.Tracer
	{
		otTracer = stdopentracing.GlobalTracer()
	}

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), names.F_SVC_FUNCTION, osenv.GetMicroPortString())
	defer reporter.Close()

	session, err := db.CreateSession(osenv.GetMongoDBAddr())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//插入默认功能
	insertDef(session)

	rs, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), names.F_SVC_FUNCTION)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	service := function.NewService(session, rs)
	endpoints := function.NewEndpoints(service, zipkinTracer, logger)
	svc := function.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_function.RegisterFunctionServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, names.F_SVC_FUNCTION, osenv.GetConsulAddr())

	go func() {
		grpcListener, err := net.Listen("tcp", osenv.GetMicroPortString())
		if err != nil {
			fmt.Println(err)
			errc <- err
		}
		errc <- gs.Serve(grpcListener)
	}()

	logger.Log("exit", <-errc)
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

func upsert(c *mgo.Collection, f *fs_pkg_model.APIFunction) {
	c.Upsert(bson.M{"api": f.Function.Api}, &function.FunctionModel{
		Func:  f.Function.Func,
		ZH:    f.Function.Zh,
		Level: f.Function.Level,
		Fcv:   f.Function.Fcv,
		EN:    f.Function.En,
		API:   f.Prefix + f.Infix,
	})
}
