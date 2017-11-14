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

> hosts + dev.local

Will add/change dev.local in your /etc/hosts with ip 127.0.0.1.

> hosts + test.local 127.0.0.2

Will add/change test.local to 127.0.0.2 in your /etc/hosts.

> hosts + db.local 127.0.0.3 ~/etc/hosts

Will add/change db.local to 127.0.0.3 in your
~/etc/hosts.

> hosts - db.local

Will remove all rows with a "domain" matching db.local from your /etc/hosts.

> hosts - dev.local ~/etc/hosts

Will remove all rows with a "domain" matching dev.local from your ~/etc/hosts.

If the user executing are lacking permissions, there shall be errors (and they shall be pretty...ish).