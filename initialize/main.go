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
	"zskparker.com/foundation/base/strategy/cmd/strategycli"
	"zskparker.com/foundation/base/strategy/pb"
	"zskparker.com/foundation/base/user/cmd/usercli"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/base/veds/cmd/vedscli"
	"zskparker.com/foundation/base/veds/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/match"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/serv"
	"zskparker.com/foundation/pkg/tool/veds"
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

	vedssvc := vedscli.NewClient(zipkinTracer)
	usersrv := usercli.NewClient(zipkinTracer)
	facesrv := facecli.NewClient(zipkinTracer)
	projectsrv := projectcli.NewClient(zipkinTracer)
	strategysrv := strategycli.NewClient(zipkinTracer)
	functionsrv := functioncli.NewClient(zipkinTracer)

	userId := bson.NewObjectId().Hex()
	userAlreadyRegister := false

	{
		//test
		vr, err := vedssvc.Encrypt(context.Background(), &fs_base_veds.CryptRequest{
			Value: "test",
		})
		if err != nil {
			panic(err)
		}
		if !vr.State.Ok {
			panic(vr.State.Message)
		}
	}

	pr, err := projectsrv.Init(context.Background(), &fs_base_project.InitRequest{
		Desc:   "root project",
		En:     "SSORoute",
		Zh:     "SSORoute",
		UserId: userId,
	})
	if err != nil {
		fmt.Println("add project err", err)
		return
	}
	if !pr.State.Ok {
		fmt.Println("add project err", pr.State.Message)
		return
	} else {

		v := fs_service_veds.Encrypt(vedssvc, pr.Session, pr.AndroidId, pr.IosId, pr.MacOSId, pr.WebId, pr.WindowsId)

		if !v.State.Ok {
			panic(v.State.Message)
		}

		fmt.Println("def project session [set it to the API Service environment variable]:", v.Values[0])
		fmt.Println("android client id:", v.Values[1])
		fmt.Println("iOS client id:", v.Values[2])
		fmt.Println("macOS client id:", v.Values[3])
		fmt.Println("web client id:", v.Values[4])
		fmt.Println("windows client id:", v.Values[5])

		sr, err := strategysrv.Init(context.Background(), &fs_base_strategy.InitRequest{
			Session: pr.Session,
			Creator: userId,
		})

		if err != nil {
			fmt.Println("add root strategy err", err)
			return
		}

		if !sr.State.Ok {
			fmt.Println("add root strategy", sr.State.Message)
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

		fur, err := functionsrv.Init(context.Background(), &fs_base_function.InitRequest{
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
	}

	fmt.Println("system initialize ok.")
}

func usage() {
	fmt.Fprintf(os.Stderr, `foundation initlizate command
Options:
`)
	flag.PrintDefaults()
}
