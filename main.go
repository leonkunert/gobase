package main

func main() {
    // Start Profiler
    startProfiler()
    // Print Welcome Message
    PrintWelcome()
    // Initialize the database connection
    InitDBConnection()
    // Defer Closing of SQL Connection
    defer DBC.Close()
    // Starting router
    Router()
}