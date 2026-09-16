package time

//说明，这个是看B站的这个博主：https://www.bilibili.com/video/BV18M411C7pY/?spm_id_from=333.1391.0.0&vd_source=c20c79c492d5d1dba8eee9983daff245，写的
import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	C := make(chan int)
	go func() {
		fmt.Println("提交订单。。。")
		// t.Log("提交订单。。。")
		time.Sleep(10 * time.Second)
		fmt.Println("开始支付")
		// t.Log("开始支付")
		C <- 1
	}()
loop:
	for {
		select {
		case <-C:
			fmt.Println("支付成功")
			break loop
		case <-time.After(5 * time.Second):
			fmt.Println("订单超时，取消订单...")
			break loop
		}
	}
}

func TestChannel(t *testing.T) {
	C := make(chan int, 8)
	// defer close(C)
	for i := 0; i < 8; i++ {
		C <- i
	}
	close(C)
	fmt.Println(C)
	t.Logf("len=%d cap=%d", len(C), cap(C))
	TesatChannel(t)
}

func TesatChannel(t *testing.T) {
	C := make(chan string, 8)
	// defer close(C)
	for _, v := range []string{"a", "b", "c", "d", "e", "f", "g", "h"} {
		C <- v
	}
	close(C)
	fmt.Println(C)
	t.Logf("len=%d cap=%d", len(C), cap(C))
}

type Cinema struct {
	rw *sync.RWMutex
}

type Person struct {
	Name string
}

type Admin struct {
	Name string
}

func (p *Person) SitDown(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s观众坐下\n", p.Name)
}
func (p *Person) StartWatch() {
	fmt.Printf("%s开始观看\n", p.Name)
	i := time.Duration(rand.Intn(5) + 1)
	time.Sleep(i * time.Second)
	fmt.Printf("%s 观看结束\n", p.Name)
}
func (p *Person) WatchMovie(c *Cinema, wg *sync.WaitGroup, b chan int) {
	c.rw.RLock()
	// 11先坐下
	p.SitDown(wg)
	<-b
	// 看电影
	p.StartWatch()
	c.rw.RUnlock()
}
func (a *Admin) ChangeMovie(c *Cinema, wg *sync.WaitGroup) {
	defer wg.Done()
	c.rw.Lock()
	// for Ic.rw.TryLock(){
	// 11
	// fmt.Printf("还有人看电影n")
	// 11
	// time.Sleep(1 * time.Second)
	// 11了
	fmt.Printf("所有观众看完电影，%s管理员切换影片\n", a.Name)
	c.rw.Unlock()
}

// 电影院
func example1() {
	c := &Cinema{
		rw: &sync.RWMutex{},
	}
	a := &Admin{
		Name: "强人",
	}

	b := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int, c *Cinema, wg *sync.WaitGroup, b chan int) {
			p := Person{
				Name: fmt.Sprintf("%d号观众", i),
			}
			p.WatchMovie(c, wg, b)

		}(i, c, &wg, b)
	}
	wg.Wait()
	fmt.Printf("所有观众入座，开始播放电影\n")
	close(b)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		a.ChangeMovie(c, wg)
	}(&wg)
	wg.Wait()
}

type Desk struct {
	rw     *sync.RWMutex
	Snacks int
}

func example2() {
	d := Desk{
		rw: &sync.RWMutex{},
	}
	wg := sync.WaitGroup{}
	wg.Add(12)

	go d.PlaceSnacks(&wg)
	for i := 0; i < 11; i++ {
		p := Person{
			Name: fmt.Sprintf("%d", i),
		}
		go p.GetSnacks(&d, &wg)
	}

	wg.Wait()
}

func (d *Desk) PlaceSnacks(wg *sync.WaitGroup) {
	defer wg.Done()
	d.rw.Lock()
	d.Snacks = 10
	time.Sleep(3 * time.Second)
	fmt.Printf("零食准备完毕...\n")
	d.rw.Unlock()
}

func (p *Person) GetSnacks(d *Desk, wg *sync.WaitGroup) {
	defer wg.Done()

	timeout := time.After(6 * time.Second) // 最多等待5秒
	for {
		select {
		case <-timeout:
			fmt.Printf("%s 等待超时，放弃抢零食\n", p.Name)
			return
		default:
			if d.rw.TryLock() {
				// 成功获取写锁，继续执行
				defer d.rw.Unlock()
				if d.Snacks > 0 {
					fmt.Printf("%s抢到零食，开心!\n", p.Name)
					d.Snacks--
				} else {
					fmt.Printf("%s没抢到零食，难受!\n", p.Name)
				}
				return
			}
			// 没获取到锁，短暂等待后再试
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func TestWMutex(t *testing.T) {
	example2()
}
func (p *Person) GetSnacksError(d *Desk, wg *sync.WaitGroup) {
	for !d.rw.TryLock() {
		fmt.Printf("桌上没零食，%s望眼欲穿\n", p.Name)
		time.Sleep(2 * time.Second)
	}

	d.rw.RUnlock()
	defer wg.Done()
	defer d.rw.Unlock()
	d.rw.Lock()
	if d.Snacks > 0 {
		fmt.Printf("%s抢到零食，开心!\n", p.Name)
		d.Snacks--
		return
	}
	fmt.Printf("%s没抢到零食，难受!\n", p.Name)
}
