package pulsar

import "time"

// @Author: Feng
// @Date: 2022/3/28 14:29

// 构造pulsar的选项结构体
type pulsarOptions struct {
	url, topic           string
	consumerIntervalTime time.Duration
}

type pulsarOptionFunc interface {
	apply(*pulsarOptions)
}

type pulsarOption struct {
	f func(*pulsarOptions)
}

func (p *pulsarOption) apply(option *pulsarOptions) {
	p.f(option)
}

//newPulsarOptionFunc 新参数
func newPulsarOptionFunc(f func(options *pulsarOptions)) pulsarOptionFunc {
	return &pulsarOption{
		f: f,
	}
}

//WithUrl 带上url
func WithUrl(url string) pulsarOptionFunc {
	return newPulsarOptionFunc(func(options *pulsarOptions) {
		options.url = url
	})
}

//WithTopic 带上topic
func WithTopic(topic string) pulsarOptionFunc {
	return newPulsarOptionFunc(func(options *pulsarOptions) {
		options.topic = topic
	})
}

//WithConsumerIntervalTime 消费者守护协程重试间隔时间
func WithConsumerIntervalTime(intervalTime time.Duration) pulsarOptionFunc {
	if intervalTime < time.Second {
		intervalTime = time.Second
	}
	return newPulsarOptionFunc(func(options *pulsarOptions) {
		options.consumerIntervalTime = intervalTime
	})
}

//NewPulsarOptions 新建pulsar选项
func NewPulsarOptions(opts ...pulsarOptionFunc) *pulsarOptions {
	p := &pulsarOptions{
		consumerIntervalTime: 5 * time.Second,
	}
	// 应用传入参数
	for _, opt := range opts {
		opt.apply(p)
	}
	return p
}
