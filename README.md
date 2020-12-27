# goServerGenRandNumber

## Compile the project

```
$ compileit.sh
```


## Build the docker image
```
$ docker build -t sample .
```

## Run the docker image

```
$ docker run -p8080:8080 sample
```

## Use

### Input

```
curl http://localhost:8080/
```

### Output

```
{"randomnumber": 20.000000" }
```

## References

[GO by example](https://gobyexample.com/)
[A tour of GO](https://tour.golang.org/welcome/1)
[An intro to program in GO](https://www.golang-book.com/books/intro)
[Go Project Structure](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/)
[First Rest API Go](https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj)


