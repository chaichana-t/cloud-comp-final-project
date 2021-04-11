package sync_service

func Increase(key string, pubsubChannel string, payloadConstructor func (string, int64) string) {
	result := client.Incr(key)
	client.Publish(pubsubChannel, payloadConstructor(key, result.Val()))
}

func Decrease(key string, pubsubChannel string, payloadConstructor func (string, int64) string) {
	result := client.Decr(key)
	client.Publish(pubsubChannel, payloadConstructor(key, result.Val()))
}
