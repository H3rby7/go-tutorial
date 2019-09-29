package main

import (
	"testing"
)

func TestCompareMaps(t *testing.T) {
	{
		m1 := map[string]int{
			"steve": 12000,
			"Maria": 15000,
			"Mike":  9000,
		}
		m2 := map[string]int{
			"Maria": 15000,
			"Mike":  9000,
		}
		if compareMaps(m1, m2) {
			t.Error(m1, m2, "Should not be equal")
		}
	}
	{
		m1 := map[string]int{
			"steve": 12000,
			"Maria": 15000,
			"Mike":  9000,
		}
		m2 := map[string]int{
			"Maria": 15000,
			"steve": 12001,
			"Mike":  9000,
		}
		if compareMaps(m1, m2) {
			t.Error(m1, m2, "Should not be equal")
		}
	}
	{
		m1 := map[string]int{
			"steve": 12000,
			"Maria": 15000,
			"Mike":  9000,
		}
		var m2 map[string]int
		if compareMaps(m1, m2) {
			t.Error(m1, m2, "Should not be equal")
		}
	}
	{
		var m1 map[string]int
		m2 := map[string]int{
			"steve": 12000,
			"Maria": 15000,
			"Mike":  9000,
		}
		if compareMaps(m1, m2) {
			t.Error(m1, m2, "Should not be equal")
		}
	}
	{
		var m1 map[string]int
		m2 := map[string]int{}
		if compareMaps(m1, m2) {
			t.Error(m1, m2, "Should not be equal")
		}
	}
	{
		m1 := map[string]int{}
		var m2 map[string]int
		if compareMaps(m1, m2) {
			t.Error(m1, m2, "Should not be equal")
		}
	}
	{
		m1 := map[string]int{
			"Clara": 12000,
			"Maria": 15000,
			"Mike":  9000,
		}
		m2 := map[string]int{
			"Maria": 15000,
			"steve": 12001,
			"Mike":  9000,
		}
		if compareMaps(m1, m2) {
			t.Error(m1, m2, "Should not be equal")
		}
	}
	{
		m1 := map[string]int{
			"steve": 12000,
			"Maria": 15000,
			"Mike":  9000,
		}
		m2 := map[string]int{
			"Maria": 15000,
			"steve": 12000,
			"Mike":  9000,
		}
		if !compareMaps(m1, m2) {
			t.Error(m1, m2, "Should be equal")
		}
	}
	{
		var m1 map[string]int
		var m2 map[string]int
		if !compareMaps(m1, m2) {
			t.Error(m1, m2, "Should be equal")
		}
	}
	{
		m1 := map[string]int{}
		m2 := map[string]int{}
		if !compareMaps(m1, m2) {
			t.Error(m1, m2, "Should be equal")
		}
	}

}
