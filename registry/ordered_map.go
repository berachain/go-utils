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

package registry

import (
	libtypes "github.com/berachain/go-utils/types"

	"github.com/elliotchance/orderedmap/v2"
)

// orderedMapRegistry is a simple implementation of `Registry` that uses a map as the underlying
// data structure.
type orderedMapRegistry[K comparable, T libtypes.Registrable[K]] struct {
	// items is the map of items in the registry.
	items *orderedmap.OrderedMap[K, T]
}

// NewOrderedMap creates and returns a new `orderedMapRegistry`. Maintains order of registration
// (insertion).
//
//nolint:revive // only used as Registry interface.
func NewOrderedMap[K comparable, T libtypes.Registrable[K]]() *orderedMapRegistry[K, T] {
	return &orderedMapRegistry[K, T]{
		items: orderedmap.NewOrderedMap[K, T](),
	}
}

// Get returns an item using its ID.
func (mr *orderedMapRegistry[K, T]) Get(id K) T {
	val, _ := mr.items.Get(id)
	return val
}

// Register adds an item to the registry.
func (mr *orderedMapRegistry[K, T]) Register(item T) error {
	_ = mr.items.Set(item.RegistryKey(), item)
	return nil
}

// Remove removes an item from the registry.
func (mr *orderedMapRegistry[K, T]) Remove(id K) {
	_ = mr.items.Delete(id)
}

// Has returns true if the item exists in the registry.
func (mr *orderedMapRegistry[K, T]) Has(id K) bool {
	_, ok := mr.items.Get(id)
	return ok
}

// Iterate returns the underlying map (unordered).
func (mr *orderedMapRegistry[K, T]) Iterate() map[K]T {
	ret := make(map[K]T, mr.items.Len())
	for el := mr.items.Front(); el != nil; el = el.Next() {
		ret[el.Key] = el.Value
	}
	return ret
}

// IterateInOrder returns the underlying map (ordered).
func (mr *orderedMapRegistry[K, T]) IterateInOrder() (*orderedmap.OrderedMap[K, T], error) {
	return mr.items.Copy(), nil
}
