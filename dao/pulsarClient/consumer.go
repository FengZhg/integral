package pulsarClient

import (
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
	"time"
)

// @Author: Feng
// @Date: 2022/3/26 16:55

type pulsarConsumerDaemon struct {
	option          *pulsarOptions
	procMsgCallback func(pulsar.ConsumerMessage) error
}

//NewPulsarConsumerDaemon 新pulsar消费者守护协程
func NewPulsarConsumerDaemon(option *pulsarOptions,
	callback func(pulsar.ConsumerMessage) error) *pulsarConsumerDaemon {
	return &pulsarConsumerDaemon{
		option:          option,
		procMsgCallback: callback,
	}
}

//Start 启动消费者守护协程
func (p *pulsarConsumerDaemon) Start() {
	for {
		log.Infof("Start Daemon Do Pulsar url:%v  topic:%v", p.option.url, p.option.topic)
		err := p.doConsume()
		if err != nil {
			log.Errorf("Consumer Daemon Quit With Error err:%v", err)
		}
		time.Sleep(p.option.consumerRestartInterval)
	}
}

//doConsume 启动消费者
func (p *pulsarConsumerDaemon) doConsume() error {
	if p.option == nil || p.procMsgCallback == nil {
		return fmt.Errorf("pulsarClient Consumer Lack Of Param")
	}
	// 获取客户端连接
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: p.option.url,
	})
	if err != nil {
		log.Errorf("Pulsar New Client Error err:%v", err)
		return err
	}
	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
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
		err := p.procMsgCallback(msg)
		if err != nil {
			log.Errorf("Process Message Error %v", err)
		}
		consumer.Ack(msg)
	}
	return nil
}
