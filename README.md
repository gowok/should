# must
Assertion utilities for testing in Go

Write your test as usual.
```go
func TestAdd(t *testing.T) {
	actual := Add(1, 2)
	expected := 3
	if actual != expected {
		t.Errorf("Add should return %d, not %d", expected, actual)
	}
}
```

Run the test with `-v` flag, you will see the output like bellow.
```
=== RUN   TestAdd
    example_test.go:9: Add should return 3, not 2
--- FAIL: TestAdd (0.00s)
```

Refactor by using golang must!
```go
func TestAdd(t *testing.T) {
	actual := Add(1, 2)
	expected := 3

	must.New(t).Equal(expected, actual)
}
```

Run the test with `-v` flag, you will see the output like bellow.
```
=== RUN   TestAdd
          example_test.go:13: need 3, got 2
--- FAIL: TestAdd (0.00s)
```

Your test more clean and clear, now 👍
