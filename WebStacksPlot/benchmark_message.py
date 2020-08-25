import matplotlib.pyplot as plt

#折线图
x = [100,500,1000,2000]#点的横坐标
k1 = [13,30.68,31.9,32.11]#线1的纵坐标
k2 = [16.63,41.4,41.17,41.61]#线2的纵坐标

plt.plot(x,k1,'s-',color = 'r',label="GRPC call")#s-:方形
plt.plot(x,k2,'o-',color = 'g',label="HTTP call")#o-:圆形

plt.xlabel("Concurrence")  #横坐标名字
plt.ylabel("Latency mean(ms)")  #纵坐标名字
plt.xticks(x)  #x轴刻度
plt.legend(loc = "best")#图例
plt.show()