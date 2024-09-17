package usecase

import "time"

func login(track_id string) (string error) {
	if track.get(id=track_id).repeat_time > time.Now() {
		return nil, error
	}
	send_sms(phone)
	return "", nil
}

func
