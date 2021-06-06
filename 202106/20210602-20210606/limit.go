package _0210602_20210606

import (
	"math"
	"sync"
	"time"
)

//计数器限流
//在一段时间间隔内，对请求进行计数，与阀值进行比较判断是否需要限流，一旦到了时间临界点，将计数器清零
//如果有个需求对于某个接口 /query 每分钟最多允许访问 200 次，
//假设有个用户在第 59 秒的最后几毫秒瞬间发送 200 个请求，
//当 59 秒结束后 Counter 清零了，他在下一秒的时候又发送 200 个请求。
// 那么在 1 秒钟内这个用户发送了 2 倍的请求，这个是符合我们的设计逻辑的，
// 这也是计数器方法的设计缺陷，系统可能会承受恶意用户的大量请求，甚至击穿系统
type Counter struct {
	rate  int           //计数周期内最多允许的请求数
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //计数周期内累计收到的请求数
	lock  sync.Mutex
}

func (c *Counter) Allow() bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.count >= c.rate {
		now := time.Now()
		if now.Sub(now) >= c.cycle {
			c.reset(now)
			return true
		} else {
			return false
		}
	} else {
		c.count++
	}
}
func (c *Counter) reset(t time.Time) {
	c.begin = t
	c.count = 0
}

//漏桶
//漏桶限制的是常量流出速率（即流出速率是一个固定常量值），所以最大的速率就是出水的速率，不能出现突发流量
type LeakyBucket struct {
	rate       float64 //固定每秒出水速率
	capacity   float64 //桶的容量
	water      float64 //桶中当前水量
	lastLeakMs int64   //桶上次漏水时间戳 ms

	lock sync.Mutex
}

func (l *LeakyBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().UnixNano() / 1e6
	eclipse := float64((now - l.lastLeakMs)) * l.rate / 1000 //先执行漏水
	l.water = l.water - eclipse                              //计算剩余水量
	l.water = math.Max(0, l.water)                           //桶干了
	l.lastLeakMs = now
	if (l.water + 1) < l.capacity {
		// 尝试加水,并且水还未满
		l.water++
		return true
	} else {
		// 水满，拒绝加水
		return false
	}
}

func (l *LeakyBucket) Set(r, c float64) {
	l.rate = r
	l.capacity = c
	l.water = 0
	l.lastLeakMs = time.Now().UnixNano() / 1e6
}

// 令牌桶
// 我们有一个固定的桶，桶里存放着令牌（token）。一开始桶是空的，系统按固定的时间（rate）往桶里添加令牌，
// 直到桶里的令牌数满，多余的请求会被丢弃。
// 当请求来的时候，从桶里移除一个令牌，如果桶是空的则拒绝请求或者阻塞
type TokenBucket struct {
	rate         int64 //固定的token放入速率, r/s
	capacity     int64 //桶的容量
	tokens       int64 //桶中当前token数量
	lastTokenSec int64 //桶上次放token的时间戳 s

	lock sync.Mutex
}

func (l *TokenBucket) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now().Unix()
	l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate // 先添加令牌
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastTokenSec = now
	if l.tokens > 0 {
		// 还有令牌，领取令牌
		l.tokens--
		return true
	} else {
		// 没有令牌,则拒绝
		return false
	}
}

func (l *TokenBucket) Set(r, c int64) {
	l.rate = r
	l.capacity = c
	l.tokens = 0
	l.lastTokenSec = time.Now().Unix()
}
