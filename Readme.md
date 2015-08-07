[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

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
[license-img]: https://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: https://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/pkg4go/glob
