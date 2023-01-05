package widget

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/goccy/go-json"
)

type emptyInterface struct {
	typ  *struct{}
	data unsafe.Pointer
}

type BindField struct {
	fieldNum   int
	jsonName   string
	validators map[string]interface{}
}

func (f *BindField) MinLength(length int) *BindField {
	f.validators["minLength"] = length
	return f
}

func (f *BindField) MaxLength(length int) *BindField {
	f.validators["maxLength"] = length
	return f
}

func (f *BindField) Required() *BindField {
	f.validators["required"] = true
	return f
}

func (f *BindField) ServerSide() *BindField {
	f.validators["serverSide"] = true
	return f
}

func (f *BindField) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.jsonName)
}

type Bind func(fieldPtr any) *BindField

type fieldMeta struct {
	fieldNum int
	jsonName string
}

type binder struct {
	// offset -> fieldMeta
	metaData  map[uintptr]*fieldMeta
	structPtr uintptr
}

func newBinder(fields reflect.Type, structPtr uintptr) *binder {
	metaData := make(map[uintptr]*fieldMeta)
	val := fields.Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		jsonTags, ok := field.Tag.Lookup("json")
		if !ok {
			continue
		}

		for _, opt := range strings.Split(jsonTags, ",") {
			if jsonName := strings.TrimSpace(opt); jsonName != "omitempty" {
				metaData[field.Offset] = &fieldMeta{
					fieldNum: i,
					jsonName: jsonName,
				}
				break
			}
		}
	}

	return &binder{
		metaData:  metaData,
		structPtr: structPtr,
	}
}

func (b *binder) Bind(fieldPtr any) *BindField {
	eface := (*emptyInterface)(unsafe.Pointer(&fieldPtr))
	offset := uintptr(eface.data) - b.structPtr
	if md, ok := b.metaData[offset]; ok {
		return &BindField{
			fieldNum:   md.fieldNum,
			jsonName:   md.jsonName,
			validators: make(map[string]interface{}),
		}
	}

	panic("bind name not found")
}
