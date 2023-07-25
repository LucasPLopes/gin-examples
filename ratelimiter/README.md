# Ratelimit in Gin

This project is a sample for ratelimit using Leaky Bucket. Although the golang official pkg provide a implement with Token Bucket [time/rate](https://pkg.go.dev/golang.org/x/time/rate?tab=doc),

you can also make your owns via gin's functional `Use()` to integrate extra middlewares.

## Effect

```go
// You can assign the ratelimit of the server
// rps: requests per second
go run rate.go -rps=100
```

- Let's install ab tool for testing using apache2-utils

```
sudo apt install apache2-utils
```


- Let's have a simple test by ab with 3000 mock requests, not surprisingly，it will takes 10ms each request.

```bash
ab -n 3000 -c 1 http://localhost:8080/rate
```

- Gin Log Output

```bash
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |    9.189998ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |    9.128622ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |   10.175002ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |   10.166607ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |    9.196049ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |   10.180232ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |    9.199595ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |   10.174602ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |    9.200807ms |       127.0.0.1 | GET      "/rate"
[GIN] 10ms
[GIN] 2023/07/25 - 03:15:00 | 200 |   10.174603ms |       127.0.0.1 | GET      "/rate"
```

- AB Test Reporter

```java
Concurrency Level:      1
Time taken for tests:   29.991 seconds
Complete requests:      3000
Failed requests:        0
Total transferred:      420000 bytes
HTML transferred:       51000 bytes
Requests per second:    100.03 [#/sec] (mean)
Time per request:       9.997 [ms] (mean)
Time per request:       9.997 [ms] (mean, across all concurrent requests)
Transfer rate:          13.68 [Kbytes/sec] received
```