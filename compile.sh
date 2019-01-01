#!/usr/bin/env bash


protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. analysis/statistics/pb/statistics.proto

protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/pb/base.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/pb/strategy1.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/authenticate/pb/authenticate.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/reporter/pb/reporter.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/face/pb/face.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/project/pb/project.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/message/pb/message.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/strategy/pb/strategy1.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/user/pb/user.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/validate/pb/validate.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/function/pb/function.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/interceptor/pb/interceptor.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/review/pb/review.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/state/pb/state.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/userinfo/pb/userinfo.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/usersync/pb/usersync.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. base/invite/pb/invite.proto

protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. entry/login/pb/login.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. entry/register/pb/register.proto

protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. safety/blacklist/pb/blacklist.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. safety/froze/pb/froze.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. safety/trustdevice/pb/trustdevice.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. safety/unblock/pb/unblock.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. safety/verification/pb/verification.proto
protoc -I /Users/Richard/Desktop/Development/Golang/src/ --proto_path=. --go_out=plugins=grpc:. safety/update/pb/update.proto
