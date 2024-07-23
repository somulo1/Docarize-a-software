# Security Policy

## Supported Versions

This ascii-art-web software is compatible with the following operating system versions:

| Version      | Supported                     |
| ------------ | ----------------------------- |
| 21.3         | :linux:                       |
| 22.04.4      | :ubuntu:                      |
| 6.1          | :Parrot OS                    |
| 9.2.23       | :NetBSD                       |
| 7.1.23       | :OpenBSD                      |
| r151036.23   | :OmniOS                       |
| 13.1.23      | :FreeBSD                      |
| 7.2.23       | :DragonFly BSD                |
| 12.21        | :macOS (formerly OS X)        |
| 13           | :Android                      |
| AIX 7.2 TL5  | :IBM AIX                      |
| 11.23        | :Microsoft Windows            |

## Reporting a Vulnerability

Older versions of Linux, Ubuntu, and Parrot OS can also use this software version provided they address the newline handling correctly. Newline handling varies across different machines. When handling Thinkertoy banner files, ensure that both `\n` (LF) and `\r\n` (CRLF) are handled seamlessly, depending on the platform. Windows machines use `\r\n` for line endings.
The listed versions above are officially supported. Newer versions may also be compatible but should adhere to the same newline handling standards discussed.


