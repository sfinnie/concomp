using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

namespace ConComp
{
    class Program2
    {
        static void Main(string[] args)
        {
            if (args.Length == 0)
            {
                Console.WriteLine("Please specify one or more filenames.");
            }
            else
            {
                var missingFileName = args.FirstOrDefault(f => !File.Exists(f));
                if (missingFileName != null)
                {
                    Console.WriteLine("There is no file named {0}.", missingFileName);
                }
                else
                {
                    var tasks = args.Select(f => Task.Run(() => new FileInfo(f).Length)).ToArray();
                    Task.WaitAll(tasks);

                    var results = tasks.Select((t, i) => new { Filename = args[i], Length = t.Result });
                    var lengthOfLargest = results.Max(r => r.Length);
                    var namesOfLargestFiles = results.Where(r => r.Length == lengthOfLargest).Select(r => r.Filename).ToList();

                    if (namesOfLargestFiles.Count == 1)
                    {
                        Console.WriteLine("The largest file (at {0} bytes) is {1}", lengthOfLargest, namesOfLargestFiles[0]);
                    }
                    else
                    {
                        Console.WriteLine("The largest files (at {0} bytes) are {1}", lengthOfLargest, string.Join(", ", namesOfLargestFiles));
                    }
                }
            }

            Console.ReadLine();
        }
    }
}