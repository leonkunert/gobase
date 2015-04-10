package main

import (
    "runtime/pprof"
    "flag"
    "os"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var heapprofile = flag.String("heapprofile", "", "write heap profile to file")

func startProfiler() {
    flag.Parse()
    if *cpuprofile != "" {
        // Start Profiler
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log("Start CPU Profiler", err.Error())
        }
        pprof.StartCPUProfile(f)
    }
    if *heapprofile != "" {
        // Start Profiler
        f, err := os.Create(*heapprofile)
        if err != nil {
            log("Start Heap Profiler", err.Error())
        }
        pprof.Lookup("heap").WriteTo(f, 0)
    }
}