package main

import (
	poke_cache "github.com/muszkin/pokedex/poke-cache"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input              string
		expected           []string
		expectedReturnSize int
	}{
		{
			input:              "  hello  world  ",
			expected:           []string{"hello", "world"},
			expectedReturnSize: 2,
		}, {
			input:              "TY to Te ",
			expected:           []string{"ty", "to", "te"},
			expectedReturnSize: 3,
		}, {
			input:              "PiKaChU iS SUPER",
			expected:           []string{"pikachu", "is", "super"},
			expectedReturnSize: 3,
		}, {
			input:              "8799yv97chsh9348904r20u24",
			expected:           []string{"8799yv97chsh9348904r20u24"},
			expectedReturnSize: 1,
		}, {
			input:              "Y^(*&(*& U(",
			expected:           []string{"y^(*&(*&", "u("},
			expectedReturnSize: 2,
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != c.expectedReturnSize {
			t.Errorf("return slice different from expected input: %v, output: %v", c.expectedReturnSize, len(actual))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("return value different from expected input: %v, output %v", expectedWord, word)
				return
			}
		}
	}
}

func TestCache(t *testing.T) {
	const testString = "test string"
	const testKey = "test"
	const duration4sString = "4s"
	const duration5sString = "5s"
	const duration6sString = "6s"
	duration, _ := time.ParseDuration(duration5sString)
	longDuration, _ := time.ParseDuration(duration6sString)
	shorDuration, _ := time.ParseDuration(duration4sString)
	cache := poke_cache.NewCache(duration)

	cache.Add(testKey, []byte(testString))
	time.Sleep(longDuration)
	_, received := cache.Get(testKey)
	if received {
		t.Errorf("cache should delete key after passing duration")
		return
	}

	cache.Add(testKey, []byte(testString))
	time.Sleep(shorDuration)
	val, received := cache.Get(testKey)
	if !received {
		t.Errorf("cache should containt key when duration is not passed")
		return
	}
	if string(val) != testString {
		t.Errorf("value in cache should be the same like value before cache")
		return
	}
}
