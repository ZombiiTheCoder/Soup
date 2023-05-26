package runtime

type null *struct{}
var nul null

type RuntimeVal interface{
	GetType() string
}