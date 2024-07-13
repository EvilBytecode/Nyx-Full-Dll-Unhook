package main

import (
	"github.com/EvilBytecode/Nyx-Full-Dll-Unhook"
)

func main() {
  Nyx.AutoUnhook()   // Unhooks based on users Arch, (Automatically Gets it) and patches Dlls.
 // Nyx.Unhook64() // Unhooks on Amd64 Archs
 // Nyx.Unhook386() // Unhooks on I386 Archs
}
