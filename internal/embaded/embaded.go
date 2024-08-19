package embaded17


import _"embed"

//go:embed templates

var embaded []byte

func GetTemplates()[]byte{
	return embaded
}