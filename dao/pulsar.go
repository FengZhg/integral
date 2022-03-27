package dao

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"time"
)

// @Author: Feng
// @Date: 2022/3/26 16:55

const (
	pulsarUrl   = "pulsar://210.36.22.32:6650"
	pulsarTopic = "normal"
)

func init() {

}

//startConsumer 启动消费者守护协程
func startConsumer() {
	for {
		doConsume()
		time.Sleep(5 * time.Second)
	}
}

//doConsume 启动消费者
func doConsume() error {
	// 获取客户端连接
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: pulsarUrl,
	})
	if err != nil {
		log.Errorf("Pulsar New Client Error err:%v", err)
		return err
	}
	defer client.Close()

	rd, err := client.CreateReader(pulsar.ReaderOptions{
		Topic:          pulsarTopic,
		StartMessageID: pulsar.EarliestMessageID(),
	})
	if err != nil {
		log.Errorf("Pulsar Create Reader Error err:%v", err)
		return err
	}

	for rd.HasNext() {
		msg, err := rd.Next(context.Background())
		if err != nil {
			log.Errorf("Receive Error Msg err:%v", err)
			return err
		}

		// 写数据库

	}
	return nil
}
