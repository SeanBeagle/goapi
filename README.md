# Simple Golang REST API
A simple template for a REST API written in go using:
- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [github.com/jinzhu/gorm](https://github.com/jinzhu/gorm)

## model
### Person
#### Person
| PROPERTY  | TYPE   |
|-----------|--------|
| id        | int    |
| firstName | string |
| lastName  | string |

#### CreatePersonDTO
| PROPERTY  | TYPE   |
|-----------|--------|
| id        | int    |
| firstName | string |
| lastName  | string |



## controller
### Person
| METHOD | ENDPOINT        | STATUS CODES  |
|--------|-----------------|---------------|
| GET    | /v1/person      | 200, 404      |
| GET    | /v1/person/{id} | 200, 404      |
| POST   | /v1/person.     | 201, 400      |
| PATCH  | /v1/person/{id} | 200, 400, 404 |

## service
### db
