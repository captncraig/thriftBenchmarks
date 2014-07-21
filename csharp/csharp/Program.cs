using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Text;
using bench;
using Thrift.Protocol;
using Thrift.Transport;

namespace csharp
{
    class Program
    {
        private static readonly string Person = "{'1':{'str':'Bob'},'2':{'i32':42}}".Replace('\'','"');
        private static string _noStrings = "{'1':{'i32':56743},'2':{'i64':123456789066},'3':{'lst':['i16',11,12345,4567,1,1,1,1,1,42,56,65,456]}}".Replace('\'', '"');
        private static readonly string Complex = "{'1':{'str':'ABCDEFGHIJKLMNOP'},'2':{'i32':12345},'3':{'lst':['i32',9,1,2,3,4,5,87654,12345,999999,765433211]},'4':{'map':['str','str',3,'A','BBBB','ASDASFASF','ERWERWER','23423SDGSDG','SFASFASFKLASJGLKASJGLKASFJALKSFJALK']},'5':{'lst':['rec',8,{'1':{'str':'Bob'},'2':{'i32':42}},{'1':{'str':'Bob'},'2':{'i32':42}},{'1':{'str':'Bob'},'2':{'i32':42}},{'1':{'str':'Bob'},'2':{'i32':42}},{'1':{'str':'Bob'},'2':{'i32':42}},{'1':{'str':'Bob'},'2':{'i32':42}},{'1':{'str':'Bob'},'2':{'i32':42}},{'1':{'str':'Bob'},'2':{'i32':42}}]}}".Replace('\'', '"');
        static void Main(string[] args)
        {
            var inStream = new MemoryStream();
            var buffer = new TStreamTransport(inStream,inStream);
            var json = new TJSONProtocol(buffer);
            var compact = new TCompactProtocol(buffer);
            var binary = new TBinaryProtocol(buffer);

            var person = new Person {Age = 42, Name = "Bob"};
            var complex = new Complex
            {
                Foo = "ABCDEFGHIJKLMNOP",
                Bar = 12345,
                Baz = new List<int> {1, 2, 3, 4, 5, 87654, 12345, 999999, 765433211},
                People = new List<Person> {person, person, person, person, person, person, person, person},
                Abc = new Dictionary<string, string>
                {
                    {"A", "BBBB"},
                    {"ASDASFASF", "ERWERWER"},
                    {"23423SDGSDG", "SFASFASFKLASJGLKASJGLKASFJALKSFJALK"},
                }
            };
            var noStrings = new NoStrings
            {
                Abc = 56743,
                Def = 123456789066,
                Xyz = new List<short> {12345, 4567, 1, 1, 1, 1, 1, 42, 56, 65, 456}
            };
            var objects = new Dictionary<string, TBase>
            {
                {"person", person},
                {"complex", complex},
                {"numeric", noStrings}
            };
            var serializers = new Dictionary<string, ISerializer>
            {
                {"json", new ProtocolSerializer(json)},
                {"binary", new ProtocolSerializer(binary)},
                {"compact", new ProtocolSerializer(compact)},
                {"json.net", new JSerializer()},
            };

            const int iterations = 1000000;
            foreach (var o in objects)
            {
                Console.WriteLine("----------"+o.Key);
                foreach (var serializer in serializers)
                {
                    var watch = Stopwatch.StartNew();
                    for (int i = 0; i < iterations; i++)
                    {
                        inStream.SetLength(0);
                        inStream.Position = 0;
                        serializer.Value.Serialize(o.Value);
                    }
                    watch.Stop();
                    var millisPer = (double)watch.ElapsedMilliseconds/iterations;
                    var microsPer = millisPer*1000;
                    Console.WriteLine("{0} -Serialize- {1} ({2})", serializer.Key, watch.ElapsedMilliseconds, microsPer);

                    watch = Stopwatch.StartNew();
                    for (int i = 0; i < iterations; i++)
                    {
                        inStream.Position = 0;
                        serializer.Value.Deserialize(o.Value);
                    }
                    watch.Stop();
                    millisPer = (double)watch.ElapsedMilliseconds / iterations;
                    microsPer = millisPer * 1000;
                    Console.WriteLine("{0} -Deserialize- {1} ({2})", serializer.Key, watch.ElapsedMilliseconds, microsPer);
                    
                }
            }
            Console.ReadLine();
        }
    }
}
