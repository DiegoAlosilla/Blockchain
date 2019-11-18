using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;

namespace Blockchain.API
{
    public class TransactionApi
    {
        public HttpClient Initial()
        {
            var Client = new HttpClient();
            Client.BaseAddress = new Uri("http://localhost:9000/app/");
            return Client;
        }
    }
}
