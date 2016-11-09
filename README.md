# Spinarak : The Cute Web Crawler

![](http://cpl.mwisely.xyz/hw/7/spinarak.png)

This tool is designed to assist Meowth with his new desk job: watching the news for any mention of Team Rocket.

Documentation for its behavior can be found publicly here: http://cpl.mwisely.xyz
To verify that the code works properly, compare its output against provided samples.

**Note: Follow the design specifications EXACTLY.**
Not doing so will hurt your grade.

**Note: the following commands will work on campus machines.**
**If you use your own machine or editors, you are on your own.**

## Setup Go 1.7.3

~~~shell
$ bash setup.sh
$ GOROOT="$(pwd)/go" ./go/bin/go version
go version go1.7.3 linux/amd64
~~~

## Check and Correct Style

~~~shell
$ GOROOT="$(pwd)/go" ./go/bin/go fmt spinarak.go
~~~~

This will run the `go fmt` tool to properly format your Go code.

**Warning** `go fmt` will modify your `.go` file(s).
You should close your text editors before you run it.
Not all editors are smart enough to handle the "file changed beneath its feet" situation.

## Run the Program

Build and run with `go run`

~~~shell
# List links
$ GOROOT="$(pwd)/go" ./go/bin/go run spinarak.go -word "meowth" http://www.npr.org http://www.pbs.org/
~~~

Or build an executable with `go build`.
Don't forget to recompile!

~~~shell
$ GOROOT="$(pwd)/go" ./go/bin/go build spinarak.go
$ ./spinarak http://www.npr.org http://www.pbs.org/
~~~

**DO NOT** commit compiled files to your git repository.

**DO NOT** add the go compiler (or its `.tar.gz`) to your git repository.
