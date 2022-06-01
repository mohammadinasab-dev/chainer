package config

import "strconv"

func GetRabbitMQConnectionString() string {
	return getConfigString("rabbitmq.connection_string")
}
func GetRabbitMQExchangeTypeString() string {
	return getConfigString("rabbitmq.exchange_type")
}
func GetRabbitMQExchangeNameString() string {
	return getConfigString("rabbitmq.exchange_name")
}
func GetRabbitMQRoutingKeyString() string {
	return getConfigString("rabbitmq.routing_key")
}
func GetRabbitMQQeueNameString() string {
	return getConfigString("rabbitmq.qeue_name")
}
func GetRabbitMQConsumerCountInt() int {
	consumerCount := getConfigString("rabbitmq.consumer_count")
	count, _ := strconv.Atoi(consumerCount)
	return count
}
