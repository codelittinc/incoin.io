package main

import (
  "bytes"
  "net/http"
  "os"
  "text/template"

  "github.com/gorilla/mux"
  "github.com/mattbaird/gochimp"
  log "github.com/sirupsen/logrus"
)

var (
  _   gochimp.MandrillError
  tpl *template.Template
  api *gochimp.MandrillAPI
)

func main() {
  var err error

  defer func() {
    log.Info("Shutting down")
  }()

  log.SetLevel(log.DebugLevel)
  log.SetFormatter(&log.JSONFormatter{})

  parseTemplate()
  if api, err = gochimp.NewMandrill(os.Getenv("MANDRILL_KEY")); err != nil {
    log.Panic("Unable to create mandrill API")
  }

  router := mux.NewRouter()

  // These GET routes will be handled via the Marionette frontend
  router.HandleFunc("/", http.HandlerFunc(index)).Methods("GET")
  router.HandleFunc("/features", http.HandlerFunc(features)).Methods("GET")
  router.HandleFunc("/about", http.HandlerFunc(about)).Methods("GET")
  router.HandleFunc("/contact", http.HandlerFunc(contact)).Methods("GET")
  router.HandleFunc("/pricing", http.HandlerFunc(pricing)).Methods("GET")
  router.HandleFunc("/privacy", http.HandlerFunc(privacy)).Methods("GET")
  // Accept post data and fake-register the product
  router.HandleFunc("/email", http.HandlerFunc(send_email)).Methods("POST")

  http.HandleFunc("/assets/", http.HandlerFunc(assets))
  http.Handle("/", router)

  log.Info("Listening on port 4000")
  http.ListenAndServe(":4000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "public/index.html")
}

func features(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "public/features.html")
}

func about(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "public/about.html")
}

func pricing(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "public/pricing.html")
}

func contact(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "public/contact_us.html")
}

func privacy(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "public/privacy-terms-conditions.html")
}

func send_email(w http.ResponseWriter, r *http.Request) {
  // 1. Accept incoming post data
  // 2. Send an email
  r.ParseForm()

  go func() {
    defer func() {
      if e := recover(); e != nil {
        log.WithFields(log.Fields{
          "error": e,
        }).Error("Error sending email")
      }
    }()

    data := map[string]string{
      "Name":    r.Form["name"][0],
      "Email":   r.Form["email"][0],
      "Comment": r.Form["comment"][0],
    }

    buf := &bytes.Buffer{}
    if err := tpl.Execute(buf, data); err != nil {
      log.Error("Error executing template")
    }

    recipients := []gochimp.Recipient{
      gochimp.Recipient{Email: "hello@incoin.io"},
    }

    message := gochimp.Message{
      Html:      buf.String(),
      Subject:   "Contact Form Submission from " + r.Form["name"][0],
      FromEmail: r.Form["email"][0],
      FromName:  r.Form["name"][0],
      To:        recipients,
    }

    if _, err := api.MessageSend(message, false); err != nil {
      log.Error("Error sending email")
    }
  }()

  w.Header().Add("location", "/contact")
  w.WriteHeader(303)
}

func parseTemplate() {
  var err error
  tpl, err = template.New("email").Parse(`<p>name={{.Name}}</p><br>
<p>email={{.Email}}</p><br>
<p>comment={{.Comment}}</p>`)

  if err != nil {
    panic(err.Error())
  }
}

func redirect(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("location", "/")
  w.WriteHeader(303)
}

func assets(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "public/"+r.URL.Path[1:])
}