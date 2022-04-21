[![Build Status Badge]][Build Status]
[![Go Docs Badge]][Go Docs]

## Glob

- __glob__ for Golang, add `**/*` support.

### Usage

```go
import "github/pkg4go/glob"

func main() {
  matches, err = Glob(".", "**/*.yaml")
  Expect(err).To(BeNil())
  Expect(len(matches)).To(Equal(1))
  Expect(matches[0]).To(Equal(".github/workflows/test.yaml"))
}
```

[Build Status Badge]: https://github.com/haoxins/glob/actions/workflows/test.yaml/badge.svg
[Build Status]: https://github.com/haoxins/glob/actions/workflows/test.yaml
[Go Docs Badge]: https://pkg.go.dev/badge/github.com/haoxins/glob
[Go Docs]: https://pkg.go.dev/github.com/haoxins/glob
