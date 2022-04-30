package pulsarClient

import (
	"github.com/apache/pulsar-client-go/pulsar"
	log "github.com/sirupsen/logrus"
	"integral/dao"
	"integral/model"
)

// @Author: Feng
// @Date: 2022/4/21 13:19

var (
	PulsarCfg *pulsarConfig
	//PulsarOpt 参数情况
	PulsarOpt = NewPulsarOptions(
		WithUrl(model.PulsarUrl),
		WithTopic(model.PulsarTopic),
		WithConsumeRateLimit(2000),
	)
)

// 初始化
func init() {
	PulsarCfg = NewPulsarConsumerDaemon(PulsarOpt, dao.FlowConsumeCallback)
}

type pulsarConfig struct {
	option       *pulsarOptions
	producer     pulsar.Producer
	consumeMsgCB func(pulsar.ConsumerMessage) error
	client       pulsar.Client
}

//NewPulsarConsumerDaemon 新pulsar消费者守护协程
func NewPulsarConsumerDaemon(opt *pulsarOptions, cb func(pulsar.ConsumerMessage) error) *pulsarConfig {
	// 初始化其他配置
	cfg := &pulsarConfig{
		option:       opt,
		consumeMsgCB: cb,
	}
	// 获取客户端连接
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: cfg.option.url,
	})
	if err != nil {
		log.Errorf("Pulsar New Client Error err:%v", err)
		panic(err)
	}
	cfg.client = client
	// 初始化生产者和消费者
	go cfg.startConsumer()
	cfg.initProducer()
	return cfg
}
