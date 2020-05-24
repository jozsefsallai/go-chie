# go-chie

`chie` is a Go port of the [chie Node.js library][1] which aims to be more
performant and cross-platform. It serves a decryption/encryption library and
command-line tool for Cave Story's script file format (TSC).

## Getting Started

**Get the package:**

```sh
go get -u github.com/jozsefsallai/go-chie
```

**Use it:**

```go
package main

import "github.com/jozsefsallai/go-chie"

func main() {
  tsc := chie.NewTSCParser()

  err := tsc.FromFile("./Ballo1.tsc")
  if err != nil {
    panic(err)
  }

  err = tsc.Decrypt().ToFile("./Ballo1.txt")
  if err != nil {
    panic(err)
  }
}
```

**Full docs and reference:** https://pkg.go.dev/github.com/jozsefsallai/go-chie

## CLI

`chie` also comes with a CLI. You can either [grab a release][2]* or install via
`go get`:

```sh
go get github.com/jozsefsallai/go-chie/cmd/chie
```

_*) if you do this, you might want to download it somewhere inside your PATH_

**Usage:**

```sh
chie tsc decrypt Ballo1.tsc
# will decrypt to output.txt

chie tsc decrypt Ballo1.tsc -o Ballo1.txt
# will decrypt to Ballo1.txt

chie tsc encrypt Ballo1.txt
# will encrypt to output.tsc

chie tsc encrypt Ballo1.txt -o Ballo1.tsc
# will encrypt to Ballo1.tsc
```

or decrypt/encrypt multiple files at once:

```sh
chie tsc decrypt "Stage/*.tsc" -o "Stage/*.txt"
# will decrypt all files with the extension ".tsc" in the Stage folder
# the quotes are necessary
```

## License

MIT.

[1]: https://github.com/jozsefsallai/chie
[2]: https://github.com/jozsefsallai/go-chie/releases/latest
