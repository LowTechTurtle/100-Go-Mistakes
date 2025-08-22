package main

type cache struct {
	m map[string]int
}

func (c *cache) get1(bytes []byte) (v int, contains bool) {
	key := string(bytes)
	v, contains = c.m[key]
	return
}

func (c *cache) get2(bytes []byte) (v int, contains bool) {
	// this seem insignificant, but when we write it like this
	// compiler will know and optimize, it will use the bytes directly
	// instead of converting to string and then pass the string and use
	// the underlying []bytes
	v, contains = c.m[string(bytes)]
	return
}
