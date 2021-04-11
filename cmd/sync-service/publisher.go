package sync_service

func publishValueChange(key string, value int64) {
	client.Publish(key, value)
}

func Increase(key string) {
	result := client.Incr(key)
	publishValueChange(key, result.Val())
}

func Decrease(key string) {
	result := client.Decr(key)
	publishValueChange(key, result.Val())
}
