import matplotlib.pyplot as plt

#折线图
x = [100,500,1000,2000]#点的横坐标
k1 = [2088.95,1978.9,1959.99,1862.65]#线1的纵坐标
k2 = [2486.62,2348.25,2276.49,2217.46]#线2的纵坐标
k3 = [1931.34,1992.8,1897.38,1859.62]
k4 = [2196.88,2124.59,2148.29,2202.77]
plt.plot(x,k1,'s-',color = 'r',label="GRPC_ORM")#s-:方形
plt.plot(x,k2,'o-',color = 'g',label="GRPC_SQL")#o-:圆形
plt.plot(x,k3,'g+-',color = 'b',label="HTTP_ORM")#o-:圆形
plt.plot(x,k4,'b^-',color = 'y',label="HTTP_SQL")#o-:圆形
plt.xlabel("Concurrence")  #横坐标名字
plt.ylabel("QPS(Requests/sec)")  #纵坐标名字
plt.xticks(x)  #x轴刻度
plt.legend(loc = "best")#图例
plt.show()