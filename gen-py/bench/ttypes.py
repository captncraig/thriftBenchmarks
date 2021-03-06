#
# Autogenerated by Thrift Compiler (1.0.0-dev)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#
#  options string: py
#

from thrift.Thrift import TType, TMessageType, TException, TApplicationException

from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol, TProtocol
try:
  from thrift.protocol import fastbinary
except:
  fastbinary = None



class Person:
  """
  Attributes:
   - Name
   - Age
  """

  thrift_spec = (
    None, # 0
    (1, TType.STRING, 'Name', None, None, ), # 1
    (2, TType.I32, 'Age', None, None, ), # 2
  )

  def __init__(self, Name=None, Age=None,):
    self.Name = Name
    self.Age = Age

  def read(self, iprot):
    if iprot.__class__ == TBinaryProtocol.TBinaryProtocolAccelerated and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastbinary is not None:
      fastbinary.decode_binary(self, iprot.trans, (self.__class__, self.thrift_spec))
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break
      if fid == 1:
        if ftype == TType.STRING:
          self.Name = iprot.readString();
        else:
          iprot.skip(ftype)
      elif fid == 2:
        if ftype == TType.I32:
          self.Age = iprot.readI32();
        else:
          iprot.skip(ftype)
      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if oprot.__class__ == TBinaryProtocol.TBinaryProtocolAccelerated and self.thrift_spec is not None and fastbinary is not None:
      oprot.trans.write(fastbinary.encode_binary(self, (self.__class__, self.thrift_spec)))
      return
    oprot.writeStructBegin('Person')
    if self.Name is not None:
      oprot.writeFieldBegin('Name', TType.STRING, 1)
      oprot.writeString(self.Name)
      oprot.writeFieldEnd()
    if self.Age is not None:
      oprot.writeFieldBegin('Age', TType.I32, 2)
      oprot.writeI32(self.Age)
      oprot.writeFieldEnd()
    oprot.writeFieldStop()
    oprot.writeStructEnd()

  def validate(self):
    return


  def __repr__(self):
    L = ['%s=%r' % (key, value)
      for key, value in self.__dict__.iteritems()]
    return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

  def __eq__(self, other):
    return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

  def __ne__(self, other):
    return not (self == other)

class Complex:
  """
  Attributes:
   - Foo
   - Bar
   - Baz
   - Abc
   - People
  """

  thrift_spec = (
    None, # 0
    (1, TType.STRING, 'Foo', None, None, ), # 1
    (2, TType.I32, 'Bar', None, None, ), # 2
    (3, TType.LIST, 'Baz', (TType.I32,None), None, ), # 3
    (4, TType.MAP, 'Abc', (TType.STRING,None,TType.STRING,None), None, ), # 4
    (5, TType.LIST, 'People', (TType.STRUCT,(Person, Person.thrift_spec)), None, ), # 5
  )

  def __init__(self, Foo=None, Bar=None, Baz=None, Abc=None, People=None,):
    self.Foo = Foo
    self.Bar = Bar
    self.Baz = Baz
    self.Abc = Abc
    self.People = People

  def read(self, iprot):
    if iprot.__class__ == TBinaryProtocol.TBinaryProtocolAccelerated and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastbinary is not None:
      fastbinary.decode_binary(self, iprot.trans, (self.__class__, self.thrift_spec))
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break
      if fid == 1:
        if ftype == TType.STRING:
          self.Foo = iprot.readString();
        else:
          iprot.skip(ftype)
      elif fid == 2:
        if ftype == TType.I32:
          self.Bar = iprot.readI32();
        else:
          iprot.skip(ftype)
      elif fid == 3:
        if ftype == TType.LIST:
          self.Baz = []
          (_etype3, _size0) = iprot.readListBegin()
          for _i4 in xrange(_size0):
            _elem5 = iprot.readI32();
            self.Baz.append(_elem5)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 4:
        if ftype == TType.MAP:
          self.Abc = {}
          (_ktype7, _vtype8, _size6 ) = iprot.readMapBegin()
          for _i10 in xrange(_size6):
            _key11 = iprot.readString();
            _val12 = iprot.readString();
            self.Abc[_key11] = _val12
          iprot.readMapEnd()
        else:
          iprot.skip(ftype)
      elif fid == 5:
        if ftype == TType.LIST:
          self.People = []
          (_etype16, _size13) = iprot.readListBegin()
          for _i17 in xrange(_size13):
            _elem18 = Person()
            _elem18.read(iprot)
            self.People.append(_elem18)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if oprot.__class__ == TBinaryProtocol.TBinaryProtocolAccelerated and self.thrift_spec is not None and fastbinary is not None:
      oprot.trans.write(fastbinary.encode_binary(self, (self.__class__, self.thrift_spec)))
      return
    oprot.writeStructBegin('Complex')
    if self.Foo is not None:
      oprot.writeFieldBegin('Foo', TType.STRING, 1)
      oprot.writeString(self.Foo)
      oprot.writeFieldEnd()
    if self.Bar is not None:
      oprot.writeFieldBegin('Bar', TType.I32, 2)
      oprot.writeI32(self.Bar)
      oprot.writeFieldEnd()
    if self.Baz is not None:
      oprot.writeFieldBegin('Baz', TType.LIST, 3)
      oprot.writeListBegin(TType.I32, len(self.Baz))
      for iter19 in self.Baz:
        oprot.writeI32(iter19)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.Abc is not None:
      oprot.writeFieldBegin('Abc', TType.MAP, 4)
      oprot.writeMapBegin(TType.STRING, TType.STRING, len(self.Abc))
      for kiter20,viter21 in self.Abc.items():
        oprot.writeString(kiter20)
        oprot.writeString(viter21)
      oprot.writeMapEnd()
      oprot.writeFieldEnd()
    if self.People is not None:
      oprot.writeFieldBegin('People', TType.LIST, 5)
      oprot.writeListBegin(TType.STRUCT, len(self.People))
      for iter22 in self.People:
        iter22.write(oprot)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    oprot.writeFieldStop()
    oprot.writeStructEnd()

  def validate(self):
    return


  def __repr__(self):
    L = ['%s=%r' % (key, value)
      for key, value in self.__dict__.iteritems()]
    return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

  def __eq__(self, other):
    return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

  def __ne__(self, other):
    return not (self == other)

class CalculationData:
  """
  Attributes:
   - a
   - b
  """

  thrift_spec = (
    None, # 0
    (1, TType.I32, 'a', None, None, ), # 1
    (2, TType.I32, 'b', None, None, ), # 2
  )

  def __init__(self, a=None, b=None,):
    self.a = a
    self.b = b

  def read(self, iprot):
    if iprot.__class__ == TBinaryProtocol.TBinaryProtocolAccelerated and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastbinary is not None:
      fastbinary.decode_binary(self, iprot.trans, (self.__class__, self.thrift_spec))
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break
      if fid == 1:
        if ftype == TType.I32:
          self.a = iprot.readI32();
        else:
          iprot.skip(ftype)
      elif fid == 2:
        if ftype == TType.I32:
          self.b = iprot.readI32();
        else:
          iprot.skip(ftype)
      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if oprot.__class__ == TBinaryProtocol.TBinaryProtocolAccelerated and self.thrift_spec is not None and fastbinary is not None:
      oprot.trans.write(fastbinary.encode_binary(self, (self.__class__, self.thrift_spec)))
      return
    oprot.writeStructBegin('CalculationData')
    if self.a is not None:
      oprot.writeFieldBegin('a', TType.I32, 1)
      oprot.writeI32(self.a)
      oprot.writeFieldEnd()
    if self.b is not None:
      oprot.writeFieldBegin('b', TType.I32, 2)
      oprot.writeI32(self.b)
      oprot.writeFieldEnd()
    oprot.writeFieldStop()
    oprot.writeStructEnd()

  def validate(self):
    return


  def __repr__(self):
    L = ['%s=%r' % (key, value)
      for key, value in self.__dict__.iteritems()]
    return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

  def __eq__(self, other):
    return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

  def __ne__(self, other):
    return not (self == other)
