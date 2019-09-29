package computer

//exported struct
type Spec struct {
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}
