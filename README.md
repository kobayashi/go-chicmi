## go-chicmi [![Build Status](https://travis-ci.org/kobayashi/go-chicmi.svg?branch=master)](https://travis-ci.org/kobayashi/go-chicmi)

go-chicmi is rest client to get data from chicmi API

### samples

- GET Events

```
func main() {
	httpClient := http.DefaultClient

	client := chicmi.NewClient(httpClient)

	events, resp, err := client.Search(&chicmi.SearchParams{
		City:       "new-york",
		DateFrom:   "2017/12/11",
		DateTo:     "2017/12/12",
		MaxResults: 1,
		Source:     "embed",
	})
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println(events)
}
```

- GET Event Detail

```
func main() {
	httpClient := http.DefaultClient

	client := chicmi.NewClient(httpClient)

	event, resp, err := client.GetEvent(&chicmi.EventSearchParams{
		EventID: 11878,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
}
```

## License

[MIT](https://github.com/kobayashi/go-chicmi/blob/master/LICENSE)

## Author

[kobayashi](https://github.com/kobayashi)
