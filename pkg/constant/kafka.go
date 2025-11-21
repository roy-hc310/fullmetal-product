package constant

const (
	DefaultTopic = "default-topic"

	ProductCreateTopic = "product.create"
	ProductUpdateTopic = "product.update"
	ProductDeleteTopic = "product.delete"
)

var KafkaTopics = []string{
	DefaultTopic,
	ProductCreateTopic,
	ProductUpdateTopic,
	ProductDeleteTopic,
}
