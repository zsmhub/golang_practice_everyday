package hello

import (
    "bytes"
    "fmt"
    "net/http"
    "net/http/httptest"
    "reflect"
    "strconv"
    "sync"
    "testing"
    "time"
)

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper() // 告诉测试套件这个方法是辅助函数。通过这样做，当测试失败时所报告的行号将在 函数调用中 而不是在辅助函数内部。这将帮助其他开发人员更容易地跟踪问题。
        if got != want {
            t.Errorf("got %q wang %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        name := "Chris"
        got := Hello(name)
        want := englishHelloPrefix + name

        assertCorrectMessage(t, got, want)
    })

    t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("")
        want := englishHelloPrefix + "World"

        assertCorrectMessage(t, got, want)
    })
}

func TestAdd(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}

func BenchmarkSprintf(b *testing.B) {
    num := 10
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("%d", num)
    }
}

func BenchmarkFormat(b *testing.B) {
    num := int64(10)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        strconv.FormatInt(num, 10)
    }
}

func BenchmarkItoa(b *testing.B) {
    num := 10
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        strconv.Itoa(num)
    }
}

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Chris")

    got := buffer.String()
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func slowStubWebsiteChecker(_ string) bool {
    time.Sleep(20 * time.Millisecond)
    return true
}

func BenchmarkCheckWebsites(b *testing.B) {
    urls := make([]string, 100)
    for i := 0; i < len(urls); i++ {
        urls[i] = "a url"
    }

    for i := 0; i < b.N; i++ {
        CheckWebsites(slowStubWebsiteChecker, urls)
    }
}

func TestRacer(t *testing.T) {
    slowServer := makeDelayedServer(20 * time.Millisecond)
    fastServer := makeDelayedServer(0 * time.Millisecond)

    defer slowServer.Close()
    defer fastServer.Close()

    slowURL := slowServer.URL
    fastURL := fastServer.URL
    t.Logf("%s, %s", slowURL, fastURL)

    want := fastURL
    got := Racer(slowURL, fastURL)

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}

func TestWalk(t *testing.T) {
    cases := []struct {
        Name          string
        Input         interface{}
        ExpectedCalls []string
    }{
        {
            "Struct with one string field",
            struct {
                Name string
            }{"Chris"},
            []string{"Chris"},
        },
        {
            "Struct with two string fields",
            struct {
                Name string
                City string
            }{"Chris", "London"},
            []string{"Chris", "London"},
        },
    }

    for _, test := range cases {
        t.Run(test.Name, func(t *testing.T) {
            var got []string
            walk(test.Input, func(input string) {
                got = append(got, input)
            })

            if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %v, want %v", got, test.ExpectedCalls)
            }
        })
    }
}

func TestCounter(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        counter := Counter{}
        counter.Inc()
        counter.Inc()
        counter.Inc()

        assertCounter(t, counter, 3)
    })

    t.Run("test2", func(t *testing.T) {
        wantedCount := 1000
        counter := Counter{}

        var wg sync.WaitGroup
        wg.Add(wantedCount)

        for i := 0; i < wantedCount; i++ {
            go func(w *sync.WaitGroup) {
                counter.Inc()
                w.Done()
            }(&wg)
        }
        wg.Wait()

        assertCounter(t, counter, wantedCount)
    })
}

func assertCounter(t *testing.T, got Counter, want int) {
    t.Helper()
    if got.Value() != want {
        t.Errorf("got %d, want %d", got.Value(), want)
    }
}
