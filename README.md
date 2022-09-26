# Mandatory Hand-In 2 (TCP/IP Simulator in Go)

Repository for assignment 2 by the group "Cat Squish Gang".

## Answered Questions

a) *What are packages in your implementation? What data structure do you use to transmit data and meta-data?*

The packages in our implementation are protocol buffers. 

protocol buffers are developed by Google as a "language-neutral, platform-neutral, extensible mechanism for serializing structured data".

Conveniently, it is useful for developing programs to communicate with each other over a network. 

-------------------------------------------

b) *Does your implementation use threads or processes? Why is it not realistic to use threads?*

Our implementation uses processes. 

The reason why we don't use threads is that it doesn't take into consideration package loss and networks delay, for example. 

Threads are rather one-dimensional and will not go through the entire network layer. 

a thread would only pass through the Application layer instead of all the 7 layers, such as application, presentation, session etc. 

If you were accurately simulating a package delivery (like we do with a process), then we will go through all the layers. 

-------------------------------------------

c) *How do you handle message re-ordering?*

Packets can be fragmented and arrive in out-of-sequence order.
The job of TCP is that it receives the packets, buffer and reorder them
before presenting the data to an application.
I.e. a streaming service needs the data ordered to show the movie chronologically.

-------------------------------------------

d) *How do you handle message loss?*

Transmission Control Protocol (TCP) will detect packet loss
in which it will try and perform a retransmission to make sure that messaging is reliable.
In TCP, packet loss affects the time it takes for certain items to arrive to the other connection.
I.e. in real-time applications, such as online games, packet loss will affect the quality of experience (QoE).

-------------------------------------------

e) *Why is the 3-way handshake important?*

A 3-way handshake is important to establish a connection between client and server:
First the client will try and SYN with the server,
the server thereafter will send back a SYN/ACK in which it tries to establish connection.
Lastly the client sends back an ACK which concludes the 3-way handshake.

## Example Output

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
