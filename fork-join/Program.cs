using System;
using System.Threading;

namespace fork_join {
    class Program {
        static void Main(string[] args) {
            Thread t = new Thread(DoWork);
            t.Start();           // Fork
            Thread.Sleep(500);
            Console.WriteLine("Job done");
            t.Join();            // Join
            Console.WriteLine("Everything done");
        }

        static void DoWork() {
            Thread.Sleep(1000);
            Console.WriteLine("Other job done");
        }
    }
}