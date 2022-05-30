#   Remember to BET

<ul>
    <li>Benchmark</li>
    <li>Example</li>
    <li>Test</li>
</ul>


# Funcs
```
BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.T)
```

# Commands

```
go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
go tool cover -html=c.out -o coverage.html
```