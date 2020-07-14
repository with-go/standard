// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

/*
Arrays are slices of interface{} which has functions to perform traversal
and mutation operations. Length of an array is fixed, but the types of
its elements are not because the use of interface{} type.

Arrays cannot use strings as element indexes but must use integers. Setting
or accessing via non-integers using bracket notation will not set or retrieve
an element from the array list itself.
*/
package array
