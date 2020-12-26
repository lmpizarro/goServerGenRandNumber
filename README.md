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

[Go Project Structure](https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/)


