package introp

import "runtime"

//KeepAlive is batch invoke runtime.KeepAlive() with many vars.
func KeepAlive(vars ...interface{}) {
	for v := range vars {
		runtime.KeepAlive(v)
	}
}
