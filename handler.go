package messenger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct {
	handle      func(*entry, *Messaging)
	verifyToken string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()

		if mode := r.Form.Get("hub.mode"); mode != "subscribe" {
			log.Printf("unknown mode: %s", mode)
			return
		}

		if verifyToken := r.Form.Get("hub.verify_token"); verifyToken != h.verifyToken {
			log.Printf("invalid verify token: %s", verifyToken)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		fmt.Fprint(w, r.Form.Get("hub.challenge"))

		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("could not read request body: %v", err)
		return
	}

	var res response
	if err := json.Unmarshal(body, &res); err != nil {
		log.Printf("could not unmarshal request body: %v", err)
		return
	}

	if res.Object != "page" {
		log.Printf("unknown object: %s", res.Object)
		return
	}

	for _, entry := range res.Entry {
		for _, messaging := range entry.Messaging {
			h.handle(entry, messaging)
		}
	}
}

// Handler func
func (ms *Messenger) Handler() http.Handler {
	return handler{
		handle:      ms.handle,
		verifyToken: ms.verifyToken,
	}
}
