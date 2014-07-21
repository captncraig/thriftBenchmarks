package main

import (
	ejson "encoding/json"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	//"strings"
	"thriftBenchmarks/gen-go/bench"
	"time"
)

var buf = thrift.NewTMemoryBuffer()
var json thrift.TProtocol
var binary thrift.TProtocol
var compact thrift.TProtocol
var person writable
var complicated writable
var noStrings writable

func init() {
	p := bench.NewPerson()
	p.Age = 42
	p.Name = "Bob"
	person = p

	c := bench.NewComplex()
	complicated = c
	c.Foo = "ABCDEFGHIJKLMNOP"
	c.Bar = 12345
	c.Baz = []int32{1, 2, 3, 4, 5, 87654, 12345, 999999, 765433211}
	c.Abc = map[string]string{
		"A":           "BBBB",
		"ASDASFASF":   "ERWERWER",
		"23423SDGSDG": "SFASFASFKLASJGLKASJGLKASFJALKSFJALK",
	}
	c.People = []*bench.Person{p, p, p, p, p, p, p, p}

	n := bench.NewNoStrings()
	n.Abc = 56743
	n.Def = 123456789066
	n.Xyz = []int16{12345, 4567, 1, 1, 1, 1, 1, 42, 56, 65, 456}
	noStrings = n

	json = thrift.NewTJSONProtocol(buf)
	compact = thrift.NewTCompactProtocol(buf)
	binary = thrift.NewTBinaryProtocol(buf, true, true)
}

type writable interface {
	Read(thrift.TProtocol) error
	Write(thrift.TProtocol) error
}

func main() {
	sizeTest(person, "Simple Person")
	sizeTest(complicated, "Complex object")
	sizeTest(noStrings, "Numeric Object")
	//speedTest(person, "Simple Person")
	//speedTest(complicated, "Complex object")
	//speedTest(noStrings, "No Strings object")
}

var iterations = 1000000

func timeThing(n string, f func()) {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		f()
	}
	duration := time.Now().Sub(start)
	fmt.Printf("%v: %v (%v)\n", n, duration.String(), (duration / time.Duration(iterations)).String())
}
func speedTest(o writable, name string) {
	fmt.Printf("------------%v Speed\n", name)

	timeThing("json", func() {
		buf.Reset()
		o.Write(json)
	})

	by := buf.Bytes()
	timeThing("json - d", func() {
		buf.Reset()
		buf.Write(by)
		o.Read(json)
	})

	timeThing("binary", func() {
		buf.Reset()
		o.Write(binary)
	})

	by = buf.Bytes()
	timeThing("binary - d", func() {
		buf.Reset()
		buf.Write(by)
		o.Read(binary)
	})

	timeThing("compact", func() {
		buf.Reset()
		o.Write(compact)
	})

	by = buf.Bytes()
	timeThing("compact - d", func() {
		buf.Reset()
		buf.Write(by)
		o.Read(compact)
	})

	timeThing("encoding/json", func() {
		by, _ = ejson.Marshal(o)
	})

	timeThing("encoding/json - d", func() {
		ejson.Unmarshal(by, o)
	})
}

func sizeTest(o writable, name string) {
	fmt.Printf("--------%v Size\n", name)
	buf.Reset()
	o.Write(json)
	json.Flush()
	fmt.Printf("JSON: %d\n", buf.Len())
	fmt.Println(buf.String())

	buf.Reset()
	o.Write(binary)
	fmt.Printf("Binary: %d\n", buf.Len())
	//fmt.Println(b.String())

	buf.Reset()
	o.Write(compact)
	fmt.Printf("Compact: %d\n", buf.Len())
	//fmt.Println(b.String())

	x, _ := ejson.Marshal(o)
	fmt.Printf("Standard json: %d\n", len(x))
	//fmt.Println(string(x))
}
