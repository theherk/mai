mai
===

Given a MAC address, return company name.

Installation
------------

    go get -u github.com/theherk/mai
    # Or clone.
    git clone https://github.com/theherk/mai.git
    cd mai && go install

Testing
-------

If testing or contributing best to clone outside GOPATH.

    make test

Usage
-----

    export MAI_APIKEY=[api key]
    mai [mac]...
    # Or using the docker image built with `make image`.
    docker run -ti -e MAI_APIKEY=$MAI_APIKEY mai 44:38:39:ff:ef:57

Notes
-----

- Library coverage is 87.5%; only lines not covered are std lib calls.
- Provides binary `mai` and docker implementation.
- Makefile for some handy shortcuts.

Of course many optimizations could be made. In addition to documentation notes in the code, here are some things that come to mind.

- API key as an environment variable may be insufficient in some contexts. There are many options.
- The container could be smaller. Build stage container and alpine would do that.
- Unpacking the full json response into types would be better, but only the name was requested.
- Coverage could be better, but it would be superfluous in the current state.
- This is a lot of foundation for a popsicle stick house.
    - A few lines could yield the same result.
    - But this seems like an appropriate start to something that could expand without being over-engineered.
