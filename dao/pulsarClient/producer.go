package pulsarClient

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"time"
)

// @Author: Feng
// @Date: 2022/3/28 13:52

func Send(option *pulsarOptions, msgStr string) error {
	// 获取客户端连接
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: option.url,
	})
	if err != nil {
		log.Errorf("Pulsar New Client Error err:%v", err)
		return err
	}
	defer client.Close()

	// 获取消费者
	pd, err := client.CreateProducer(pulsar.ProducerOptions{Topic: option.topic})
	if err != nil {
		log.Errorf("Pulsar New Producer Error err:%v", err)
		return err
	}
	defer pd.Close()

	// 发送消息
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = pd.Send(ctxWithTimeout, &pulsar.ProducerMessage{Payload: []byte(msgStr)})
	if err != nil {
		log.Errorf("Pulsar Send Message Error err:%v", err)
		return err
	}
	return nil
}
