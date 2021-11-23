# gcint

Simple gc using integer vectors to simulate

Iterate primarily over what should be the shorter vector (readers) removing unused references in from versions vector

### TODO

**Note** There are still optimizations to reduce the inner loop against a Slice of the `versions` vector since we do NOT need to start at the beginning EACH time since reader entries are sorted...

### Run 

```shell
go run main.go
```

Sample output:
```shell
Running: gcint v0.1
======================================
versions=[1 5 10 14]
readers=[7 11 12]
EVAL: reader[0]=7
EVAL: version[0]=1
EVAL: version[1]=5
EVAL: version[2]=10
  >> APPEND: version=[5]; reader=[7]
NEW: versions=[10 14]
EVAL: reader[1]=11
EVAL: version[0]=10
EVAL: version[1]=14
  >> APPEND: version=[10]; reader=[11]
NEW: versions=[14]
EVAL: reader[2]=12
EVAL: version[0]=14
preserve=[5 10 14]
expected=[5 10 14]; result=[5 10 14]
```


### Test

Run all unit tests (i.e., [main_test.go](./main_test.go))

```
go test -v
```

Sample output:
```shell
=== RUN   TestSimple
versions=[1 5 10 14]
readers=[7 11]
EVAL: reader[0]=7
EVAL: version[0]=1
EVAL: version[1]=5
EVAL: version[2]=10
  >> APPEND: version=[5]; reader=[7]
NEW: versions=[10 14]
EVAL: reader[1]=11
EVAL: version[0]=10
EVAL: version[1]=14
  >> APPEND: version=[10]; reader=[11]
NEW: versions=[14]
preserve=[5 10 14]
expected=[5 10 14]; result=[5 10 14]
--- PASS: TestSimple (0.00s)

=== RUN   TestSimple2
versions=[1 5 10 14]
readers=[7 11 12]
EVAL: reader[0]=7
EVAL: version[0]=1
EVAL: version[1]=5
EVAL: version[2]=10
  >> APPEND: version=[5]; reader=[7]
NEW: versions=[10 14]
EVAL: reader[1]=11
EVAL: version[0]=10
EVAL: version[1]=14
  >> APPEND: version=[10]; reader=[11]
NEW: versions=[14]
EVAL: reader[2]=12
EVAL: version[0]=14
preserve=[5 10 14]
expected=[5 10 14]; result=[5 10 14]
--- PASS: TestSimple2 (0.00s)
PASS
ok  	github.com/mrutkows/gc	0.147s
```
