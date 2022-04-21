package pulsarClient

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/3/28 13:52

//initProducer 消息生产事件循环
func (p *pulsarConfig) initProducer() error {
	// 获取消费者
	pd, err := p.client.CreateProducer(pulsar.ProducerOptions{Topic: p.option.topic})
	if err != nil {
		log.Errorf("Pulsar New Producer Error err:%v", err)
		return err
	}
	p.producer = pd
	return nil
}

//Produce 生产者生产消息
func (p *pulsarConfig) Produce(ctx context.Context, msg []byte) error {
	_, err := p.producer.Send(ctx, &pulsar.ProducerMessage{Payload: msg})
	if err != nil {
		log.Errorf("Pulsar Send Msg Error %v", err)
		return err
	}
	return nil
}
