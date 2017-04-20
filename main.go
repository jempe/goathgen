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

package goathgen

import "bufio"
import "fmt"
import "os"
import "time"

import "github.com/w8rbt/goathgen"

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

	if goathgen.Debug {
		fmt.Println("----- main -----")
		fmt.Printf("stdin %s\n", stdin.Text())
		fmt.Printf("unix_now: %d\n", unix_now)
		fmt.Printf("unix_epoch: %d\n", unix_epoch)
		fmt.Printf("time_step: %d\n", time_step)
		fmt.Println("----- main -----")
	}

	the_totp := goathgen.Truncate(goathgen.ToBinary(goathgen.Totp(stdin.Text(), unix_now, time_step, unix_epoch)))
	fmt.Printf("%06d\n", the_totp)
}
