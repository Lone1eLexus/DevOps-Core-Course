# Lab 18 — Reproducible Builds with Nix

## Task 1 — Build Reproducible Python App

```bash
$ curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install
info: downloading the Determinate Nix Installer
 INFO nix-installer v3.20.0
`nix-installer` needs to run as `root`, attempting to escalate now via `sudo`...
[sudo] password for lord: 
 INFO nix-installer v3.20.0
Nix install plan (v3.20.0)
Planner: linux (with default settings)

Planned actions:
* Create directory `/nix`
* Install Determinate Nixd
* Extract the bundled Nix (originally from /nix/store/3h8wrgwk0rczd8rjv4fvl344mp1w765m-nix-binary-tarball-3.20.0/nix-3.20.0-x86_64-linux.tar.xz) to `/nix/temp-install-dir`
* Create a directory tree in `/nix`
* Synchronize /nix and /nix/var ownership
* Move the downloaded Nix into `/nix`
* Synchronize /nix/store ownership
* Create build users (UID 30001-30032) and group (GID 30000)
* Setup the default Nix profile
* Place the Nix configuration in `/etc/nix/nix.conf`
* Configure the shell profiles
* Configure the Determinate Nix daemon
* Cleanup


Proceed? ([Y]es/[n]o/[e]xplain): Y
 INFO Step: Create directory `/nix`
 INFO Step: Install Determinate Nixd
 INFO Step: Provision Nix
 INFO Step: Create build users (UID 30001-30032) and group (GID 30000)
 INFO Step: Configure Nix
 INFO Step: Create directory `/etc/tmpfiles.d`
 INFO Step: Configure the Determinate Nix daemon
 INFO Step: Cleanup
 INFO Running self test for shell sh
 INFO Running self test for shell bash
Nix was installed successfully!
To get started using Nix, open a new shell or run `. /nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh`
$ nix --version
nix (Determinate Nix 3.20.0) 2.34.6
$ nix run nixpkgs#hello
Hello, world!
```

I decided to take app from Lab 2.



