using System;
using System.Diagnostics;
using System.Threading;

namespace benchm {
    class Program {
        private const int NUM_RUNS = 10000;
        
        private static void Main(string[] args) {
            //RunThreadsStart();
            RunNThreads(int.Parse(args[0]));
        }

        private static void RunThreadsStart() {
            Console.WriteLine("Benchmarking threads in C#");

            BenchmarkThreadStart(1);    // Warmup

            for (int i = 0; i < 5; i++) {
                var t = BenchmarkThreadStart(NUM_RUNS);
                long singleRun = (t.Ticks*100) / NUM_RUNS;
                Console.WriteLine($"Starting a single thread took {singleRun}ns");
            }
        }

        private static TimeSpan BenchmarkThreadStart(int n) {
            var watch = System.Diagnostics.Stopwatch.StartNew();
            for (int i = 0; i < n; i++) {
                Thread t = new Thread(() => {});
                t.Start();
                t.Join();
            }
            return watch.Elapsed;
        }

        private static void RunNThreads(int n) {
            ManualResetEvent waitEvent = new ManualResetEvent(false);
            int started = 0;

            Console.WriteLine($"Starting {n} threads");

            try {
                for (int i = 0; i < n; i++) {
                    Thread t = new Thread(() => {
                        Interlocked.Increment(ref started);
                        waitEvent.WaitOne();
                    });
                    t.Start();
                }
            }
            catch (OutOfMemoryException ex) {
                Console.WriteLine(ex.Message);
                Console.WriteLine($"After {started} threads started");
                Environment.Exit(1);
            }

            Console.WriteLine("All threads started");
            Console.WriteLine("Press ENTER to exit");

            Console.ReadLine();
            waitEvent.Set();
        }
    }
}