# Mandatory Hand-in 2 (TCP/IP Simulator in Go)

Repository for assignment 2 by the group "Cat Squish Gang".


## Answered Questions

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?


b) Does your implementation use threads or processes? Why is it not realistic to use threads?


c) How do you handle message re-ordering?

Packets can be fragmented and arrive in out-of-sequence order.
The job of TCP is that it receives the packets, buffer them and reorder them,
before presenting the data to an application.
I.e. a streaming service needs the data ordered in order to show the movie cronologically.

d) How do you handle message loss?

Transmission Control Protocol (TCP) will detect packet loss
in which it will try and perform a retransmission to make sure that messaging is reliable.
In TCP, packet loss affects the time it takes for certain items to arrive to the other connection.
I.e. in real-time applications, such as online games, packet loss will affect the quality of experience (QoE).

e) Why is the 3-way handshake important?

A 3-way handshake is important to establish a connection between client and server.
First the client will try and SYN with the server,
the server thereafter will send back a SYN/ACK in which it tries to establish connection.
Lastly the client sends back an ACK which concludes the 3-way handshake.

## Output

// Server

2022/09/26 11:33:17 server listening at 127.0.0.1:50051
2022/09/26 11:33:21 Received: 0
2022/09/26 11:33:21 Received: 1
2022/09/26 11:33:21 Received: 2

// Client

2022/09/26 11:33:21 Establishing TCP connection with server...
2022/09/26 11:33:21 handshake 0
2022/09/26 11:33:21 handshake 1
2022/09/26 11:33:21 handshake 2
2022/09/26 11:33:21 TCP handshake successfull
