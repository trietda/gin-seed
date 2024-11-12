package repository

import "gin-seed/app/user/model"

var sessions map[string]model.Session = make(map[string]model.Session)

func SaveSession(session model.Session) {
	sessions[session.Id] = session
}
