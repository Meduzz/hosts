# hosts
Manipulate hosts files.

# Get the software

Either 

## Clone and run

> go build hosts.go

or to cross compile run:

> ./build.sh

build.sh will generate binaries for:

* MacOS 64-bit
* Linux 64-bit
* Linux arm (Raspberry Pi) untested :(

Tweak the build script to roll your own.

## Or download

* MacOS x86_64 [Here](https://storage.googleapis.com/nixutils/hosts/hosts_osx_x86_64)
* Linux x86_64 [Here](https://storage.googleapis.com/nixutils/hosts/hosts_linux_x86_64)
* Linux Pi [Here](https://storage.googleapis.com/nixutils/hosts/hosts_linux_pi)

# Usage

> hosts + domain

Will add/change domain in your /etc/hosts with ip 127.0.0.1.

> hosts + domain 127.0.0.2

Will add/change domain to 127.0.0.2 in your /etc/hosts.

> hosts + domain 127.0.0.3 ~/etc/hosts

Will add/change domain to 127.0.0.3 in your
~/etc/hosts.

> hosts - domain

Will remove all rows with a "domain" matching domain from your /etc/hosts.

> hosts - domain ~/etc/hosts

Will remove all rows with a "domain" matching domain from your ~/etc/hosts.

If the user executing are lacking permissions, there shall be errors.