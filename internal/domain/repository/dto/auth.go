package dto

type CreateTrack struct {
	phone string
}

type Verify struct {
	id          string
	verify_code string
}

type CreateUser struct {
	phone string
}
