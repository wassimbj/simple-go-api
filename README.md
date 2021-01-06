> a simple pure go web api

first run the file (wich is the server)
```bash
go run main.go
```

there is two routes there
/products => wich is a GET route, that will return the products  
/create => wich is a POST route, it will create a product; the data is stored in memory no DB is used.

open the cmd and write
```bash
curl http://localhost:1234/products
```
to get the products (its a slice of structs)

and
```bash
curl -X POST --data "{ \"name\": \"Iphone\", \"price\": 999}" http://localhost:1234/create
```
to create a products, it just append the data to the slice.


this is just for learning purpose.
