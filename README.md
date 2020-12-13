# mystore
There are two main paackages inside
  my_store_server
  my_store_client
both includes package redis, which includes REDIS specifications required for server and client.

# how to build
1. Create folder my_store under $GOPATH and files and direct
-->src/my_store
  -->my_store_client
    -->my_store_client.go
  -->my_store_server
    -->my_store_server.go
  -->redis
    -->redis.go
    -->redis_test.go
2. export GO111MODULE="off"
3. run `go build`and `go install` in direcory redis
4. run 'go get my_store/redis` and `go build` in directory my_store_server
5. run 'go get my_store/redis` and `go build` in directory my_store_client

# how to run the program
Following example will demostrate how tcp server and client can communicate,
store keys with values, retreive and delete. Please read the comments in the files and the notes below
to look in for implimentation details.

./my_store_client SET NAME MYNAME
dialing,  localhost:3000
->: OK

 ./my_store_client GET NAME
 dialing,  localhost:3000
->:
MYNAME


CLM-C02Y209CJG5J:my_store_client mnaduviledath$ ./my_store_client DEL NAME
dialing,  localhost:3000
->: 1

# Notes
I chose to name it as my_store since rediscl already exists in my system
Port 3000 was chosen over the defualt tcp port for no reason

Not many tests were written, how ever the encoding to redis have a test to show the framework.
Code speak for itslef, however some notes as below:-

redis package has the useful constants and methods on REDIS protocol
my_store_server uses go routine for each client to serve cuncurrent access to multiple clients

The server uses just a map declared globally to store the values and loses its values on exists.
No persistance storeage was developed since that was not specified as a requirement.

