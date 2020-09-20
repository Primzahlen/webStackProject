import matplotlib.pyplot as plt


x = [100,500,1000,2000]
k1 = [47.68,125.3,124.24,127.71]
k2 = [40.08,105.61,106.65,107.65]
k3 = [51.75,125.27,131.44,133.87]
k4 = [45.47,117.71,116.16,113.3]
plt.plot(x,k1,'s-',color = 'r',label="GRPC_ORM")
plt.plot(x,k2,'o-',color = 'g',label="GRPC_SQL")
plt.plot(x,k3,'g+-',color = 'b',label="HTTP_ORM")
plt.plot(x,k4,'b^-',color = 'y',label="HTTP_SQL")
plt.xlabel("Concurrence")
plt.ylabel("Latency mean(ms)")
plt.xticks(x)
plt.legend(loc = "best")
plt.show()