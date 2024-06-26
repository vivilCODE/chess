package dbhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vivilCODE/chess/chessapi/log"

	"github.com/vivilCODE/chess/db/models"
)

type DBHandler struct {
	Config *Config
}

type Config struct {
	Address string
}

func (h *DBHandler) GetUser(id string) (*models.User, error) {
	reqString := fmt.Sprintf("http://localhost:8082/db/users/%s", id)

	res, err := http.Get(reqString)
	if err != nil {
		return nil, fmt.Errorf("unable to get user from database service: %v", err)
	}

	// If we get a 404 response that means nothing went wrong, but there simply was not an ID match for that user
	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read getuser response body: %v", err)	
	}

	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("unable to unmarshal getuser response: %v", err)
	}


	log.Logger.Debug("fetched user", "user", user)

	return &user, nil
}

func (h *DBHandler) PostUser(user *models.User) error {
	// Marshal user struct
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("unable to marshal user: %v", err)
	}

	res, err := http.Post("http://localhost:8082/db/users", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		return fmt.Errorf("unable to post user to database: %v", err)
	}

	log.Logger.Debug("posted user to database", "response", res)

	return nil
}



