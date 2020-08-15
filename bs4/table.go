package kcg

type Bs4Table struct {
	s string
}

func NewBs4Table() *Bs4Table {
	t := new(Bs4Table)
	return t
}

func (t *Bs4Table) Start(name string, id string) {
	t.s = "\t" + "\t" + `<div class="container-fluid" style="margin-top:80px" id="` + id + `">` + "\n"
	t.s += "\t" + "\t" + "\t" + `<h2>` + name + `</h2>` + "\n"
	t.s += "\t" + "\t" + "\t" + `<table class="table table-striped table-bordered">` + "\n"
}

func (t *Bs4Table) SetHeader(items []string) {
	t.s += "\t" + "\t" + "\t" + "\t" + `<thead class="thead-dark">` + "\n"
	t.s += "\t" + "\t" + "\t" + "\t" + "\t" + `<tr>` + "\n"
	for k := 0; k < len(items); k++ {
		t.s += "\t" + "\t" + "\t" + "\t" + "\t" + "\t" + `<th>` + items[k] + `</th>` + "\n"
	}
	t.s += "\t" + "\t" + "\t" + "\t" + "\t" + `</tr>` + "\n"
	t.s += "\t" + "\t" + "\t" + "\t" + `</thead>` + "\n"
}

func (t *Bs4Table) AddRowWithStyle(items []string, styles []string) {
	t.s += "\t" + "\t" + "\t" + "\t" + "\t" + `<tr>` + "\n"
	for k := 0; k < len(items); k++ {
		if styles[k] != "" {
			t.s += "\t" + "\t" + "\t" + "\t" + "\t" + "\t" + `<td ` + styles[k] + `>` + items[k] + `</td>` + "\n"
		} else {
			t.s += "\t" + "\t" + "\t" + "\t" + "\t" + "\t" + `<td>` + items[k] + `</td>` + "\n"
		}
	}
	t.s += "\t" + "\t" + "\t" + "\t" + "\t" + `</tr>` + "\n"
}

func (t *Bs4Table) AddRow(items []string) {
	t.s += "\t" + "\t" + "\t" + "\t" + "\t" + `<tr>` + "\n"
	for k := 0; k < len(items); k++ {
		t.s += "\t" + "\t" + "\t" + "\t" + "\t" + "\t" + `<td>` + items[k] + `</td>` + "\n"
	}
	t.s += "\t" + "\t" + "\t" + "\t" + "\t" + `</tr>` + "\n"
}
func (t *Bs4Table) NewBody() {
	t.s += "\t" + "\t" + "\t" + "\t" + `<tbody>` + "\n"
}

func (t *Bs4Table) EndBody() {
	t.s += "\t" + "\t" + "\t" + "\t" + `</tbody>` + "\n"
}

func (t *Bs4Table) End() string {
	t.s += "\t" + "\t" + "\t" + `</table>` + "\n"
	t.s += "\t" + "\t" + `</div>` + "\n"
	return t.s
}
