//go:build windows

package auth

import (
	"syscall"
	"unsafe"
)

var (
	advapi32       = syscall.NewLazyDLL("advapi32.dll")
	procLogonUserW = advapi32.NewProc("LogonUserW")
)

func Authenticate(user, pass string) (bool, error) {
	var token syscall.Handle
	userPtr, _ := syscall.UTF16PtrFromString(user)
	domainPtr, _ := syscall.UTF16PtrFromString(".")
	passPtr, _ := syscall.UTF16PtrFromString(pass)

	ret, _, err := procLogonUserW.Call(
		uintptr(unsafe.Pointer(userPtr)),
		uintptr(unsafe.Pointer(domainPtr)),
		uintptr(unsafe.Pointer(passPtr)),
		2, 0,
		uintptr(unsafe.Pointer(&token)),
	)

	if ret == 0 {
		return false, err
	}
	syscall.CloseHandle(token)
	return true, nil
}
