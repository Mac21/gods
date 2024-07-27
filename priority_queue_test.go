package gods

import (
    "testing"
)

func TestIntMinPQ(t *testing.T) {
    input := map[int]int{
        0: 5,
        1: 2,
        2: 3,
        3: 4,
        4: 1,
    }
    pq := NewPriorityQueue[int]()
    for v, p := range input {
        pq.Push(v, p)
    }

    i := 0
    answer := []int{ 4, 1, 2, 3, 0, }
    for pq.Len() > 0 {
        res, err := pq.Pop()
        expectedAnswer := answer[i]
        i++
        if err != nil {
            t.Fatal(err)
        }

        if res != expectedAnswer {
            t.Fatalf("pq.Pop() result doesn't match expected %v got %v", res, expectedAnswer)
        }

        expectedLen := len(answer[i:])
        if pq.Len() != expectedLen {
            t.Fatalf("priority queue Len() doesn't match expected %v got %v", expectedLen, pq.Len())
        }

    }
}

func TestStructMinPQ(t *testing.T) {
    type CustomValue struct {
        Number int
        Weight int
    }
    input := []CustomValue{
        {
            Number: 0,
            Weight: 7,
        },

        {
            Number: 1,
            Weight: 6,
        },
        {
            Number: 2,
            Weight: 5,
        },
        {
            Number: 3,
            Weight: 1,
        },
        {
            Number: 4,
            Weight: 2,
        },
    }
    pq := NewPriorityQueue[CustomValue]()
    for _, cv := range input {
        pq.Push(cv, cv.Weight)
    }

    i := 0
    answer := []CustomValue{
        {
            Number: 3,
            Weight: 1,
        },

        {
            Number: 4,
            Weight: 2,
        },
        {
            Number: 2,
            Weight: 5,
        },
        {
            Number: 1,
            Weight: 6,
        },
        {
            Number: 0,
            Weight: 7,
        },

    }
    for pq.Len() > 0 {
        res, err := pq.Pop()
        expectedAnswer := answer[i]
        i++
        if err != nil {
            t.Fatal(err)
        }

        if res != expectedAnswer {
            t.Fatalf("pq.Pop() result doesn't match expected %v got %v", res, expectedAnswer)
        }

        expectedLen := len(answer[i:])
        if pq.Len() != expectedLen {
            t.Fatalf("priority queue Len() doesn't match expected %v got %v", expectedLen, pq.Len())
        }

    }
}
