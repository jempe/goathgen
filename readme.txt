Goathgen - A simple TOTP implementation written in Go.

# How to use goathgen with your encrypted TOTP secret

    1. Obtain a hex encoded software TOTP secret.
    2. Save the TOTP secret in a plain text file.
    3. Encrypt the TOTP secret with your PGP key.
    4. Securely delete the plaintext TOTP secret.
    5. Pipe the decrypted TOTP secret to goathgen.

# Example CLI usage

    gpg -d secret.pgp | goathgen

# Notes

    * goathgen only works in TOTP mode using the TOTP defaults

        * HMAC-SHA1
        * Six digit codes
        * 30 second time step

    * goathgen only works with hex encoded secrets.

    * If you need other OATH functionality, edit the source.

        * Read C++ oathgen source code for full OATH functionality
        * https://github.com/w8rbt/oathgen

    * goathgen is statically linked and should run anywhere.

    * To build goathgen from source, run 'go build goathgen.go'

        * https://github.com/w8rbt/goathgen

