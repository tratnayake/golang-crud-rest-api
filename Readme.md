# GoLang-CRUD-REST-API
A learning journey to build a Go REST API for a webapp.

# Helpers
#### MYSQL SETUP
* `brew install mysql`
* `brew services start mysql`
* `mysql -u root`


# Log:

##  03 Feb 23 -  Tutorial 
* Link: https://codewithmukesh.com/blog/implementing-crud-in-golang-rest-api/
* Created a very basic CRUD API using GoLang that will CRUD a single entity into a MySQL database.

## 03 Feb 23 - Tutorial - 002-Add-Gin-Framework
* Link: https://medium.com/codepubcast/learn-to-develop-restful-apis-like-a-pro-with-go-and-gin-cfa36b28a464
* Switched from Mux + raw handlers to using Gin API framework.
    * Also helpful link: https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
    *  Ref for next time:
        * [Grouping Routes](https://pkg.go.dev/github.com/gin-gonic/gin#readme-grouping-routes) - how to do `api/v1` or `/v2/*` .
        * [Using Middleware](https://pkg.go.dev/github.com/gin-gonic/gin#readme-using-middleware) - for auth, and logging
        * [Model Binding and Validation](https://pkg.go.dev/github.com/gin-gonic/gin#readme-model-binding-and-validation) - if we wanna do custom binding / validation for requests.
        * [Graceful Shutdowns](https://pkg.go.dev/github.com/gin-gonic/gin#readme-graceful-shutdown-or-restart)

## 03 Feb 23 - 003-Add-Associations-Relations
    * Link: https://gorm.io/docs/associations.html#Association-Mode
    * Goal: Create a relation between Menu -> (has many) -> Products.
    TODO: 
        * Wire the API endpoints for menu creation
        * Constraints
        * How to delete from DB when there are constraints?