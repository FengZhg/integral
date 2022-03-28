package pulsar

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"integral/model"
	"time"
)

// @Author: Feng
// @Date: 2022/3/26 16:55

func init() {
	// 启动消费者守护协程
	newPulsarConsumerDaemon(
		NewPulsarOptions(
			WithUrl(model.PulsarUrl),
			WithTopic(model.PulsarTopic)),
		nil,
	).start()

}

type pulsarConsumerDaemon struct {
	option          *pulsarOptions
	procMsgCallback func(msg *pulsar.Message)
}

//newPulsarConsumerDaemon 新pulsar消费者守护协程
func newPulsarConsumerDaemon(option *pulsarOptions, callback func(msg *pulsar.Message)) *pulsarConsumerDaemon {
	return &pulsarConsumerDaemon{
		option:          option,
		procMsgCallback: callback,
	}
}

//startConsumerDaemon 启动消费者守护协程
func (p *pulsarConsumerDaemon) start() {
	for {
		log.Infof("Start Daemon Do Pulsar url:%v  topic:%v", p.option.url, p.option.topic)
		err := p.doConsume()
		if err != nil {
			log.Errorf("Consumer Daemon Quit With Error err:%v", err)
		}
		time.Sleep(5 * time.Second)
	}
}

//doConsume 启动消费者
func (p *pulsarConsumerDaemon) doConsume() error {
	if p.option == nil || p.procMsgCallback == nil {
		return fmt.Errorf("pulsar Consumer Lack Of Param")
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

	rd, err := client.CreateReader(pulsar.ReaderOptions{
		Topic:          p.option.topic,
		StartMessageID: pulsar.EarliestMessageID(),
	})
	if err != nil {
		log.Errorf("Pulsar Create Reader Error err:%v", err)
		return err
	}

	for rd.HasNext() {
		// 阻塞等待
		msg, err := rd.Next(context.Background())
		if err != nil {
			log.Errorf("Receive Error Msg err:%v", err)
			return err
		}

		// 写数据库
		if p.procMsgCallback != nil {
			p.procMsgCallback(&msg)
		}
	}
	return nil
}
