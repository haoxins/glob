[![Build status][travis-img]][travis-url]
[![PkgGoDev](https://pkg.go.dev/badge/pkg4go/glob)](https://pkg.go.dev/pkg4go/glob)

### glob

* `glob` for golang, add `**/*` support.

### Usage

```go
import "github/pkg4go/glob"

func main() {
  glob.Glob("rootDir", "**/*.go")
  // return
  // matches []string, err error
}
```

### License
MIT

[travis-img]: https://img.shields.io/travis/pkg4go/glob.svg?style=flat-square
[travis-url]: https://travis-ci.org/pkg4go/glob
