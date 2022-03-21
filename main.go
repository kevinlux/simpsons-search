package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

type result struct {
	Title   string
	Content template.HTML
	Snippet template.HTML
}

func main() {
	godotenv.Load()
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dburl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal(err)
	}

	tplHome := template.Must(template.New(".").Parse(tplStrHome))
	tplResults := template.Must(template.New(".").Parse(tplStrResults))
	tplEpisode := template.Must(template.New(".").Parse(tplStrEpisode))

	http.HandleFunc("/episode", func(w http.ResponseWriter, r *http.Request) {
		e := r.FormValue("e")
		var title string
		var content template.HTML
		if err := db.QueryRow(sqlGetEpisode, e).Scan(&title, &content); err != nil {
			fmt.Println(title)
			http.Error(w, "not found", 404)
			return
		} else {
			tplEpisode.Execute(w, map[string]interface{}{
				"Title":   e,
				"Content": content,
			})
		}

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.FormValue("q")
		if q == "" {
			tplHome.Execute(w, nil)
			return
		}
		if len(q) > 100 {
			q = q[:100]
		}
		rows, err := db.Query(SqlSearch, q)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		defer rows.Close()
		results := make([]result, 0, 10)
		for rows.Next() {
			var r result
			if err := rows.Scan(&r.Title, &r.Snippet); err != nil {
				fmt.Println(r.Title)
				http.Error(w, "not found", 404)
				return
			}
			results = append(results, r)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "not found", 404)
		}
		tplResults.Execute(w, map[string]interface{}{
			"Results": results,
			"Query":   q,
		})
	})

	http.ListenAndServe(":"+port, nil)

}
