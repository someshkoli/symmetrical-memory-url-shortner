# symmetrical-memory-url-shortner

Shorted a url using `/set` route and use the resulting link to get back to original url.

## Store
### Inmemory
All the url and shortened are stored in memory
### File storage
All the url and shortened are stored in memory which is used as cache and is also stored in a file for persistent storage

## Installation
Start the server with binary:

    $ ./bin/symmetrical-memory-url-shorner

Start with docker [someshkoli/symmetrical-memory-url-shortner](https://hub.docker.com/repository/docker/someshkoli/symmetrical-memory-url-shortner)

    $ docker run -p 8888:8888 someshkoli/symmetrical-memory-url-shortner

## Usage
    $ http localhost:8888/set?url=google.com
    HTTP/1.1 202 Accepted
    Content-Length: 48
    Content-Type: text/plain; charset=utf-8
    Date: Sun, 16 May 2021 12:45:17 GMT

    {
        "Data": "localhost:8888/g/aGsxi",
        "Status": true
    }

    $ http "localhost:8888/g/aGsxi" --follow
    HTTP/1.1 200 OK
    Alt-Svc: h3-29=":443"; ma=2592000,h3-T051=":443"; ma=2592000,h3-Q050=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000,quic=":443"; ma=2592000; v="46,43"
    Cache-Control: private, max-age=0
    Content-Encoding: gzip
    Content-Type: text/html; charset=ISO-8859-1
    Date: Sun, 16 May 2021 12:57:21 GMT
    Expires: -1
    P3P: CP="This is not a P3P policy! See g.co/p3phelp for more info."
    Server: gws
    Set-Cookie: 1P_JAR=2021-05-16-12; expires=Tue, 15-Jun-2021 12:57:21 GMT; path=/; domain=.google.com; Secure
    Set-Cookie: NID=215=cX-lRHKA5frSNfil3LmKeDIWyIE3keiP3HgZ8nKJrJznBwEMfGTTZRtMuG4y_RmE5fn0wobDY4FI4BYOYCyIulnllFeZv83LlsaoCAN7cm7owQ2Z8NG3PCuK75A1bT5CzVtdtjus32JIIh5o5x2h8kkIg5ZGLRItxr8nhcWREcs; expires=Mon, 15-Nov-2021 12:57:21 GMT; path=/; domain=.google.com; HttpOnly
    Transfer-Encoding: chunked
    X-Frame-Options: SAMEORIGIN
    X-XSS-Protection: 0
