# Nyx-Unhooker

# Nyx Unhooker Offers:
- Full Dll Unhooking for:
- kernel32.dll
- kernelbase.dll
- ntdll.dll
- user32.dll
- apphelp.dll
- msvcrt.dll

# A bit info what are you going to use

Endpoint Detection and Response (EDR) tools utilize hooking to monitor critical system functions within DLLs of running processes. Hooking involves dynamically modifying the initial instructions of DLL functions, redirecting program execution flow to code injected by the EDR. This allows the EDR to intercept and analyze program behavior for legitimacy or potential threats based on function arguments.

To evade detection by EDRs, malware employs "unhooking," a process that restores the original code section (.text) of DLLs. Malware can achieve this by acquiring unmodified DLLs through various means:

1. Direct System Access: Obtaining DLLs directly from the system, potentially raising suspicion due to open handles.

2. Remote File Access: Downloading a DLL remotely that matches the OS version of the target system, requiring the malware author to host such a file.

3. Suspended Process Access: Extracting DLL content from a suspended process before it gets hooked.

While NTDLL.dll is commonly targeted due to its proximity to the kernel, EDRs may also hook APIs in higher-level DLLs like kernel32.dll and user32.dll to monitor broader system interactions.

This technique illustrates the ongoing cat-and-mouse game between cybersecurity defenders utilizing EDRs and malware creators employing evasion tactics like unhooking to avoid detection and maintain persistence on compromised systems.
