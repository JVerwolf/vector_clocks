package vector_clock

import (
    "errors"
)

type VectorClock struct {
    v  []int
    id int
}

func NewVectorClock(size, id int) (*VectorClock) {
    vec := new(VectorClock)
    vec.v = make([]int, size)
    vec.id = id
    vec.v[id] = 1
    return vec
}

/*
Increments a vector clocks internal value to account for work being done.
 */
func (this *VectorClock) Increment() (*VectorClock) {
    this.v[this.id] += 1
    return this
}

/*
Performs the updates required when a vector clock receives a message
from another.
 */
func (this *VectorClock) SendMsg(that *VectorClock) error {
    if len(this.v) == len(that.v) {
        this.v[this.id] += 1
        that.v[that.id] += 1
        for i := range that.v {
            if i != that.id {
                that.v[i] = maxInt(this.v[i], that.v[i])
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
    if len(this.v) == len(that.v) {
        less := true
        greater := true
        equal := true
        for i, val := range this.v {
            if val > that.v[i] {
                less = false
            }
            if val < that.v[i] {
                greater = false
            }
            if val != that.v[i] {
                equal = false
            }
        }
        if equal {
            return "Identical"
        }
        if less {
            return "Before"
        }
        if greater {
            return "After"
        }
        return "Concurrent"

    }
    return "Error: Cannot compare vectors of unequal length"
}

/*
Helper function to find max of ints
 */
func maxInt(a, b int) int {
    if (a > b) {
        return a
    }
    return b
}
