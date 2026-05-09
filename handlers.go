package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func landingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodPost {
		defer r.Body.Close()

		name := strings.TrimSpace(r.FormValue("name"))

		if name == "" {
			name = "Anonymous"
		}

		sid := generateID()

		mu.Lock()
		sessions[sid] = &Session{
			Name:      name,
			MaxScore:  len(questions) * 10,
			CreatedAt: time.Now(),
		}
		mu.Unlock()

		http.SetCookie(w, &http.Cookie{
			Name:  "sid",
			Value: sid,
			Path:  "/",
		})

		http.Redirect(w, r, "/quiz", http.StatusSeeOther)
		return
	}

	landingTmpl.Execute(w, nil)
}

func quizHandler(w http.ResponseWriter, r *http.Request) {
	sess, _ := getSession(r)

	if sess == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if sess.CurrentQuestion >= len(questions) {
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		defer r.Body.Close()

		idx, err := strconv.Atoi(r.FormValue("answer"))

		if err == nil &&
			idx >= 0 &&
			idx < len(questions[sess.CurrentQuestion].Options) {

			mu.Lock()
			sess.TotalScore += questions[sess.CurrentQuestion].Options[idx].Score
			sess.CurrentQuestion++
			mu.Unlock()
		}

		if sess.CurrentQuestion >= len(questions) {
			http.Redirect(w, r, "/result", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/quiz", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Name":     sess.Name,
		"Question": questions[sess.CurrentQuestion],
		"Number":   sess.CurrentQuestion + 1,
		"Total":    len(questions),
		"Progress": ((sess.CurrentQuestion + 1) * 100) / len(questions),
	}

	quizTmpl.Execute(w, data)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	sess, sid := getSession(r)

	if sess == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pct := (sess.TotalScore * 100) / sess.MaxScore

	var (
		label string
		desc  string
		emoji string
		color template.CSS
	)

	switch {
	case pct >= 85:
		label = "Day zeroes tong tropa ko"
		desc = "You lead with compassion and integrity."
		emoji = "🌟"
		color = "#4ade80"

	case pct >= 65:
		label = "Sheesh parang temple hood ka ya"
		desc = "You have a good heart."
		emoji = "💚"
		color = "#86efac"

	case pct >= 40:
		label = "Talaga bang may bitaw 'yan?"
		desc = "There's good in you."
		emoji = "🌫️"
		color = "#fbbf24"

	default:
		label = "Hindi ka day one sah"
		desc = "Empathy needs practice."
		emoji = "🧊"
		color = "#94a3b8"
	}

	mu.Lock()
	delete(sessions, sid)
	mu.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:   "sid",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	resultTmpl.Execute(w, Result{
		DashOffset: 339 - (sess.TotalScore * 339 / sess.MaxScore),
		Name:       sess.Name,
		Score:      sess.TotalScore,
		Max:        sess.MaxScore,
		Percent:    pct,
		Label:      label,
		Desc:       desc,
		Emoji:      emoji,
		Color:      color,
	})
}
