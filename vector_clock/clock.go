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
func (this *VectorClock) RecvMsg(that *VectorClock) error {
    if len(this.v) == len(that.v) {
        for i := range this.v {
            if i == this.id {
                this.v[i] += 1
            } else {
                this.v[i] = maxInt(this.v[i], that.v[i])
            }
        }
        return nil
    }
    return errors.New("cannot compare vectors of unequal length")
}

/*
Compares a vector clock to another to determine causal ordering.
 */
func (this *VectorClock) Compare(vec2 *VectorClock) (string, error) {
    if len(vec2.v) == len(this.v) {
        less := true
        greater := true
        equal := true
        for i, val := range this.v {
            if val > this.v[i] {
                less = false
            }
            if val < this.v[i] {
                greater = false
            }
            if val != this.v[i] {
                equal = false
            }
        }
        if less && greater {
            return "Concurrent", nil
        }
        if less {
            return "Less", nil
        }
        if greater {
            return "Greater", nil
        }
        if equal {
            return "Identical", nil
        }
    }
    return "", errors.New("cannot compare vectors of unequal length")
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
