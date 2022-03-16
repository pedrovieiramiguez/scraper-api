# scraper-api
Repo for a Golang-based API that scrapes supermarket websites for product information.

# How to use
Currently you can run this locally to scrape products from Pingo Doce's website.

Start the server by running
```shell
go run main.go
```

Curl it to get a JSON response containing the results
```shell
curl localhost:9090/pingodoce/massa
```

For spaces, use %20
```shell
curl localhost:9090/pingodoce/massa-espirais
```
