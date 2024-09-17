package repository

type DB string

type TrackUserRepository struct {
	db *DB
}

func NewTrackUserRepository(db *DB) *TrackUserRepository {
	return &TrackUserRepository{db: db}
}
func (t *TrackUserRepository) Create()          {}
func (t *TrackUserRepository) Update(id string) {}

type UserRepository struct {
	db *DB
}

func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create()                       {}
func (u *UserRepository) Get(id string)                 {}
func (u *UserRepository) Update(id string)              {}
func (u *UserRepository) GetByUsername(username string) {}

type SessionRepository struct {
	db *DB
}

func NewSessionRepository(db *DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (s *SessionRepository) Create()          {}
func (s *SessionRepository) Get(id string)    {}
func (s *SessionRepository) Delete(id string) {}
