# hello_rkt

This is a repository to document the first steps with
[coreos/rkt](https://github.com/coreos/rkt). It shall eventually evolve to being
a set of minimal working examples of what I and maybe others have been doing
with this new piece of technology.

## Background Information
Let's have an overview of which pieces build the puzzle.

### [ACI (App Container Image)](https://github.com/appc/spec/blob/master/SPEC.md#app-container-specification)
TODO

### Pods
TODO

### Signing
TODO


## *rkt* (and related tools) Usage
The docs around *rkt* seem to be very detailed about the implementation and
specifications. This section extracts the information needed to be simply use
*rkt* and tools like *actool*.

### Run an ACI

```
$ sudo rkt --insecure-skip-verify run hello_go.aci
```

## Examples Usage

### Requirements
The `rkt` binary needs to be in your $PATH, and the `actool` binary is assumed
to be at `/usr/share/rkt/bin/actool`.

If you are running Gentoo, you can use the [steveeJ
overlay](https://github.com/steveeJ/personal-portage-overlay) to install
*app-emulation/rkt* which should meet the requirements.

### hello\_go
Checkout this repository and change your terminal to its path.

To build and run the *hello\_go* example run these two simple commands.
```
$ ./build hello_go
$ sudo rkt --insecure-skip-verify run hello_go.aci
```

## ACI Overview

### hello\_go
Static compilation of small webserver application in Go.
