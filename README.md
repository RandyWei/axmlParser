# Android binary manifest XML file parser library for Golang

This is a golang port of [github.com/xgouchet/AXML](http://github.com/xgouchet/AXML).

This package aims to parse an android binary mainfest xml file on an apk file.

## Installation

```
go get github.com/RandyWei/axmlParser
```

## Usage

```Go
package main

import (
	"fmt"

	"github.com/RandyWei/axmlParser"
)

func main() {
	listener := new(axmlParser.AppNameListener)
	axmlParser.ParseApk("./myApp.apk", listener)

	fmt.Printf("ActivityName: %v\n", listener.ActivityName)
	fmt.Printf("VersionCode: %v\n", listener.VersionCode)
	fmt.Printf("VersionName: %v\n", listener.VersionName)
}
```
