# gorawgit
A simple Go &amp; Docker project to give you raw git files with proper content types

## Building

A PowerShell script is included if you want to automatically build the image. It will simply listen on port 80 and serve you whatever
you feed it in the URL path. ex:

`localhost/golang/go/master/src/expvar/expvar.go`

## Running

ca-certificates.crt is a simple hack since Alpine doesn't include it, it's taken from the default Ubuntu 16.04 install so https works for major sites(like GitHub).

To run this image simply run: `docker run -p 80:80 imvexxed/gorawgit`
