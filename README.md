# Shafran

Made for Secware <3

SharpHound 2.6.0
Running command Sharphound.exe -c All, so it should work with BloodHound 5.0.0 (CE version).

Encoding SharpHound:

    AES → Base64

Executed as shellcode using the MalDev Golang library.

You can download Shafran from the source or release.

Problem:

    Freezes after execution , but you will still get a .zip file for BloodHound.

# ShaRphic_patch

ShaRphic_patch.exe should be more stealthy against AV/EDR. - "maldev func of syscall doesn't work silently, so i try NtQueueApcThreadEx syscall for bypass ntdll.dll"

