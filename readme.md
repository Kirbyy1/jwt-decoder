# JWT Decoder

This is a lightweight Go package, `jwtdecoder`, designed to decode JSON Web Tokens (JWT) without signature verification. The primary function, `DecodeJWT`, takes a JWT string, decodes its payload, and returns the claims as a JSON-encoded string.

## How it Works

The package operates through the following steps:

1. **Split JWT into Parts:**
   - The JWT is divided into three parts: header, payload, and signature.

2. **Decode Payload:**
   - The payload (second part) undergoes base64 URL-decoding.

3. **Unmarshal Payload into Map:**
   - The payload is then unmarshalled into a map, representing the JWT claims.

4. **Marshal Claims into JSON String:**
   - The claims are marshalled back into a JSON-encoded string.

5. **Return Result:**
   - The final output is the JSON-encoded string representing the claims.

## How to Use

To utilize this package, import it into your Go project and invoke the `DecodeJWT` function with the JWT string as an argument. The function will return the claims as a JSON-encoded string.

```go
import "https://github.com/Kirbyy1/jwt-decoder"

func main() {
    token := "your.jwt.token.here"
    claimsJSON, err := jwtdecoder.DecodeJWT(token)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Decoded JWT Claims:", claimsJSON)
}
```
