# Calculator
Nothing interesting, Just learning :D

### Create Module
```shell
go mod init github.com/gunjanpatel/calculator
```

### Refer Local package
```go
module helloworld

go 1.14

require github.com/gunjanpatel/calculator v0.0.0

replace github.com/gunjanpatel/calculator => ../calculator
```

## Publish Package

```shell
git tag v0.1.0
git push origin v0.1.0
```