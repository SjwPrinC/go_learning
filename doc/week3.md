###	goroutine

select {} 可以永远阻塞

keep yourself busy or do the work yourself

leave concurrency to the caller

注意一些代码的产生的二义性：如返回空是结束了还是报错了

```go
func ListDir(dir string) ([]string, error)  //同步调用获取目录的文件列表
func ListDir(dir string) chan string     //通过chan异步获取文件列表，读到一个取一个，Q：不知道什么时候读完，也没有报错

//比较好的处理方式,设置回调
func ListDir(dir string, fn func(string))  // filepath.WalkDir

```

goroutine泄露情况：

1. channel阻塞(读写都会存在)，外部没有消费channel的地方了
2. 外部、调用方无法掌握何时stop goroutine

goroutine 生命周期管理（Never start a goroutine without knowing when it will stop）：

1. 外部需要知道这个goroutine什么时候结束

2. 外部有没有办法能够让他结束，传入stop chan 由外部发送信号进行停止

   

log.Fatal(err)会调用os.Exit(), 一旦出现，整个程序退出，defer不会执行

Fatal 最好只能在init()和mian()函数调用，在后台服务端会出现进程结束的问题

channal 安全的关闭方式：应该由写入方去关闭，因为往关闭的channel中写数据会发生panic



### memory model

reference：https://golang.org/ref/mem

Happens-Before

```go
x:=0
for i in range(100){
    x = 1
    fmt.Print(x)
}

//经过编译器重排，可能会优化成这样的代码逻辑
x:=1
for i in range(100){
    fmt.Print(x)
}
```

编译器重排，内存重排

CPU级别的锁支持，barrier、fence指令，要求在访问内存的时候都要把CPU cache上的值拷贝到memory

atomic 的cas，WaitGroup等同步方法就是基于这个原理的实现

![image-20210728212634290](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20210728212634290.png)



single machine word ：在x64机器上做8个字节的拷贝是原子的，指针的长度正好是8个字节，所以指针的赋值是原子的

并发的两个条件：原子性和可见性，原子性是指某一个操作是CPU指令执行的单元，可见性是指操作完以后其他人都能看得见，指针的赋值满足原子性，但是不一定满足可见性，因为在CPU Cache Line上

slice的赋值不是原子的，因为他的占用内存长度不止8位，包括对interface{}的赋值

cow （copy on write）写时拷贝：新老数据进行拷贝，确保修改前后的值都有

atomic.Value  atomic.Store 来拷贝对象，这时是原子的。

MESI：多核cpu通知其他cpu更新cache line 的过期的值

https://blog.csdn.net/muxiqingyang/article/details/6615199



### sync

Don't communicate by share memory, share memory by communicating

chan 适合做任务分发

data race 

i++的操作不是原子的，汇编变成三条指令

interface的赋值不是原子的,他的底层实现是有两个指针

```go
//interface的赋值不是原子的,因为他内部是两个指针,在赋值的时候会出现Type是A，Data是B
type interface struct{
    Type uintptr
    Data uintptr
}
```

slice的赋值也不是原子的，它是由sliceHeader，cap，len字段，所以也可能不会原子赋值

hmap，map是原子赋值的

Dave：要么有data race ，要么没有data race ，没有安全的data race

要么伟大的生，要么壮烈的死，没有苟且偷生



atomic.Value既满足原子性又满足可见性

mutex实现原理

等待锁的队列FIFO

![image-20210729204658972](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20210729204658972.png)



errgroup：管理并行的工作流（goroutine管理），做错误处理

sync.Pool

sync.Once



kratos

### Q

COW

mesi

map，slice，interface底层实现

atomic.Value 

redis BGSave   cow实现

```go
func main() {
	s := make([]int, 0, 10)
	s = append(s, 12)

	ChangeSliceValue(s)

	fmt.Println(s)
	fmt.Printf("outer s: %p\n", s)
}

func ChangeSliceValue(ss []int) {
	for i := 1; i < 5; i++ {
		ss = append(ss, i)
	}

	fmt.Println(ss)
	fmt.Printf("inner s: %p\n", s)
}
//输出结果是outer s [12]  inner s [12,1,2,3,4],两个s的地址是一样的
//这样的结果说明slice函数传参是指针传值，内外两个切片指向的第一个值恰好是同一个内存地址
```

map 内部是用Hmap实现的，C++map是用红黑树实现的， hmap里面有bucket

interface内部实现reference https://studygolang.com/articles/18941?fr=sidebar