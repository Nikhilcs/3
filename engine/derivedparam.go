package engine

// parameter derived from others (not directly settable)
type derivedParam struct {
	lut
	updater  func(*derivedParam)
	uptodate bool // cleared if parents' value change
	parents  []updater
}

func (p *derivedParam) init(nComp int, parents []updater, updater func(*derivedParam)) {
	p.lut.init(nComp, p)
	p.updater = updater
	p.parents = parents
}

func (p *derivedParam) invalidate() {
	p.uptodate = false
}

func (p *derivedParam) update() {
	for _, par := range p.parents {
		par.update() // may invalidate me they-re time dependent
	}
	if !p.uptodate {
		p.updater(p)
		p.gpu_ok = false
		p.uptodate = true
	}
}
