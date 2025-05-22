# Nyx Unhooker

Nyx Unhooker is a tool designed to facilitate DLL unhooking, focusing on critical system DLLs to allow evasion of Endpoint Detection and Response (EDR) mechanisms. By restoring original DLL code sections, Nyx Unhooker aims to counteract monitoring and detection by security tools, offering a method for software to operate stealthily on compromised systems.
![image](https://github.com/user-attachments/assets/4d1731b5-501e-4c2c-93d6-44030aaafaa3)

## Features
- Simple to use.
- **DLL Unhooking**: Enables the restoration of original code sections for the following DLLs:
  - kernel32.dll
  - kernelbase.dll
  - ntdll.dll
  - user32.dll
  - apphelp.dll
  - msvcrt.dll

## Purpose

Endpoint Detection and Response (EDR) tools rely on hooking mechanisms to monitor and intercept system-level function calls within DLLs. This technique allows for real-time monitoring of program behavior, helping to identify and mitigate potential security threats. However, malware and certain software applications can utilize DLL unhooking to evade detection. Unhooking involves reverting the modifications made by EDR tools, thereby allowing malicious code to execute undetected.

## How It Works

Nyx Unhooker facilitates the unhooking process by restoring the original functionality of specified DLLs. This allows software utilizing Nyx Unhooker to operate without triggering alerts from EDR tools that rely on hooking mechanisms.

## Example Scenario

In a scenario where a software application needs to avoid detection by EDR tools, Nyx Unhooker can be used to restore the original DLL code sections, thus enabling the application to operate stealthily.

## Security Considerations

- **Legal and Ethical Use**: Ensure Nyx Unhooker is used in compliance with legal and ethical standards. Unauthorized use or deployment may violate terms of service or regulatory requirements.

## Contributing

Contributions to Nyx Unhooker are welcome. Please follow the established contribution guidelines and code of conduct.

## License

This project is licensed under the [MIT]. See the LICENSE file for details.

# credits?
- https://github.com/timwhitez
---
# Examples & Info:
### ```Nyx.AutoUnhook()```
- Automatically unhooks DLLs based on the user's system architecture. This function detects whether the system is running on amd64 or i386 architecture and patches the DLLs accordingly.

### ```Nyx.Unhook64()```
- Unhooks DLLs specifically for systems running on the amd64 architecture. This function is commented out in the main example but can be used if you know your system architecture in advance.

### ```Nyx.Unhook386()```
- Unhooks DLLs specifically for systems running on the i386 architecture. This function is commented out in the main example but can be used if you know your system architecture in advance.

### Example Snippet:
```go
package main

import (
	"github.com/EvilBytecode/Nyx-Full-Dll-Unhook/pkg/nyx"
)

func main() {
	Nyx.AutoUnhook()   // Unhooks based on users Arch, (Automatically Gets it) and patches Dlls.
	// Nyx.Unhook64() Unhooks on Amd64 Archs
	// Nyx.Unhook386()  Unhooks on I386 Archs
}
```

### PoC:
- (mcafee flags everything)
![image](https://github.com/user-attachments/assets/0cd89e02-1bfc-4000-8c6f-34e769f8c302)

## Info:
gets flagged (Runtime)
```go
package main

import (
    "fmt"
    //"github.com/EvilBytecode/Nyx-Full-Dll-Unhook/pkg/nyx"
)

func main() {
    //Nyx.Unhook64()
    fmt.Println("wowowo")
    fmt.Scanln()
}
```
```go
package main

import (
    "fmt"
    "github.com/EvilBytecode/Nyx-Full-Dll-Unhook/pkg/nyx"
)

func main() {
    Nyx.Unhook64()
    fmt.Println("wowowo")
    fmt.Scanln()
}
```
- credits to timwhitez i forgot 
