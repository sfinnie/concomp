using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;

namespace ConComp
{
    internal class Program
    {
        private static void Main(string[] args)
        {
            var tasks = new Dictionary<string, Task<long>>();

            foreach (var arg in args)
            {
                var task = Task<long>.Factory.StartNew(() => new FileInfo(arg).Length);
                tasks.Add(arg, task);
            }

            Task.WaitAll(tasks.Values.Cast<Task>().ToArray());

            var maxFileSize = tasks.Max(t => t.Value.Result);
            var biggests = tasks.Where(t => t.Value.Result == maxFileSize).ToList();

            if (biggests.Count == tasks.Count)
            {
                Console.WriteLine("All files are even");
            }
            else if (biggests.Count > 1)
            {
                var all = string.Join(", ", biggests.Select(b => b.Key));
                Console.WriteLine($"{all} are the biggest");
            }
            else
            {
                var biggest = biggests.Single();
                Console.WriteLine($"{biggest.Key} is the biggest");
            }
        }
    }
}