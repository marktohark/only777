# only777
This is based on gin web framework.

# Install
```
$> git clone https://github.com/marktohark/only777 
$> cd only777
$> go mod tidy
```

# Contents
- [Route](#Route)
- [Controller](#Controller)
- [Middleware](#Middleware)
- [Env](#Env)
- [View](#View)
- [Access static directory](#Access)

## Route
1.use default route function   
path: route/custom/default.go
```
func Default(r *gin.Engine) error {
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
          "message": "pong",
        })
    })
    return nil
}
```
2.how to custom route function   
step.1 make directory in route or new a function in default.go.
```
func Test(r *gin.Engine) error {
	r.GET("/123", func(c *gin.Context) {
        c.JSON(200, gin.H{
          "message": "pong",
  })
	return nil
}
```
step.2 registe in route/reg.go
```
func RegList() []func(*gin.Engine)(error) {
	return []func(*gin.Engine)(error) {
		custom.Default,
		custom.Test, // this is your route function
	}
}
```
3.how to use route   
https://github.com/gin-gonic/gin#using-get-post-put-patch-delete-and-options
## Controller
make a directory in controller, name is your controller Name.   
>-controller   
>>-HomeController   
>>>main.go   
```
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
```

route call controller   
```
func Default(r *gin.Engine) error {
	r.GET("/", HomeController.Index)
	return nil
}
```
## Middleware
1.how to make custom middleware function   
make directory and *.go in middleware(directory)
```
package MyMiddleware
func TestMiddleware(c *gin.Context) {}
```
2.regist global middleware   
path: middleware/global/reg.go
```
func RegGlobalMiddleware(r *gin.Engine) {
	/**
		r.Use or func(r)
	 */
	_default.SessionMiddleware(r)
	_default.CsrfMiddleware(r)
	r.Use(MyMiddleware.TestMiddleware) //this is your middleware
}
```
3.regist route middleware   
```
r.Get("middleware", 
      MyMiddleware.xxxxMiddleware, //this is your middleware
      HomeController.Index)
```
## Env
Threr is a .env.example under the root directory.   
copy this file and rename file to ".env".
```
PORT => port
COOKIE_SECRET => 
CSRF_SECRET => 
SESSION_NAME => 
STATIC_URL_NAME => static url name
```
you can custom parameter in .env and call config.Get("parameter key").
```
val, ok := config.Get("port")
if ok {}
fmt.print(val)
```
## View
>-views
>>index.html
```
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
```
html template:
https://github.com/gin-gonic/gin#html-rendering

## Access static directory
