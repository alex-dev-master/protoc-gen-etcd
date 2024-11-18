package metadata

import (
	"fmt"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"regexp"
	"strings"
)

type EtcdMethodMetadata struct {
	EtcdKeyOptions      *ipb.EtcdKeyOptions
	EtcdKeyParamOptions map[string]*FieldWithEtcdKeyParamOptions
	ValueOfKey          protogen.GoIdent
	InputRequest        protogen.GoIdent
}

func NewEtcdMethodMetadata(
	etcdKeyOptions *ipb.EtcdKeyOptions,
	etcdKeyParamOptions map[string]*FieldWithEtcdKeyParamOptions,
	valueOfKey protogen.GoIdent,
	inputRequest protogen.GoIdent,
) *EtcdMethodMetadata {
	return &EtcdMethodMetadata{
		EtcdKeyOptions:      etcdKeyOptions,
		EtcdKeyParamOptions: etcdKeyParamOptions,
		ValueOfKey:          valueOfKey,
		InputRequest:        inputRequest,
	}
}

type (
	FieldWithEtcdKeyParamOptions struct {
		GoName            string
		ProtoName         string
		EnumName          string
		Kind              protoreflect.Kind
		Cardinality       protoreflect.Cardinality
		FieldInputOptsExt *ipb.EtcdKeyParamOptions
		GoTypeStr         string
	}

	RequestParams struct {
		key    string
		params []string
	}
)

func (f *FieldWithEtcdKeyParamOptions) GetVariablePlaceholder() (string, error) {
	switch f.Kind {
	case protoreflect.StringKind,
		protoreflect.EnumKind,
		protoreflect.BytesKind:
		return "%s", nil
	case protoreflect.Int32Kind,
		protoreflect.Sint32Kind,
		protoreflect.Uint32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sint64Kind,
		protoreflect.Uint64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Fixed64Kind:
		return "%d", nil
	case
		protoreflect.FloatKind,
		protoreflect.DoubleKind:
		return "%f", nil
	case protoreflect.BoolKind:
		return "%t", nil
	default:
		return "", fmt.Errorf(`unsupported type %s for path variable: "%s"`, f.Kind, f.GoName)
	}
}

// ProtoTypeToGoTypeField преобразует тип из Protocol Buffers в Go-тип
func ProtoTypeToGoTypeField(fieldType protoreflect.Kind) string {
	switch fieldType {
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64"
	case protoreflect.FloatKind:
		return "float32"
	case protoreflect.DoubleKind:
		return "float64"
	case protoreflect.StringKind:
		return "string"
	case protoreflect.BytesKind:
		return "[]byte"
	case protoreflect.EnumKind:
		return "int32"
	case protoreflect.MessageKind:
		return "*" + fieldType.String() // указатель на тип сообщения
	default:
		return "unknown"
	}
}

func (m *EtcdMethodMetadata) GetMethodRequestParams() (formatKey string, params []string, err error) {
	keyParametersRegexp := regexp.MustCompile(`\{(\w+)\}`)
	formatKey = m.EtcdKeyOptions.GetKeyPath()
	var placeholder string
	for _, match := range keyParametersRegexp.FindAllStringSubmatch(formatKey, -1) {
		if f, ok := m.EtcdKeyParamOptions[match[1]]; ok {
			if placeholder, err = f.GetVariablePlaceholder(); err != nil {
				return "", nil, err
			}

			formatKey = strings.ReplaceAll(formatKey, match[0], placeholder)
			params = append(params, fmt.Sprintf("%s%s", "rq.", f.GoName))
		}
	}
	return formatKey, params, nil
}
