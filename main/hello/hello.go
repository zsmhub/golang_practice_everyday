package hello

import (
    "bytes"
    "fmt"
    "net/http"
    "reflect"
    "sync"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
    if name == "" {
        name = "World"
    }
    return englishHelloPrefix + name
}

func Add(x, y int) int {
    return x + y
}

func Repeat(character string) string {
    var repeated string
    for i := 0; i < 5; i++ {
        repeated += character
    }
    return repeated
}

func main() {
    fmt.Println(Hello("Chris2"))
}

func Greet(writer *bytes.Buffer, name string) {
    _, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

type WebsiteChecker func(string) bool
type result struct {
    string
    bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
    resultChannel := make(chan result)

    for _, url := range urls {
        go func(u string) {
            resultChannel <- result{u, wc(u)}
        }(url)
    }

    for i := 0; i < len(urls); i++ {
        result := <-resultChannel
        results[result.string] = result.bool
    }

    return results
}

/*func Racer(a, b string) (winner string) {
    aDuration := measureResponseTime(a)
    bDuration := measureResponseTime(b)

    if aDuration < bDuration {
        return a
    }
    return b
}

func measureResponseTime(url string) (time.Duration) {
    start := time.Now()
    _, _ = http.Get(url)
    return time.Since(start)
}*/

func ping(url string) chan struct{} {
    ch := make(chan struct{})
    go func() {
        _, _ = http.Get(url)
        close(ch)
    }()
    return ch
}

func Racer(a, b string) (winner string) {
    select {
    case <-ping(a):
        return a
    case <-ping(b):
        return b
    }
}

func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fn(field.String())
    }
}

type Counter struct {
    mu sync.Mutex
    value int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    return c.value
}