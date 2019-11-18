using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Blockchain.Models
{
    public class Transaction
    {
        public string monto { get; set; }
        public string origen { get; set; }
        public string destino { get; set; }

        [JsonConstructor]
        public Transaction(string monto, string origen, string destino)
        {
            this.monto = monto;
            this.origen = origen;
            this.destino = destino;
        }

        public Transaction()
        {
        }
    }
}
