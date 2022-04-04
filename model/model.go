package model

import "integral/dao/pulsarClient"

// @Author: Feng
// @Date: 2022/4/4 20:09

var (
	PulsarOpt = pulsarClient.NewPulsarOptions(
		pulsarClient.WithUrl(PulsarUrl),
		pulsarClient.WithTopic(PulsarTopic),
	)
)
