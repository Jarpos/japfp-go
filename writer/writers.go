package writer

var FUNCTIONS = map[string]Writer{
	"ct": {Writer: ChanneledTiling, Help: "Channeled Tiling"},
	"nt": {Writer: WriteTiling, Help: "Normal Tiling"},
}
