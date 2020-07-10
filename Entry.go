package main

import (
	"errors"
	"fmt"
	"os/exec"
	"syscall"
	"time"
	"unsafe"
)

type CatchHandler func(error)

func main() {
	Throw := func(format string, a ...interface{}) error {
		if a == nil {
			panic(errors.New(format))
		}
		panic(errors.New(fmt.Sprintf(format, a)))
	}
	Try := func(fun func(), handler CatchHandler) {
		defer func() {
			if e := recover(); e != nil {
				var err error = e.(error)
				handler(err)
			}
		}()
		fun()
	}
	callW32Api := func(dllName string, procName string, params ...uintptr) uintptr {
		var result uintptr
		Try(func() {
			dlllib, err := syscall.LoadLibrary(dllName)
			if err != nil {
				Throw(fmt.Sprintf(`faild to LoadLibrary("%s")`, dllName))
			}
			defer syscall.FreeLibrary(dlllib)

			fnProc, err := syscall.GetProcAddress(dlllib, procName)
			if err != nil {
				Throw(fmt.Sprintf(`faild to GetProcAddress["%s"]`, procName))
			}
			if len(params) > 6 {
				Throw(fmt.Sprintf(` too many arguments in call ["%s"]`, procName))
			}
			switch len(params) {
			case 0:
				result, _, _ = syscall.Syscall6(uintptr(fnProc), 0, 0, 0, 0, 0, 0, 0)
				break
			case 1:
				result, _, _ = syscall.Syscall6(uintptr(fnProc), 1, params[0], 0, 0, 0, 0, 0)
				break
			case 2:
				result, _, _ = syscall.Syscall6(uintptr(fnProc), 2, params[0], params[1], 0, 0, 0, 0)
				break
			case 3:
				result, _, _ = syscall.Syscall6(uintptr(fnProc), 3, params[0], params[1], params[2], 0, 0, 0)
				break
			case 4:
				result, _, _ = syscall.Syscall6(uintptr(fnProc), 4, params[0], params[1], params[2], params[3], 0, 0)
				break
			case 5:
				result, _, _ = syscall.Syscall6(uintptr(fnProc), 5, params[0], params[1], params[2], params[3], params[4], 0)
				break
			case 6:
				result, _, _ = syscall.Syscall6(uintptr(fnProc), 6, params[0], params[1], params[2], params[3], params[4], params[5])
				break
			}
			return
		}, func(err error) {
			fmt.Println(err.Error())
		})
		return result
	}
	strptr := func(text string) uintptr {
		p1, _ := syscall.UTF16PtrFromString(text)
		return uintptr(unsafe.Pointer(p1))
	}

	go func() {
		exec.Command(`C:\WINDOWS\system32\notepad.exe`).Run()
	}()

	hwnd := callW32Api("user32.dll", "FindWindowExW", 0, 0, strptr("Notepad"), 0)
	const SW_SHOW = 5
	const SW_SHOWMAXIMIZED = 3
	callW32Api("user32.dll", "ShowWindow", hwnd, SW_SHOWMAXIMIZED, 0, 0)
	callW32Api("user32.dll", "SetForegroundWindow", hwnd, 0, 0, 0)
	const KEYEVENTF_KEYUP = 0x2
	const KEYEVENTF_EXTENDEDKEY = 0x1
	const WM_KEYDOWN = 0X100
	const WM_KEYUP = 0x101
	const VK_CONTROL = 0x11
	const VK_SHIFT = 0x10
	const VK_DELETE = 0x2E
	const VK_MENU = 0x12
	const WM_SYSKEYDOWN = 0x104
	const WM_SYSKEYUP = 0x105
	const WM_SYSCHAR = 0x106
	time.Sleep(time.Millisecond * 100)
	callW32Api("user32.dll", "keybd_event", 'H', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", 'H', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", 'I', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", 'I', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", ' ', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", ' ', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", 'F', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", 'F', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", 'R', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", 'R', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", 'O', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", 'O', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", 'M', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", 'M', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", ' ', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", ' ', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", 'G', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", 'G', 0x53, KEYEVENTF_KEYUP, 0)
	callW32Api("user32.dll", "keybd_event", 'O', 0x53, 0, 0)
	callW32Api("user32.dll", "keybd_event", 'O', 0x53, KEYEVENTF_KEYUP, 0)
	const MB_OKCANCEL = 0x1
	callW32Api("user32.dll", "MessageBoxW", 0, strptr("have fun"), strptr("title"), MB_OKCANCEL)
}
