import matplotlib.pyplot as plt


x = [100,500,1000,2000]
k1 = [2088.95,1978.9,1959.99,1862.65]
k2 = [2486.62,2348.25,2276.49,2217.46]
k3 = [1931.34,1992.8,1897.38,1859.62]
k4 = [2196.88,2124.59,2148.29,2202.77]
plt.plot(x,k1,'s-',color = 'r',label="GRPC_ORM")
plt.plot(x,k2,'o-',color = 'g',label="GRPC_SQL")
plt.plot(x,k3,'g+-',color = 'b',label="HTTP_ORM")
plt.plot(x,k4,'b^-',color = 'y',label="HTTP_SQL")
plt.xlabel("Concurrence")
plt.ylabel("QPS(Requests/sec)")
plt.xticks(x)
plt.legend(loc = "best")
plt.show()