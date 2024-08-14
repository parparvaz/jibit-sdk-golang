
### go-jibit

A Golang SDK for [jibit](https://jibit.ir) API.

All the REST APIs listed in jibit API document are implemented.

For best compatibility, please use Go >= 1.22.

Make sure you have read jibit API document before continuing.

### Installation

```shell
go get github.com/parparvaz/jibit-sdk-golang
```

### REST API

#### Setup

Init client for API services. Get ApiKey from your jibit account.

```golang
var (
	apiKey = "your api key"
        secretKey = "your secret key"
)
client := jibit.NewClient(apiKey, secretKey)
```

A service instance stands for a REST API endpoint and is initialized by client.NewXXXService function.

Simply call API in chain style. Call Do() in the end to send HTTP request.

If you have any questions, please refer to the specific reference definitions or usage methods

##### Proxy Client

```
proxyUrl := "http://127.0.0.1:7890" // Please replace it with your exact proxy URL.
client := binance.NewProxyClient(apiKey, secretKey, proxyUrl)
```


#### Send Lookup SMS

```golang
res, err := client.NewMatchCardNumberWithNameService().
	CardNumber("card number").
	Name("name").
	Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)

```
