package main

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		name := r.FormValue("name")
		body := r.FormValue("body")
		note := &Note{Name: name, Body: body}
		err := Save(note)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", 302)
		return
	}

	//http://localhost:8080/?action=delete&id=605482d9d3f174da49ba0a33
	if r.URL.Query().Get("action") == "delete" {
		URLid := r.URL.Query().Get("id")
		if _id, err := primitive.ObjectIDFromHex(URLid); err == nil {
			Delete(_id)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}

	html := `<!DOCTYPE html><html><head><title>Guestbook</title>
	<style>
    body{
        background-color: aquamarine;
    }
    h2{
        text-align: center;
        color: red;
        font-weight: 600;
    }
    form{
        border: 1px solid blue;
    }
	</style>
	</head>
	<body>
    <h2 for="welcome">Welcome to my guestbook!</h2>`

	html += `
	<form action="/" method="POST">
    <div>
        <textarea id="name" name="name" rows="1" cols="50" placeholder="Leave your name here"></textarea>
    </div>
    <div>
        <textarea id="body" name="body" rows="20" cols="50" placeholder="Leave your note here"></textarea>
    </div>
    <div>
        <input type="submit" value="save">
    </div>
    </form>
	`

	notes := getList()

	html += "<hr/>"
	for _, note := range notes {
		html += `
		<div>
			<h3 style="padding:5px 10px 5px 10px ; background-color:green; color:#fff">` + note.Name + `</h3>
		
			<p>` + note.Body + note.Time.String() + `</p>

			<a href="/?action=delete&id=` + note.ID.Hex() + `">REMOVE</a>
		</div>
		`
	}
	html += "<hr/>"

	html += `</body></html>`

	w.Write([]byte(html))
}
