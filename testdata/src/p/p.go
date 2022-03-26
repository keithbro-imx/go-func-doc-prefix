package p

// Does a thing.
func NoPrefix() {} // want "should start with function name"

// Prefixdoes a thing.
func PrefixButNoSpace() {} // want "should start with function name"

// Prefix does a thing.
func Prefix() {}
