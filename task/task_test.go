// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package task

import (
	"testing"
	"time"
)

// TestTask tests the CancellablePeriodicTask function
// by creating a task that increments a variable every
// delayUs microseconds, and then cancelling the task
// after delayUs * 10 microseconds
// The variable should be incremented 10 times
func TestTask(t *testing.T) {
	dummyFunc := func(t *int) {
		*t += 1
	}

	v := 0
	delayUs := uint64(1000)
	_, cancel := CancellablePeriodicTask(delayUs, &v, dummyFunc)

	// Sleep for delayUs * 10
	time.Sleep(time.Duration(delayUs*10) * time.Microsecond)

	// Cancel the context
	cancel()

	// Check the value
	if !(9 <= v && v <= 11) {
		t.Errorf("Expected value in range [9, 11], got %v", v)
	}
}
