package repository

type TrackUser interface {
	Create()
	Update()
}

type User interface {
	Create()
	Get()
	Update()
	GetByUsername()
}

type SessionRepository interface {
	Create()
	Get()
	Delete()
}
