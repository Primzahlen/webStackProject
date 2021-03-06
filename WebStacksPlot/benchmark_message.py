import matplotlib.pyplot as plt


x = [100,500,1000,2000]
k1 = [13,30.68,31.9,32.11]
k2 = [16.63,41.4,41.17,41.61]

plt.plot(x,k1,'s-',color = 'r',label="GRPC call")
plt.plot(x,k2,'o-',color = 'g',label="HTTP call")

plt.xlabel("Concurrence")
plt.ylabel("Latency mean(ms)")
plt.xticks(x)
plt.legend(loc = "best")
plt.show()