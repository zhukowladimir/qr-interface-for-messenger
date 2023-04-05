```
❯ go run . -mode=encode -text="Never gonna give you up" -qr-path="./mvp-demo/pure-qr-1.jpg"

❯ go run . -mode=encode -text="Never gonna let you down" -qr-path="./mvp-demo/pure-qr-2.jpg"

❯ go run . -mode=encode -text="Never gonna run around and desert you" -qr-path="./mvp-demo/pure-qr-3.jpg"

❯ go run . -mode=decode -qr-path="./mvp-demo/pure-qr-3.jpg"
Never gonna run around and desert you

❯ go run . -mode=decode -qr-path="./mvp-demo/mirror-qrs.jpg"
Never gonna let you down
Never gonna give you up
Never gonna run around and desert you
```