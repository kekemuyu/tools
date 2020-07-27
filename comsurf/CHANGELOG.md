## 2016-07-21

- **fileserver** fileserver tries to open the following files before responding
  with 404:

    * $uri
    * $uri.html
    * $uri.htm
    * $uri/index.html
    * $uri/index.htm
