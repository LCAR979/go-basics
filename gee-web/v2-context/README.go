func (r *Request) FormValue(key string) string

/*
FormValue returns the first value for the named component of the query.
POST and PUT body parameters take precedence over URL query string values.
FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions.
If key is not present, FormValue returns the empty string.
To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.
*/

func (r *Request) PostFormValue(key string) string

/*
PostFormValue returns the first value for the named component of the POST, PATCH, or PUT request body.
URL query parameters are ignored.
PostFormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions.
If key is not present, PostFormValue returns the empty string.
*/