package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HUEBRTeam/PrimeServer/ProfileManager"
	"github.com/HUEBRTeam/PrimeServer/tools"
	"github.com/gorilla/mux"
	"github.com/quan-to/slog"
)

var log = slog.Scope("REST")

type Server struct {
	r    *mux.Router
	port int
	pm   *ProfileManager.ProfileManager
}

func MakeRestServer(port int, pm *ProfileManager.ProfileManager) *Server {
	r := mux.NewRouter()

	s := &Server{
		r:    r,
		port: port,
		pm:   pm,
	}

	r.HandleFunc("/profile", s.getProfile).Methods("GET")
	r.HandleFunc("/createProfile", s.createProfile).Methods("POST")
	r.HandleFunc("/getProfileFile", s.getProfileFile).Methods("GET")

	r.HandleFunc("/change", s.change).Methods("POST")

	return s
}

func (s *Server) Listen() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.r)
}

func (s *Server) getProfileFile(w http.ResponseWriter, r *http.Request) {
	tools.InitHTTPTimer(r)

	defer func() {
		if rec := recover(); rec != nil {
			tools.CatchAllError(rec, w, r, log)
		}
	}()

	q := r.URL.Query()

	ac := q.Get("AccessCode")

	if ac == "" {
		tools.InvalidFieldData("AccessCode", "You should provide a \"AccessCode\" query parameter.", w, r, log)
		return
	}

	data, err := hex.DecodeString(ac)

	if err != nil {
		tools.InvalidFieldData("AccessCode", fmt.Sprintf("Invalid AccessCode: %s", err.Error()), w, r, log)
		return
	}

	w.Header().Set("Content-Type", tools.MimeBinary)
	w.Header().Set("Content-Disposition", "attachment; filename=\"prime.bin\"")
	w.WriteHeader(200)
	n, _ := w.Write(data)
	tools.LogExit(log, r, 200, n)
}

func (s *Server) createProfile(w http.ResponseWriter, r *http.Request) {
	tools.InitHTTPTimer(r)

	defer func() {
		if rec := recover(); rec != nil {
			tools.CatchAllError(rec, w, r, log)
		}
	}()

	var cp CreateProfileRequest

	if !tools.UnmarshalBodyOrDie(&cp, w, r, log) {
		return
	}

	accessCode, err := s.pm.Create(cp.Name, cp.CountryID, cp.Avatar, cp.Modifiers)

	if err != nil {
		tools.InternalServerError("There was an error processing your request", map[string]interface{}{
			"error":   err,
			"message": err.Error(),
		}, w, r, log)
		return
	}

	var pr CreateProfileResponse
	pr.Name = cp.Name
	pr.AccessCode = accessCode

	b, _ := json.Marshal(pr)
	w.Header().Set("Content-Type", tools.MimeJSON)
	w.WriteHeader(200)
	n, _ := w.Write(b)
	tools.LogExit(log, r, 200, n)
}

func (s *Server) change(w http.ResponseWriter, r *http.Request) {
	tools.InitHTTPTimer(r)

	defer func() {
		if rec := recover(); rec != nil {
			tools.CatchAllError(rec, w, r, log)
		}
	}()

	var cp CreateProfileChange

	if !tools.UnmarshalBodyOrDie(&cp, w, r, log) {
		return
	}

	p, err := s.pm.Load(cp.AccessCode, 0)

	if err != nil {
		tools.NotFound("AccessCode", fmt.Sprintf("No user with access code %s has been found", cp.AccessCode), w, r, log)
	}

	p.CountryID = uint8(cp.CountryID)
	p.Avatar = uint8(cp.Avatar)
	p.Modifiers = uint64(cp.Modifiers)

	var pr CreateStatus
	pr.Status = "success"

	b, _ := json.Marshal(pr)
	w.Header().Set("Content-Type", tools.MimeJSON)
	w.WriteHeader(200)
	n, _ := w.Write(b)
	tools.LogExit(log, r, 200, n)
}

func (s *Server) getProfile(w http.ResponseWriter, r *http.Request) {
	tools.InitHTTPTimer(r)

	defer func() {
		if rec := recover(); rec != nil {
			tools.CatchAllError(rec, w, r, log)
		}
	}()

	q := r.URL.Query()

	ac := q.Get("AccessCode")

	if ac == "" {
		tools.InvalidFieldData("AccessCode", "You should provide a \"AccessCode\" query parameter.", w, r, log)
		return
	}

	p, err := s.pm.Load(ac, 0)

	if err != nil {
		tools.NotFound("AccessCode", fmt.Sprintf("No user with access code %s has been found", ac), w, r, log)
	}

	b, _ := json.Marshal(p)

	w.Header().Set("Content-Type", tools.MimeJSON)
	w.WriteHeader(200)
	n, _ := w.Write(b)
	tools.LogExit(log, r, 200, n)
}
