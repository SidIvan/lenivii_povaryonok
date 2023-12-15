package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var db *sql.DB

const version = "1.0"

type ingredient struct {
	Id          int64
	ProductName string
}

type output struct {
	Version     string
	Ingredients []ingredient
	Timestamp   time.Time
}

type inputRecipe struct {
	Ids []int `json:"ingredients_ids" validate:"required"`
}

type Recipe struct {
	id          int
	name        string
	instruction string
	ingredients pq.Int32Array
}

type finalRecipe struct {
	Id          int
	Name        string
	Instruction string
	Ingredients []int32
}

func createConnection() *sql.DB {
	connStr := "postgres://povaryonok:password@db:5432/povaryonok?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func GetIngredients(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // проверяем GET ли метод
		http.Error(w, "use of wrong method", http.StatusMethodNotAllowed)
		return
	}
	db = createConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	rows, err := db.Query("SELECT * FROM ingredients")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ingredients []ingredient
	for rows.Next() {
		var alb ingredient
		if err := rows.Scan(&alb.Id, &alb.ProductName); err != nil {
			log.Fatal(err)
		}
		ingredients = append(ingredients, alb)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	var getJson output
	getJson.Version = version
	getJson.Ingredients = ingredients
	getJson.Timestamp = time.Now()
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(getJson)
	if err != nil {
		http.Error(w, "failed to make output struct", http.StatusBadRequest)
		return
	}
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // проверяем GET ли метод
		http.Error(w, "use of wrong method", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" { // проверяем заголовок
		http.Error(w, "wrong content-type header or content-type header not set", http.StatusBadRequest)
		return
	}
	var body inputRecipe
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil { // проверяем получилось ли открыть json
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(body)
	if err != nil { // проверяем, что поле json действительно называется ingredients_ids
		http.Error(w, "failed to validate struct", http.StatusBadRequest)
		return
	}
	array := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(body.Ids)), ", "), "[]")
	db = createConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	query := fmt.Sprintf(`
		WITH myconstants (user_have) as (
			values ('{%s}'::int[])
		)
		SELECT id, name, instruction, ingredients
		FROM recipes, myconstants
		WHERE icount(ingredients | user_have) / icount(user_have) >= 0.75
		ORDER BY icount(ingredients - user_have)
		LIMIT 10
		;`, array)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var data_to_return []finalRecipe
	for rows.Next() {
		var row Recipe
		err := rows.Scan(&row.id, &row.name, &row.instruction, &row.ingredients)
		if err != nil {
			log.Fatal(err)
		}
		elem := finalRecipe{row.id, row.name, row.instruction, []int32(row.ingredients)}
		data_to_return = append(data_to_return, elem)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data_to_return)
	if err != nil {
		http.Error(w, "failed to make output struct", http.StatusBadRequest)
		return
	}
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // проверяем GET ли метод
		http.Error(w, "use of wrong method", http.StatusMethodNotAllowed)
		return
	}
	_, err := io.WriteString(w, "ok\n")
	if err != nil {
		log.Fatal(err)
	}
}
