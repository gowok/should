# should
Assertion utilities for testing in Go

Write your test as usual.
```go
func TestAdd(t *testing.T) {
	a := 1
	b := 2
	actual := Add(a, b)
	expected := 3

	if actual != expected {
		t.Errorf("need %d, got %d", expected, actual)
	}

	if actual <= a {
		t.Errorf("%d < %d is false", actual, a)
	}

	if actual <= b {
		t.Errorf("%d < %d is false", actual, b)
	}
}
```

Run the test with `-v` flag, you will see the output like bellow.
```
=== RUN   TestAdd
    example_test.go:14: need 3, got 2
    example_test.go:22: 2 < 2 is false
--- FAIL: TestAdd (0.00s)
```

Refactor by using golang must!
```go
func TestAdd(t *testing.T) {
	a := 1
	b := 2
	actual := Add(a, b)
	expected := 3

	must := should.New(t)
	should.Equal(expected, actual)
	should.True(actual <= a)
	should.True(actual <= b)
}
```

Run the test with `-v` flag, you will see the output like bellow.
```
=== RUN   TestAdd
          example_test.go:16: need 3, got 2
          example_test.go:17: need true, got false
--- FAIL: TestAdd (0.00s)
```

Your test more clean and clear, now 👍
