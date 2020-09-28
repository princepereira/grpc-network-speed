# grpc-network-speed
This will help to calculate the network speed in grpc

# How to build docker image
$ make client
$ make server
$ sudo docker save grpcclient:latest -o /tmp/dual_home/grpcclient.tar          
$ sudo docker save grpcserver:latest -o /tmp/dual_home/grpcserver.tar 

# How to start server and client

// At server
$ sudo docker load -i grpcserver.tar
$ sudo docker run -p 2000:2000 -it grpcserver:latest

// At client
$ sudo docker load -i grpcclient.tar
$ sudo docker run -it grpcclient:latest
