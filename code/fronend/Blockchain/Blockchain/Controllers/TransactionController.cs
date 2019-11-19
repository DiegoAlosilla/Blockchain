using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading.Tasks;

using Blockchain.API;
using Blockchain.Models;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace Blockchain.Controllers
{
    public class TransactionController : Controller
    {
        TransactionApi client = new TransactionApi();
        // GET: Transaction
        public ActionResult Index()
        {
            IEnumerable<Transaction> transacciones = null;
            var response = client.Initial().GetAsync("transacciones");
            response.Wait();
            var result = response.Result;            
            if (result.IsSuccessStatusCode)
            {
                var readTask = result.Content.ReadAsAsync<IList<Transaction>>();
               // var readTask = result.Content.ReadAsStringAsync();
                readTask.Wait();
                //transacciones = JsonConvert.DeserializeObject<IList<Transaction>>(readTask.Result);
                transacciones = readTask.Result;
            }
            else //web api sent error response 
            {
                transacciones = Enumerable.Empty<Transaction>();
                ModelState.AddModelError(string.Empty, "Server error. Please contact administrator.");
            }
            return View(transacciones);
        }

        // GET: Transaction/Details/5
        public ActionResult Details(int id)
        {

            return View();
        }

        // GET: Transaction/Create
        public ActionResult Create()
        {
            return View();
        }

        // POST: Transaction/Create
        [HttpPost]
        [ValidateAntiForgeryToken]
        public ActionResult Create(Transaction transaction)
        {

            var postTask = client.Initial().PostAsJsonAsync("transaccion", transaction);
            postTask.Wait();
            var result = postTask.Result;
            if (result.IsSuccessStatusCode)
            {
                return RedirectToAction("Index");
            }

            ModelState.AddModelError(string.Empty, "Server Error. Please contact administrator.");
            return View(transaction);

        }

        // GET: Transaction/Edit/5
        public ActionResult Edit(int id)
        {
            return View();
        }

        // POST: Transaction/Edit/5
        [HttpPost]
        [ValidateAntiForgeryToken]
        public ActionResult Edit(int id, IFormCollection collection)
        {
            try
            {
                // TODO: Add update logic here

                return RedirectToAction(nameof(Index));
            }
            catch
            {
                return View();
            }
        }

        // GET: Transaction/Delete/5
        public ActionResult Delete(int id)
        {
            return View();
        }

        // POST: Transaction/Delete/5
        [HttpPost]
        [ValidateAntiForgeryToken]
        public ActionResult Delete(int id, IFormCollection collection)
        {
            try
            {
                // TODO: Add delete logic here

                return RedirectToAction(nameof(Index));
            }
            catch
            {
                return View();
            }
        }
    }
}