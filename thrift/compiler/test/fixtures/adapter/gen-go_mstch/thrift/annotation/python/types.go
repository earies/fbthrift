// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package python // [[[ program thrift source path ]]]

import (
    "fmt"
    "strings"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = strings.Split
var _ = thrift.ZERO


type Py3Hidden struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &Py3Hidden{}

func NewPy3Hidden() *Py3Hidden {
    return (&Py3Hidden{})
}


// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewPy3Hidden().Set<FieldNameFoo>().Set<FieldNameBar>()
type Py3HiddenBuilder struct {
    obj *Py3Hidden
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewPy3Hidden().Set<FieldNameFoo>().Set<FieldNameBar>()
func NewPy3HiddenBuilder() *Py3HiddenBuilder {
    return &Py3HiddenBuilder{
        obj: NewPy3Hidden(),
    }
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewPy3Hidden().Set<FieldNameFoo>().Set<FieldNameBar>()
func (x *Py3HiddenBuilder) Emit() *Py3Hidden {
    var objCopy Py3Hidden = *x.obj
    return &objCopy
}

func (x *Py3Hidden) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Py3Hidden"); err != nil {
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

func (x *Py3Hidden) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch id {
        default:
            if err := p.Skip(wireType); err != nil {
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

func (x *Py3Hidden) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Py3Hidden({")
    sb.WriteString("})")

    return sb.String()
}

type Flags struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &Flags{}

func NewFlags() *Flags {
    return (&Flags{})
}


// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewFlags().Set<FieldNameFoo>().Set<FieldNameBar>()
type FlagsBuilder struct {
    obj *Flags
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewFlags().Set<FieldNameFoo>().Set<FieldNameBar>()
func NewFlagsBuilder() *FlagsBuilder {
    return &FlagsBuilder{
        obj: NewFlags(),
    }
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewFlags().Set<FieldNameFoo>().Set<FieldNameBar>()
func (x *FlagsBuilder) Emit() *Flags {
    var objCopy Flags = *x.obj
    return &objCopy
}

func (x *Flags) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Flags"); err != nil {
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

func (x *Flags) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch id {
        default:
            if err := p.Skip(wireType); err != nil {
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

func (x *Flags) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Flags({")
    sb.WriteString("})")

    return sb.String()
}

type Name struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Name{}

func NewName() *Name {
    return (&Name{}).
        SetNameNonCompat("")
}

func (x *Name) GetNameNonCompat() string {
    return x.Name
}

func (x *Name) GetName() string {
    return x.Name
}

func (x *Name) SetNameNonCompat(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) SetName(value string) *Name {
    x.Name = value
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

func (x *Name) readField1(p thrift.Protocol) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetNameNonCompat(result)
    return nil
}

func (x *Name) toString1() string {  // Name
    return fmt.Sprintf("%v", x.GetNameNonCompat())
}


// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewName().Set<FieldNameFoo>().Set<FieldNameBar>()
type NameBuilder struct {
    obj *Name
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewName().Set<FieldNameFoo>().Set<FieldNameBar>()
func NewNameBuilder() *NameBuilder {
    return &NameBuilder{
        obj: NewName(),
    }
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewName().Set<FieldNameFoo>().Set<FieldNameBar>()
func (x *NameBuilder) Name(value string) *NameBuilder {
    x.obj.Name = value
    return x
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewName().Set<FieldNameFoo>().Set<FieldNameBar>()
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
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            expectedType := thrift.Type(thrift.STRING)
            if wireType == expectedType {
                if err := x.readField1(p); err != nil {
                   return err
                }
            } else {
                if err := p.Skip(wireType); err != nil {
                    return err
                }
            }
        default:
            if err := p.Skip(wireType); err != nil {
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

func (x *Name) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Name({")
    sb.WriteString(fmt.Sprintf("Name:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}

type Adapter struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
    TypeHint string `thrift:"typeHint,2" json:"typeHint" db:"typeHint"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Adapter{}

func NewAdapter() *Adapter {
    return (&Adapter{}).
        SetNameNonCompat("").
        SetTypeHintNonCompat("")
}

func (x *Adapter) GetNameNonCompat() string {
    return x.Name
}

func (x *Adapter) GetName() string {
    return x.Name
}

func (x *Adapter) GetTypeHintNonCompat() string {
    return x.TypeHint
}

func (x *Adapter) GetTypeHint() string {
    return x.TypeHint
}

func (x *Adapter) SetNameNonCompat(value string) *Adapter {
    x.Name = value
    return x
}

func (x *Adapter) SetName(value string) *Adapter {
    x.Name = value
    return x
}

func (x *Adapter) SetTypeHintNonCompat(value string) *Adapter {
    x.TypeHint = value
    return x
}

func (x *Adapter) SetTypeHint(value string) *Adapter {
    x.TypeHint = value
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

func (x *Adapter) writeField2(p thrift.Protocol) error {  // TypeHint
    if err := p.WriteFieldBegin("typeHint", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetTypeHintNonCompat()
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

func (x *Adapter) readField2(p thrift.Protocol) error {  // TypeHint
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetTypeHintNonCompat(result)
    return nil
}

func (x *Adapter) toString1() string {  // Name
    return fmt.Sprintf("%v", x.GetNameNonCompat())
}

func (x *Adapter) toString2() string {  // TypeHint
    return fmt.Sprintf("%v", x.GetTypeHintNonCompat())
}


// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewAdapter().Set<FieldNameFoo>().Set<FieldNameBar>()
type AdapterBuilder struct {
    obj *Adapter
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewAdapter().Set<FieldNameFoo>().Set<FieldNameBar>()
func NewAdapterBuilder() *AdapterBuilder {
    return &AdapterBuilder{
        obj: NewAdapter(),
    }
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewAdapter().Set<FieldNameFoo>().Set<FieldNameBar>()
func (x *AdapterBuilder) Name(value string) *AdapterBuilder {
    x.obj.Name = value
    return x
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewAdapter().Set<FieldNameFoo>().Set<FieldNameBar>()
func (x *AdapterBuilder) TypeHint(value string) *AdapterBuilder {
    x.obj.TypeHint = value
    return x
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewAdapter().Set<FieldNameFoo>().Set<FieldNameBar>()
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

func (x *Adapter) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            expectedType := thrift.Type(thrift.STRING)
            if wireType == expectedType {
                if err := x.readField1(p); err != nil {
                   return err
                }
            } else {
                if err := p.Skip(wireType); err != nil {
                    return err
                }
            }
        case 2:  // typeHint
            expectedType := thrift.Type(thrift.STRING)
            if wireType == expectedType {
                if err := x.readField2(p); err != nil {
                   return err
                }
            } else {
                if err := p.Skip(wireType); err != nil {
                    return err
                }
            }
        default:
            if err := p.Skip(wireType); err != nil {
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

func (x *Adapter) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Adapter({")
    sb.WriteString(fmt.Sprintf("Name:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("TypeHint:%s", x.toString2()))
    sb.WriteString("})")

    return sb.String()
}

type MarshalCapi struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &MarshalCapi{}

func NewMarshalCapi() *MarshalCapi {
    return (&MarshalCapi{})
}


// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewMarshalCapi().Set<FieldNameFoo>().Set<FieldNameBar>()
type MarshalCapiBuilder struct {
    obj *MarshalCapi
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewMarshalCapi().Set<FieldNameFoo>().Set<FieldNameBar>()
func NewMarshalCapiBuilder() *MarshalCapiBuilder {
    return &MarshalCapiBuilder{
        obj: NewMarshalCapi(),
    }
}

// Deprecated: Use "New" constructor and setters to build your structs.
// e.g NewMarshalCapi().Set<FieldNameFoo>().Set<FieldNameBar>()
func (x *MarshalCapiBuilder) Emit() *MarshalCapi {
    var objCopy MarshalCapi = *x.obj
    return &objCopy
}

func (x *MarshalCapi) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("MarshalCapi"); err != nil {
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

func (x *MarshalCapi) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch id {
        default:
            if err := p.Skip(wireType); err != nil {
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

func (x *MarshalCapi) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("MarshalCapi({")
    sb.WriteString("})")

    return sb.String()
}

// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
	  RegisterType(name string, initializer func() any)
}) {
    registry.RegisterType("facebook.com/thrift/annotation/python/Py3Hidden", func() any { return NewPy3Hidden() })
    registry.RegisterType("facebook.com/thrift/annotation/python/Flags", func() any { return NewFlags() })
    registry.RegisterType("facebook.com/thrift/annotation/python/Name", func() any { return NewName() })
    registry.RegisterType("facebook.com/thrift/annotation/python/Adapter", func() any { return NewAdapter() })
    registry.RegisterType("facebook.com/thrift/annotation/python/MarshalCapi", func() any { return NewMarshalCapi() })

}
