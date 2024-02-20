# Simple Responder

The tiny web application that responds to any URI in predefined way.  
Usefull for api endpoint mocking.

## How to run
The simplest way just to run docker container
```
docker run --rm -p 8080:8080 maxon755/simple-responder -h
```

## Flags
`--status`: `int`  
The http status code to respond with. Default `200`.  
```
simple-responder --status 302
```

`--body`: `string`  
The response body. Default `OK`.
```
simple-responder --body "Internal server error"
```

`--body-file`: `string`  
The path to file with response body. Usefull when body content is too big and could not be specifed as single sting.
```
simple-responder --body-file /tmp/content.txt
```

`--delay`: `int`
The delay before response in seconds. Usefull for timeout emulation.

### Example
```bash
docker run --rm \
    -p 8080:8080 \
    -v content-on-your-machine:/tmp/content.txt \
    maxon755/simple-responder \
    -status 201 \
    -delay 3 \
    -body-file /tmp/content.txt 
```
> Note. The content file must be inside the container. Can be passed via volumes.

The app will response for any URI with with status `201` and content specifeid in `/tmp/content.txt` with delay of `3 seconds`.

```
curl -i localhost:8080/any-uri

HTTP/1.1 201 Created
Content-Type: text/plain; charset=utf-8
Date: Mon, 19 Feb 2024 17:46:34 GMT
Content-Length: 10

test body
```