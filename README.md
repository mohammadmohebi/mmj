# mmj
JSON Creator library. This library helps to easly create JSON dynamicaly.

## To install
```bash
go get github.com/mohammadmohebi/mmj
```

## Create an object
```go
import "github.com/mohammadmohebi/mmj"

...
obj := mmj.NewObj()

obj.SetP("Honda", "car.brand")
obj.SetP("Civic", "car.model")
obj.SetP(2017, "car.year")

obj.Set(12000, "car", "price")
fmt.Println(obj.String())
```
Result
```json
{
  "car": {
    "brand":"Honda",
    "model":"Civic",
    "year":2017,
    "price":12000
  }
}
```

## Create an array
```go
import "github.com/mohammadmohebi/mmj"

...
arr := mmj.NewArray()
arr.Append("Honda")
arr.Append("Toyota")
arr.Append("Mazda")
fmt.Println(arr.String())
```
Result
```json
[
  "Honda",
  "Toyota",
  "Mazda"
]
```

## Insert object inside an object
```go
import "github.com/mohammadmohebi/mmj"

...
articles := mmj.NewObj()

obj := mmj.NewObj()
obj.SetP("vegetable", "type")
obj.SetP("radish", "name")
obj.SetP(2.5, "price")

articles.SetP(obj, "basket.article")

fmt.Println(articles.String())

```
Result
```json
{
  "basket": {
    "article": {
      "type":"vegetable",
      "name":"radish",
      "price": 2.5
    }
  }
}
```
## Insert objects inside an array
```go
import "github.com/mohammadmohebi/mmj"

...
articles := mmj.NewArray()

obj := mmj.NewObj()
obj.SetP("vegetable", "type")
obj.SetP("radish", "name")
obj.SetP(2.5, "price")
articles.Append(obj)

obj2 := mmj.NewObj()
obj2.SetP("fruit", "type")
obj2.SetP("apple", "name")
obj2.SetP(4, "price")
articles.Append(obj)

fmt.Println(articles.String())

```
Result
```json
[
  {
    "type":"vegetable",
    "name":"radish",
    "price": 2.5
  },
  {
    "type":"fruit",
    "name":"apple",
    "price": 4
  }
]
```

## Insert arrays inside objects
```go
import "github.com/mohammadmohebi/mmj"

...
articles := mmj.NewArray()

obj := mmj.NewObj()
obj.SetP("vegetable", "type")
obj.SetP("radish", "name")
obj.SetP(2.5, "price")
articles.Append(obj)

obj2 := mmj.NewObj()
obj2.SetP("fruit", "type")
obj2.SetP("apple", "name")
obj2.SetP(4, "price")
articles.Append(obj)

basket := mmj.NewObj()
basket.SetP(articles, "basket")

fmt.Println(basket.String())

```
Result
```json
{
  "basket" : [
    {
      "type":"vegetable",
      "name":"radish",
      "price": 2.5
    },
    {
      "type":"fruit",
      "name":"apple",
      "price": 4
    }
  ]
}
```
