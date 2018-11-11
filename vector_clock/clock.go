package vector_clock

import (
    "errors"
    "strconv"
    "strings"
    "fmt"
)

type VectorClock struct {
    V  []int
    Id int
}

func NewVectorClock(size, id int) (*VectorClock) {
    vec := new(VectorClock)
    vec.V = make([]int, size)
    vec.Id = id
    vec.V[id] = 1
    return vec
}

/*
Increments a vector clocks internal value to account for work being done.
 */
func (this *VectorClock) Increment() (*VectorClock) {
    this.V[this.Id] += 1
    return this
}

/*
Performs the updates required when a vector clock receives a message
from another.
 */
func (this *VectorClock) SendMsg(that *VectorClock) error {
    if len(this.V) == len(that.V) {
        this.V[this.Id] += 1
        that.V[that.Id] += 1
        for i := range that.V {
            if i != that.Id {
                that.V[i] = maxInt(this.V[i], that.V[i])
            }
        }
        return nil
    }
    return errors.New("cannot compare vectors of unequal length")
}

/*
Compares a vector clock to another to determine causal ordering.
 */
func (this *VectorClock) Compare(that *VectorClock) (string) {
    if len(this.V) == len(that.V) {
        less := true
        greater := true
        equal := true
        for i, val := range this.V {
            if val > that.V[i] {
                less = false
            }
            if val < that.V[i] {
                greater = false
            }
            if val != that.V[i] {
                equal = false
            }
        }
        if equal {
            return "Is Identical With"
        }
        if less {
            return "Happened Before"
        }
        if greater {
            return "Happened After"
        }
        return "Is Concurrent With"

    }
    return "Error: Cannot compare vectors of unequal length"
}

/*
Pretty prints the vector.
 */
func (this *VectorClock) ToString() string {
    output := strconv.Itoa(this.Id) + ":["
    output += strings.Trim(strings.Join(strings.Fields(fmt.Sprint(this.V)), ","), "[]")
    output += "]"
    return output
}

/*
Helper function to find max of ints
 */
func maxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}
