package permission

type Permission map[Asset]Operator

func (p Permission) Has(asset Asset, operator Operator) bool {
	return p[asset]&operator == operator
}

func (p Permission) AddOperator(asset Asset, operators ...Operator) {
	if p[asset] == 0 {
		p[asset] = 0
	}
	for _, operator := range operators {
		p[asset] |= operator
	}
}

func (p Permission) RemoveOperator(asset Asset, operators ...Operator) {
	for _, operator := range operators {
		p[asset] &= ^operator
	}
	if p[asset] == 0 {
		delete(p, asset)
	}
}

func (p Permission) RemoveAsset(asset Asset) {
	delete(p, asset)
}

func (p Permission) Clear() {
	p = make(Permission)
}
