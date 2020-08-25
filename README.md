## QPS 
(1) Query QPS 

![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/QPS_query.jpg)

(2) Send fix size message 

![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/QPS_message.jpg)
## Latency 
(1) Query Latency 

![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/Latency_mean_query.jpg)
(2) Send fix size message latency 

![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/Latency_mean_message.jpg)

## table
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/table.png)
## grpc server
Server implement by grpc with grom
## grpc v1 client
Client implement by grpc, it has two request methods.
wrk -t8 -c100 -d30s http://127.0.0.1:8080/login_orm
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p3.png)
wrk -t8 -c100 -d30s http://127.0.0.1:8080/login_sql
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p4.png)
wrk -t8 -c100 -d30s http://127.0.0.1:8080/send_message
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p5.png)
## http v1 server
Server implement by http protol
## http client
Client implement by http protol
wrk -t8 -c 100 -d30s http://127.0.0.1:8080/login_orm
wrk -t8 -c 100 -d30s http://127.0.0.1:8080/login_sql
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p1.png)
wrk -t8 -c100 -d30s http://127.0.0.1:8080/send_message
![avatar](https://github.com/Primzahlen/webStackProject/blob/master/benchmark/p2.png)
