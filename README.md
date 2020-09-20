## QPS 
(1) Query QPS 

According to the experimental results, the number of connections on the scale of [200,500,1000,2000] was tested in one thread respectively, with a duration of 10 seconds. QPS is the number of requests processed by the server per second. It is clear that GRPC-SQL has the best experimental results and the best performance. The second is the performance of HTTP-SQL, which maintains good stability under different high concurrency and is similar to GRPC-SQL at 2,000 connections. Using a combination of ORM frameworks, the performance is not as good as the original SQL, and the performance of both GRPC and HTTP protocols is close. 

![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/QPS_query.jpg)

(2) Send fixed size message 
 
The client sends a packet of fixed size (600byte) to the server, which updates the status field after receiving it: OK, and returns the data to the client through re-encapsulation. According to the experimental results, the combination of GRPC protocol is significantly superior to the combination of HTTP protocol. Using GRPC has obvious advantages in transmitting larger packets, but performance is slightly reduced as the number of connections increases. Although the combination of HTTP protocol is lagging behind in QPS value, with the increase of the number of connections, the performance does not fluctuate greatly and is relatively stable. 

![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p6.png) 

## Latency 
(1) Query Latency 

In Latency data, the number of connections with a scale of [200,500,1000,2000] was tested under 1 thread respectively, with a duration of 10 seconds. Latency is the average response time, which performs best when grPC-SQL combination is used, and the average response time of HTTP-SQL is also close to that of GRPC-SQL group. The Latency was higher in the remaining two groups which used ORM as the database connection mode. 

![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/Latency_mean_query.jpg) 

(2) Send fixed size message latency 
 
The client sends a packet of fixed size (600byte) to the server, which updates the status field after receiving it: OK, and returns the data to the client through re-encapsulation. According to the experimental results, GRPC protocol has the lowest average response time, which is consistent with the previous experimental results of QPS. HTTP is not as efficient at transmitting packets as GRPC protocol. 

![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/Latency_mean_message.jpg)

## table
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/table.png)
## grpc server
Server implement by grpc with grom
## grpc client
Client implement by grpc, it has two request methods.
wrk -t8 -c100 -d30s http://127.0.0.1:8080/login_orm
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p3.png)
wrk -t8 -c100 -d30s http://127.0.0.1:8080/login_sql
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p4.png)
wrk -t8 -c100 -d30s http://127.0.0.1:8080/send_message
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p5.png)
## http server
Server implement by http protol
## http client
Client implement by http protol
wrk -t8 -c 100 -d30s http://127.0.0.1:8080/login_orm
wrk -t8 -c 100 -d30s http://127.0.0.1:8080/login_sql
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p1.png)
wrk -t8 -c100 -d30s http://127.0.0.1:8080/send_message
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p2.png)
