package util

import (
	"fmt"
	"syscall"

	"github.com/MiyamuraMiyako/go-util/introp"
)

type LUID struct {
	LowPart  uint32
	HighPart int32
}

type LUID_AND_ATTRIBUTES struct {
	Luid       LUID
	Attributes uint32
}

type TokPriv1Luid struct {
	PrivilegeCount uint32
	Privileges     [1]LUID_AND_ATTRIBUTES
}

//PowerOff CAN NOT USE YET!!!!
func PowerOff(force, reboot bool) {
	advapi32 := syscall.NewLazyDLL("advapi32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	LookupPrivilegeValue := advapi32.NewProc("LookupPrivilegeValueW")
	AdjustTokenPrivileges := advapi32.NewProc("AdjustTokenPrivileges")
	ExitWindowsEx := user32.NewProc("ExitWindowsEx")
	GetCurrentProcess := kernel32.NewProc("GetCurrentProcess")

	r1, _, err := GetCurrentProcess.Call()
	fmt.Println(err.Error())
	var tk syscall.Token

	syscall.OpenProcessToken((syscall.Handle)(r1), 40, &tk)

	laa := LUID_AND_ATTRIBUTES{
		Luid:       LUID{},
		Attributes: 2}
	toPrivLuid := TokPriv1Luid{
		PrivilegeCount: 1,
		Privileges:     [1]LUID_AND_ATTRIBUTES{laa}}

	LookupPrivilegeValue.Call(0, introp.ToPtr("SeShutdownPrivilege"), introp.ToPtr(toPrivLuid.Privileges[0].Luid))
	AdjustTokenPrivileges.Call((uintptr)(tk), 0, introp.ToPtr(toPrivLuid), 0, 0, 0)
	if reboot {
		flg := 6
		if force {
			flg = 18
		}
		ExitWindowsEx.Call(uintptr(flg), 0)
	} else {
		ExitWindowsEx.Call(1, 0)
	}
}
