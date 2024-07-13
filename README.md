# Nyx Unhooker

Nyx Unhooker is a tool designed to facilitate DLL unhooking, focusing on critical system DLLs to allow evasion of Endpoint Detection and Response (EDR) mechanisms. By restoring original DLL code sections, Nyx Unhooker aims to counteract monitoring and detection by security tools, offering a method for software to operate stealthily on compromised systems.

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

## Usage

To use Nyx Unhooker effectively, follow these steps:

1. **Compile**: Compile the Nyx Unhooker code using your preferred compiler.
2. **Run**: Execute Nyx Unhooker with administrative privileges to perform DLL unhooking.
3. **Monitor**: Verify the effectiveness of DLL unhooking using appropriate testing and monitoring tools.

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
