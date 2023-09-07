// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package module // [[[ program thrift source path ]]]


import (
    "context"
    "fmt"
    "strings"
    "sync"


    "thrift/lib/go/thrift"
)


// (needed to ensure safety because of naive import list construction)
var _ = context.Background
var _ = fmt.Printf
var _ = thrift.ZERO
var _ = strings.Split
var _ = sync.Mutex{}



type C interface {
    F(ctx context.Context) (error)
    Thing(ctx context.Context, a int32, b string, c []int32) (string, error)
}

// Deprecated: Use C instead.
type CClientInterface interface {
    thrift.ClientInterface
    F() (error)
    Thing(a int32, b string, c []int32) (string, error)
}

type CChannelClient struct {
    ch thrift.RequestChannel
}
// Compile time interface enforcer
var _ C = &CChannelClient{}

func NewCChannelClient(channel thrift.RequestChannel) *CChannelClient {
    return &CChannelClient{
        ch: channel,
    }
}

func (c *CChannelClient) Close() error {
    return c.ch.Close()
}

func (c *CChannelClient) IsOpen() bool {
    return c.ch.IsOpen()
}

func (c *CChannelClient) Open() error {
    return c.ch.Open()
}

// Deprecated: Use CChannelClient instead.
type CClient struct {
    chClient *CChannelClient
    Mu       sync.Mutex
}
// Compile time interface enforcer
var _ CClientInterface = &CClient{}

// Deprecated: Use NewCChannelClient() instead.
func NewCClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *CClient {
    return &CClient{
        chClient: NewCChannelClient(
            thrift.NewSerialChannel(iprot),
        ),
    }
}

func (c *CClient) Close() error {
    return c.chClient.Close()
}

func (c *CClient) IsOpen() bool {
    return c.chClient.IsOpen()
}

func (c *CClient) Open() error {
    return c.chClient.Open()
}

// Deprecated: Use CChannelClient instead.
type CThreadsafeClient = CClient

// Deprecated: Use NewCChannelClient() instead.
func NewCThreadsafeClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *CThreadsafeClient {
    return NewCClient(t, iprot, oprot)
}

// Deprecated: Use NewCChannelClient() instead.
func NewCClientProtocol(prot thrift.Protocol) *CClient {
  return NewCClient(prot.Transport(), prot, prot)
}

// Deprecated: Use NewCChannelClient() instead.
func NewCThreadsafeClientProtocol(prot thrift.Protocol) *CClient {
  return NewCClient(prot.Transport(), prot, prot)
}

// Deprecated: Use NewCChannelClient() instead.
func NewCClientFactory(t thrift.Transport, pf thrift.ProtocolFactory) *CClient {
  iprot := pf.GetProtocol(t)
  oprot := pf.GetProtocol(t)
  return NewCClient(t, iprot, oprot)
}

// Deprecated: Use NewCChannelClient() instead.
func NewCThreadsafeClientFactory(t thrift.Transport, pf thrift.ProtocolFactory) *CThreadsafeClient {
  return NewCClientFactory(t, pf)
}


func (c *CChannelClient) F(ctx context.Context) (error) {
    in := &reqCF{
    }
    out := newRespCF()
    err := c.ch.Call(ctx, "f", in, out)
    if err != nil {
        return err
    }
    return nil
}

func (c *CClient) F() (error) {
    return c.chClient.F(context.TODO())
}


func (c *CChannelClient) Thing(ctx context.Context, a int32, b string, c []int32) (string, error) {
    in := &reqCThing{
        A: a,
        B: b,
        C: c,
    }
    out := newRespCThing()
    err := c.ch.Call(ctx, "thing", in, out)
    if err != nil {
        return "", err
    } else if out.Bang != nil {
        return "", out.Bang
    }
    return out.Value, nil
}

func (c *CClient) Thing(a int32, b string, c []int32) (string, error) {
    return c.chClient.Thing(context.TODO(), a, b, c)
}


type reqCF struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &reqCF{}

type CFArgs = reqCF

func newReqCF() *reqCF {
    return (&reqCF{})
}


// Deprecated: Use reqCF.Set* methods instead or set the fields directly.
type reqCFBuilder struct {
    obj *reqCF
}

func newReqCFBuilder() *reqCFBuilder {
    return &reqCFBuilder{
        obj: newReqCF(),
    }
}

func (x *reqCFBuilder) Emit() *reqCF {
    var objCopy reqCF = *x.obj
    return &objCopy
}

func (x *reqCF) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("reqCF"); err != nil {
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

func (x *reqCF) Read(p thrift.Protocol) error {
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
    }

    if err := p.ReadFieldEnd(); err != nil {
        return err
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *reqCF) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("reqCF({")
    sb.WriteString("})")

    return sb.String()
}
type respCF struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &respCF{}
var _ thrift.WritableResult = &respCF{}

func newRespCF() *respCF {
    return (&respCF{})
}


// Deprecated: Use respCF.Set* methods instead or set the fields directly.
type respCFBuilder struct {
    obj *respCF
}

func newRespCFBuilder() *respCFBuilder {
    return &respCFBuilder{
        obj: newRespCF(),
    }
}

func (x *respCFBuilder) Emit() *respCF {
    var objCopy respCF = *x.obj
    return &objCopy
}

func (x *respCF) Exception() thrift.WritableException {
    return nil
}

func (x *respCF) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("respCF"); err != nil {
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

func (x *respCF) Read(p thrift.Protocol) error {
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
    }

    if err := p.ReadFieldEnd(); err != nil {
        return err
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *respCF) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("respCF({")
    sb.WriteString("})")

    return sb.String()
}
type reqCThing struct {
    A int32 `thrift:"a,1" json:"a" db:"a"`
    B string `thrift:"b,2" json:"b" db:"b"`
    C []int32 `thrift:"c,3" json:"c" db:"c"`
}
// Compile time interface enforcer
var _ thrift.Struct = &reqCThing{}

type CThingArgs = reqCThing

func newReqCThing() *reqCThing {
    return (&reqCThing{}).
        SetANonCompat(0).
        SetBNonCompat("").
        SetCNonCompat(nil)
}

func (x *reqCThing) GetANonCompat() int32 {
    return x.A
}

func (x *reqCThing) GetA() int32 {
    return x.A
}

func (x *reqCThing) GetBNonCompat() string {
    return x.B
}

func (x *reqCThing) GetB() string {
    return x.B
}

func (x *reqCThing) GetCNonCompat() []int32 {
    return x.C
}

func (x *reqCThing) GetC() []int32 {
    if !x.IsSetC() {
        return nil
    }

    return x.C
}

func (x *reqCThing) SetANonCompat(value int32) *reqCThing {
    x.A = value
    return x
}

func (x *reqCThing) SetA(value int32) *reqCThing {
    x.A = value
    return x
}

func (x *reqCThing) SetBNonCompat(value string) *reqCThing {
    x.B = value
    return x
}

func (x *reqCThing) SetB(value string) *reqCThing {
    x.B = value
    return x
}

func (x *reqCThing) SetCNonCompat(value []int32) *reqCThing {
    x.C = value
    return x
}

func (x *reqCThing) SetC(value []int32) *reqCThing {
    x.C = value
    return x
}

func (x *reqCThing) IsSetC() bool {
    return x.C != nil
}

func (x *reqCThing) writeField1(p thrift.Protocol) error {  // A
    if err := p.WriteFieldBegin("a", thrift.I32, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetANonCompat()
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *reqCThing) writeField2(p thrift.Protocol) error {  // B
    if err := p.WriteFieldBegin("b", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetBNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *reqCThing) writeField3(p thrift.Protocol) error {  // C
    if !x.IsSetC() {
        return nil
    }

    if err := p.WriteFieldBegin("c", thrift.SET, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetCNonCompat()
    if err := p.WriteSetBegin(thrift.I32, len(item)); err != nil {
    return thrift.PrependError("error writing set begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteI32(item); err != nil {
    return err
}
    }
}
if err := p.WriteSetEnd(); err != nil {
    return thrift.PrependError("error writing set end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *reqCThing) readField1(p thrift.Protocol) error {  // A
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.SetANonCompat(result)
    return nil
}

func (x *reqCThing) readField2(p thrift.Protocol) error {  // B
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetBNonCompat(result)
    return nil
}

func (x *reqCThing) readField3(p thrift.Protocol) error {  // C
    _ /* elemType */, size, err := p.ReadSetBegin()
if err != nil {
    return thrift.PrependError("error reading set begin: ", err)
}

setResult := make([]int32, 0, size)
for i := 0; i < size; i++ {
    var elem int32
    {
        result, err := p.ReadI32()
if err != nil {
    return err
}
        elem = result
    }
    setResult = append(setResult, elem)
}

if err := p.ReadSetEnd(); err != nil {
    return thrift.PrependError("error reading set end: ", err)
}
result := setResult

    x.SetCNonCompat(result)
    return nil
}

func (x *reqCThing) toString1() string {  // A
    return fmt.Sprintf("%v", x.GetANonCompat())
}

func (x *reqCThing) toString2() string {  // B
    return fmt.Sprintf("%v", x.GetBNonCompat())
}

func (x *reqCThing) toString3() string {  // C
    return fmt.Sprintf("%v", x.GetCNonCompat())
}


// Deprecated: Use reqCThing.Set* methods instead or set the fields directly.
type reqCThingBuilder struct {
    obj *reqCThing
}

func newReqCThingBuilder() *reqCThingBuilder {
    return &reqCThingBuilder{
        obj: newReqCThing(),
    }
}

func (x *reqCThingBuilder) A(value int32) *reqCThingBuilder {
    x.obj.A = value
    return x
}

func (x *reqCThingBuilder) B(value string) *reqCThingBuilder {
    x.obj.B = value
    return x
}

func (x *reqCThingBuilder) C(value []int32) *reqCThingBuilder {
    x.obj.C = value
    return x
}

func (x *reqCThingBuilder) Emit() *reqCThing {
    var objCopy reqCThing = *x.obj
    return &objCopy
}

func (x *reqCThing) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("reqCThing"); err != nil {
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

func (x *reqCThing) Read(p thrift.Protocol) error {
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
        case 1:  // a
            expectedType := thrift.Type(thrift.I32)
            if wireType == expectedType {
                if err := x.readField1(p); err != nil {
                   return err
                }
            } else {
                if err := p.Skip(wireType); err != nil {
                    return err
                }
            }
        case 2:  // b
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
        case 3:  // c
            expectedType := thrift.Type(thrift.SET)
            if wireType == expectedType {
                if err := x.readField3(p); err != nil {
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
    }

    if err := p.ReadFieldEnd(); err != nil {
        return err
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *reqCThing) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("reqCThing({")
    sb.WriteString(fmt.Sprintf("A:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("B:%s ", x.toString2()))
    sb.WriteString(fmt.Sprintf("C:%s", x.toString3()))
    sb.WriteString("})")

    return sb.String()
}
type respCThing struct {
    Value string `thrift:"value,0" json:"value" db:"value"`
    Bang *Bang `thrift:"bang,1,optional" json:"bang,omitempty" db:"bang"`
}
// Compile time interface enforcer
var _ thrift.Struct = &respCThing{}
var _ thrift.WritableResult = &respCThing{}

func newRespCThing() *respCThing {
    return (&respCThing{}).
        SetValueNonCompat("")
}

func (x *respCThing) GetValueNonCompat() string {
    return x.Value
}

func (x *respCThing) GetValue() string {
    return x.Value
}

func (x *respCThing) GetBangNonCompat() *Bang {
    return x.Bang
}

func (x *respCThing) GetBang() *Bang {
    if !x.IsSetBang() {
        return nil
    }

    return x.Bang
}

func (x *respCThing) SetValueNonCompat(value string) *respCThing {
    x.Value = value
    return x
}

func (x *respCThing) SetValue(value string) *respCThing {
    x.Value = value
    return x
}

func (x *respCThing) SetBangNonCompat(value Bang) *respCThing {
    x.Bang = &value
    return x
}

func (x *respCThing) SetBang(value *Bang) *respCThing {
    x.Bang = value
    return x
}

func (x *respCThing) IsSetBang() bool {
    return x.Bang != nil
}

func (x *respCThing) writeField0(p thrift.Protocol) error {  // Value
    if err := p.WriteFieldBegin("value", thrift.STRING, 0); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetValueNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *respCThing) writeField1(p thrift.Protocol) error {  // Bang
    if !x.IsSetBang() {
        return nil
    }

    if err := p.WriteFieldBegin("bang", thrift.STRUCT, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetBangNonCompat()
    if err := item.Write(p); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *respCThing) readField0(p thrift.Protocol) error {  // Value
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetValueNonCompat(result)
    return nil
}

func (x *respCThing) readField1(p thrift.Protocol) error {  // Bang
    result := *NewBang()
err := result.Read(p)
if err != nil {
    return err
}

    x.SetBangNonCompat(result)
    return nil
}

func (x *respCThing) toString0() string {  // Value
    return fmt.Sprintf("%v", x.GetValueNonCompat())
}

func (x *respCThing) toString1() string {  // Bang
    return fmt.Sprintf("%v", x.GetBangNonCompat())
}

// Deprecated: Use newRespCThing().GetBang() instead.
var respCThing_Bang_DEFAULT = newRespCThing().GetBang()

// Deprecated: Use newRespCThing().GetBang() instead.
func (x *respCThing) DefaultGetBang() *Bang {
    if !x.IsSetBang() {
        return NewBang()
    }
    return x.Bang
}


// Deprecated: Use respCThing.Set* methods instead or set the fields directly.
type respCThingBuilder struct {
    obj *respCThing
}

func newRespCThingBuilder() *respCThingBuilder {
    return &respCThingBuilder{
        obj: newRespCThing(),
    }
}

func (x *respCThingBuilder) Value(value string) *respCThingBuilder {
    x.obj.Value = value
    return x
}

func (x *respCThingBuilder) Bang(value *Bang) *respCThingBuilder {
    x.obj.Bang = value
    return x
}

func (x *respCThingBuilder) Emit() *respCThing {
    var objCopy respCThing = *x.obj
    return &objCopy
}

func (x *respCThing) Exception() thrift.WritableException {
    if x.Bang != nil {
        return x.Bang
    }
    return nil
}

func (x *respCThing) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("respCThing"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField0(p); err != nil {
        return err
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

func (x *respCThing) Read(p thrift.Protocol) error {
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
        case 0:  // value
            expectedType := thrift.Type(thrift.STRING)
            if wireType == expectedType {
                if err := x.readField0(p); err != nil {
                   return err
                }
            } else {
                if err := p.Skip(wireType); err != nil {
                    return err
                }
            }
        case 1:  // bang
            expectedType := thrift.Type(thrift.STRUCT)
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
    }

    if err := p.ReadFieldEnd(); err != nil {
        return err
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *respCThing) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("respCThing({")
    sb.WriteString(fmt.Sprintf("Value:%s ", x.toString0()))
    sb.WriteString(fmt.Sprintf("Bang:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}


type CProcessor struct {
    processorMap       map[string]thrift.ProcessorFunctionContext
    functionServiceMap map[string]string
    handler            C
}
// Compile time interface enforcer
var _ thrift.ProcessorContext = &CProcessor{}

func (p *CProcessor) AddToProcessorMap(key string, processor thrift.ProcessorFunctionContext) {
    p.processorMap[key] = processor
}

func (p *CProcessor) AddToFunctionServiceMap(key, service string) {
    p.functionServiceMap[key] = service
}

func (p *CProcessor) GetProcessorFunctionContext(key string) (processor thrift.ProcessorFunctionContext, err error) {
    if processor, ok := p.processorMap[key]; ok {
        return processor, nil
    }
    return nil, nil
}

func (p *CProcessor) ProcessorMap() map[string]thrift.ProcessorFunctionContext {
    return p.processorMap
}

func (p *CProcessor) FunctionServiceMap() map[string]string {
    return p.functionServiceMap
}

func NewCProcessor(handler C) *CProcessor {
    p := &CProcessor{
        handler:            handler,
        processorMap:       make(map[string]thrift.ProcessorFunctionContext),
        functionServiceMap: make(map[string]string),
    }
    p.AddToProcessorMap("f", &procFuncCF{handler: handler})
    p.AddToProcessorMap("thing", &procFuncCThing{handler: handler})
    p.AddToFunctionServiceMap("f", "C")
    p.AddToFunctionServiceMap("thing", "C")

    return p
}


type procFuncCF struct {
    handler C
}
// Compile time interface enforcer
var _ thrift.ProcessorFunctionContext = &procFuncCF{}

func (p *procFuncCF) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
    args := newReqCF()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncCF) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
    var err2 error
    messageType := thrift.REPLY
    switch result.(type) {
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("f", messageType, seqId); err2 != nil {
        err = err2
    }
    if err2 = result.Write(oprot); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.Flush(); err == nil && err2 != nil {
        err = err2
    }
    return err
}

func (p *procFuncCF) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    result := newRespCF()
    err := p.handler.F(ctx)
    if err != nil {
        x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing F: " + err.Error(), err)
        return x, x
    }

    return result, nil
}


type procFuncCThing struct {
    handler C
}
// Compile time interface enforcer
var _ thrift.ProcessorFunctionContext = &procFuncCThing{}

func (p *procFuncCThing) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
    args := newReqCThing()
    if err := args.Read(iprot); err != nil {
        return nil, err
    }
    iprot.ReadMessageEnd()
    return args, nil
}

func (p *procFuncCThing) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
    var err2 error
    messageType := thrift.REPLY
    switch v := result.(type) {
    case *Bang:
        result = &respCThing{
            Bang: v,
        }
    case thrift.ApplicationException:
        messageType = thrift.EXCEPTION
    }

    if err2 = oprot.WriteMessageBegin("thing", messageType, seqId); err2 != nil {
        err = err2
    }
    if err2 = result.Write(oprot); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
        err = err2
    }
    if err2 = oprot.Flush(); err == nil && err2 != nil {
        err = err2
    }
    return err
}

func (p *procFuncCThing) RunContext(ctx context.Context, reqStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
    args := reqStruct.(*reqCThing)
    result := newRespCThing()
    retval, err := p.handler.Thing(ctx, args.A, args.B, args.C)
    if err != nil {
        switch v := err.(type) {
        case *Bang:
            result.Bang = v
        default:
            x := thrift.NewApplicationExceptionCause(thrift.INTERNAL_ERROR, "Internal error processing doRaise: " + err.Error(), err)
            return x, x
        }
    }

    result.Value = retval
    return result, nil
}


