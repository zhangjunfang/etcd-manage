package etcd_queue

import (
	"os"

	"github.com/sony/sonyflake"
)

const (
	prefix = "/queue/"
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
func addPrefix() string {
	id, _ := sf.NextID()
	return prefix + id
}
