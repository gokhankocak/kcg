package kcg

type Bs4Nav struct {
	s string
}

type Bs4NameHref struct {
	Name string
	Href string
}

func NewBs4Nav() *Bs4Nav {
	nav := new(Bs4Nav)
	return nav
}

func (nav *Bs4Nav) Start() {
	nav.s = "\t" + "\t" + `<nav class="navbar navbar-expand-sm bg-dark navbar-dark fixed-top">` + "\n"
}

func (nav *Bs4Nav) SetLogo(logo string) {
	nav.s += "\t" + "\t" + "\t" + `<a class="navbar-brand" href="#">` + logo + `</a>` + "\n"
}

func (nav *Bs4Nav) AddDropDownMenu(menu Bs4NameHref, items []Bs4NameHref) {
	nav.s += "\t" + "\t" + "\t" + "\t" + `<li class="nav-item dropdown">` + "\n"
	nav.s += "\t" + "\t" + "\t" + "\t" + "\t" + `<a class="nav-link dropdown-toggle" data-toggle="dropdown" href="#` + menu.Href + `">` + menu.Name + `</a>` + "\n"
	nav.s += "\t" + "\t" + "\t" + "\t" + "\t" + `<div class="dropdown-menu">` + "\n"

	for k := 0; k < len(items); k++ {
		nav.s += "\t" + "\t" + "\t" + "\t" + "\t" + "\t" + `<a class="dropdown-item" href="#` + items[k].Href + `">` + items[k].Name + `</a>` + "\n"
	}

	nav.s += "\t" + "\t" + "\t" + "\t" + "\t" + `</div>` + "\n"
	nav.s += "\t" + "\t" + "\t" + "\t" + `</li>` + "\n"
}

func (nav *Bs4Nav) AddItem(item Bs4NameHref) {
	nav.s += "\t" + "\t" + "\t" + "\t" + `<li class="nav-item">` + "\n"
	nav.s += "\t" + "\t" + "\t" + "\t" + "\t" + `<a class="nav-link" href="#` + item.Href + `">` + item.Name + `</a>` + "\n"
	nav.s += "\t" + "\t" + "\t" + "\t" + `</li>` + "\n"
}

func (nav *Bs4Nav) End() string {
	nav.s += "\t" + "\t" + "\t" + `</ul>` + "\n"
	nav.s += "\t" + "\t" + `</nav>` + "\n"
	return nav.s
}
