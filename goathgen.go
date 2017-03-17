/*
 * Copyright (c) 2016 Richard B Tilley <brad@w8rbt.org>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import "bufio"
import "bytes"
import "crypto/sha1"
import "crypto/hmac"
import "encoding/base32"
import "encoding/binary"
import "encoding/hex"
import "fmt"
import "log"
import "os"
import "strings"
import "time"

var debug bool = false

func hotp(secret string, counter int64) []byte {
	// The secret must be hex or base32 encoded!

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, counter)
	if err != nil {
		log.Fatal(err)
	}

	// ToUpper needed here for base32
	encoded := []byte(strings.ToUpper(secret))
	decoded := make([]byte, 64)

	var num_bytes int

	num_bytes, err = hex.Decode(decoded, encoded)
	if err != nil {
		num_bytes, err = base32.StdEncoding.Decode(decoded, encoded)
		if err != nil {
			log.Fatal("The secret must be hex or base32 encoded!")
		}
	}

	mac_160 := hmac.New(sha1.New, decoded)
	mac_160.Write(buf.Bytes())

	if debug {
		fmt.Println("----- hotp -----")
		fmt.Printf("counter: %d\n", counter)
		fmt.Printf("encoded secret: %s\n", secret)
		fmt.Printf("encoded array: %v\n", encoded)
		fmt.Printf("encoded len: %v\n", len(encoded))
		fmt.Printf("decoded array: %v\n", decoded)
		fmt.Printf("decoded len: %v\n", len(decoded))
		fmt.Printf("decoded secret: %s\n", decoded[:num_bytes])
		fmt.Printf("mac_160: %v\n", mac_160.Sum(nil))
		fmt.Printf("mac_160 len: %v\n", mac_160.Size())
		fmt.Println("----- hotp -----")
	}

	return mac_160.Sum(nil)
}

func totp(secret string, time_now int64, time_step int64, unix_epoch int64) []byte {
	// The secret must be hex or base32 encoded!

	var time_counter int64
	time_counter = (time_now - unix_epoch) / time_step

	if debug {
		fmt.Println("----- totp -----")
		fmt.Printf("time_now: %d\n", time_now)
		fmt.Printf("time_step: %d\n", time_step)
		fmt.Printf("unix_epoch: %d\n", unix_epoch)
		fmt.Printf("time_counter: %d\n", time_counter)
		fmt.Println("----- totp -----")
	}

	return hotp(secret, time_counter)
}

func to_binary(hmac []byte) int {
	offset := hmac[len(hmac)-1] & 0xf

	bincode := int(((int(hmac[offset]) & 0x7f) << 24) |
		((int(hmac[offset+1]) & 0xff) << 16) |
		((int(hmac[offset+2]) & 0xff) << 8) |
		((int(hmac[offset+3]) & 0xff) << 0))

	if debug {
		fmt.Println("----- to_binary -----")
		fmt.Printf("len: %d\n", len(hmac))
		fmt.Printf("offset: %d\n", offset)
		fmt.Printf("bincode: %d\n", bincode)
		fmt.Println("----- to_binary -----")
	}

	return bincode
}

func truncate(bincode int) int {
	modulo_by := 1000000

	otp := bincode % modulo_by

	if debug {
		fmt.Println("--- truncate ---")
		fmt.Printf("bincode: %d\n", bincode)
		fmt.Printf("modulo by: %d\n", modulo_by)
		fmt.Printf("otp: %d\n", otp)
		fmt.Println("--- truncate ---")
	}

	return otp
}

func main() {

	// Secret must be sent by stdin
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	now := time.Now()
	unix_now := now.Unix()

	var unix_epoch int64
	unix_epoch = 0

	var time_step int64
	time_step = 30

	if debug {
		fmt.Println("----- main -----")
		fmt.Printf("stdin %s\n", stdin.Text())
		fmt.Printf("unix_now: %d\n", unix_now)
		fmt.Printf("unix_epoch: %d\n", unix_epoch)
		fmt.Printf("time_step: %d\n", time_step)
		fmt.Println("----- main -----")
	}

	the_totp := truncate(to_binary(totp(stdin.Text(), unix_now, time_step, unix_epoch)))
	fmt.Printf("%06d\n", the_totp)
}
