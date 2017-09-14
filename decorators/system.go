package proto

import (
	"github.com/slavaromanov/sysinfo"
)

var (
	system  = "Linux"
	device  = "Laptop"
	version = "0.0.2"
)

func format(arg1, arg2 string) string {
	return arg1 + " (" + args2 + ")"
}

func Init() error {
	info, err := sysinfo.NewInfo()
	if err != nil {
		return err
	}
	system = format(info.OS, info.Arch)
	device = format(info.Name, info.Version)
	return nil
}
