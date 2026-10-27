## Unreleased

### Improvements

- http: raise the floor of `MaxIdleConnsPerHost` in `DefaultPooledTransport` to 5
  (previously `GOMAXPROCS(0)+1`, which underestimates concurrency on low-CPU
  machines and causes connection churn) [#15](https://github.com/hashicorp/go-cleanhttp/issues/15)

### Changes

### Fixed

### Security