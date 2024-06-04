package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

var database *sql.DB

// db connection
func init() {
	var err error

	database, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS games (id INTEGER PRIMARY KEY, name TEXT, genre TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func CheckIDExists(id int) bool {
	var exists bool
	_ = database.QueryRow("SELECT EXISTS(SELECT 1 FROM games WHERE id = ?)", id).Scan(&exists)
	return exists
}

func AddGame(name, genre string) error {
	statement, err := database.Prepare("INSERT INTO games (name, genre) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(name, genre)
	if err != nil {
		return err
	}

	if err := statement.Close(); err != nil {
		return err
	}

	return err
}

func ViewAll() error {
	rows, err := database.Query("SELECT id, name, genre FROM games")
	if err != nil {
		return err
	}

	var id int
	var name string
	var genre string

	for rows.Next() {
		err := rows.Scan(&id, &name, &genre)
		if err != nil {
			return err
		}

		fmt.Println(strconv.Itoa(id) + ": " + name + " " + genre)
	}

	if err := rows.Close(); err != nil {
		return err
	}

	return err
}

func Update(id int, name, genre string) error {
	if !CheckIDExists(id) {
		return errors.New("ID does not exist")
	}

	statement, err := database.Prepare("UPDATE games SET name = ?, genre = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = statement.Exec(name, genre, id)
	if err != nil {
		return err
	}

	if err := statement.Close(); err != nil {
		return err
	}
	return err
}

func Delete(id int) error {
	if !CheckIDExists(id) {
		return errors.New("ID does not exist")
	}

	statement, err := database.Prepare("DELETE FROM games WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	if err := statement.Close(); err != nil {
		return err
	}
	return err
}

func main() {
	var menuCode int

	for {
		fmt.Println("Select an option:")
		fmt.Println("1.) Add Game")
		fmt.Println("2.) View Games")
		fmt.Println("3.) Update Game")
		fmt.Println("4.) Delete Game")
		fmt.Println("5.) Quit")

		// Read user input
		fmt.Print(">> ")
		_, err := fmt.Scan(&menuCode)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch menuCode {
		// Add game ------------>
		case 1:
			var name, genre string

			fmt.Print("Enter Name & Genre >> ")

			_, err := fmt.Scan(&name, &genre)
			if err != nil {
				fmt.Println("[Error]:", err)
				break
			}

			err = AddGame(name, genre)
			if err != nil {
				fmt.Println("[Error]:", err)
				break
			}
			fmt.Println("Game added!")

		// View All Games ------------>
		case 2:
			err = ViewAll()
			if err != nil {
				fmt.Println("[Error]:", err)
				break
			}
			fmt.Println("")

		// Update Game ------------>
		case 3:
			var id int
			var name, genre string

			fmt.Print("Enter ID, Name & Genre >> ")

			_, err := fmt.Scan(&id, &name, &genre)
			if err != nil {
				fmt.Println("[Error]:", err)
				break
			}

			err = Update(id, name, genre)
			if err != nil {
				fmt.Println("[Error]:", err)
				break
			}
			fmt.Println("Game Updated!")

		// Delete Game ------------>
		case 4:
			var id int

			fmt.Print("Enter ID >> ")

			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("[Error]:", err)
				break
			}

			err = Delete(id)
			if err != nil {
				fmt.Println("[Error]:", err)
				break
			}
			
			fmt.Println("Game Deleted!")

		// Exit ------------>
		case 5:
			fmt.Println("Exiting...")
			goto exit

		default:
			fmt.Println("\nError: Only values from the given menu are accepted")
		}

	}
exit:
	fmt.Println("\n------------------")
}
