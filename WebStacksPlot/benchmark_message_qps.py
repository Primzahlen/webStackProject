import matplotlib.pyplot as plt

#折线图
x = [100,500,1000,2000]#点的横坐标
k1 = [7867.05,8092.85,7711.7,7565.4]#线1的纵坐标
k2 = [6020.37,6044.66,6074.21,6007.59]#线2的纵坐标

plt.plot(x,k1,'s-',color = 'r',label="GRPC call")#s-:方形
plt.plot(x,k2,'o-',color = 'g',label="HTTP call")#o-:圆形

plt.xlabel("Concurrence")  #横坐标名字
plt.ylabel("QPS(Requests/sec)")  #纵坐标名字
plt.xticks(x)  #x轴刻度
plt.legend(loc = "best")#图例
plt.show()