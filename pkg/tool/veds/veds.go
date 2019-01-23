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
	Values []string
}

func Decrypt(client fs_base_veds.VEDSServer, values []string) Values {
	return process(client, false, values)
}

func Encrypt(client fs_base_veds.VEDSServer, values ...string) Values {
	return process(client, true, values)
}

func process(client fs_base_veds.VEDSServer, ref bool, values []string) Values {
	var wg sync.WaitGroup

	v := Values{
		State:  errno.Ok,
		Values: make([]string, len(values)),
	}

	errc := func(s *fs_base.State) {
		if !v.State.Ok {
			v.State = s
		}
		wg.Done()
	}

	wg.Add(len(values))

	for i := 0; i < len(values); i++ {
		go func(index int) {
			var r *fs_base_veds.CryptResponse
			var e error
			if ref {
				r, e = client.Encrypt(fs_metadata_transport.BuildInnerServiceAuthToContext(), &fs_base_veds.CryptRequest{
					Value: values[index],
				})
			} else {
				r, e = client.Decrypt(fs_metadata_transport.BuildInnerServiceAuthToContext(), &fs_base_veds.CryptRequest{
					Value: values[index],
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
			v.Values[index] = r.Value
			errc(errno.Ok)
		}(i)
	}

	wg.Wait()

	if len(values) != len(v.Values) {
		v.State = errno.ErrEncrypt
	}

	return v
}
