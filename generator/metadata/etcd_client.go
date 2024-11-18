package metadata

type EtcdClientMetadata struct {
	KeyPrefix string
}

func NewEtcdClientMetadata(keyPrefix string) *EtcdClientMetadata {
	return &EtcdClientMetadata{
		KeyPrefix: keyPrefix,
	}
}
