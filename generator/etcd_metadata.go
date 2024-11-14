package generator

import (
	"fmt"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"regexp"
	"strings"
)

type (
	EtcdMetadata struct {
		KeyPath       string
		Fields        map[string]*Field
		RequestParams *RequestParams
	}

	Field struct {
		goName      string // name in go generated files
		protoName   string // name in proto file
		enumName    string
		kind        protoreflect.Kind
		cardinality protoreflect.Cardinality
		optional    bool
		goTypeStr   string
	}

	RequestParams struct {
		key    string
		params []string
	}
)

func (f *Field) getVariablePlaceholder() (string, error) {
	switch f.kind {
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
		return "", fmt.Errorf(`unsupported type %s for path variable: "%s"`, f.kind, f.goName)
	}
}

func GetEtcdMetadataFromMessage(
	message *protogen.Message,
	etcdKeyParams *ipb.EtcdKeyParams,
) (meta *EtcdMetadata, err error) {
	meta = &EtcdMetadata{
		KeyPath: etcdKeyParams.GetKeyPath(),
		Fields:  GetMethodFields(message),
	}
	if meta.RequestParams, err = getMethodRequestParams(meta); err != nil {
		return nil, err
	}
	return meta, nil
}

func GetMethodFields(
	message *protogen.Message,
) (fields map[string]*Field) {
	fields = make(map[string]*Field)
	for _, fieldMessage := range message.Fields {
		f := Field{
			goName:      fieldMessage.GoName,
			protoName:   fieldMessage.Desc.JSONName(),
			kind:        fieldMessage.Desc.Kind(),
			cardinality: fieldMessage.Desc.Cardinality(),
			optional:    fieldMessage.Desc.HasOptionalKeyword(),
			goTypeStr:   protoTypeToGoTypeField(fieldMessage.Desc.Kind()),
		}
		if fieldMessage.Desc.Kind() == protoreflect.EnumKind {
			f.enumName = fieldMessage.Enum.GoIdent.GoName
		}

		fields[f.protoName] = &f
	}
	return fields
}

// protoTypeToGoTypeField преобразует тип из Protocol Buffers в Go-тип
func protoTypeToGoTypeField(fieldType protoreflect.Kind) string {
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

func getMethodRequestParams(m *EtcdMetadata) (_ *RequestParams, err error) {
	var (
		placeholder         string
		keyParametersRegexp = regexp.MustCompile(`\{(\w+)\}`)
	)
	formatKey := m.KeyPath
	params := make([]string, 0, len(m.Fields))
	for _, match := range keyParametersRegexp.FindAllStringSubmatch(m.KeyPath, -1) {
		if f, ok := m.Fields[match[1]]; ok {
			if placeholder, err = f.getVariablePlaceholder(); err != nil {
				return nil, err
			}
			parameterName := strcase.ToLowerCamel(f.goName)
			formatKey = strings.ReplaceAll(formatKey, match[0], placeholder)
			params = append(params, fmt.Sprintf("%s %s", parameterName, f.goTypeStr))
		}
	}
	return &RequestParams{
		key:    formatKey,
		params: params,
	}, nil
}
