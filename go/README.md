# Go hacks

## numgen.go
Generates numbers - Yeah there are better ways and no reason to do this but I like go and didn't want to download a new tool
```
▶ numgen
Generate number list incremented by value
numgen start end increment (optional) precision
Ex: numgen 0 10 .01
Default precision is 0 ie no decimals
```

## substowords.go
Generates words from a large list of subdomains (using for s3 bucket bashing)

```
cat testfiles/substowords.test | go run substowords.go
```

## Discord message
Sends messages to a webhook defined as D_NOTIFICATION_WH env var
```
export D_NOTIFICATION_WH=[your webhook url]
discordmessage "hello world" "test" "etc"
```
