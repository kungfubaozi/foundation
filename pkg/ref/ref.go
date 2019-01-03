package ref

import (
	"reflect"
	"zskparker.com/foundation/base/pb"
)

func GetMetaInfo(request interface{}) *fs_base.Meta {
	m := reflect.ValueOf(request).Elem().FieldByName("Meta")
	if !m.CanSet() {
		return &fs_base.Meta{}
	}
	return m.Interface().(*fs_base.Meta)
}
