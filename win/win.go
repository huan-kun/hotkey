// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// +build windows

// Package hotkey_win is win32api wrapper for hotkey.
package hotkey_win

import (
	. "github.com/lxn/win"
	"golang.org/x/sys/windows"
	"syscall"
)

var (
	libuser32   *windows.LazyDLL
	libkernel32 *windows.LazyDLL

	registerHotKey    *windows.LazyProc
	unregisterHotKey  *windows.LazyProc
	postThreadMessage *windows.LazyProc

	getCurrentThread *windows.LazyProc
	getThreadId      *windows.LazyProc
)

func init() {
	// Library
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	libkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	// Functions
	registerHotKey = libuser32.NewProc("RegisterHotKey")
	unregisterHotKey = libuser32.NewProc("UnregisterHotKey")
	postThreadMessage = libuser32.NewProc("PostThreadMessageW")

	getCurrentThread = libkernel32.NewProc("GetCurrentThread")
	getThreadId = libkernel32.NewProc("GetThreadId")
}

func RegisterHotKey(hwnd HWND, id int32, fsModifiers, vk uint32) bool {
	ret, _, _ := syscall.Syscall6(registerHotKey.Addr(), 4,
		uintptr(hwnd),
		uintptr(id),
		uintptr(fsModifiers),
		uintptr(vk),
		0, 0)

	return ret != 0
}

func PostThreadMessage(idThread uint32, msg uint32, wParam, lParam int32) bool {
	ret, _, _ := syscall.Syscall6(postThreadMessage.Addr(), 4,
		uintptr(idThread),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0, 0)
	return ret != 0
}

func UnregisterHotKey(hwnd HWND, id int32) bool {
	ret, _, _ := syscall.Syscall(unregisterHotKey.Addr(), 2,
		uintptr(hwnd),
		uintptr(id),
		0)

	return ret != 0
}

func GetCurrentThread() HANDLE {
	ret, _, _ := syscall.Syscall(getCurrentThread.Addr(), 0, 0, 0, 0)
	return HANDLE(ret)
}

func GetThreadId(thread HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(getThreadId.Addr(), 1, uintptr(thread), 0, 0)
	return uint32(ret)
}
