package query

var UserQuery = struct {
	SelectListUser string
	SelectAll      string
}{
	SelectListUser: `
		select nama, alamat, pendidikan, count(*) OVER() as count from users
`,
	SelectAll: `
		select *, count(*) OVER() as count from users
	`,
}
