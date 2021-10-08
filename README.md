## hookrelay

Minimal go library to relay webhook events back to an arbitrary service. With
the use of a primary HTTP mux router, we are able to register endpoints (e.g.,
`/hook`) with a corresponding `Response` function to handle the aggregation of
incoming requests.

First, there is a `Hook` interface we must satisfy by implementing it's
functions for later usage.

```go
type Hook interface {
	Response(i interface{}) func(w http.ResponseWriter, r *http.Request)
	LoadConfig(config string) error
	Endpoint() string
	Options() byte
}
```

* `Response`: Contains the incoming webhook event request data and defines  
how to handle it.
* `LoadConfig`: Used for any special configuration that may be used by  
the hook.
* `Endpoint`: The endpoint path that events should be sent to.
* `Options`:  Optional middleware a hook may find useful, such as limiting  
the endpoint to only accept `POST` methods.

Examples can be found in the [examples](examples) directory for practical usage
and a better understanding. There is not a great deal behind this, and was more
of a thin-wrapper to use in smaller sized, adhoc projects.

## License

[BSD 2-Clause](LICENSE)
