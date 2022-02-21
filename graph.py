import matplotlib.pyplot as plt
import sys


dir1 = sys.argv[1]
dir2 = sys.argv[2]

with open(dir1, "r") as r:
    data1 = [float(i) for i in r.readlines()]

with open(dir2, "r") as r:
    data2 = [float(i) for i in r.readlines()]

print(len(data1))
plt.plot(range(100),data1,data2)
plt.show()
plt.savefig("out.png")