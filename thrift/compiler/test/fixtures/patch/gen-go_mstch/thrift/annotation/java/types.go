// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package java // [[[ program thrift source path ]]]

import (
    "fmt"

    scope "thrift/annotation/scope"
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

var _ = scope.GoUnusedProtection__

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = thrift.ZERO


type Adapter struct {
    AdapterClassName string `thrift:"adapterClassName,1" json:"adapterClassName" db:"adapterClassName"`
    TypeClassName string `thrift:"typeClassName,2" json:"typeClassName" db:"typeClassName"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Adapter{}

func NewAdapter() *Adapter {
    return (&Adapter{})
}

func (x *Adapter) GetAdapterClassNameNonCompat() string {
    return x.AdapterClassName
}

func (x *Adapter) GetAdapterClassName() string {
    return x.AdapterClassName
}

func (x *Adapter) GetTypeClassNameNonCompat() string {
    return x.TypeClassName
}

func (x *Adapter) GetTypeClassName() string {
    return x.TypeClassName
}

func (x *Adapter) SetAdapterClassName(value string) *Adapter {
    x.AdapterClassName = value
    return x
}

func (x *Adapter) SetTypeClassName(value string) *Adapter {
    x.TypeClassName = value
    return x
}



func (x *Adapter) writeField1(p thrift.Protocol) error {  // AdapterClassName
    if err := p.WriteFieldBegin("adapterClassName", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetAdapterClassNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Adapter) writeField2(p thrift.Protocol) error {  // TypeClassName
    if err := p.WriteFieldBegin("typeClassName", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetTypeClassNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Adapter) readField1(p thrift.Protocol) error {  // AdapterClassName
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetAdapterClassName(result)
    return nil
}

func (x *Adapter) readField2(p thrift.Protocol) error {  // TypeClassName
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetTypeClassName(result)
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

func (x *AdapterBuilder) AdapterClassName(value string) *AdapterBuilder {
    x.obj.AdapterClassName = value
    return x
}

func (x *AdapterBuilder) TypeClassName(value string) *AdapterBuilder {
    x.obj.TypeClassName = value
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
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // adapterClassName
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // typeClassName
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

type Wrapper struct {
    WrapperClassName string `thrift:"wrapperClassName,1" json:"wrapperClassName" db:"wrapperClassName"`
    TypeClassName string `thrift:"typeClassName,2" json:"typeClassName" db:"typeClassName"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Wrapper{}

func NewWrapper() *Wrapper {
    return (&Wrapper{})
}

func (x *Wrapper) GetWrapperClassNameNonCompat() string {
    return x.WrapperClassName
}

func (x *Wrapper) GetWrapperClassName() string {
    return x.WrapperClassName
}

func (x *Wrapper) GetTypeClassNameNonCompat() string {
    return x.TypeClassName
}

func (x *Wrapper) GetTypeClassName() string {
    return x.TypeClassName
}

func (x *Wrapper) SetWrapperClassName(value string) *Wrapper {
    x.WrapperClassName = value
    return x
}

func (x *Wrapper) SetTypeClassName(value string) *Wrapper {
    x.TypeClassName = value
    return x
}



func (x *Wrapper) writeField1(p thrift.Protocol) error {  // WrapperClassName
    if err := p.WriteFieldBegin("wrapperClassName", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetWrapperClassNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) writeField2(p thrift.Protocol) error {  // TypeClassName
    if err := p.WriteFieldBegin("typeClassName", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetTypeClassNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Wrapper) readField1(p thrift.Protocol) error {  // WrapperClassName
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetWrapperClassName(result)
    return nil
}

func (x *Wrapper) readField2(p thrift.Protocol) error {  // TypeClassName
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetTypeClassName(result)
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

func (x *WrapperBuilder) WrapperClassName(value string) *WrapperBuilder {
    x.obj.WrapperClassName = value
    return x
}

func (x *WrapperBuilder) TypeClassName(value string) *WrapperBuilder {
    x.obj.TypeClassName = value
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
        case 1:  // wrapperClassName
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // typeClassName
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
