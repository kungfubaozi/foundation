package fs_service_veds

import (
	"sync"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/veds/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/transport"
)

type Values struct {
	State  *fs_base.State
	Values map[int]string
}

func Decrypt(client fs_base_veds.VEDSServer, values map[int]string) Values {
	return process(client, false, values)
}

func Encrypt(client fs_base_veds.VEDSServer, values map[int]string) Values {
	return process(client, true, values)
}

func process(client fs_base_veds.VEDSServer, ref bool, values map[int]string) Values {
	var wg sync.WaitGroup

	v := Values{
		State:  errno.Ok,
		Values: make(map[int]string),
	}

	errc := func(s *fs_base.State) {
		if !v.State.Ok {
			v.State = s
		}
		wg.Done()
	}

	wg.Add(len(values))

	for k, va := range values {
		go func(key int) {
			var r *fs_base_veds.CryptResponse
			var e error
			if ref {
				r, e = client.Encrypt(fs_metadata_transport.BuildInnerServiceAuthToContext(), &fs_base_veds.CryptRequest{
					Value: values[k],
				})
			} else {
				r, e = client.Decrypt(fs_metadata_transport.BuildInnerServiceAuthToContext(), &fs_base_veds.CryptRequest{
					Value: values[k],
				})
			}
			if e != nil {
				errc(errno.ErrSystem)
				return
			}
			if !r.State.Ok {
				errc(r.State)
				return
			}
			v.Values[k] = r.Value
			errc(errno.Ok)
		}(k)
	}

	wg.Wait()

	if len(values) != len(v.Values) {
		v.State = errno.ErrEncrypt
	}

	return v
}
