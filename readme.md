> This program was originally created by Richard Tilley and was available at github.com/w8rbt/goathgen. Since that repository is no longer available, all URLs have been updated to point to the jempe organization.

# goathgen

A simple TOTP library and command line executable written in Go.

## Executable usage examples

```bash
$ gpg -d encrypted-secret.txt | goathgen
$ cat plaintext-secret.txt | goathgen
```

## Installation via Go

```bash
$ go get github.com/jempe/goathgen
```

## To build and install the executable

```bash
$ cd main && go build -o goathgen
$ sudo mv goathgen /usr/local/bin
```

## To use the library in your go program
```bash
Read main.go to see a simple example
```

## To run unit tests against the library
```bash
$ go test -v
    or
$ go test
```

## Notes

1. To use the goathgen executable with a PGP encrypted TOTP secret
    * Obtain a hex or base32 encoded TOTP secret from your service provider.
    * Encrypt the TOTP secret with your PGP key and save the encrypted secret in a file.
    * Securely delete and overwrite the plain-text TOTP secret.
    * Decrypt the TOTP secret and pipe the output to goathgen.
2. goathgen only works in TOTP mode using the TOTP defaults
    * HMAC-SHA1
    * Six digit codes
    * 30 second time step
3. If you need other OATH functionality, edit the goathgen source code.
    * Read the C++ oathgen source code for full OATH functionality
    * https://github.com/jempe/oathgen
4. Run NTP when using goathgen.
