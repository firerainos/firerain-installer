package config

type Package struct {
	PkgList []string
}

func (p *Package) AddPackage(pkg string) {
	p.PkgList = append(p.PkgList, pkg)
}

func (p *Package) RemovePackage(pkg string) {
	var list []string
	for _, name := range p.PkgList {
		if name != pkg {
			list = append(list, name)
		}
	}
	p.PkgList = list
}

func (p *Package) SearchPackage(pkg string) bool {
	for _, name := range p.PkgList {
		if name == pkg {
			return true
		}
	}
	return false
}
