// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package bench

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Person struct {
	Name string `thrift:"Name,1" json:"Name"`
	Age  int32  `thrift:"Age,2" json:"Age"`
}

func NewPerson() *Person {
	return &Person{}
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int32 {
	return p.Age
}
func (p *Person) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Person) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *Person) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Age = v
	}
	return nil
}

func (p *Person) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Person"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Person) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Name", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:Name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return fmt.Errorf("%T.Name (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:Name: %s", p, err)
	}
	return err
}

func (p *Person) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Age", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:Age: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Age)); err != nil {
		return fmt.Errorf("%T.Age (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:Age: %s", p, err)
	}
	return err
}

func (p *Person) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Person(%+v)", *p)
}

type NoStrings struct {
	Abc int32   `thrift:"Abc,1" json:"Abc"`
	Def int64   `thrift:"Def,2" json:"Def"`
	Xyz []int16 `thrift:"Xyz,3" json:"Xyz"`
}

func NewNoStrings() *NoStrings {
	return &NoStrings{}
}

func (p *NoStrings) GetAbc() int32 {
	return p.Abc
}

func (p *NoStrings) GetDef() int64 {
	return p.Def
}

func (p *NoStrings) GetXyz() []int16 {
	return p.Xyz
}
func (p *NoStrings) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *NoStrings) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Abc = v
	}
	return nil
}

func (p *NoStrings) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Def = v
	}
	return nil
}

func (p *NoStrings) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]int16, 0, size)
	p.Xyz = tSlice
	for i := 0; i < size; i++ {
		var _elem0 int16
		if v, err := iprot.ReadI16(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_elem0 = v
		}
		p.Xyz = append(p.Xyz, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *NoStrings) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("NoStrings"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *NoStrings) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Abc", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:Abc: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Abc)); err != nil {
		return fmt.Errorf("%T.Abc (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:Abc: %s", p, err)
	}
	return err
}

func (p *NoStrings) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Def", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:Def: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Def)); err != nil {
		return fmt.Errorf("%T.Def (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:Def: %s", p, err)
	}
	return err
}

func (p *NoStrings) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Xyz", thrift.LIST, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:Xyz: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.I16, len(p.Xyz)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.Xyz {
		if err := oprot.WriteI16(int16(v)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:Xyz: %s", p, err)
	}
	return err
}

func (p *NoStrings) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NoStrings(%+v)", *p)
}

type Complex struct {
	Foo    string            `thrift:"Foo,1" json:"Foo"`
	Bar    int32             `thrift:"Bar,2" json:"Bar"`
	Baz    []int32           `thrift:"Baz,3" json:"Baz"`
	Abc    map[string]string `thrift:"Abc,4" json:"Abc"`
	People []*Person         `thrift:"People,5" json:"People"`
}

func NewComplex() *Complex {
	return &Complex{}
}

func (p *Complex) GetFoo() string {
	return p.Foo
}

func (p *Complex) GetBar() int32 {
	return p.Bar
}

func (p *Complex) GetBaz() []int32 {
	return p.Baz
}

func (p *Complex) GetAbc() map[string]string {
	return p.Abc
}

func (p *Complex) GetPeople() []*Person {
	return p.People
}
func (p *Complex) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Complex) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Foo = v
	}
	return nil
}

func (p *Complex) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Bar = v
	}
	return nil
}

func (p *Complex) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]int32, 0, size)
	p.Baz = tSlice
	for i := 0; i < size; i++ {
		var _elem1 int32
		if v, err := iprot.ReadI32(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_elem1 = v
		}
		p.Baz = append(p.Baz, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Complex) ReadField4(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]string, size)
	p.Abc = tMap
	for i := 0; i < size; i++ {
		var _key2 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key2 = v
		}
		var _val3 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val3 = v
		}
		p.Abc[_key2] = _val3
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *Complex) ReadField5(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*Person, 0, size)
	p.People = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &Person{}
		if err := _elem4.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem4, err)
		}
		p.People = append(p.People, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Complex) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Complex"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Complex) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Foo", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:Foo: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Foo)); err != nil {
		return fmt.Errorf("%T.Foo (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:Foo: %s", p, err)
	}
	return err
}

func (p *Complex) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Bar", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:Bar: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Bar)); err != nil {
		return fmt.Errorf("%T.Bar (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:Bar: %s", p, err)
	}
	return err
}

func (p *Complex) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Baz", thrift.LIST, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:Baz: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.I32, len(p.Baz)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.Baz {
		if err := oprot.WriteI32(int32(v)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:Baz: %s", p, err)
	}
	return err
}

func (p *Complex) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Abc", thrift.MAP, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:Abc: %s", p, err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Abc)); err != nil {
		return fmt.Errorf("error writing map begin: %s", err)
	}
	for k, v := range p.Abc {
		if err := oprot.WriteString(string(k)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return fmt.Errorf("error writing map end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:Abc: %s", p, err)
	}
	return err
}

func (p *Complex) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("People", thrift.LIST, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:People: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.People)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.People {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:People: %s", p, err)
	}
	return err
}

func (p *Complex) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Complex(%+v)", *p)
}

type CalculationData struct {
	A int32 `thrift:"a,1" json:"a"`
	B int32 `thrift:"b,2" json:"b"`
}

func NewCalculationData() *CalculationData {
	return &CalculationData{}
}

func (p *CalculationData) GetA() int32 {
	return p.A
}

func (p *CalculationData) GetB() int32 {
	return p.B
}
func (p *CalculationData) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CalculationData) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.A = v
	}
	return nil
}

func (p *CalculationData) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.B = v
	}
	return nil
}

func (p *CalculationData) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CalculationData"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *CalculationData) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("a", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:a: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.A)); err != nil {
		return fmt.Errorf("%T.a (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:a: %s", p, err)
	}
	return err
}

func (p *CalculationData) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("b", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:b: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.B)); err != nil {
		return fmt.Errorf("%T.b (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:b: %s", p, err)
	}
	return err
}

func (p *CalculationData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CalculationData(%+v)", *p)
}
