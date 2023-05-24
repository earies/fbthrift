// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package hack // [[[ program thrift source path ]]]

import (
    "fmt"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)


// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = thrift.ZERO


type FieldWrapper struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &FieldWrapper{}

func NewFieldWrapper() *FieldWrapper {
    return (&FieldWrapper{}).
        SetNameNonCompat("")
}

func (x *FieldWrapper) GetNameNonCompat() string {
    return x.Name
}

func (x *FieldWrapper) GetName() string {
    return x.Name
}

func (x *FieldWrapper) SetNameNonCompat(value string) *FieldWrapper {
    x.Name = value
    return x
}

func (x *FieldWrapper) SetName(value string) *FieldWrapper {
    x.Name = value
    return x
}

func (x *FieldWrapper) writeField1(p thrift.Protocol) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *FieldWrapper) readField1(p thrift.Protocol) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetNameNonCompat(result)
    return nil
}

func (x *FieldWrapper) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use FieldWrapper.Set* methods instead or set the fields directly.
type FieldWrapperBuilder struct {
    obj *FieldWrapper
}

func NewFieldWrapperBuilder() *FieldWrapperBuilder {
    return &FieldWrapperBuilder{
        obj: NewFieldWrapper(),
    }
}

func (x *FieldWrapperBuilder) Name(value string) *FieldWrapperBuilder {
    x.obj.Name = value
    return x
}

func (x *FieldWrapperBuilder) Emit() *FieldWrapper {
    var objCopy FieldWrapper = *x.obj
    return &objCopy
}

func (x *FieldWrapper) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("FieldWrapper"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *FieldWrapper) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type Wrapper struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
    UnderlyingName string `thrift:"underlyingName,2" json:"underlyingName" db:"underlyingName"`
    ExtraNamespace string `thrift:"extraNamespace,3" json:"extraNamespace" db:"extraNamespace"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Wrapper{}

func NewWrapper() *Wrapper {
    return (&Wrapper{}).
        SetNameNonCompat("").
        SetUnderlyingNameNonCompat("").
        SetExtraNamespaceNonCompat("thrift_adapted_types")
}

func (x *Wrapper) GetNameNonCompat() string {
    return x.Name
}

func (x *Wrapper) GetName() string {
    return x.Name
}

func (x *Wrapper) GetUnderlyingNameNonCompat() string {
    return x.UnderlyingName
}

func (x *Wrapper) GetUnderlyingName() string {
    return x.UnderlyingName
}

func (x *Wrapper) GetExtraNamespaceNonCompat() string {
    return x.ExtraNamespace
}

func (x *Wrapper) GetExtraNamespace() string {
    return x.ExtraNamespace
}

func (x *Wrapper) SetNameNonCompat(value string) *Wrapper {
    x.Name = value
    return x
}

func (x *Wrapper) SetName(value string) *Wrapper {
    x.Name = value
    return x
}

func (x *Wrapper) SetUnderlyingNameNonCompat(value string) *Wrapper {
    x.UnderlyingName = value
    return x
}

func (x *Wrapper) SetUnderlyingName(value string) *Wrapper {
    x.UnderlyingName = value
    return x
}

func (x *Wrapper) SetExtraNamespaceNonCompat(value string) *Wrapper {
    x.ExtraNamespace = value
    return x
}

func (x *Wrapper) SetExtraNamespace(value string) *Wrapper {
    x.ExtraNamespace = value
    return x
}

func (x *Wrapper) writeField1(p thrift.Protocol) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) writeField2(p thrift.Protocol) error {  // UnderlyingName
    if err := p.WriteFieldBegin("underlyingName", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetUnderlyingNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) writeField3(p thrift.Protocol) error {  // ExtraNamespace
    if err := p.WriteFieldBegin("extraNamespace", thrift.STRING, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetExtraNamespaceNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) readField1(p thrift.Protocol) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetNameNonCompat(result)
    return nil
}

func (x *Wrapper) readField2(p thrift.Protocol) error {  // UnderlyingName
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetUnderlyingNameNonCompat(result)
    return nil
}

func (x *Wrapper) readField3(p thrift.Protocol) error {  // ExtraNamespace
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetExtraNamespaceNonCompat(result)
    return nil
}

func (x *Wrapper) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use Wrapper.Set* methods instead or set the fields directly.
type WrapperBuilder struct {
    obj *Wrapper
}

func NewWrapperBuilder() *WrapperBuilder {
    return &WrapperBuilder{
        obj: NewWrapper(),
    }
}

func (x *WrapperBuilder) Name(value string) *WrapperBuilder {
    x.obj.Name = value
    return x
}

func (x *WrapperBuilder) UnderlyingName(value string) *WrapperBuilder {
    x.obj.UnderlyingName = value
    return x
}

func (x *WrapperBuilder) ExtraNamespace(value string) *WrapperBuilder {
    x.obj.ExtraNamespace = value
    return x
}

func (x *WrapperBuilder) Emit() *Wrapper {
    var objCopy Wrapper = *x.obj
    return &objCopy
}

func (x *Wrapper) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Wrapper"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // underlyingName
            if err := x.readField2(p); err != nil {
                return err
            }
        case 3:  // extraNamespace
            if err := x.readField3(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type Adapter struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Adapter{}

func NewAdapter() *Adapter {
    return (&Adapter{}).
        SetNameNonCompat("")
}

func (x *Adapter) GetNameNonCompat() string {
    return x.Name
}

func (x *Adapter) GetName() string {
    return x.Name
}

func (x *Adapter) SetNameNonCompat(value string) *Adapter {
    x.Name = value
    return x
}

func (x *Adapter) SetName(value string) *Adapter {
    x.Name = value
    return x
}

func (x *Adapter) writeField1(p thrift.Protocol) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Adapter) readField1(p thrift.Protocol) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetNameNonCompat(result)
    return nil
}

func (x *Adapter) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use Adapter.Set* methods instead or set the fields directly.
type AdapterBuilder struct {
    obj *Adapter
}

func NewAdapterBuilder() *AdapterBuilder {
    return &AdapterBuilder{
        obj: NewAdapter(),
    }
}

func (x *AdapterBuilder) Name(value string) *AdapterBuilder {
    x.obj.Name = value
    return x
}

func (x *AdapterBuilder) Emit() *Adapter {
    var objCopy Adapter = *x.obj
    return &objCopy
}

func (x *Adapter) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Adapter"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Adapter) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type SkipCodegen struct {
    Reason string `thrift:"reason,1" json:"reason" db:"reason"`
}
// Compile time interface enforcer
var _ thrift.Struct = &SkipCodegen{}

func NewSkipCodegen() *SkipCodegen {
    return (&SkipCodegen{}).
        SetReasonNonCompat("")
}

func (x *SkipCodegen) GetReasonNonCompat() string {
    return x.Reason
}

func (x *SkipCodegen) GetReason() string {
    return x.Reason
}

func (x *SkipCodegen) SetReasonNonCompat(value string) *SkipCodegen {
    x.Reason = value
    return x
}

func (x *SkipCodegen) SetReason(value string) *SkipCodegen {
    x.Reason = value
    return x
}

func (x *SkipCodegen) writeField1(p thrift.Protocol) error {  // Reason
    if err := p.WriteFieldBegin("reason", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetReasonNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *SkipCodegen) readField1(p thrift.Protocol) error {  // Reason
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetReasonNonCompat(result)
    return nil
}

func (x *SkipCodegen) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use SkipCodegen.Set* methods instead or set the fields directly.
type SkipCodegenBuilder struct {
    obj *SkipCodegen
}

func NewSkipCodegenBuilder() *SkipCodegenBuilder {
    return &SkipCodegenBuilder{
        obj: NewSkipCodegen(),
    }
}

func (x *SkipCodegenBuilder) Reason(value string) *SkipCodegenBuilder {
    x.obj.Reason = value
    return x
}

func (x *SkipCodegenBuilder) Emit() *SkipCodegen {
    var objCopy SkipCodegen = *x.obj
    return &objCopy
}

func (x *SkipCodegen) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("SkipCodegen"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *SkipCodegen) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // reason
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type Name struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
    Reason string `thrift:"reason,2" json:"reason" db:"reason"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Name{}

func NewName() *Name {
    return (&Name{}).
        SetNameNonCompat("").
        SetReasonNonCompat("")
}

func (x *Name) GetNameNonCompat() string {
    return x.Name
}

func (x *Name) GetName() string {
    return x.Name
}

func (x *Name) GetReasonNonCompat() string {
    return x.Reason
}

func (x *Name) GetReason() string {
    return x.Reason
}

func (x *Name) SetNameNonCompat(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) SetName(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) SetReasonNonCompat(value string) *Name {
    x.Reason = value
    return x
}

func (x *Name) SetReason(value string) *Name {
    x.Reason = value
    return x
}

func (x *Name) writeField1(p thrift.Protocol) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Name) writeField2(p thrift.Protocol) error {  // Reason
    if err := p.WriteFieldBegin("reason", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetReasonNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Name) readField1(p thrift.Protocol) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetNameNonCompat(result)
    return nil
}

func (x *Name) readField2(p thrift.Protocol) error {  // Reason
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetReasonNonCompat(result)
    return nil
}

func (x *Name) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use Name.Set* methods instead or set the fields directly.
type NameBuilder struct {
    obj *Name
}

func NewNameBuilder() *NameBuilder {
    return &NameBuilder{
        obj: NewName(),
    }
}

func (x *NameBuilder) Name(value string) *NameBuilder {
    x.obj.Name = value
    return x
}

func (x *NameBuilder) Reason(value string) *NameBuilder {
    x.obj.Reason = value
    return x
}

func (x *NameBuilder) Emit() *Name {
    var objCopy Name = *x.obj
    return &objCopy
}

func (x *Name) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Name"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Name) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // reason
            if err := x.readField2(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type UnionEnumAttributes struct {
    Attributes []string `thrift:"attributes,1" json:"attributes" db:"attributes"`
}
// Compile time interface enforcer
var _ thrift.Struct = &UnionEnumAttributes{}

func NewUnionEnumAttributes() *UnionEnumAttributes {
    return (&UnionEnumAttributes{}).
        SetAttributesNonCompat(make([]string, 0))
}

func (x *UnionEnumAttributes) GetAttributesNonCompat() []string {
    return x.Attributes
}

func (x *UnionEnumAttributes) GetAttributes() []string {
    if !x.IsSetAttributes() {
        return make([]string, 0)
    }

    return x.Attributes
}

func (x *UnionEnumAttributes) SetAttributesNonCompat(value []string) *UnionEnumAttributes {
    x.Attributes = value
    return x
}

func (x *UnionEnumAttributes) SetAttributes(value []string) *UnionEnumAttributes {
    x.Attributes = value
    return x
}

func (x *UnionEnumAttributes) IsSetAttributes() bool {
    return x.Attributes != nil
}

func (x *UnionEnumAttributes) writeField1(p thrift.Protocol) error {  // Attributes
    if !x.IsSetAttributes() {
        return nil
    }

    if err := p.WriteFieldBegin("attributes", thrift.LIST, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetAttributesNonCompat()
    if err := p.WriteListBegin(thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *UnionEnumAttributes) readField1(p thrift.Protocol) error {  // Attributes
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]string, 0, size)
for i := 0; i < size; i++ {
    var elem string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.SetAttributesNonCompat(result)
    return nil
}

func (x *UnionEnumAttributes) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use UnionEnumAttributes.Set* methods instead or set the fields directly.
type UnionEnumAttributesBuilder struct {
    obj *UnionEnumAttributes
}

func NewUnionEnumAttributesBuilder() *UnionEnumAttributesBuilder {
    return &UnionEnumAttributesBuilder{
        obj: NewUnionEnumAttributes(),
    }
}

func (x *UnionEnumAttributesBuilder) Attributes(value []string) *UnionEnumAttributesBuilder {
    x.obj.Attributes = value
    return x
}

func (x *UnionEnumAttributesBuilder) Emit() *UnionEnumAttributes {
    var objCopy UnionEnumAttributes = *x.obj
    return &objCopy
}

func (x *UnionEnumAttributes) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("UnionEnumAttributes"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *UnionEnumAttributes) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // attributes
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type StructTrait struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &StructTrait{}

func NewStructTrait() *StructTrait {
    return (&StructTrait{}).
        SetNameNonCompat("")
}

func (x *StructTrait) GetNameNonCompat() string {
    return x.Name
}

func (x *StructTrait) GetName() string {
    return x.Name
}

func (x *StructTrait) SetNameNonCompat(value string) *StructTrait {
    x.Name = value
    return x
}

func (x *StructTrait) SetName(value string) *StructTrait {
    x.Name = value
    return x
}

func (x *StructTrait) writeField1(p thrift.Protocol) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *StructTrait) readField1(p thrift.Protocol) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetNameNonCompat(result)
    return nil
}

func (x *StructTrait) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use StructTrait.Set* methods instead or set the fields directly.
type StructTraitBuilder struct {
    obj *StructTrait
}

func NewStructTraitBuilder() *StructTraitBuilder {
    return &StructTraitBuilder{
        obj: NewStructTrait(),
    }
}

func (x *StructTraitBuilder) Name(value string) *StructTraitBuilder {
    x.obj.Name = value
    return x
}

func (x *StructTraitBuilder) Emit() *StructTrait {
    var objCopy StructTrait = *x.obj
    return &objCopy
}

func (x *StructTrait) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("StructTrait"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *StructTrait) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type Attributes struct {
    Attributes []string `thrift:"attributes,1" json:"attributes" db:"attributes"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Attributes{}

func NewAttributes() *Attributes {
    return (&Attributes{}).
        SetAttributesNonCompat(make([]string, 0))
}

func (x *Attributes) GetAttributesNonCompat() []string {
    return x.Attributes
}

func (x *Attributes) GetAttributes() []string {
    if !x.IsSetAttributes() {
        return make([]string, 0)
    }

    return x.Attributes
}

func (x *Attributes) SetAttributesNonCompat(value []string) *Attributes {
    x.Attributes = value
    return x
}

func (x *Attributes) SetAttributes(value []string) *Attributes {
    x.Attributes = value
    return x
}

func (x *Attributes) IsSetAttributes() bool {
    return x.Attributes != nil
}

func (x *Attributes) writeField1(p thrift.Protocol) error {  // Attributes
    if !x.IsSetAttributes() {
        return nil
    }

    if err := p.WriteFieldBegin("attributes", thrift.LIST, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetAttributesNonCompat()
    if err := p.WriteListBegin(thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Attributes) readField1(p thrift.Protocol) error {  // Attributes
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]string, 0, size)
for i := 0; i < size; i++ {
    var elem string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.SetAttributesNonCompat(result)
    return nil
}

func (x *Attributes) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use Attributes.Set* methods instead or set the fields directly.
type AttributesBuilder struct {
    obj *Attributes
}

func NewAttributesBuilder() *AttributesBuilder {
    return &AttributesBuilder{
        obj: NewAttributes(),
    }
}

func (x *AttributesBuilder) Attributes(value []string) *AttributesBuilder {
    x.obj.Attributes = value
    return x
}

func (x *AttributesBuilder) Emit() *Attributes {
    var objCopy Attributes = *x.obj
    return &objCopy
}

func (x *Attributes) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Attributes"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Attributes) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // attributes
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type StructAsTrait struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &StructAsTrait{}

func NewStructAsTrait() *StructAsTrait {
    return (&StructAsTrait{})
}

func (x *StructAsTrait) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use StructAsTrait.Set* methods instead or set the fields directly.
type StructAsTraitBuilder struct {
    obj *StructAsTrait
}

func NewStructAsTraitBuilder() *StructAsTraitBuilder {
    return &StructAsTraitBuilder{
        obj: NewStructAsTrait(),
    }
}

func (x *StructAsTraitBuilder) Emit() *StructAsTrait {
    var objCopy StructAsTrait = *x.obj
    return &objCopy
}

func (x *StructAsTrait) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("StructAsTrait"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *StructAsTrait) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type ModuleInternal struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &ModuleInternal{}

func NewModuleInternal() *ModuleInternal {
    return (&ModuleInternal{})
}

func (x *ModuleInternal) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use ModuleInternal.Set* methods instead or set the fields directly.
type ModuleInternalBuilder struct {
    obj *ModuleInternal
}

func NewModuleInternalBuilder() *ModuleInternalBuilder {
    return &ModuleInternalBuilder{
        obj: NewModuleInternal(),
    }
}

func (x *ModuleInternalBuilder) Emit() *ModuleInternal {
    var objCopy ModuleInternal = *x.obj
    return &objCopy
}

func (x *ModuleInternal) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("ModuleInternal"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *ModuleInternal) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

