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

package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	i int
}

func (s *TestStruct) Fun1() {}

type TestInterface interface {
	Fun1()
}

type TestAlias []TestInterface

func TestInPlaceAppend(t *testing.T) {
	bytes := []byte{}
	InPlaceAppend(&bytes, 'x')
	require.Equal(t, "x", string(bytes))
	InPlaceAppend(&bytes, 'y', 'z')
	require.Equal(t, "xyz", string(bytes))
	InPlaceAppend(&bytes, []byte("abc")...)
	require.Equal(t, "xyzabc", string(bytes))

	strings := []string{}
	InPlaceAppend(&strings, "x")
	require.Equal(t, []string{"x"}, strings)
	InPlaceAppend(&strings, "y", "z")
	require.Equal(t, []string{"x", "y", "z"}, strings)
	InPlaceAppend(&strings, []string{"a", "b", "c"}...)
	require.Equal(t, []string{"x", "y", "z", "a", "b", "c"}, strings)

	structs := []TestStruct{}
	InPlaceAppend(&structs, TestStruct{1})
	require.Equal(t, []TestStruct{{1}}, structs)
	InPlaceAppend(&structs, TestStruct{2}, TestStruct{3})
	require.Equal(t, []TestStruct{{1}, {2}, {3}}, structs)
	InPlaceAppend(&structs, []TestStruct{{4}, {5}}...)
	require.Equal(t, []TestStruct{{1}, {2}, {3}, {4}, {5}}, structs)

	structPtrs := []*TestStruct{}
	InPlaceAppend(&structPtrs, &TestStruct{1})
	require.Equal(t, []*TestStruct{{1}}, structPtrs)
	InPlaceAppend(&structPtrs, &TestStruct{2}, &TestStruct{3})
	require.Equal(t, []*TestStruct{{1}, {2}, {3}}, structPtrs)
	InPlaceAppend(&structPtrs, []*TestStruct{{4}, {5}}...)
	require.Equal(t, []*TestStruct{{1}, {2}, {3}, {4}, {5}}, structPtrs)

	alias := TestAlias([]TestInterface{})
	InPlaceAppend[TestAlias, TestInterface](&alias, &TestStruct{1})
	require.Equal(t, TestAlias([]TestInterface{&TestStruct{1}}), alias)
}

func TestImmutableAppend(t *testing.T) {
	// no capacity
	prefix := []byte("abc")
	require.Equal(t, 3, cap(prefix))
	require.Equal(t, "abc", string(ImmutableAppend(prefix)))
	require.Equal(t, "abcd", string(ImmutableAppend(prefix, 'd')))
	require.Equal(t, "abce", string(ImmutableAppend(prefix, 'e')))
	require.Equal(t, "abc", string(prefix))

	// has capacity
	prefix = []byte("ab")
	prefix = append(prefix, 'c')
	require.Greater(t, cap(prefix), 3)
	require.Equal(t, "abc", string(ImmutableAppend(prefix)))
	require.Equal(t, "abcd", string(ImmutableAppend(prefix, 'd')))
	require.Equal(t, "abce", string(ImmutableAppend(prefix, 'e')))
	require.Equal(t, "abc", string(prefix))

	// no capacity
	prefix2 := []string{"a", "b", "c"}
	require.Equal(t, 3, cap(prefix2))
	require.Equal(t, []string{"a", "b", "c"}, ImmutableAppend(prefix2))
	require.Equal(t, []string{"a", "b", "c", "d"}, ImmutableAppend(prefix2, "d"))
	require.Equal(t, []string{"a", "b", "c", "e"}, ImmutableAppend(prefix2, "e"))
	require.Equal(t, []string{"a", "b", "c"}, prefix2)

	// has capacity
	prefix2 = []string{"a", "b"}
	prefix2 = append(prefix2, "c")
	require.Greater(t, cap(prefix2), 3)
	require.Equal(t, []string{"a", "b", "c"}, ImmutableAppend(prefix2))
	require.Equal(t, []string{"a", "b", "c", "d"}, ImmutableAppend(prefix2, "d"))
	require.Equal(t, []string{"a", "b", "c", "e"}, ImmutableAppend(prefix2, "e"))
	require.Equal(t, []string{"a", "b", "c"}, prefix2)

	// no capacity
	prefix3 := []TestStruct{{1}, {2}, {3}}
	require.Equal(t, 3, cap(prefix3))
	require.Equal(t, []TestStruct{{1}, {2}, {3}}, ImmutableAppend(prefix3))
	require.Equal(t, []TestStruct{{1}, {2}, {3}, {4}}, ImmutableAppend(prefix3, TestStruct{4}))
	require.Equal(t, []TestStruct{{1}, {2}, {3}, {5}}, ImmutableAppend(prefix3, TestStruct{5}))
	require.Equal(t, []TestStruct{{1}, {2}, {3}}, prefix3)

	// has capacity
	prefix3 = []TestStruct{{1}, {2}}
	prefix3 = append(prefix3, TestStruct{3})
	require.Greater(t, cap(prefix3), 3)
	require.Equal(t, []TestStruct{{1}, {2}, {3}}, ImmutableAppend(prefix3))
	require.Equal(t, []TestStruct{{1}, {2}, {3}, {4}}, ImmutableAppend(prefix3, TestStruct{4}))
	require.Equal(t, []TestStruct{{1}, {2}, {3}, {5}}, ImmutableAppend(prefix3, TestStruct{5}))
	require.Equal(t, []TestStruct{{1}, {2}, {3}}, prefix3)

	// no capacity
	prefix4 := []*TestStruct{{1}, {2}, {3}}
	require.Equal(t, 3, cap(prefix4))
	require.Equal(t, []*TestStruct{{1}, {2}, {3}}, ImmutableAppend(prefix4))
	require.Equal(t, []*TestStruct{{1}, {2}, {3}, {4}}, ImmutableAppend(prefix4, &TestStruct{4}))
	require.Equal(t, []*TestStruct{{1}, {2}, {3}, {5}}, ImmutableAppend(prefix4, &TestStruct{5}))
	require.Equal(t, []*TestStruct{{1}, {2}, {3}}, prefix4)

	// has capacity
	prefix4 = []*TestStruct{{1}, {2}}
	prefix4 = append(prefix4, &TestStruct{3})
	require.Greater(t, cap(prefix4), 3)
	require.Equal(t, []*TestStruct{{1}, {2}, {3}}, ImmutableAppend(prefix4))
	require.Equal(t, []*TestStruct{{1}, {2}, {3}, {4}}, ImmutableAppend(prefix4, &TestStruct{4}))
	require.Equal(t, []*TestStruct{{1}, {2}, {3}, {5}}, ImmutableAppend(prefix4, &TestStruct{5}))
	require.Equal(t, []*TestStruct{{1}, {2}, {3}}, prefix4)

	// no capacity
	prefix5 := TestAlias([]TestInterface{})
	require.Equal(t, TestAlias([]TestInterface{&TestStruct{1}}), ImmutableAppend[TestAlias, TestInterface](prefix5, &TestStruct{1}))
	require.Equal(t, TestAlias([]TestInterface{&TestStruct{2}}), ImmutableAppend[TestAlias, TestInterface](prefix5, &TestStruct{2}))

	// has capacity
	prefix5 = TestAlias([]TestInterface{})
	prefix5 = append(prefix5, &TestStruct{0})
	require.Equal(t, TestAlias([]TestInterface{&TestStruct{0}, &TestStruct{1}}), ImmutableAppend[TestAlias, TestInterface](prefix5, &TestStruct{1}))
	require.Equal(t, TestAlias([]TestInterface{&TestStruct{0}, &TestStruct{2}}), ImmutableAppend[TestAlias, TestInterface](prefix5, &TestStruct{2}))
}