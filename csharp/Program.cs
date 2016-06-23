using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

namespace ConComp
{
    class Program
    {
        static void Main(string[] args)
        {
            var fname1 = args[0];
            var fname2 = args[1];
            var task1 = Task<long>.Factory.StartNew(() => new FileInfo(fname1).Length);
            var task2 = Task<long>.Factory.StartNew(() => new FileInfo(fname2).Length);

            Task.WaitAll(task1, task2);

            if (task1.Result > task2.Result)
            {
                Console.WriteLine($"{fname1} is bigger");
            }
            else if (task2.Result > task1.Result)
            {
                Console.WriteLine($"{fname2} is bigger");
            }
            else
            {
                Console.WriteLine("The files are the same size");
            }
        }
    }
}