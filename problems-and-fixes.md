# Problem - `Scam error, unspported scan, storing driver.Value into *time.Time`:
* When attempting to get all menu's from the GET /api/menus endpoint. I got this error:

```
2023/02/04 11:11:43 /Users/tratnayake/go/src/github.com/tratnayake/golang-crud-rest-api/controllers/menucontroller.go:22 sql: Scan error on column index 1, name "created_at": unsupported Scan, storing driver.Value type []uint8 into type *time.Time; sql: Scan error on column index 1, name "created_at": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
[2.057ms] [rows:2] SELECT * FROM `menus` WHERE `menus`.`deleted_at` IS NULL
[GIN] 2023/02/04 - 11:11:43 | 200 |    2.685708ms |       127.0.0.1 | GET      "/api/menus
```
## Fix:
* This is beacuse `parseTime` needs to be set to true. See: [Connecting to the Database](https://gorm.io/docs/connecting_to_the_database.html)
* Changed `config.json` to `"connection_string": "root@tcp(127.0.0.1:3306)/crud_demo?parseTime=True",`

```
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2023/02/04 - 11:16:15 | 200 |    1.799667ms |       127.0.0.1 | GET      "/api/menus"
```

# Problem - Querying for Menu's shows Associations `null`
* When we query for Menu's, we should also be presented with the associated products.

GET http://{{host}}/api/menus HTTP/1.1
content-type: application/json

```
[
  {
    "ID": 1,
    "CreatedAt": "2023-01-04T18:09:57.054Z",
    "UpdatedAt": "2023-01-04T18:09:57.056Z",
    "DeletedAt": null,
    "Products": null
  },
  {
    "ID": 2,
    "CreatedAt": "2023-01-11T18:09:57.054Z",
    "UpdatedAt": "2023-01-11T18:09:57.056Z",
    "DeletedAt": null,
    "Products": null
  }
]
```

## Fix:
~https://blog.depa.do/post/gorm-gotchas use .Migrate()?~
* When querying for `GetMenus()`  do a Preload. 
  * `database.DB.Model(&models.Menu{}).Preload("Products").Find(&menus)` 