package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/tonky/shipping/PortDomainService/domain"
)

type pgStorage struct {
	db *sql.DB
}

func (pg pgStorage) Add(ports []domain.Port) error {
	log.Printf("pgStorage Add: adding %d ports", len(ports))
	ts := time.Now()
	err := pg.upsertCopy(ports)

	log.Printf("Upserted %d ports in %dms\n", len(ports), time.Now().Sub(ts).Milliseconds())

	return err

	/*
		// sample batching implementation
		batchStart := time.Now()

		batchSize := 10000
		i := 0

		for {
			bs := batchSize * i
			be := batchSize * (i + 1)

			if be > len(ports) {
				be = len(ports)
			}

			st := time.Now()

			batch := ports[bs:be]

			if err := pg.upsertCopy(batch); err != nil {
				log.Println("Error in upsertCopy: ", err)
				return err
			}

			log.Printf("Upserted %d batch in %dms\n", len(batch), time.Now().Sub(st).Milliseconds())

			if be == len(ports) {
				break
			}

			i++
		}
	*/
}

func (pg pgStorage) Get(name string) (*domain.Port, error) {
	log.Printf("pgStorage Get: %s\n", name)

	row := pg.db.QueryRow(`SELECT name, data FROM ports WHERE name = $1`, name)

	var key string
	var data string

	if err := row.Scan(&key, &data); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		log.Println("Error querying row: ", err)
		return nil, err
	}

	log.Printf("search name: %s, found key %s, data %s\n", name, key, data)

	var p domain.Port

	err := json.Unmarshal([]byte(data), &p)

	if err != nil {
		log.Println("Error unmarshalling data: ", err)
		return nil, err
	}

	return &p, nil
}

func New(user, password, host, port, dbname string) (pgStorage, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	log.Println("Connecting to db: ", dbUrl)

	for {
		db, err := sql.Open("postgres", dbUrl)

		if err != nil {
			log.Println("Connection error: ", err)
			time.Sleep(300 * time.Millisecond)
			continue
		}

		log.Println("sql.Open() ok")

		if err := db.Ping(); err != nil {
			log.Println("Couldn't ping postgres: ", err)
			time.Sleep(300 * time.Millisecond)
			continue
		}

		log.Println("Ping ok, DB connected")

		if _, err := db.Exec(`CREATE TABLE if not exists ports (name TEXT PRIMARY KEY, data JSONB)`); err != nil {
			log.Println("Error creating table: ", err)
			continue
		}

		if _, err := db.Exec(`CREATE TABLE if not exists ports_copy (name TEXT PRIMARY KEY, data JSONB)`); err != nil {
			log.Println("Error creating table: ", err)
			continue
		}

		return pgStorage{db}, nil
	}
}

func (pg pgStorage) upsertCopy(ports []domain.Port) error {
	if _, err := pg.db.Exec(`TRUNCATE TABLE ports_copy`); err != nil {
		log.Println("Error dropping table: ", err)
		return err
	}

	st := time.Now()

	txn, err := pg.db.Begin()
	if err != nil {
		log.Println("Error beginning txn")
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("ports_copy", "name", "data"))
	if err != nil {
		log.Println("Error preparing statement")
		return err
	}

	defer stmt.Close()

	for _, p := range ports {
		data, err := json.Marshal(p)
		if err != nil {
			log.Println("Error marshalling port: ", err)
			return err
		}

		_, err = stmt.Exec(p.Key, string(data))
		if err != nil {
			log.Println("Error executing statement: ", p.Key, string(data), err)
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error executing final statement", err)
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println("Error closing final statement", err)
		return err
	}

	if err := txn.Commit(); err != nil {
		log.Println("Error committing COPY IN transaction", err)
		return err
	}

	log.Printf("COPY IN successful, added %d ports in %dms, begin upserting data into `ports`\n", len(ports), time.Now().Sub(st).Milliseconds())

	q := `insert into ports (name, data) select pc.name, pc.data from ports_copy pc on conflict (name) do update set data = excluded.data`

	if _, err := pg.db.Exec(q); err != nil {
		log.Println("Error upserting records from `ports_copy` to `ports`: ", err)
		return err
	}

	return nil
}
