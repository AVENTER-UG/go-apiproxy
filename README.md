# go-apiproxy

This is a api proxy written in go. To use it is very simple.

```bash
API_TOKEN=thisisapitokenforaservice
API_URL=https://api.service.nothing

docker run -e API_TOKEN=$API_TOKEN -e API_URL=$API_URL -p 10777:10777 avhost/go-apiproxy:latest
```

After that, you can use this container as api endpoint without authentication header. It will add the authentication header and forward all requests to the API_URL.

Why this container! Well, in some cases its not a good idea to give a service direct access to a outside API. And sometimes there are some services they cannot handle token authentication.