package file

type innerWindowsPlatform struct{}

func (inst *innerWindowsPlatform) Roots() []string {

	//	str := os.Environ()
	return []string{}
}

func (inst *innerWindowsPlatform) PathSeparatorChar() rune {
	return ';'
}

func (inst *innerWindowsPlatform) SeparatorChar() rune {
	return '\\'
}
