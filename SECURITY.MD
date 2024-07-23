# Security Policy

## Supported Versions

This ascii-art-web software can work best with the listed operating systems upto the most recent versions lised bellow;

| Version      | Supported                     |
| ------------ | ----------------------------- |
| 21.3         | :linux:                       |
| 22.04.4      | :ubuntu:                      |
| 6.1          | :ParotOS                      |
| 9.2.23       | :NetBSD                       |
| 7.1.23       | :OpenBSD                      |
| r151036.23   | :OmniOS                       |
| 13.1 .23     | :FreeBSD                      |
| 7.2.23       | :DragonFly BSD                |
| 12.21        | :macOS (formerly OS X)        |
| 13           | :Android                      |
| AIX 7.2 TL5  | :IBM AIX                      |
| 11.23        | :Microsoft Windows            |

## Reporting a Vulnerability

Other older versions of Linux, ubuntu and ParotOS can also be use this code version provided the concept of `\n\r` are well addresed. the concept of new line differs in diferent machines. <br> When handling the thinkertoy banner files you must therefore consider `\n` (LF) and `\r\n` (CRLF) handled seamlessly, depending on the platform windows machine uses.for this reason therefore, while the linux versions recorcorgnises "\n" the code bellow can be used to solve this vulnerability;

``` bash

func checkOS() []string {
	var sep []string
	os := runtime.GOOS // Identify operating system

	// Set separators according to operating system
	if os == "windows" {
		sep = []string{"\n", "\r\n"}
	} else {
		sep = []string{"\r\n", "\n"}
	}

	return sep
} 
}

 ```
 The versions addressed above stands to be the versions addressed. other newer versions can also therefore qualify within the above addressed issue.
