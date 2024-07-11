func main() {
    filename := "example2.txt" // Replace with the path to your input file

    farm, err := functions.ReadFile(filename)
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    paths, err := functions.FindAllPaths(farm)
    if err != nil {
        log.Fatalf("Error finding paths: %v", err)
    }

    // The rest of the code is not needed for testing FindAllPaths
}
