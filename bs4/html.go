package kcg

type Bs4Html struct {
	s string
}

func NewBs4Html() *Bs4Html {
	h := new(Bs4Html)
	return h
}

func (h *Bs4Html) New(title string) {
	h.s = `<!DOCTYPE html>` + "\n"
	h.s += `<html lang="en">` + "\n"
	h.s += "\t" + `<head>` + "\n"
	h.s += "\t" + "\t" + `<title>` + title + `</title>` + "\n"
	h.s += "\t" + "\t" + `<meta charset="utf-8">` + "\n"
	h.s += "\t" + "\t" + `<meta name="viewport" content="width=device-width, initial-scale=1">` + "\n"
	h.s += "\t" + "\t" + `<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css">` + "\n"
	h.s += "\t" + "\t" + `<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>` + "\n"
	h.s += "\t" + "\t" + `<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>` + "\n"
	h.s += "\t" + "\t" + `<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.0/js/bootstrap.min.js"></script>` + "\n"
	h.s += "\t" + `</head>` + "\n"
	h.s += "\t" + `<body>` + "\n"
}

func (h *Bs4Html) End() string {
	h.s += "\t" + `</body>` + "\n"
	h.s += `</html>` + "\n"
	return h.s
}

func (h *Bs4Html) Insert(s string) {
	h.s += s
}
