# Go Password Generator

This is a simple library to generate random passwords with a given length and a given set of characters. 
It is written in Go and is based on the [LastPass Password Generator](https://www.lastpass.com/features/password-generator)

Example of password generated with this library:

```
LyKcKyBM
_zhy?e>x#hfox?{;
s2H)LkH;+;aBMPCi<;TFeVw0%!49EX,9
wnsgsqnestjtrwosltbqwihdlkvjvwiviicvbdiytebrunvl
```

### Installation

```
go get -u github.com/viniciussouzao/go-password-gen
```

### Usage

```go

package main

import (
	"fmt"
	"log"

	"github.com/viniciussouzao/go-password-gen/password"
)

func main() {
	//Generate a simple password with 16 characters, without numbers and symbols, but with uppercase letters
	pass, err := password.Generate(16, true, true, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pass)

	//Generate a custom password with 22 characters, with 3 numbers and 1 symbol and 4 uppercase letters
	pass, err = password.GenerateCustom(22, 3, 1, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pass)
}
    
```

### License
---
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details