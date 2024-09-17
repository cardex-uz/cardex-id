package model

type User struct {
	id         string
	username   string
	phone      string
	first_name string
	last_name  string
	avatar     string
}

type Track struct {
	id          string
	phone       string
	is_verified string
	verify_time string
	verify_code string
}
