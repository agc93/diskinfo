// +build windows

package main

// this is currently pointless. Aside from the fact that Windows doesn't
// really have an equivalent to ata link numbers, the `ghw` library
// doesn't actually support Windows. oh well.

func GetDeviceNumber(deviceName string) string {

	return ""

}
