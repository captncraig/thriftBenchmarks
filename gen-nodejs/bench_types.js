//
// Autogenerated by Thrift Compiler (1.0.0-dev)
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//
var thrift = require('thrift');
var Thrift = thrift.Thrift;
var Q = thrift.Q;


var ttypes = module.exports = {};
if (typeof bench === 'undefined') {
  bench = {};
}
bench.Person = module.exports.Person = function(args) {
  this.Name = null;
  this.Age = null;
  if (args) {
    if (args.Name !== undefined) {
      this.Name = args.Name;
    }
    if (args.Age !== undefined) {
      this.Age = args.Age;
    }
  }
};
bench.Person.prototype = {};
bench.Person.prototype.read = function(input) {
  input.readStructBegin();
  while (true)
  {
    var ret = input.readFieldBegin();
    var fname = ret.fname;
    var ftype = ret.ftype;
    var fid = ret.fid;
    if (ftype == Thrift.Type.STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
      if (ftype == Thrift.Type.STRING) {
        this.Name = input.readString();
      } else {
        input.skip(ftype);
      }
      break;
      case 2:
      if (ftype == Thrift.Type.I32) {
        this.Age = input.readI32();
      } else {
        input.skip(ftype);
      }
      break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

bench.Person.prototype.write = function(output) {
  output.writeStructBegin('Person');
  if (this.Name !== null && this.Name !== undefined) {
    output.writeFieldBegin('Name', Thrift.Type.STRING, 1);
    output.writeString(this.Name);
    output.writeFieldEnd();
  }
  if (this.Age !== null && this.Age !== undefined) {
    output.writeFieldBegin('Age', Thrift.Type.I32, 2);
    output.writeI32(this.Age);
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

bench.NoStrings = module.exports.NoStrings = function(args) {
  this.Abc = null;
  this.Def = null;
  this.Xyz = null;
  if (args) {
    if (args.Abc !== undefined) {
      this.Abc = args.Abc;
    }
    if (args.Def !== undefined) {
      this.Def = args.Def;
    }
    if (args.Xyz !== undefined) {
      this.Xyz = args.Xyz;
    }
  }
};
bench.NoStrings.prototype = {};
bench.NoStrings.prototype.read = function(input) {
  input.readStructBegin();
  while (true)
  {
    var ret = input.readFieldBegin();
    var fname = ret.fname;
    var ftype = ret.ftype;
    var fid = ret.fid;
    if (ftype == Thrift.Type.STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
      if (ftype == Thrift.Type.I32) {
        this.Abc = input.readI32();
      } else {
        input.skip(ftype);
      }
      break;
      case 2:
      if (ftype == Thrift.Type.I64) {
        this.Def = input.readI64();
      } else {
        input.skip(ftype);
      }
      break;
      case 3:
      if (ftype == Thrift.Type.LIST) {
        var _size0 = 0;
        var _rtmp34;
        this.Xyz = [];
        var _etype3 = 0;
        _rtmp34 = input.readListBegin();
        _etype3 = _rtmp34.etype;
        _size0 = _rtmp34.size;
        for (var _i5 = 0; _i5 < _size0; ++_i5)
        {
          var elem6 = null;
          elem6 = input.readI16();
          this.Xyz.push(elem6);
        }
        input.readListEnd();
      } else {
        input.skip(ftype);
      }
      break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

bench.NoStrings.prototype.write = function(output) {
  output.writeStructBegin('NoStrings');
  if (this.Abc !== null && this.Abc !== undefined) {
    output.writeFieldBegin('Abc', Thrift.Type.I32, 1);
    output.writeI32(this.Abc);
    output.writeFieldEnd();
  }
  if (this.Def !== null && this.Def !== undefined) {
    output.writeFieldBegin('Def', Thrift.Type.I64, 2);
    output.writeI64(this.Def);
    output.writeFieldEnd();
  }
  if (this.Xyz !== null && this.Xyz !== undefined) {
    output.writeFieldBegin('Xyz', Thrift.Type.LIST, 3);
    output.writeListBegin(Thrift.Type.I16, this.Xyz.length);
    for (var iter7 in this.Xyz)
    {
      if (this.Xyz.hasOwnProperty(iter7))
      {
        iter7 = this.Xyz[iter7];
        output.writeI16(iter7);
      }
    }
    output.writeListEnd();
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

bench.Complex = module.exports.Complex = function(args) {
  this.Foo = null;
  this.Bar = null;
  this.Baz = null;
  this.Abc = null;
  this.People = null;
  if (args) {
    if (args.Foo !== undefined) {
      this.Foo = args.Foo;
    }
    if (args.Bar !== undefined) {
      this.Bar = args.Bar;
    }
    if (args.Baz !== undefined) {
      this.Baz = args.Baz;
    }
    if (args.Abc !== undefined) {
      this.Abc = args.Abc;
    }
    if (args.People !== undefined) {
      this.People = args.People;
    }
  }
};
bench.Complex.prototype = {};
bench.Complex.prototype.read = function(input) {
  input.readStructBegin();
  while (true)
  {
    var ret = input.readFieldBegin();
    var fname = ret.fname;
    var ftype = ret.ftype;
    var fid = ret.fid;
    if (ftype == Thrift.Type.STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
      if (ftype == Thrift.Type.STRING) {
        this.Foo = input.readString();
      } else {
        input.skip(ftype);
      }
      break;
      case 2:
      if (ftype == Thrift.Type.I32) {
        this.Bar = input.readI32();
      } else {
        input.skip(ftype);
      }
      break;
      case 3:
      if (ftype == Thrift.Type.LIST) {
        var _size8 = 0;
        var _rtmp312;
        this.Baz = [];
        var _etype11 = 0;
        _rtmp312 = input.readListBegin();
        _etype11 = _rtmp312.etype;
        _size8 = _rtmp312.size;
        for (var _i13 = 0; _i13 < _size8; ++_i13)
        {
          var elem14 = null;
          elem14 = input.readI32();
          this.Baz.push(elem14);
        }
        input.readListEnd();
      } else {
        input.skip(ftype);
      }
      break;
      case 4:
      if (ftype == Thrift.Type.MAP) {
        var _size15 = 0;
        var _rtmp319;
        this.Abc = {};
        var _ktype16 = 0;
        var _vtype17 = 0;
        _rtmp319 = input.readMapBegin();
        _ktype16 = _rtmp319.ktype;
        _vtype17 = _rtmp319.vtype;
        _size15 = _rtmp319.size;
        for (var _i20 = 0; _i20 < _size15; ++_i20)
        {
          var key21 = null;
          var val22 = null;
          key21 = input.readString();
          val22 = input.readString();
          this.Abc[key21] = val22;
        }
        input.readMapEnd();
      } else {
        input.skip(ftype);
      }
      break;
      case 5:
      if (ftype == Thrift.Type.LIST) {
        var _size23 = 0;
        var _rtmp327;
        this.People = [];
        var _etype26 = 0;
        _rtmp327 = input.readListBegin();
        _etype26 = _rtmp327.etype;
        _size23 = _rtmp327.size;
        for (var _i28 = 0; _i28 < _size23; ++_i28)
        {
          var elem29 = null;
          elem29 = new ttypes.Person();
          elem29.read(input);
          this.People.push(elem29);
        }
        input.readListEnd();
      } else {
        input.skip(ftype);
      }
      break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

bench.Complex.prototype.write = function(output) {
  output.writeStructBegin('Complex');
  if (this.Foo !== null && this.Foo !== undefined) {
    output.writeFieldBegin('Foo', Thrift.Type.STRING, 1);
    output.writeString(this.Foo);
    output.writeFieldEnd();
  }
  if (this.Bar !== null && this.Bar !== undefined) {
    output.writeFieldBegin('Bar', Thrift.Type.I32, 2);
    output.writeI32(this.Bar);
    output.writeFieldEnd();
  }
  if (this.Baz !== null && this.Baz !== undefined) {
    output.writeFieldBegin('Baz', Thrift.Type.LIST, 3);
    output.writeListBegin(Thrift.Type.I32, this.Baz.length);
    for (var iter30 in this.Baz)
    {
      if (this.Baz.hasOwnProperty(iter30))
      {
        iter30 = this.Baz[iter30];
        output.writeI32(iter30);
      }
    }
    output.writeListEnd();
    output.writeFieldEnd();
  }
  if (this.Abc !== null && this.Abc !== undefined) {
    output.writeFieldBegin('Abc', Thrift.Type.MAP, 4);
    output.writeMapBegin(Thrift.Type.STRING, Thrift.Type.STRING, Thrift.objectLength(this.Abc));
    for (var kiter31 in this.Abc)
    {
      if (this.Abc.hasOwnProperty(kiter31))
      {
        var viter32 = this.Abc[kiter31];
        output.writeString(kiter31);
        output.writeString(viter32);
      }
    }
    output.writeMapEnd();
    output.writeFieldEnd();
  }
  if (this.People !== null && this.People !== undefined) {
    output.writeFieldBegin('People', Thrift.Type.LIST, 5);
    output.writeListBegin(Thrift.Type.STRUCT, this.People.length);
    for (var iter33 in this.People)
    {
      if (this.People.hasOwnProperty(iter33))
      {
        iter33 = this.People[iter33];
        iter33.write(output);
      }
    }
    output.writeListEnd();
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

bench.CalculationData = module.exports.CalculationData = function(args) {
  this.a = null;
  this.b = null;
  if (args) {
    if (args.a !== undefined) {
      this.a = args.a;
    }
    if (args.b !== undefined) {
      this.b = args.b;
    }
  }
};
bench.CalculationData.prototype = {};
bench.CalculationData.prototype.read = function(input) {
  input.readStructBegin();
  while (true)
  {
    var ret = input.readFieldBegin();
    var fname = ret.fname;
    var ftype = ret.ftype;
    var fid = ret.fid;
    if (ftype == Thrift.Type.STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
      if (ftype == Thrift.Type.I32) {
        this.a = input.readI32();
      } else {
        input.skip(ftype);
      }
      break;
      case 2:
      if (ftype == Thrift.Type.I32) {
        this.b = input.readI32();
      } else {
        input.skip(ftype);
      }
      break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

bench.CalculationData.prototype.write = function(output) {
  output.writeStructBegin('CalculationData');
  if (this.a !== null && this.a !== undefined) {
    output.writeFieldBegin('a', Thrift.Type.I32, 1);
    output.writeI32(this.a);
    output.writeFieldEnd();
  }
  if (this.b !== null && this.b !== undefined) {
    output.writeFieldBegin('b', Thrift.Type.I32, 2);
    output.writeI32(this.b);
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};
