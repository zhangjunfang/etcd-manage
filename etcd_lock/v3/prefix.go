package v3

import (
	"os"
	"strings"

	"github.com/sony/sonyflake"
)

const (
	prefix = "/etcd-lock"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	name, _ := os.Hostname()
	st.MachineID = name
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}
func addPrefix(key string) string {
	if !strings.HasPrefix(key, "/") {
		key = "/" + key
	}
	id, _ := sf.NextID()
	return prefix + key + id
}
