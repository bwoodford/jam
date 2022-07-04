package cache

import "testing"

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
			"SetTwoKeys",
			[]string{
				"Test1",
				"Test2",
			},
			2,
		},
		{
			"SetThreeKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
			},
			3,
		},
		{
			"SetFourKeys",
			[]string{
				"Test1",
				"Test2",
				"Test3",
				"Test4",
			},
			4,
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
