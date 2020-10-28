package main
import (
    "fmt"
)
func main() {
    var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt", "Emma", "Isabella", "Emily", "Madison", "Ava", "Olivia", "Sophia", "Abigail", "Elizabeth", "Chloe", "Samantha", "Addison", "Natalie", "Mia", "Alexis"}
    matrix := [][]string{}
    i, j, count := 1, 0, 0
    getHighestLength := 0 // for getting the longest string in slice
    var row []string
    for len(names) > 0 {
        if j == 0 {
            row = []string{}
        }
        if len(names[j]) == i {
            row = append(row, names[j])
            if len(names) == j {
                names = append(names[:j])
            }
            if len(names) < j {
                names = append(names[:j], names[j+1:]...)
            }
        }
        if len(names[j]) > getHighestLength && count==0 {
            getHighestLength = len(names[j])
            }
        j++
        if j == len(names) {
            matrix = append(matrix, row)
            j = 0
            count=1
            i++
        }
        if i > getHighestLength {
            break
        }
    }
    fmt.Println(matrix)
}
