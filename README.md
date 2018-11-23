# Algorithm
## 目录
[排序算法](#排序算法)  
- [冒泡排序](#冒泡排序)
- [选择排序](#选择排序)

[缓存算法](LRU)
- [LRU](#LRU)
- [Two-Queues](#two-queues)

[其他](#其他)
- [单向链表](#单向链表)
## 排序算法

### 冒泡排序

 它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果他们的顺序（如从大到小、首字母从A到Z）错误就把他们交换过来。走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素已经排序完成。
这个算法的名字由来是因为越大的元素会经由交换慢慢“浮”到数列的顶端（升序或降序排列），就如同碳酸饮料中二氧化碳的气泡最终会上浮到顶端一样，故名“冒泡排序”。

##### 算法原理

1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
2. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
3. 针对所有的元素重复以上的步骤，除了最后一个。
4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较



### 选择排序

 选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是每一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到全部待排序的数据元素排完。 选择排序是不稳定的排序方法。

## 缓存算法

### LRU

LRU全称是Least Recently Used，即最近最久未使用的意思。
LRU算法的设计原则是：如果一个数据在最近一段时间没有被访问到，那么在将来它被访问的可能性也很小。也就是说，当限定的空间已存满数据时，应当把最久没有被访问到的数据淘汰。

##### 实现

利用链表和hashmap。当需要插入新的数据项的时候，如果新数据项在链表中存在（一般称为命中），则把该节点移到链表头部，如果不存在，则新建一个节点，放到链表头部，存储后返回的节点地址作为map的value,若缓存满了，则把链表最后一个节点删除即可。在访问数据的时候，如果数据项在链表中存在，则把该节点移到链表头部，否则返回-1。这样一来在链表尾部的节点就是最近最久未访问的数据项。

### Two-Queues
Two queues算法类似于LRU-2，不同点在于2Q将LRU-2中的访问历史队列（注意这不是缓存数据的）改为一个FIFO缓存队列，即2Q算法=一个是FIFO队列，一个是LRU队列。命中率高于LRU
##### 实现
当数据第一次访问时，2Q算法将数据缓存在FIFO队列里面，当数据第二次被访问时，则将数据从FIFO队列移到LRU队列里，两个队列各自按照自己的方法淘汰数据。

1. 新访问的数据插入到FIFO队列；

2. 如果数据在FIFO队列中一直没有被再次访问，则最终按照FIFO规则淘汰；

3. 如果数据在FIFO队列中被再次访问，则将数据移到LRU队列头部；

4. 如果数据在LRU队列再次被访问，则将数据移到LRU队列头部；

5. LRU队列淘汰末尾的数据。

## 其他

### 单向链表

单向链表（单链表）是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要通过顺序读取从头部开始；链表是使用指针进行构造的列表；又称为结点列表，因为链表是由一个个结点组装起来的；其中每个结点都有指针成员变量指向列表中的下一个结点；
列表是由结点构成，head指针指向第一个成为表头结点，而终止于最后一个指向nuLL的指针。
单向链表包含两个域，一个是信息域，一个是指针域。也就是单向链表的节点被分成两部分，一部分是保存或显示关于节点的信息，第二部分存储下一个节点的地址，而最后一个节点则指向一个空值。

##### 实现
 
1. 定义节点
2. 定义链表，链表只有一个Head节点
3. 创建一个值域为0，指针域指向nil的节点作为头部节点
4. 新建节点的指针域都指向nil，循环链表，找到最后一个节点，并把指针域指向当前新建的节点（Append） 
