package string_calculator_go

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func String_Calculator (numbers string) int {
    return 0
}


func Test_String_Calculator(t *testing.T) {
    type testCase struct {
        name            string
        numbers         string
        expectedOutput  int
    }

    cases := []*testCase {
        {
            name: "Should return 0 on empty string",
            numbers: "",
            expectedOutput: 0,
        },
        {
            name: "Should return correct sum of 2 numbers",
            numbers: "1,2",
            expectedOutput: 3,
        },
        {
            name: "Should return correct sum of 3 numbers",
            numbers: "1,2,3",
            expectedOutput: 6,
        },
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            result := String_Calculator(c.numbers)
            assert.Equal(t, c.expectedOutput, result)
        })
    }
}

