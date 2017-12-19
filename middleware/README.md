# Middleware Folder

Contains all middleware files

### Basic Auth
 Usage

 ```go
 basicAuthConfig := middleware.NewConfig("Wuriyanto", "12345")

 http.Handle("/api/fuck", middleware.BasicAuth(basicAuthConfig, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
   utils.JsonResponse(res, "Holla", http.StatusOK)
 })))
 ```
