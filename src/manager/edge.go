package manager

import "utils"

type Edge struct {
	Addr string
	Role string
	Desc string
	Loc  *utils.Loc
}

func NewEdge(ip, role, desc string) (e *Edge) {
	e = new(Edge)
	e.Addr = ip
	e.Role = role
	e.Desc = desc
	e.Loc = utils.NewLoc(ip)

	return
}