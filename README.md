# expandhost

Expand host pattern to host list.

E.g.:

Expand `foo[01-03,06,10-12].bar.com`, the result will be as follows:

```go
[]string{
    "foo01.bar.com",
    "foo02.bar.com",
    "foo03.bar.com",
    "foo06.bar.com",
    "foo010.bar.com",
    "foo011.bar.com",
    "foo012.bar.com",
}
```

## Usage

```go
import "github.com/go-project-pkg/expandhost"

func main() {
    pattern := "foo[01-03,06,10-12].bar.com"

    hosts, err := expandhost.PatternToHosts(pattern)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%v\n", hosts)
}
```

Output:

```text
[foo01.bar.com foo02.bar.com foo03.bar.com foo06.bar.com foo10.bar.com foo11.bar.com foo12.bar.com]
```

## License

This project is under the MIT License.
See the [LICENSE](LICENSE) file for the full license text.
