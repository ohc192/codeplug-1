// Copyright 2017 Dale Farnsworth. All rights reserved.

// Dale Farnsworth
// 1007 W Mendoza Ave
// Mesa, AZ  85210
// USA
//
// dale@farnsworth.org

// This file is part of Radio.
//
// Radio is free software: you can redistribute it and/or modify
// it under the terms of version 3 of the GNU Lesser General Public
// License as published by the Free Software Foundation.
//
// Radio is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Radio.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
)

func dprint(v ...interface{}) {
	skip := 1

	_, filename, line, ok := runtime.Caller(skip)

	str := ""
	if ok {
		dir := filepath.Base(filepath.Dir(filename))
		filename = filepath.Base(filename)
		filename = filepath.Join(dir, filename)
		str = fmt.Sprintf("%s:%d", filename, line)
	}

	v = append([]interface{}{str}, v...)

	fmt.Fprintln(os.Stderr, v...)
}

func printStack() {
	fmt.Fprintln(os.Stderr, "start stack trace")
	debug.PrintStack()
	fmt.Fprintln(os.Stderr)
}
