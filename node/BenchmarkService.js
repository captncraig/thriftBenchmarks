//
// Autogenerated by Thrift Compiler (1.0.0-dev)
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//
var thrift = require('thrift');
var Thrift = thrift.Thrift;


var ttypes = require('./bench_types');
//HELPER FUNCTIONS AND STRUCTURES

bench.BenchmarkService_Add_args = function(args) {
  this.d = null;
  if (args) {
    if (args.d !== undefined) {
      this.d = args.d;
    }
  }
};
bench.BenchmarkService_Add_args.prototype = {};
bench.BenchmarkService_Add_args.prototype.read = function(input) {
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
      if (ftype == Thrift.Type.STRUCT) {
        this.d = new ttypes.CalculationData();
        this.d.read(input);
      } else {
        input.skip(ftype);
      }
      break;
      case 0:
        input.skip(ftype);
        break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

bench.BenchmarkService_Add_args.prototype.write = function(output) {
  output.writeStructBegin('BenchmarkService_Add_args');
  if (this.d !== null && this.d !== undefined) {
    output.writeFieldBegin('d', Thrift.Type.STRUCT, 1);
    this.d.write(output);
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

bench.BenchmarkService_Add_result = function(args) {
  this.success = null;
  if (args) {
    if (args.success !== undefined) {
      this.success = args.success;
    }
  }
};
bench.BenchmarkService_Add_result.prototype = {};
bench.BenchmarkService_Add_result.prototype.read = function(input) {
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
      case 0:
      if (ftype == Thrift.Type.I32) {
        this.success = input.readI32();
      } else {
        input.skip(ftype);
      }
      break;
      case 0:
        input.skip(ftype);
        break;
      default:
        input.skip(ftype);
    }
    input.readFieldEnd();
  }
  input.readStructEnd();
  return;
};

bench.BenchmarkService_Add_result.prototype.write = function(output) {
  output.writeStructBegin('BenchmarkService_Add_result');
  if (this.success !== null && this.success !== undefined) {
    output.writeFieldBegin('success', Thrift.Type.I32, 0);
    output.writeI32(this.success);
    output.writeFieldEnd();
  }
  output.writeFieldStop();
  output.writeStructEnd();
  return;
};

bench.BenchmarkServiceClient = exports.Client = function(output, pClass) {
    this.output = output;
    this.pClass = pClass;
    this._seqid = 0;
    this._reqs = {};
};
bench.BenchmarkServiceClient.prototype = {};
bench.BenchmarkServiceClient.prototype.seqid = function() { return this._seqid; }
bench.BenchmarkServiceClient.prototype.new_seqid = function() { return this._seqid += 1; }
bench.BenchmarkServiceClient.prototype.Add = function(d, callback) {
  this._seqid = this.new_seqid();
  if (callback === undefined) {
    var _defer = Q.defer();
    this._reqs[this.seqid()] = function(error, result) {
      if (error) {
        _defer.reject(error);
      } else {
        _defer.resolve(result);
      }
    };
    this.send_Add(d);
    return _defer.promise;
  } else {
    this._reqs[this.seqid()] = callback;
    this.send_Add(d);
  }
};

bench.BenchmarkServiceClient.prototype.send_Add = function(d) {
  var output = new this.pClass(this.output);
  output.writeMessageBegin('Add', Thrift.MessageType.CALL, this.seqid());
  var args = new bench.BenchmarkService_Add_args();
  args.d = d;
  args.write(output);
  output.writeMessageEnd();
  return this.output.flush();
};

bench.BenchmarkServiceClient.prototype.recv_Add = function(input,mtype,rseqid) {
  var callback = this._reqs[rseqid] || function() {};
  delete this._reqs[rseqid];
  if (mtype == Thrift.MessageType.EXCEPTION) {
    var x = new Thrift.TApplicationException();
    x.read(input);
    input.readMessageEnd();
    return callback(x);
  }
  var result = new bench.BenchmarkService_Add_result();
  result.read(input);
  input.readMessageEnd();

  if (null !== result.success) {
    return callback(null, result.success);
  }
  return callback('Add failed: unknown result');
};
bench.BenchmarkServiceProcessor = exports.Processor = function(handler) {
  this._handler = handler
}
bench.BenchmarkServiceProcessor.prototype.process = function(input, output) {
  var r = input.readMessageBegin();
  if (this['process_' + r.fname]) {
    return this['process_' + r.fname].call(this, r.rseqid, input, output);
  } else {
    input.skip(Thrift.Type.STRUCT);
    input.readMessageEnd();
    var x = new Thrift.TApplicationException(Thrift.TApplicationExceptionType.UNKNOWN_METHOD, 'Unknown function ' + r.fname);
    output.writeMessageBegin(r.fname, Thrift.MessageType.EXCEPTION, r.rseqid);
    x.write(output);
    output.writeMessageEnd();
    output.flush();
  }
}

bench.BenchmarkServiceProcessor.prototype.process_Add = function(seqid, input, output) {
  var args = new bench.BenchmarkService_Add_args();
  args.read(input);
  input.readMessageEnd();
  if (this._handler.Add.length === 1) {
    Q.fcall(this._handler.Add, args.d)
      .then(function(result) {
        var result = new bench.BenchmarkService_Add_result({success: result});
        output.writeMessageBegin("Add", Thrift.MessageType.REPLY, seqid);
        result.write(output);
        output.writeMessageEnd();
        output.flush();
      }, function (err) {
        var result = new bench.BenchmarkService_Add_result(err);
        output.writeMessageBegin("Add", Thrift.MessageType.REPLY, seqid);
        result.write(output);
        output.writeMessageEnd();
        output.flush();
      });
  } else {
    this._handler.Add(args.d,  function (err, result) {
      var result = new bench.BenchmarkService_Add_result((err != null ? err : {success: result}));
      output.writeMessageBegin("Add", Thrift.MessageType.REPLY, seqid);
      result.write(output);
      output.writeMessageEnd();
      output.flush();
    });
  }
}

