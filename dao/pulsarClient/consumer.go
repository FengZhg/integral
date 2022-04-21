package pulsarClient

import (
	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
	"time"
)

// @Author: Feng
// @Date: 2022/3/26 16:55

//startConsumer 启动消费者守护协程
func (p *pulsarConfig) startConsumer() {
	for {
		log.Infof("startConsumer Daemon Do Pulsar url:%v  topic:%v", p.option.url, p.option.topic)
		err := p.consumerEventLoop()
		if err != nil {
			log.Errorf("Consumer Daemon Quit With Error err:%v", err)
		}
		time.Sleep(p.option.consumerRestartInterval)
	}
}

//consumerEventLoop 启动消费者
func (p *pulsarConfig) consumerEventLoop() error {
	// 订阅
	consumer, err := p.client.Subscribe(pulsar.ConsumerOptions{
		Topic:            p.option.topic,
		SubscriptionName: "integral-normal",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Errorf("Pulsar Create Reader Error err:%v", err)
		return err
	}
	defer consumer.Close()

	// 控制消费速度
	rl := ratelimit.New(p.option.consumeRateLimit)

	// 阻塞等待管道
	for msg := range consumer.Chan() {
		rl.Take()
		// 写数据库
		err := p.consumeMsgCB(msg)
		if err != nil {
			log.Errorf("Process Message Error %v", err)
		}
		consumer.Ack(msg)
	}
	return nil
}
