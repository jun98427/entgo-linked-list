package main

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jun98427/linkedlist/ent"
	"github.com/jun98427/linkedlist/repository"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func printCmd() (int, error) {
	fmt.Print("1: add node\n2: delete node\n3: print list\n4: exit\nwrite cmd :")

	var cmd int
	_, err := fmt.Scanf("%d", &cmd)
	if err != nil {
		return 0, err
	}

	return cmd, nil
}

func openDB(dbURL string) *ent.Client {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dbURL := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", host, user, dbname, password)
	log.Println(dbURL)
	client := openDB(dbURL)
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	repo := repository.New(client)

	for {
		cmd, err := printCmd()
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch cmd {
		case 1:
			cnt, err := repo.PrintList(context.Background())
			if err != nil {
				log.Println(err)
				continue
			}

			var value, prevID int

			if cnt == 0 {
				fmt.Print("val : ")
				_, err = fmt.Scanf("%d", &value)
				if err != nil {
					fmt.Println("invalid value")
				}
				err = repo.AddNode(context.Background(), value, 0)
				if err != nil {
					log.Println(err)
					continue
				}
			} else {
				fmt.Print("prevID : ")
				_, err = fmt.Scanf("%d", &prevID)
				if err != nil {
					fmt.Println("invalid prevID")
				}
				fmt.Print("val : ")
				_, err = fmt.Scanf("%d", &value)
				if err != nil {
					fmt.Println("invalid value")
				}
				err = repo.AddNode(context.Background(), value, prevID)
				if err != nil {
					log.Println(err)
					continue
				}
			}
		case 2:
			_, err := repo.PrintList(context.Background())
			if err != nil {
				log.Println(err)
				continue
			}

			var id int
			fmt.Print("id : ")
			_, err = fmt.Scanf("%d", &id)
			if err != nil {
				fmt.Println("invalid id")
			}

			err = repo.DeleteNode(context.Background(), id)
			if err != nil {
				log.Println(err)
				continue
			}
		case 3:
			_, err := repo.PrintList(context.Background())
			if err != nil {
				log.Println(err)
				continue
			}
		case 4:
			fmt.Println("exit")
			return
		}
		fmt.Println()
	}
}
