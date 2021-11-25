# Url Rewrite middleware for gin

## Example

In this exable these urls use the same route

- `http://localhost:1234/test-me`
- `http://localhost:1234/index.php/test-me`

And

- `http://localhost:1234/v1/test-me`
- `http://localhost:1234/v1.php/test-me`


```go
package main

func main() {
	r := gin.Default()
	r.Use(
      middleware.UrlRewrite(r, "/index.php", ""),
      middleware.UrlRewrite(r, "/v1.php", "/v1"),
   )

	r.GET("/test-me", func(c *gin.Context) {
		c.Status(http.StatusTeapot)
	})

   r.GET("/v1/test-me", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.Run(":1234")
}
```

## Install

copy the contents of [gin-url-rewrite.go](/gin-url-rewrite.go) in to your project.

## License

WTFPL – Do What the Fuck You Want to Public License
