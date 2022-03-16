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

Example response
```shell
[
   {
      "name":" Massa Espirais Caçarola ",
      "price":" 1,25 € "
   },
   {
      "name":" Massa Espirais Milaneza ",
      "price":" 1,35 € "
   },
   {
      "name":" Massa Espirais Pingo Doce ",
      "price":" 0,63 € "
   },
   {
      "name":" Massa Espirais sem Glúten Milaneza ",
      "price":" 2,39 € "
   },
   {
      "name":" Massa Espirais de Trigo Sarraceno Vianeza ",
      "price":" 2,84 € "
   },
   {
      "name":" Massa Espirais Integral com Alfarroba Nacional ",
      "price":" 1,39 € "
   },
   {
      "name":" Massa Espirais Fusilli Molde de Bronze Milaneza ",
      "price":" 1,59 € "
   },
   {
      "name":" Massa Fusilli Espirais Tricolor Milaneza ",
      "price":" 1,07 € "
   },
   {
      "name":" Massa Espiral de Quinoa e Lentilhas Caçarola ",
      "price":" 1,53 € "
   }
]
```

