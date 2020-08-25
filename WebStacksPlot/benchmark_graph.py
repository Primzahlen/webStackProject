import matplotlib.pyplot as plt

#折线图
x = [100,500,1000,2000]#点的横坐标
k1 = [47.68,125.3,124.24,127.71]#线1的纵坐标
k2 = [40.08,105.61,106.65,107.65]#线2的纵坐标
k3 = [51.75,125.27,131.44,133.87]
k4 = [45.47,117.71,116.16,113.3]
plt.plot(x,k1,'s-',color = 'r',label="GRPC_ORM")#s-:方形
plt.plot(x,k2,'o-',color = 'g',label="GRPC_SQL")#o-:圆形
plt.plot(x,k3,'g+-',color = 'b',label="HTTP_ORM")#o-:圆形
plt.plot(x,k4,'b^-',color = 'y',label="HTTP_SQL")#o-:圆形
plt.xlabel("Concurrence")  #横坐标名字
plt.ylabel("Latency mean(ms)")  #纵坐标名字
plt.xticks(x)  #x轴刻度
plt.legend(loc = "best")#图例
plt.show()