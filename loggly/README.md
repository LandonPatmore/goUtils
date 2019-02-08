# Loggly
This package interfaces with the [Loggly logging service](www.loggly.com).

## Usage

```
The "LOGGLY_API_KEY" env variable must be set, or this will not work
```

There are multiple different types of log messages that can be sent, they are modeled after the log4j log messages:

```go
func Error(message interface{}){}
func Warn(message interface{}){}
func Debug(message interface{}){}
func Info(message interface{}){}
func Trace(message interface{}){}

func ErrorEcho(message interface{}){}
func WarnEcho(message interface{}){}
func DebugEcho(message interface{}){}
func InfoEcho(message interface{}){}
func TraceEcho(message interface{}){}
```

Regular log functions just send the log message to Loggly.

Echo log functions send the log message to loggly, **as well as** print out the log message to the console.

These functions can also take in any types of structs or interfaces, e.g **string, int, customo struct, etc.**

## Example

```go
package main

import (
  "github.com/landonp1203/goUtils/loggly"
)

func main() {
  loggly.EchoDebug("Hey this is a debug message!")
}
```

Console Output:
```
2019/02/08 01:37:28 Tag: Debug
 Hey this is a debug message!
```
