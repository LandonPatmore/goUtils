# Networking
This package is a wrapper over the [net/http](https://golang.org/pkg/net/http/) package.  It handles the response and will return errors back to the user to deal with once scrubbed properly.

## Usage

There are different functions:

```go
func Get(url string) ([] byte, error){}
func Post(url string, bodyData string) ([] byte, error){}
func PostJson(url string, jsonData interface{}) ([] byte, error){}
```

### Get
Takes in a URL string and then returns the response of the GET request

### Post
Takes in a URL string and data string to send, and then returns the response of the POST request

### PostJson
Takes in a URL string and an interface, and then returns the response of the POST request
