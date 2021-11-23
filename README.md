# gcint

Simple gc using integer vectors to simulate

Iterate primarily over what should be the shorter vector (readers) removing unused references in from versions vector

### Run 

```shell
go run main.go
```

Sample output:
```shell
Running: gcInt v0.1
======================================
versions=[1 5 10 14]
readers=[7 11]
EVAL: reader[0]=7
EVAL: version[0]=1
EVAL: version[1]=5
EVAL: version[2]=10
  >> APPEND: version=[5]; reader=[7]
EVAL: reader[1]=11
EVAL: version[0]=1
EVAL: version[1]=5
EVAL: version[2]=10
EVAL: version[3]=14
  >> APPEND: version=[10]; reader=[11]
preserve=[5 10 14]
```


### Test

Run all unit tests (i.e., [main_test.go](./main_test.go))

```
go test -v
```

Sample output:
```shell
expected=[5 10 14]; result=[5 10 14]
--- PASS: TestSimple (0.00s)
PASS
ok  	github.com/mrutkows/gc	0.597s

```
