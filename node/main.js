var thrift = require('thrift');
var g = require('./bench_types')

transport = new thrift.TBufferedTransport();

json = new thrift.TJSONProtocol(transport);
binary = new thrift.TBinaryProtocol(transport,true,true);
compact = new thrift.TCompactProtocol(transport);

person = new g.Person();
person.Name = "Joe";
person.Age = 21;