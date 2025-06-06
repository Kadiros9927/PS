package verify

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/jordan-wright/email"
	"go/git-ps/3-validation-api/config"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//todo type VerifyHandlerDeps struct {}

var hashStorage = make(map[string]string)

func NewVerifyHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()
		recipient := r.FormValue("email")

		if recipient == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		rand.Seed((time.Now().UnixNano()))
		randomData := fmt.Sprintf("%s:%d", recipient, rand.Int63())
		sha := sha256.Sum256([]byte(randomData))
		hash := hex.EncodeToString(sha[:])

		rec := Record{Email: recipient, Hash: hash}
		err := SaveRecord(rec)
		if err != nil {
			http.Error(w, "Ошибка сохранения данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		verifyURL := fmt.Sprintf("https://localhost:8080/verify/%s", hash)

		createAndSendVerifyEmail(recipient, verifyURL, cfg)
	}
}

func VerifyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 || parts[2] == "" {
			http.Error(w, "Invalid verification link", http.StatusBadRequest)
			return
		}
		hash := parts[2]

		email, found := DeleteByHash(hash)

		w.Header().Set("Content-Type", "application/json")

		if found {
			io.WriteString(w, `{"success": true, "email": "`+email+`"}`)
		} else {
			io.WriteString(w, `{"success": false}`)
		}

		fmt.Printf("Email %s has been successfully verified!", email)
	}
}

func createAndSendVerifyEmail(recipient, verifyURL string, cfg config.Config) {

	e := email.NewEmail()
	e.From = cfg.Email
	e.To = []string{recipient}
	e.Subject = "Verify your email"
	e.Text = []byte(fmt.Sprintf("Click the link to verify your email: %s", verifyURL))

	fmt.Println(string(e.Text))
	//err := e.Send(cfg.Address, smtp.PlainAuth("", cfg.Email, cfg.Password, cfg.Address[:strings.Index(cfg.Address, ":")]))

	//if err != nil {
	//	fmt.Println("Не удалось отправить email: "+err.Error(), http.StatusInternalServerError)
	//}
	fmt.Println("Email sent successfully to " + recipient)
}
