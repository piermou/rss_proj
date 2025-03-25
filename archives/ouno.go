package archives

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import du driver PostgreSQL
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albumsByArtist queries for albums that have the specified artist name.

func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.

func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = $1", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum ajoute l'album spécifié à la base de données,
// et renvoie l'ID de l'album nouvellement inséré.

func addAlbum(alb Album) (int64, error) {
	// Utilisation de la syntaxe PostgreSQL pour les placeholders $1, $2, $3
	query := "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id"

	// Exécution de la requête avec les valeurs pour title, artist et price
	var id int64
	err := db.QueryRow(query, alb.Title, alb.Artist, alb.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// // Capture connection properties.
// connStr := fmt.Sprintf(
// 	"host=127.0.0.1 port=5432 user=%s password=%s dbname=recodings sslmode=disable",
// 	os.Getenv("DBUSER"), // Récupère le user depuis les variables d'env
// 	os.Getenv("DBPASS"), // Récupère le password depuis les variables d'env
// )

// // Get a database handle.
// var err error
// db, err = sql.Open("postgres", connStr)
// if err != nil {
// 	log.Fatal(err)
// }

// pingErr := db.Ping()
// if pingErr != nil {
// 	log.Fatal(pingErr)
// }

// fmt.Println("Connected to PostgreSQL!")

// albums, err := albumsByArtist("John Coltrane")
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("Albums found: %v\n", albums)

// // Hard-code ID 2 here to test the query.
// alb, err := albumByID(3)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("Album found: %v\n", alb)

// albID, err := addAlbum(Album{
// 	Title:  "The Modern Sound of Betty Carter",
// 	Artist: "Betty Carter",
// 	Price:  49.99,
// })
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Printf("ID of added album: %v\n", albID)
//
//
// // string can be converted to byte slices and vice versa

// var s string = "this is a string"
// 	fmt.Println(s)
// 	var b []byte
// 	b = []byte(s)
// 	fmt.Println(b)
// 	for i := range b {
// 		fmt.Println(string(b[i]))
// 	}
// 	s = string(b)
// 	fmt.Println(s)

// func main() {
// 	str := "something"
// 	buf := bytes.NewBufferString(str)
// 	for i := 0; i < 10; i++ {
// 		buf.Write([]byte(randomString()))
// 	}
// 	fmt.Println(buf.String())
// }

// func randomString() string {
// 	ret := "pretend-this-is-random"
// 	return ret
// }

// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// func RandString(length int) string {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[r.Intn(len(charset))]
// 	}
// 	return string(b)
// }
