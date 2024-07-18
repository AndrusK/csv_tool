package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "os"
)

var (
    mode        string
    firstCSVPath string
    secondCSVPath string
    outputCSVPath string
    verbose     bool
    showHelp    bool
)

func init() {
    flag.StringVar(&mode, "m", "", "Mode to use for processing a file.\nOptions:\nsimilarities - find similarities between two CSV files.\ndifferences - find differences between two CSV files. This will show data that exists in ONLY the second CSV file (and not the first). Use it to find oddities in data.")
    flag.StringVar(&firstCSVPath, "f", "", "Your first CSV file.")
    flag.StringVar(&secondCSVPath, "s", "", "Your second CSV file.")
    flag.StringVar(&outputCSVPath, "o", "", "Your output file, including the .csv extension")
    flag.BoolVar(&verbose, "v", false, "Enable verbose output")
    flag.BoolVar(&showHelp, "h", false, "Show this help message")
    flag.BoolVar(&showHelp, "help", false, "Show this help message")
    
    flag.Parse()
}

func main() {
    if showHelp {
        flag.Usage()
        os.Exit(1)
    }

    // Check if required flags are provided
    if mode == "" || firstCSVPath == "" || secondCSVPath == "" {
        fmt.Println("Error: -m, -f, and -s flags are required")
        flag.Usage()
        os.Exit(1)
    }

    if mode != "similarities" && mode != "differences" {
        fmt.Println("Error: Invalid mode specified. Available options: similarities, differences")
        flag.Usage()
        os.Exit(1)
    }

    // Open first CSV file
    firstFile, err := os.Open(firstCSVPath)
    if err != nil {
        fmt.Println("Error opening first CSV file:", err)
        return
    }
    defer firstFile.Close()

    // Open second CSV file
    secondFile, err := os.Open(secondCSVPath)
    if err != nil {
        fmt.Println("Error opening second CSV file:", err)
        return
    }
    defer secondFile.Close()

    // Create CSV readers
    firstReader := csv.NewReader(firstFile)
    secondReader := csv.NewReader(secondFile)

    // Read all records from first CSV
    firstCSVLoaded, err := firstReader.ReadAll()
    if err != nil {
        fmt.Println("Error reading first CSV:", err)
        return
    }

    // Read all records from second CSV
    secondCSVLoaded, err := secondReader.ReadAll()
    if err != nil {
        fmt.Println("Error reading second CSV:", err)
        return
    }

    var output [][]string

    // Process based on mode
    switch mode {
    case "similarities":
        output = similarities(firstCSVLoaded, secondCSVLoaded)
    case "differences":
        output = differences(firstCSVLoaded, secondCSVLoaded)
    }

    // Print verbose output if enabled
    if verbose {
        fmt.Println("Output:")
        for _, line := range output {
            fmt.Println(line)
        }
    }

    // Write output to file if specified
    if outputCSVPath != "" {
        outputFile, err := os.Create(outputCSVPath)
        if err != nil {
            fmt.Println("Error creating output file:", err)
            return
        }
        defer outputFile.Close()

        writer := csv.NewWriter(outputFile)
        defer writer.Flush()

        for _, record := range output {
            if err := writer.Write(record); err != nil {
                fmt.Println("Error writing record to file:", err)
                return
            }
        }

        fmt.Println("Output written to", outputCSVPath)
    } else {
        fmt.Println("No output file specified.")
    }
}

// Function to find similarities between two CSVs
func similarities(firstCSV, secondCSV [][]string) [][]string {
    var output [][]string

    // Iterate through each record in firstCSV
    for _, row1 := range firstCSV {
        // Check if the record exists in secondCSV
        if contains(secondCSV, row1) {
            output = append(output, row1)
        }
    }

    return output
}

// Function to find differences between two CSVs
func differences(firstCSV, secondCSV [][]string) [][]string {
    var output [][]string

    // Iterate through each record in firstCSV
    for _, row1 := range firstCSV {
        // Check if the record exists in secondCSV
        if !contains(secondCSV, row1) {
            output = append(output, row1)
        }
    }

    return output
}

// Function to check if a record exists in a CSV
func contains(csv [][]string, record []string) bool {
    for _, row := range csv {
        if equalSlices(row, record) {
            return true
        }
    }
    return false
}

// Function to compare equality of two slices of strings
func equalSlices(slice1, slice2 []string) bool {
    if len(slice1) != len(slice2) {
        return false
    }
    for i := range slice1 {
        if slice1[i] != slice2[i] {
            return false
        }
    }
    return true
}

