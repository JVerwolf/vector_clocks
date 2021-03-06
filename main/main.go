package main

import (
    "../vector_clock"
    "strconv"
)

/*
The code in this main function is to test the causal relationships shown in
the Diagram.png, attached in this repo.
 */
func main() {
    // First, we’ll initialize vector clocks each of the 3 processes.
    vec := initVectors(3)
    printState(vec)

    // Save the original state of the vectors for a later comparison.
    initialVecs := copyVectors(vec)

    // Show that all vectors are causally concurrent.
    printCompare(vec[0], vec[1])
    printCompare(vec[1], vec[2])
    printCompare(vec[2], vec[0])

    // Send msgs from 2 to 1, and 0 to 1.
    printSendMessage(vec[2], vec[1])
    printSendMessage(vec[0], vec[1])

    // Show that 0 is now causally connected to 1.
    printCompare(vec[0], vec[1])

    // Increment vectors in accordance with work being done.
    printDoWork(vec[0])
    printDoWork(vec[2])

    // Show how 0 and 1 are now causally concurrent.
    printCompare(vec[0], vec[1])

    // Send a message from 0 to 2.
    printSendMessage(vec[0], vec[2])

    // Finally, let’s check the causality in the blue path
    // in Diagram.png. As shown in the diagram, position N
    // should depend on position C, but not A.
    printCompare(vec[0], initialVecs[0])
    printCompare(vec[0], initialVecs[2])
}

/*
Deep copy an array of vector clocks.
 */
func copyVectors(vec []*vector_clock.VectorClock) (newVec []*vector_clock.VectorClock) {
    newVec = make([]*vector_clock.VectorClock, len(vec))
    for i, v := range vec {
        newVec[i] = v.Copy()
    }
    return
}

func printCompare(v1, v2 *vector_clock.VectorClock) {
    output := "Compare Vector " + strconv.Itoa(v1.Id) + " with " + strconv.Itoa(v2.Id) + ":\n"
    output += v1.ToString() + "\n"
    output += v1.Compare(v2) + "\n"
    output += v2.ToString() + "\n"
    println(output)
}

/*
Print a formatted description of system state.
 */
func printState(vec []*vector_clock.VectorClock) {
    println("State of all Vector Clocks:")
    for i := range vec {
        println(vec[i].ToString())
    }
    println()
}

/*
Print a formatted description of incrementing a vector.
 */
func printDoWork(v *vector_clock.VectorClock) {
    output := "Vector Clock " + strconv.Itoa(v.Id) + " Does Work:\n"
    output += "Before: " + v.ToString() + "\n"
    v.Increment()
    output += "After: " + v.ToString() + "\n"
    println(output)
}

/*
Initialize a slice of vector clocks.
 */
func initVectors(numVecs int) (vec []*vector_clock.VectorClock) {
    for i := 0; i < numVecs; i++ {
        vec = append(vec, vector_clock.NewVectorClock(numVecs, i))
    }
    return
}

/*
Print a formatted description of message passing.
 */
func printSendMessage(v1, v2 *vector_clock.VectorClock) {
    output := "Sending Message from " + strconv.Itoa(v1.Id) + " to " + strconv.Itoa(v2.Id) + ":\n"
    output += "Before:\t" + v1.ToString() + ", " + v2.ToString() + "\n"
    v1.SendMsg(v2)
    output += "After:\t" + v1.ToString() + ", " + v2.ToString() + "\n"
    println(output)
}
