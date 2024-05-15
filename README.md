# golang rps (requests per second)

Simulate running a simple service where information fetched from a DB or cache and then sent to a notification worker. The simulated time is set to 100ms, I'm not sure if this is reasonable or not, it depends on a bunch of factors, the core idea is that there is some latency between the request and pushing the data into the queue.

Typically in a system design context you are given estimates of requests per day or month and then convert that into requests per second.

5M requests per day sounds like a lot, at least to me, but that's only 50 requests per second.

A notification service that has on any given day:

- 10M requests for push
- 1M for sms
- 5M for email

will have 160 req/s.

You can serve that on an entry level machine.

```sh
docker build -t go-notification-service .
docker run -it --rm -p 8080:8080 --cpus=1 --memory=512mb --name go-notification-service go-notification-service
```

```sh
λ ~/code/golang-rps: wrk -t12 -c400 -d30s -s post.lua http://localhost:8080/notify

Running 30s test @ http://localhost:8080/notify
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    29.27ms    3.24ms  79.36ms   86.59%
    Req/Sec     1.13k   113.52     1.33k    73.03%
  406397 requests in 30.10s, 55.81MB read
  Socket errors: connect 0, read 375, write 0, timeout 0
Requests/sec:  13503.26
Transfer/sec:      1.85MB
```


