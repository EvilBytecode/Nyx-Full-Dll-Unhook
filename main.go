package main

import (
	"Nyx_Unhooker/pkg/Nyx"
)

func main() {
  Nyx.AutoUnhook()   // Unhooks based on users Arch, (Automatically Gets it) and patches Dlls.
 // Nyx.Unhook64() // Unhooks on Amd64 Archs
 // Nyx.Unhook386() // Unhooks on I386 Archs
}
