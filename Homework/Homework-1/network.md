# network

## OSI 七层模型

![](./images/osi.gif)

### 物理层

实现相邻计算机节点之间**比特流**的透明传送，尽可能屏蔽掉具体传输介质和物理设备的差异。使其上面的数据链路层不必考虑网络的具体传输介质是什么。

### 数据链路层

通过各种控制协议，将有差错的物理信道变为无差错的、能可靠传输**数据帧**的数据链路。

### 网络层

数据链路层的数据在这一层被转换为**数据包**，然后通过路径选择、分段组合、顺序、进/出路由等控制，将信息从一个网络设备传送到另一个网络设备。

一般地，数据链路层是解决同一网络内节点之间的通信，而网络层主要解决不同子网间的通信。

### 传输层

OSI下3层的主要任务是数据通信，上3层的任务是数据处理。而传输层是OSI模型的第4层。因此该层是通信子网和资源子网的接口和桥梁，起到承上启下的作用。

传输层负责提供两节点之间数据的可靠传送，当两节点的联系确定之后，传输层则负责监督工作。

### 会话层

向两个实体的表示层提供建立和使用连接的方法。将不同实体之间的表示层的连接称为会话。

会话层的任务就是组织和协调两个会话进程之间的通信，并对数据交换进行管理。

### 表示层

对来自应用层的命令和数据进行解释，对各种语法赋予相应的含义，并按照一定的格式传送给会话层。如数据格式处理、数据的编码、压缩和解压缩以及数据的加密和解密。

### 应用层

计算机用户，以及各种应用程序和网络之间的接口，其功能是直接向用户提供服务，完成用户希望在网络上完成的各种工作。

## 路由器参数

### CPU

路由器 CPU 的发展大致经历如下四个阶段：

#### 通用处理器阶段

上个世纪60年代，人们曾经使用普通电脑充当路由器的角色，但通用处理器不是面向网络通信需要特殊设计的，所以转发速度一般相对较慢，可扩展性差，很难满足网络的需求。

#### 嵌入式处理器阶段

嵌入式微处理器多数工作在设备制造商自己设计的系统中，是面向应用的处理器。

第一代的路由器是基于嵌入式微处理器的嵌入式系统，有专门的电路、接口及操作系统，是一台专门的设备，已经不再是基于通用微处理器、通用接口、通用操作系统的PC了。

#### ASIC（专用集成电路）处理器阶段

当网络速度比较慢时，嵌入式处理器的路由及转发的处理速度完全赶得上数据流，后来，线路带宽变宽，数据速率变快，基于“CPU+操作系统”的路由器逐渐成了网络的瓶颈。

在这种情况下，网络设备开始采用ASIC技术。它通过把指令集或计算逻辑固化到芯片中，它把转发过程的所有细节全部采用硬件方式来实现，因而可以获得很高的处理速度。

#### 网络处理器阶段

ASIC的优点也是它的缺点，就是缺乏灵活性。一旦指令或计算逻辑固化到芯片硬件中，就很难修改升级，要增加新的功能或提高性能，就得重新设计芯片。另外，设计和制造复杂的ASIC一般需要花费周期长，研发费用较高。除此之外，当前网络的应用范围在不断扩大、新的业务不断涌现，网络的发展也不仅仅是带宽的不断提高，而更多地表现为对“智能化处理”的要求。

网络处理器较之ASIC最大的优势是灵活，开发周期相对较短。它是为优化包处理而设计的，它能把数据包以线速送到下一个节点，另外，如果需要新的功能或新的标准，设备制造商能通过给网络处理器编程来实现，以满足各种新的网络应用。

### 内存

路由器中可能有多种内存，例如Flash（闪存）、DRAM（动态内存）等，内存用作存储配置、路由器操作系统、路由协议软件等内容。在中低端路由器中，路由表可能存储在内存中。

闪存（Flash）是可读可写的存储器，在系统重新启动或关机之后仍能保存数据。Flash中存放着当前使用中的IOS。

RAM也是可读可写的存储器，但它存储的内容在系统重启或关机后将被清除。运行期间，RAM中包含路由表项目、ARP缓冲项目、日志项目和队列中排队等待发送的分组。除此之外，还包括运行配置文件（Running-config）、正在执行的代码、IOS操作系统程序和一些临时数据信息。

### 带宽

带宽（band width）又叫频宽，是指在固定的的时间可传输的资料数量，亦即在传输管道中可以传递数据的能力。在数字设备中，频宽通常以bps表示，即每秒可传输之位数。在模拟设备中，频宽通常以每秒传送周期或赫兹 (Hz)来表示。

### 天线增益

天线增益是指：在输入功率相等的条件下，实际天线与理想的辐射单元在空间同一点处所产生的信号的功率密度之比。它定量地描述一个天线把输入功率集中辐射的程度。

dBi是功率增益的单位，3dbi就是把相同距离上的功率扩大3倍，同理，5dbi就是把相同距离上的功率扩大5倍。从理论上来说，天线增益越大能够将无线信号传的更远。
