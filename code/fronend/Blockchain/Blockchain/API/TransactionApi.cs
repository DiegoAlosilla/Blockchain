﻿using System;
using System.Net.Http;

namespace Blockchain.API
{
    public class TransactionApi
    {


        public HttpClient Initial( string Port)
        {
            HttpClient Client = new HttpClient
            {
                BaseAddress = new Uri("http://localhost:"+ Port + "/app/")
            };
            return Client;
        }
    }
}
