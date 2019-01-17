package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"os"
	"time"
	"zskparker.com/foundation/base/face/cmd/facecli"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/function/cmd/functioncli"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/project/cmd/projectcli"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/base/user/cmd/usercli"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/match"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/serv"
)

var (
	facePath   string
	email      string
	enterprise string
	phone      string
	username   string
	name       string
	password   string
	consul     string
)

func init() {
	flag.StringVar(&facePath, "face", "system_admin.jpg", "def admin face local path.")
	flag.StringVar(&email, "email", "", "def admin email.")
	flag.StringVar(&enterprise, "enterprise", "", "def admin enterprise account.")
	flag.StringVar(&phone, "phone", "", "def admin phone.")
	flag.StringVar(&username, "username", "", "def admin username.")
	flag.StringVar(&name, "name", "", "def admin realName.")
	flag.StringVar(&password, "password", "", "def admin password.")
	flag.StringVar(&consul, "consul", "", "micro services consul address.")

	flag.Usage = usage
}

func main() {
	flag.Parse()

	if len(facePath) == 0 {
		fmt.Println("please set admin face local path.")
		return
	}
	if !fs_regx_match.Email(email) {
		fmt.Println("email format err.")
		return
	}
	if !fs_regx_match.Phone(phone) {
		fmt.Println("phone format err.")
		return
	}
	if len(username) < 2 {
		fmt.Println("admin username length must >= 2")
		return
	}
	if len(name) < 2 {
		fmt.Println("admin name length must >= 2")
		return
	}
	if len(password) < 6 {
		fmt.Println("admin password length must >= 6")
		return
	}
	if len(consul) < 6 {
		fmt.Println("please set micro services consul address.")
		return
	}
	fmt.Println("system initialize...")
	os.Setenv("CONSUL_ADDR", consul)
	time.Sleep(200)

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.INIT, osenv.GetMicroPortString())
	defer reporter.Close()

	usersrv := usercli.NewClient(zipkinTracer)
	facesrv := facecli.NewClient(zipkinTracer)
	projectsrv := projectcli.NewClient(zipkinTracer)

	userId := bson.NewObjectId().Hex()
	userAlreadyRegister := false

	pr, err := projectsrv.Init(context.Background(), &fs_base_project.InitRequest{
		Desc:   "def project",
		En:     "foundation",
		Zh:     "foundation",
		UserId: userId,
	})
	if err != nil {
		fmt.Println("add project err", err)
		return
	}
	if !pr.State.Ok {
		fmt.Println("add project err", pr.State.Message)
		if pr.State.Code != errno.ErrProjectAlreadyExists.Code {
			return
		}
	} else {
		fmt.Println("def project session [set it to the API Service environment variable]:", pr.Session)
		fmt.Println("android client id:", pr.AndroidId)
		fmt.Println("iOS client id:", pr.IosId)
		fmt.Println("macOS client id:", pr.MaxOSId)
		fmt.Println("web client id:", pr.WebId)
		fmt.Println("windows client id:", pr.WindowsId)
	}

	time.Sleep(100)

	ur, err := usersrv.Add(context.Background(), &fs_base_user.AddRequest{
		UserId:        userId,
		Password:      password,
		Username:      username,
		RealName:      name,
		Enterprise:    enterprise,
		Level:         fs_constants.LEVEL_ADMIN,
		Phone:         phone,
		Email:         email,
		FromProjectId: pr.ProjectId,
	})
	if err != nil {
		fmt.Println("add user err", err)
		return
	}
	if !ur.State.Ok {
		userAlreadyRegister = true
		fmt.Println("add user err", ur.State.Message)
		if pr.State != errno.ErrUserAlreadyExists {
			return
		}
	}

	time.Sleep(100)

	f := functioncli.NewClient(zipkinTracer)

	fur, err := f.Init(context.Background(), &fs_base_function.InitRequest{
		Session: pr.Session,
	})

	if err != nil {
		fmt.Println("add function err", err)
		return
	}

	if !fur.State.Ok {
		fmt.Println("add function err", fur.State.Message)
		return
	}

	if !userAlreadyRegister {
		b, err := ioutil.ReadFile("system_admin.jpg")
		if err != nil {
			panic(err)
		}

		str := base64.StdEncoding.EncodeToString(b)

		fr, err := facesrv.Upsert(context.Background(), &fs_base_face.UpsertRequest{
			UserId:     ur.Content,
			Base64Face: str,
		})

		if err != nil {
			fmt.Println("add face err", err)
			return
		}
		if !fr.State.Ok {
			fmt.Println("add face err", fr.State.Message)
			return
		}
	}

	time.Sleep(100)

	fmt.Println("system initialize ok.")
}

func usage() {
	fmt.Fprintf(os.Stderr, `foundation initlizate command
Options:
`)
	flag.PrintDefaults()
}
