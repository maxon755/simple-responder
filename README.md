# Simple Responder

The tiny web application that responds to any URI in predefined way.  
Usefull for api endpoint mocking.

## How to run
The simplest way just to run docker container
```
docker run --rm -p 8080:8080 maxon755/simple-responder -h
```

## Usage
```
  -status int
    	The status to respond with (default 200)
  -body string
    	The body to resnsond with (default "OK")
  -body-file string
    	The path to file with response content
  -delay int
    	The delay before response. Usefull for timeout emulation
```

The response content can be specified either with `-body` parameter 
as a string or via file with `-body-file` parameter.

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

This app will response for any URI with with status `201` and content specifeid in `/tmp/content.txt` with delay of `3 seconds`.

```
curl -i localhost:8080

HTTP/1.1 201 Created
Content-Type: text/plain; charset=utf-8
Date: Mon, 19 Feb 2024 17:46:34 GMT
Content-Length: 10

test body
```