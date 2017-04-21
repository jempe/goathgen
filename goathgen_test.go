/*
 * Copyright (c) 2017 Richard B Tilley <brad@w8rbt.org>
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

package goathgen_test

import "fmt"
import "testing"
import "github.com/w8rbt/goathgen"

var b32_secret string = "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
var bincode int = 0
var counter int64 = 0
var hex_secret string = "3132333435363738393031323334353637383930"
var hotp string = ""
var seconds int64 = 0
var time_step int64 = 30
var unix_epoch int64 = 0
var totp string = ""

/*
 * TestXa and TestXb test the Base32 encoded secret
 * TestXc and TestXd test the Hex encoded secret
 *
 * Tests 1 through 6 are from Appendix B of RFC 6238
 *     https://tools.ietf.org/html/rfc6238
 *
 * Tests 7 through 27 are from Appendix C of RFC 4226
 *     https://www.ietf.org/rfc/rfc4226.txt
 */

// Test 1

func Test1a(t *testing.T) {
	bincode = 1094287082
	seconds = 59
	totp = "287082"
	binary_value := goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test1b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

func Test1c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test1d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

// Test 2

func Test2a(t *testing.T) {
	bincode = 907081804
	seconds = 1111111109
	totp = "081804"
	binary_value := goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test2b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

func Test2c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test2d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

// Test 3

func Test3a(t *testing.T) {
	bincode = 414050471
	seconds = 1111111111
	totp = "050471"
	binary_value := goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test3b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

func Test3c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test3d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

// Test 4

func Test4a(t *testing.T) {
	bincode = 689005924
	seconds = 1234567890
	totp = "005924"
	binary_value := goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test4b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

func Test4c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test4d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

// Test 5

func Test5a(t *testing.T) {
	bincode = 2069279037
	seconds = 2000000000
	totp = "279037"
	binary_value := goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test5b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

func Test5c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test5d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

// Test 6

func Test6a(t *testing.T) {
	bincode = 1465353130
	seconds = 20000000000
	totp = "353130"
	binary_value := goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test6b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(b32_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

func Test6c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test6d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(hex_secret, seconds, time_step, unix_epoch)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != totp {
		t.Error("Expected totp", totp, "got", final_value)
	}
}

// Test 7

func Test7a(t *testing.T) {
	bincode = 1284755224
	counter = 0
	hotp = "755224"
	binary_value := goathgen.ToBinary(goathgen.Hotp(b32_secret, counter))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test7b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Hotp(b32_secret, counter)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != hotp {
		t.Error("Expected hotp", hotp, "got", final_value)
	}
}

func Test7c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Hotp(hex_secret, counter))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test7d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Hotp(hex_secret, counter)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != hotp {
		t.Error("Expected hotp", hotp, "got", final_value)
	}
}

// Test 8

func Test8a(t *testing.T) {
	bincode = 1094287082
	counter = 1
	hotp = "287082"
	binary_value := goathgen.ToBinary(goathgen.Hotp(b32_secret, counter))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test8b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Hotp(b32_secret, counter)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != hotp {
		t.Error("Expected hotp", hotp, "got", final_value)
	}
}

func Test8c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Hotp(hex_secret, counter))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test8d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Hotp(hex_secret, counter)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != hotp {
		t.Error("Expected hotp", hotp, "got", final_value)
	}
}

// Test 9

func Test9a(t *testing.T) {
	bincode = 137359152
	counter = 2
	hotp = "359152"
	binary_value := goathgen.ToBinary(goathgen.Hotp(b32_secret, counter))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test9b(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Hotp(b32_secret, counter)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != hotp {
		t.Error("Expected hotp", hotp, "got", final_value)
	}
}

func Test9c(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Hotp(hex_secret, counter))
	if binary_value != bincode {
		t.Error("Expected bincode", bincode, "got", binary_value)
	}
}

func Test9d(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Hotp(hex_secret, counter)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != hotp {
		t.Error("Expected hotp", hotp, "got", final_value)
	}
}
