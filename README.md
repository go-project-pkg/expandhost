# expandhost

Expand host pattern to host list.

E.g.:

Expand `foo[01-03,06,10-12].[beijing,wuhan].bar.com`, the result will be as follows:

```go
[]string{
    "foo01.beijing.bar.com",
    "foo01.wuhan.bar.com",
    "foo02.beijing.bar.com",
    "foo02.wuhan.bar.com",
    "foo03.beijing.bar.com",
    "foo03.wuhan.bar.com",
    "foo06.beijing.bar.com",
    "foo06.wuhan.bar.com",
    "foo10.beijing.bar.com",
    "foo10.wuhan.bar.com",
    "foo11.beijing.bar.com",
    "foo11.wuhan.bar.com",
    "foo12.beijing.bar.com",
    "foo12.wuhan.bar.com",
}
```

## Usage

```go
import "github.com/go-project-pkg/expandhost"

func main() {
    pattern := "foo[01-03,06,10-12].[beijing,wuhan].bar.com"

    hosts, err := expandhost.PatternToHosts(pattern)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%v\n", hosts)
}
```

Output:

```text
[foo01.beijing.bar.com foo01.wuhan.bar.com foo02.beijing.bar.com foo02.wuhan.bar.com foo03.beijing.bar.com foo03.wuhan.bar.com foo06.beijing.bar.com foo06.wuhan.bar.com foo10.beijing.bar.com foo10.wuhan.bar.com foo11.beijing.bar.com foo11.wuhan.bar.com foo12.beijing.bar.com foo12.wuhan.bar.com]
```

## License

This project is under the MIT License.
See the [LICENSE](LICENSE) file for the full license text.
