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

func Test1(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Totp("GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ", 59, 30, 0))
	if binary_value != 1094287082 {
		t.Error("Expected bincode 1094287082, got", binary_value)
	}
}

func Test2(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp("GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ", 59, 30, 0)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != "287082" {
		t.Error("Expected OTP 287082, got", final_value)
	}
}

// ---------

func Test3(t *testing.T) {
	binary_value := goathgen.ToBinary(goathgen.Totp("3132333435363738393031323334353637383930", 1111111109, 30, 0))
	if binary_value != 907081804 {
		t.Error("Expected bincode 907081804, got", binary_value)
	}
}

func Test4(t *testing.T) {
	truncated_value := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp("3132333435363738393031323334353637383930", 1111111109, 30, 0)))
	final_value := fmt.Sprintf("%06d", truncated_value)
	if final_value != "081804" {
		t.Error("Expected OTP 081804, got", final_value)
	}
}
