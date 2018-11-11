package vector_clock

import "errors"

type Causality int

const (
    EQUAL      Causality = iota
    LESS
    GREATER
    CONCURRENT
)

type VectorClock struct {
    v []int
}

func NewVectorClock(size int) (*VectorClock) {
    vec := new(VectorClock)
    vec.v = make([]int, size)
    return vec
}

func (vec *VectorClock) Set(i, val int) (vec2 *VectorClock) {
    vec.v[i] = val
    return vec
}

func (vec1 *VectorClock) Compare(vec2 *VectorClock) (Causality, error) {
    if len(vec2.v) == len(vec1.v) {
        less := true
        greater := true
        equal := true
        for i, val := range vec1.v {
            if val > vec1.v[i] {
                less = false
            }
            if val < vec1.v[i] {
                greater = false
            }
            if val != vec1.v[i] {
                equal = false
            }
        }
        if less && greater {
            return CONCURRENT, nil
        }
        if less {
            return LESS, nil
        }
        if greater {
            return GREATER, nil
        }
        if equal {
            return EQUAL, nil
        }

    }
    return 0, errors.New("cannot compare vectors of unequal length")
}
