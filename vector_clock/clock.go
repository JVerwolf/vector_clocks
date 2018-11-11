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

func (this *VectorClock) Set(i, val int) (*VectorClock) {
    this.v[i] = val
    return this
}

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
