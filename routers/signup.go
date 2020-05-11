package routers

import (
	"encoding/json"
	"net/http"

	"github.com/chrisaandes/twitter-clone/models"
)

/*SignUp ...*/
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "BAD Request"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email required!", 400)
		return
	}

	if len(user.Email) < 6 {
		http.Error(w, "Min 6 characters password", 400)
		return
	}

	_, exist, _ := db.checkUser(user.Email)
	if exist {
		http.Error(w, "User already exists", 403)
		return
	}

	_, status, err := db.insert(t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "something went wrong", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
