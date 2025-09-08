package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 接收一个整数指针，并将其指向的值增加10
func increaseBy10(ptr *int) {
	*ptr += 10 // 通过指针修改原始值
}

// 一个函数 接收整数切片的指针
func sliceMethodDouble(sliceInt *[]int) {
	for i := range *sliceInt {
		(*sliceInt)[i] *= 2
	}
}

// Task 定义任务类型
type Task func()

// TaskResult 存储任务执行结果
type TaskResult struct {
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Success   bool
	Error     error
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks     []Task
	results   []TaskResult
	wg        sync.WaitGroup
	mu        sync.Mutex
	startTime time.Time
	endTime   time.Time
}

// NewScheduler 创建新的调度器实例
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务到调度器
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Run 并发执行所有任务
func (s *Scheduler) Run() {
	s.startTime = time.Now()

	// 为每个任务启动一个协程
	for i, task := range s.tasks {
		s.wg.Add(1)
		go s.executeTask(i, task)
	}

	// 等待所有任务完成
	s.wg.Wait()
	s.endTime = time.Now()
}

// executeTask 执行单个任务并记录结果
func (s *Scheduler) executeTask(id int, task Task) {
	defer s.wg.Done()

	// 创建任务结果记录
	result := TaskResult{
		TaskID:    id,
		StartTime: time.Now(),
	}

	// 执行任务并捕获可能的异常
	defer func() {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)

		// 使用互斥锁保护共享数据
		s.mu.Lock()
		s.results = append(s.results, result)
		s.mu.Unlock()
	}()

	// 执行任务
	defer func() {
		if r := recover(); r != nil {
			result.Success = false
			result.Error = fmt.Errorf("任务执行出现异常: %v", r)
		}
	}()

	task()
	result.Success = true
}

// PrintResults 打印任务执行结果
func (s *Scheduler) PrintResults() {
	fmt.Printf("\n=== 任务执行结果汇总 ===\n")
	fmt.Printf("总任务数: %d\n", len(s.tasks))
	fmt.Printf("总执行时间: %v\n", s.endTime.Sub(s.startTime))
	fmt.Printf("开始时间: %v\n", s.startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("结束时间: %v\n", s.endTime.Format("2006-01-02 15:04:05"))

	fmt.Println("\n--- 各任务详情 ---")
	for _, result := range s.results {
		status := "成功"
		if !result.Success {
			status = "失败"
		}

		fmt.Printf("任务 %d: %s, 耗时: %v", result.TaskID, status, result.Duration)
		if result.Error != nil {
			fmt.Printf(", 错误: %v", result.Error)
		}
		fmt.Println()
	}
}

// 示例任务函数
func exampleTask1() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("任务1执行完成")
}

func exampleTask2() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("任务2执行完成")
}

func exampleTask3() {
	time.Sleep(700 * time.Millisecond)
	fmt.Println("任务3执行完成")
}

func exampleTaskWithError() {
	time.Sleep(200 * time.Millisecond)
	panic("模拟任务执行出错")
}

// 定义一个接口
type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体定义
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 辅助函数：打印形状的信息
func printShapeInfo(s shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

//输出员工的信息

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Employee 结构体，组合了Person
type Employee struct {
	Person     // 匿名字段，组合Person
	EmployeeID string
	Department string
	Position   string
}

// PrintInfo 方法输出员工信息
func (e Employee) PrintInfo() {
	fmt.Println("=== 员工信息 ===")
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("工号: %s\n", e.EmployeeID)
	fmt.Printf("部门: %s\n", e.Department)
	fmt.Printf("职位: %s\n", e.Position)
	fmt.Println("==============")
}

// 可选：为Employee添加一个晋升方法
func (e *Employee) Promote(newPosition string) {
	fmt.Printf("%s 从 %s 晋升为 %s\n", e.Name, e.Position, newPosition)
	e.Position = newPosition
}

/** 锁机制  并发线程------------------------**/
// Counter 结构体，包含一个互斥锁和一个计数器
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment 方法使用互斥锁保护计数器递增操作
func (c *Counter) Increment() {
	c.mu.Lock()         // 获取锁
	defer c.mu.Unlock() // 确保在函数返回时释放锁
	c.count++           // 递增计数器
}

// Value 方法返回当前计数器的值
func (c *Counter) Value() int {
	c.mu.Lock()         // 获取锁
	defer c.mu.Unlock() // 确保在函数返回时释放锁
	return c.count      // 返回计数器值
}

func main() {
	// 声明一个整数变量
	/*num := 5
	fmt.Println("调用函数前的值:", num)

	// 将变量的指针传递给函数
	increaseBy10(&num)

	// 输出修改后的值
	fmt.Println("调用函数后的值:", num)

	// 额外演示：展示指针地址和值的变化
	fmt.Println("\n--- 指针细节演示 ---")
	fmt.Printf("变量地址: %p\n", &num)
	fmt.Printf("当前值: %d\n", num)*/

	//声明一个指针类型
	/*	slicenum := []int{1, 2, 3}
		fmt.Println("初始化切片", slicenum)
		sliceMethodDouble(&slicenum)
		fmt.Println("经过函数运算后的值", slicenum)
		// 额外演示：展示切片长度和容量的变化
		fmt.Println("\n--- 切片详细信息 ---")
		fmt.Printf("长度: %d, 容量: %d\n", len(slicenum), cap(slicenum))
		fmt.Printf("切片地址: %p\n", &slicenum)
		fmt.Printf("底层数组地址: %p\n", slicenum)*/

	/*fmt.Println("=== 矩形 ===")
	rect := Rectangle{Width: 5, Height: 3}
	printShapeInfo(rect)
	fmt.Printf("矩形详细信息: 宽度=%.2f, 高度=%.2f\n", rect.Width, rect.Height)

	fmt.Println("\n=== 圆形 ===")
	circle := Circle{Radius: 4}
	printShapeInfo(circle)
	fmt.Printf("圆形详细信息: 半径=%.2f\n", circle.Radius)

	fmt.Println("\n=== 使用接口切片 ===")
	// 创建一个Shape接口切片
	shapes := []shape{rect, circle}

	for i, shape := range shapes {
		fmt.Printf("形状 %d: ", i+1)
		printShapeInfo(shape)
	}
	*/
	// 使用WaitGroup等待两个协程完成
	/*	var wg sync.WaitGroup
		wg.Add(2) // 等待两个协程

		fmt.Println("开始打印奇数和偶数...")

		// 启动打印奇数的协程
		go func() {
			defer wg.Done() // 协程结束时通知WaitGroup
			for i := 1; i <= 10; i += 2 {
				fmt.Printf("奇数: %d\n", i)
				time.Sleep(100 * time.Millisecond) // 稍微延迟，使输出更明显
			}
		}()

		// 启动打印偶数的协程
		go func() {
			defer wg.Done() // 协程结束时通知WaitGroup
			for i := 2; i <= 10; i += 2 {
				fmt.Printf("偶数: %d\n", i)
				time.Sleep(100 * time.Millisecond) // 稍微延迟，使输出更明显
			}
		}()

		// 等待所有协程完成
		wg.Wait()
		fmt.Println("打印完成!")
	*/

	// 创建调度器
	/*scheduler := NewScheduler()

	// 添加任务
	scheduler.AddTask(exampleTask1)
	scheduler.AddTask(exampleTask2)
	scheduler.AddTask(exampleTask3)
	scheduler.AddTask(exampleTaskWithError)

	fmt.Println("开始执行任务...")

	// 执行所有任务
	scheduler.Run()

	// 打印结果
	scheduler.PrintResults()*/

	// 创建一个Employee实例
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "E1001",
		Department: "技术部",
		Position:   "高级工程师",
	}

	// 调用PrintInfo方法输出员工信息
	emp.PrintInfo()

	// 演示晋升方法
	emp.Promote("技术主管")

	// 再次打印信息查看变化
	emp.PrintInfo()

	// 创建另一个员工实例
	emp2 := Employee{
		Person: Person{
			Name: "李四",
			Age:  25,
		},
		EmployeeID: "E1002",
		Department: "市场部",
		Position:   "市场专员",
	}

	emp2.PrintInfo()

	// 演示直接访问嵌入字段
	fmt.Printf("员工 %s 的年龄是 %d 岁\n", emp2.Name, emp2.Age)

	/**-------------------------------------------通道相关的习题=----------------------------------------------**/
	// 创建一个整数通道
	ch := make(chan int)

	// 启动发送协程
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i // 将整数i发送到通道
		}
		close(ch) // 发送完成后关闭通道
	}()

	// 在主协程中接收并打印数据
	for num := range ch {
		fmt.Println(num)
	}

	/** --------------------------------------------通道的接收与输出--------------------------------**/
	// 创建一个缓冲大小为10的整数通道
	/*	ch1 := make(chan int, 10)

		// 使用WaitGroup等待两个协程完成
		var wg sync.WaitGroup
		wg.Add(2) // 等待两个协程完成

		// 生产者协程：发送100个整数到通道
		go func() {
			defer wg.Done()  // 协程结束时通知WaitGroup
			defer close(ch1) // 发送完成后关闭通道

			for i := 1; i <= 100; i++ {
				ch1 <- i
				fmt.Printf("生产者发送: %d\n", i)
			}
			fmt.Println("生产者完成")
		}()

		// 消费者协程：从通道接收并打印整数
		go func() {
			defer wg.Done() // 协程结束时通知WaitGroup

			for num := range ch1 {
				fmt.Printf("消费者接收: %d\n", num)
			}
			fmt.Println("消费者完成")
		}()

		// 等待两个协程完成
		wg.Wait()
		fmt.Println("程序结束")
	*/
	/**----------------------------  ***/
	// 创建计数器实例
	/*var counter Counter

	// 使用WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加WaitGroup计数器

		go func(id int) {
			defer wg.Done() // 协程完成时减少WaitGroup计数器

			// 每个协程对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}

			fmt.Printf("协程 %d 完成\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终的计数器值
	fmt.Printf("最终计数器值: %d\n", counter.Value())*/

	/** --------------------------------------------------------**/
	// 使用int64类型的原子计数器
	var counter int64

	// 使用WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加WaitGroup计数器

		go func(id int) {
			defer wg.Done() // 协程完成时减少WaitGroup计数器

			// 每个协程对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				// 使用原子操作递增计数器
				atomic.AddInt64(&counter, 1)
			}

			fmt.Printf("协程 %d 完成\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 使用原子操作读取计数器的值
	finalValue := atomic.LoadInt64(&counter)
	fmt.Printf("最终计数器值: %d\n", finalValue)
}
