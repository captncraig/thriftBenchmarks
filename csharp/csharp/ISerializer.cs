using System.IO;
using System.Text;
using Newtonsoft.Json;
using Thrift.Protocol;

namespace csharp
{
    public interface ISerializer
    {
        void Serialize(TBase t);
        void Deserialize(TBase t);
    }

    class JSerializer : ISerializer
    {
        private string s;


        public void Serialize(TBase t)
        {
            s = JsonConvert.SerializeObject(t);
        }

        public void Deserialize(TBase t)
        {
            JsonConvert.DeserializeObject(s, t.GetType());
        }
    }

    class ProtocolSerializer : ISerializer
    {
        private readonly TProtocol _prot;

        public ProtocolSerializer(TProtocol prot)
        {
            this._prot = prot;
        }

        public void Serialize(TBase t)
        {
            t.Write(_prot);
        }

        public void Deserialize(TBase t)
        {
            t.Read(_prot);
        }
    }
}