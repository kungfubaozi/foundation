package errno

import "zskparker.com/foundation/base/pb"

func ErrResponse(state *fs_base.State) (*fs_base.Response, error) {
	return &fs_base.Response{}, nil
}

var (
	Ok         = &fs_base.State{Ok: true, Message: "ok", Code: 0}
	ErrRequest = &fs_base.State{Ok: false, Message: "err request", Code: 10001}
	ErrSystem  = &fs_base.State{Ok: false, Message: "err system", Code: 10002}
)
