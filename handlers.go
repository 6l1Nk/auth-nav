package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	// "bytes"
	// "math"
	// "net/url"
	// "strconv"
	//
	// "github.com/betterstack-community/go-logging/logger"
	// "go.uber.org/zap"
)

type SignUpFormData struct {
	Email    string
	Password string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl := template.Must(template.ParseFiles("templates/sign-up.html"))
		tmpl.Execute(w, nil)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		formData := SignUpFormData{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		// Generate a salted hash of the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formData.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
			return
		}

		err = saveNewUser(formData.Email, hashedPassword)
		if err != nil {
			http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
			return
		}

		// Redirect to a success page
		http.Redirect(w, r, "/success", http.StatusFound)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// func signInHandler(w http.ResponseWriter, r *http.Request) {
//     tmpl := template.Must(template.ParseFiles("templates/signin.html"))
//     tmpl.Execute(w, nil)
// // Verify a password against the hashed password
// err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(formData.Password))
// if err != nil {
// 	fmt.Println("Password does not match hash:", err)
// 	return
// }
// fmt.Println("Password matches hash!")
// //
// }

// type handlerWithError func(w http.ResponseWriter, r *http.Request) error
//
// func (fn handlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	l := logger.FromCtx(r.Context())
//
// 	err := fn(w, r)
// 	if err != nil {
// 		l.Error("an unexpected error occurred", zap.Error(err))
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
//
// 		return
// 	}
// }
//
// func indexHandler(w http.ResponseWriter, r *http.Request) error {
// 	l := logger.FromCtx(r.Context())
//
// 	l.Debug("entered indexHandler()")
//
// 	if r.URL.Path != "/" {
// 		l.Sugar().Debugf("path '%s' not found", r.URL.Path)
// 		http.NotFound(w, r)
//
// 		return nil
// 	}
//
// 	buf := &bytes.Buffer{}
//
// 	err := tpl.Execute(buf, nil)
// 	if err != nil {
// 		return err
// 	}
//
// 	_, err = buf.WriteTo(w)
//
// 	return err
// }
//
// func searchHandler(w http.ResponseWriter, r *http.Request) error {
// 	ctx := r.Context()
//
// 	l := logger.FromCtx(ctx)
//
// 	l.Debug("entered searchHandler()")
//
// 	u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		return err
// 	}
//
// 	params := u.Query()
// 	searchQuery := params.Get("q")
//
// 	pageNum := params.Get("page")
// 	if pageNum == "" {
// 		pageNum = "1"
// 	}
//
// 	l = l.With(
// 		zap.String("search_query", searchQuery),
// 		zap.String("page_num", pageNum),
// 	)
//
// 	l.Sugar().Infof(
// 		"incoming search query '%s' on page '%s'",
// 		searchQuery,
// 		pageNum,
// 	)
//
// 	nextPage, err := strconv.Atoi(pageNum)
// 	if err != nil {
// 		return err
// 	}
//
// 	pageSize := 20
//
// 	resultsOffset := (nextPage - 1) * pageSize
//
// 	searchResponse, err := searchWikipedia(
// 		ctx,
// 		searchQuery,
// 		pageSize,
// 		resultsOffset,
// 	)
// 	if err != nil {
// 		return err
// 	}
//
// 	l.Debug(
// 		"search response from Wikipedia",
// 		zap.Any("search_response", searchResponse),
// 	)
//
// 	totalHits := searchResponse.Query.SearchInfo.TotalHits
//
// 	search := &Search{
// 		Query:      searchQuery,
// 		Results:    searchResponse,
// 		TotalPages: int(math.Ceil(float64(totalHits) / float64(pageSize))),
// 		NextPage:   nextPage + 1,
// 	}
//
// 	buf := &bytes.Buffer{}
//
// 	err = tpl.Execute(buf, search)
// 	if err != nil {
// 		return err
// 	}
//
// 	_, err = buf.WriteTo(w)
// 	if err != nil {
// 		return err
// 	}
//
// 	l.Debug("search query succeeded without errors")
//
// 	return nil
// }
