[![CircleCI](https://circleci.com/gh/velvetreactor/nginxconf.svg?style=svg)](https://circleci.com/gh/velvetreactor/nginxconf)

### nginxconf
* This project aims to:
  - Parse nginx configuration files
  - Parse custom JSON files
  - Generate valid and semantically correct nginx config files from these JSON files

### Routes JSON schema
```
{
  "routes": [
    {
      "host_endpoint": "/",
      "proxy_to": "http://www.google.com",
      "append_path": true
    }
  ]
}
```
* `host_endpoint`: __String__ The path on the host machine running nginx
* `proxy_to`: __String__ The endpoint to redirect to
* `append_path`: __Boolean__ Determines whether or not the request path should be appended to the proxy path
  * Ex: `{ "append_path": true, "host_endpoint": "/peterlugers", "proxy_to": "http://www.yelp.com"`
    * Request made to `http://HOST/peterlugers` redirects to `http://www.yelp.com/peterlugers`
  * Ex: `{ "append_path": false, "host_endpoint": "/google/search_engine", "proxy_to": "http://www.google.com" }`
    * Request made to `http://HOST/google/search_engine` redirects to `http://www.google.com`

### Usage
* Run `go run main.go [input file]`
* The generated conf file will be output to `test.conf` in your working directory

### nginx testing
* After running the `go run` command and generating the `test.conf` file, run
  `make proxy-test` in order to run an nginx Docker container with the conf file
  mounted.
* The container will be listening on `0.0.0.0:80`, so pointing your browser
  toward `localhost/[path]` will allow you to test your proxy settings.
