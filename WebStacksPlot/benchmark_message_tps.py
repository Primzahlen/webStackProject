import matplotlib.pyplot as plt


x = [100,500,1000,2000]
k1 = [7867.05,8092.85,7711.7,7565.4]
k2 = [6020.37,6044.66,6074.21,6007.59]

plt.plot(x,k1,'s-',color = 'r',label="GRPC call")
plt.plot(x,k2,'o-',color = 'g',label="HTTP call")

plt.xlabel("Concurrence")
plt.ylabel("TPS(Requests/sec)")
plt.xticks(x)
plt.legend(loc = "best")
plt.show()