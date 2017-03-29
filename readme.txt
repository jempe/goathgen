Goathgen - A simple TOTP implementation written in Go.

# How to use goathgen with your encrypted TOTP secret

    1. Obtain a hex or base32 encoded TOTP secret from your service provider.
    2. Encrypt the TOTP secret with your PGP key and save the encrypted secret in a file.
    3. To generate a one-time code, decrypt the TOTP secret and pipe the output to goathgen.

# Example usage

    gpg -d encrypted-secret.txt | goathgen


# Installation via Go

    go get github.com/w8rbt/goathgen

# Installation via Source

    go build goathgen.go
    sudo mv goathgen /usr/local/bin

# Notes

    * goathgen only works in TOTP mode using the TOTP defaults

        * HMAC-SHA1
        * Six digit codes
        * 30 second time step

    * If you need other OATH functionality, edit the goathgen source code.

        * Read C++ oathgen source code for full OATH functionality
        * https://github.com/w8rbt/oathgen

    * TOTP requires precise system time so run NTP when using goathgen.

