package postgres

func ToInterfaces[T comparable](list []T) []interface{} {
	interfaces := make([]interface{}, 0, len(list))
	for _, v := range list {
		interfaces = append(interfaces, v)
	}
	return interfaces
}
