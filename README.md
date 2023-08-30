# Encrypt Decrypt Crypto

Foobar is a Python library for dealing with word pluralization.

## Installation

```bash
go get github.com/ryanmustofa/encryptdecrypt
```

## Usage

```go
package main
import(
  encryptdecrypt "github.com/ryanmustofa/encryptdecrypt"
  "fmt"
)

func main(){
  encrypt , err := encryptdecrypt.Encrypt("1234567891010234","text for encrypt")
  if err != nil {
    panic(err)
  }
  fmt.Println(encrypt)
  decrypt , err := encryptdecrypt.Decrypt("1234567891010234",encrypt)
  if err != nil {
    panic(err)
  }
  fmt.Println(decrypt)

}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
