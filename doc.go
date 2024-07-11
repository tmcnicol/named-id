package id

// Wrapper around https://pkg.go.dev/github.com/oklog/ulid which adds
// prefix support to pretty print ids.
//
// The String output contains the <prefix>_<base-32 cockford encoded> For
// aesthetic reason the ULID is converted to lowercase.
//
// Example string output 'mat_01hk6nfgxxpyy10cscqw81p3pa'
