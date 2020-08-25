#以下代码在notebook中执行
import matplotlib.pyplot as plt

#折线图
x = [100,500,1000,2000]#点的横坐标
k1 = [0.8558,0.9101,0.8884,0.8963] #线1的纵坐标
# k2 = [9.56,30.7,37.12,58.87]#线2的纵坐标
dy=0.8
# plt.plot(x,k1,'s-',color = 'r',label="GRPC_ORM")#s-:方形
plt.errorbar(x, k1, yerr=dy, fmt='.k')
# plt.errorbar(x, k2, yerr=dy, fmt='o')
# plt.plot(x,k2,'o-',color = 'g',label="GRPC_SQL")#o-:圆形

plt.xlabel("Concurrence")  #横坐标名字
plt.ylabel("Latency mean(ms)")  #纵坐标名字
plt.xticks(x)  #x轴刻度
plt.legend(loc = "best")#图例
plt.show()
#fmt是一种控制线条和点的外观的代码格式,语法与plt.plot的缩写代码一样.