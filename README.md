gophp
====

[![Build Status](https://travis-ci.com/techoner/gophp.svg?branch=master)](https://travis-ci.com/techoner/gophp)


Golang implementation for PHP's functions


## Install / Update

```
go get -u github.com/techoner/gophp
```

## Example

```golang
package main

import (
	"fmt"

	"github.com/techoner/gophp"
)

func main() {

	str := `a:1:{s:3:"php";s:24:"世界上最好的语言";}`

	// unserialize() in php
	out, _ := gophp.Unserialize([]byte(str))

	fmt.Println(out) //map[php:世界上最好的语言]

	// serialize() in php
	jsonbyte, _ := gophp.Serialize(out)

	fmt.Println(string(jsonbyte)) // a:1:{s:3:"php";s:24:"世界上最好的语言";}

}
```

## License

This project is licensed under the [MIT license](LICENSE).