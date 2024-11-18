package metadata

type EtcdClientMetadata struct {
	KeyPrefix string
	Imports   map[string]*ImportResolver
}

func NewEtcdClientMetadata(keyPrefix string, imports map[string]*ImportResolver) *EtcdClientMetadata {
	return &EtcdClientMetadata{
		KeyPrefix: keyPrefix,
		Imports:   imports,
	}
}
