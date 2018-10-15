# seat-distributor
Server side API to distribute seats on a flight. The solution is just a concept to showcase a certain way of creating a software architecture around planes and ticket reservations.
For more information take a look at the inline documentation.

Currently the server is initalised with 2 developer plane.

## Tools
#### Docker
You can start the application with the following command: `make docker-run`

It creates a new docker image then starts it on 8080
#### Makefile
You can create binaries for your favourite platform with the following commands:
* make osx
* make linux
* make windows

#### API
You can find the swagger yaml for the API [here](./swagger/swagger.yml)
You can insert the file here: https://editor.swagger.io

