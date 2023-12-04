# E-commerce API

Where you can buy everything you need.

**How to generate the secretKey**

```go
func init(){
    key := make([]byte, 64)

    if _, err := rand.Read(key); err != nil{
        log.fatal(err)
    }

    secretKey := base64.StdEncoding.EncodeToString(key)
    fmt.Println(secretKey)
}
``