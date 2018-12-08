# Example Go Service

This is a Golang application which exposes a single HTTP endpoint `/health` that will return a HTTP 200 response containing a JSON body:

    {
        "status": "UP"
    }

I use this application as a sample app when experimenting with infrastructure and provisioning tools and processes.

## Build and run the application

You will need Golang installed, this application was developed using 1.11.2. There's lots of good setup guides for whatever OS you're using e.g. (for Mac)[http://sourabhbajaj.com/mac-setup/Go/README.html]

Then just run:

    go build
    ./example-go-service (omit the './' if you're on Windows)
