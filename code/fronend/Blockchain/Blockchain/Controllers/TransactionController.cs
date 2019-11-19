using Blockchain.API;
using Blockchain.Models;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;

namespace Blockchain.Controllers
{
    public class TransactionController : Controller
    {
        private TransactionApi client = new TransactionApi();
        // GET: Transaction
        
        public ActionResult RequestPort()
        {
 
            return View();
        }

        
        public ActionResult Index(string port)
        {
            if (port == null)
            { port = TempData["puertoIndex"].ToString(); }
            ViewBag.port = port;
            
      
            IEnumerable<Transaction> transacciones = null;
            Task<HttpResponseMessage> response = client.Initial(port).GetAsync("transacciones");
            response.Wait();
            HttpResponseMessage result = response.Result;
            if (result.IsSuccessStatusCode)
            {
                Task<IList<Transaction>> readTask = result.Content.ReadAsAsync<IList<Transaction>>();
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
        public ActionResult Create(string port)
        {
            ViewBag.port = port;
            TempData["puerto"] = port;
            return View();
        }


        // POST: Transaction/Create
        [HttpPost]
        [ValidateAntiForgeryToken]
        public ActionResult Create(string port,Transaction transaction )
        {
            port = TempData["puerto"].ToString();
            Task<HttpResponseMessage> postTask = client.Initial(port).PostAsJsonAsync("transaccion", transaction);
            postTask.Wait();
            HttpResponseMessage result = postTask.Result;
            if (result.IsSuccessStatusCode)
            {
                TempData["puertoIndex"] = port;
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