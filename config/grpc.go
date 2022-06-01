package config

func GetGrpcAddressString() string {
	return getConfigString("grpc.address")
}

func GetGrpcPortString() string {
	return getConfigString("grpc.port")
}

func GetGrpcConnectionTypeString() string {
	return getConfigString("grpc.Type")
}
