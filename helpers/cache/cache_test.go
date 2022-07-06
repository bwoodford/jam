package cache

import (
	"bytes"
	"testing"
)

var global Cache

func Before() {
	global = NewCache()
}

func TestRemovePass(t *testing.T) {

	cases := []struct {
		desc     string
		input    []string
		remove   []string
		expected int
	}{
		{
			"RemoveOneKey",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
			},
			[]string{
				"Test1",
			},
			3,
		},
		{
			"RemoveTwoKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
			},
			[]string{
				"Test1",
				"Test3",
			},
			2,
		},
		{
			"RemoveThreeKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
			},
			[]string{
				"Test1",
				"Test3",
				"Test4",
			},
			1,
		},
		{
			"RemoveAllKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
			},
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
			},
			0,
		},
	}

	for _, tc := range cases {
		Before()
		for _, in := range tc.input {
			global.cMap[in] = []byte{}
		}
		for _, re := range tc.remove {
			global.Remove(re)
		}

		if len(global.cMap) != tc.expected {
			t.Fatalf("%s: expected: %d got: %d", tc.desc, tc.expected, len(global.cMap))
		}
	}
}

func TestRemoveFail(t *testing.T) {

	cases := []struct {
		desc     string
		load     []string
		input    string
		expected error
	}{
		{
			"RemoveEmptyMap",
			[]string{},
			"MissingKey",
			ErrKeyNotFound,
		},
		{
			"RemoveFilledMap",
			[]string{
				"TestKey1",
				"TestKey2",
				"TestKey3",
				"TestKey4",
				"TestKey5",
			},
			"MissingKey",
			ErrKeyNotFound,
		},
	}

	for _, tc := range cases {
		Before()
		for _, in := range tc.load {
			global.cMap[in] = []byte{}
		}

		err := global.Remove(tc.input)

		if err != tc.expected {
			t.Fatalf("%s: expected: %d got: %d", tc.desc, tc.expected, err)
		}
	}
}

func TestSetPass(t *testing.T) {

	cases := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			"SetOneKey",
			[]string{
				"Test1",
			},
			1,
		},
		{
			"SetFiveKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
			},
			5,
		},
		{
			"SetTenKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
				"Test6",
				"Test7",
				"Test8",
				"Test9",
				"Test10",
			},
			MAX_SIZE,
		},
		{
			"SetFifteenKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
				"Test6",
				"Test7",
				"Test8",
				"Test9",
				"Test10",
				"Test11",
				"Test12",
				"Test13",
				"Test14",
				"Test15",
			},
			MAX_SIZE,
		},
		{
			"SetTwentyKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
				"Test6",
				"Test7",
				"Test8",
				"Test9",
				"Test10",
				"Test11",
				"Test12",
				"Test13",
				"Test14",
				"Test15",
				"Test16",
				"Test17",
				"Test18",
				"Test19",
				"Test20",
			},
			MAX_SIZE,
		},
	}

	for _, tc := range cases {
		Before()
		for _, in := range tc.input {
			global.Set(in, []byte{})
		}
		if len(global.cMap) != tc.expected {
			t.Fatalf("%s: expected: %d got: %d", tc.desc, tc.expected, len(global.cMap))
		}
	}
}

func TestSetEviction(t *testing.T) {

	cases := []struct {
		desc     string
		input    []string
		expected []string
	}{
		{
			"SetOneKey",
			[]string{
				"Test1",
			},
			[]string{
				"Test1",
			},
		},
		{
			"SetFiveKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
			},
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
			},
		},
		{
			"SetTenKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
				"Test6",
				"Test7",
				"Test8",
				"Test9",
				"Test10",
			},
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
				"Test6",
				"Test7",
				"Test8",
				"Test9",
				"Test10",
			},
		},
		{
			"SetFifteenKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
				"Test6",
				"Test7",
				"Test8",
				"Test9",
				"Test10",
				"Test11",
				"Test12",
				"Test13",
				"Test14",
				"Test15",
			},
			[]string{
				"Test6",
				"Test7",
				"Test8",
				"Test9",
				"Test10",
				"Test11",
				"Test12",
				"Test13",
				"Test14",
				"Test15",
			},
		},
		{
			"SetTwentyKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
				"Test5",
				"Test6",
				"Test7",
				"Test8",
				"Test9",
				"Test10",
				"Test11",
				"Test12",
				"Test13",
				"Test14",
				"Test15",
				"Test16",
				"Test17",
				"Test18",
				"Test19",
				"Test20",
			},
			[]string{
				"Test11",
				"Test12",
				"Test13",
				"Test14",
				"Test15",
				"Test16",
				"Test17",
				"Test18",
				"Test19",
				"Test20",
			},
		},
	}

	for _, tc := range cases {
		Before()
		for _, in := range tc.input {
			global.Set(in, []byte{})
		}

		for _, ex := range tc.expected {
			if _, ok := global.cMap[ex]; !ok {
				t.Fatalf("%s: expected: %s got: ErrKeyNotFound", tc.desc, ex)
			}
		}
	}
}

func TestGetPass(t *testing.T) {

	cases := []struct {
		desc  string
		input map[string][]byte
	}{
		{
			"SetOneKey",
			map[string][]byte{
				"Test1": []byte{'w'},
			},
		},
		{
			"SetTenKeys",
			map[string][]byte{
				"Test1":  []byte{'w'},
				"Test2":  []byte{'e'},
				"Test3":  []byte{'e'},
				"Test4":  []byte{'w'},
				"Test5":  []byte{'o'},
				"Test6":  []byte{'o'},
				"Test7":  []byte{'w'},
				"Test8":  []byte{'a'},
				"Test9":  []byte{'r'},
				"Test10": []byte{'e'},
			},
		},
		{
			"SetTenMoreKeys",
			map[string][]byte{
				"Test6":  []byte{'o'},
				"Test7":  []byte{'w'},
				"Test8":  []byte{'a'},
				"Test9":  []byte{'r'},
				"Test10": []byte{'e'},
				"Test11": []byte{'w'},
				"Test12": []byte{'o'},
				"Test13": []byte{'l'},
				"Test14": []byte{'f'},
				"Test15": []byte{'e'},
			},
		},
	}

	for _, tc := range cases {
		Before()
		for key, val := range tc.input {
			global.cMap[key] = val
			global.queue.PushBack(key)
		}

		for key, inVal := range tc.input {
			inMap, outVal := global.Get(key)
			if !inMap || bytes.Compare(inVal, outVal) != 0 {
				t.Fatalf("%s: expected inMap: %v got: %v: expected val: %v got: %v", tc.desc, true, inMap, inVal, outVal)
			}
		}
	}
}
