package pods // All the files inside lib will have the same package name

import "fmt"

// This func must be Exported, Capitalized, and comment added.
func GetPod() {
    fmt.Println("Fetching Pods...")
}