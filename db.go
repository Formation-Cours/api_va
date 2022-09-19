package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	username = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	hostname = os.Getenv("HOSTNAME") + ":" + os.Getenv("PORT")
	dbname   = os.Getenv("DB_NAME")
)

func dsn(dbName string) string {
	if strings.TrimSpace(username) == "" ||
		strings.TrimSpace(password) == "" ||
		strings.TrimSpace(hostname) == "" ||
		strings.TrimSpace(dbname) == "" {
		log.Panicln("Problem Environment variables")
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func connecDB() *sql.DB {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Println(err)
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil
	}

	db.Close()

	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Println(err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
	}
	log.Printf("Connected to DB %s successfully\n", dbname)

	err = dropUserTable(db)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = createUserTable(db)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = insertUserTable(db)
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}

func createUserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS user(
		id int auto_increment primary key,
		firstname varchar(50),
		lastname varchar(50),
		email varchar(150)
	)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	//rows, err := res.RowsAffected()
	//if err != nil {
	//	log.Printf("Error %s when getting rows affected", err)
	//	return err
	//}
	//log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func dropUserTable(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS user`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	//rows, err := res.RowsAffected()
	//if err != nil {
	//	log.Printf("Error %s when getting rows affected", err)
	//	return err
	//}
	//log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func insertUserTable(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := db.ExecContext(ctx, sqlDB)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	//rows, err := res.RowsAffected()
	//if err != nil {
	//	log.Printf("Error %s when getting rows affected", err)
	//	return err
	//}
	//log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

const sqlDB = `insert into user (firstname, lastname, email) 
values 	('Robbie', 'McIlveen', 'rmcilveen0@blinklist.com'),
		('Hubey', 'Balden', 'hbalden1@narod.ru'),
		('Gaspard', 'Cosgreave', 'gcosgreave2@parallels.com'),
		('Brandi', 'Morcombe', 'bmorcombe3@mediafire.com'),
		('Bernhard', 'Mainland', 'bmainland4@prweb.com'),
		('Livy', 'Brownscombe', 'lbrownscombe5@narod.ru'),
		('Artair', 'Maccrae', 'amaccrae6@networkadvertising.org'),
		('Gustaf', 'Blaker', 'gblaker7@google.ca'),
		('Elwin', 'Steggals', 'esteggals8@gmpg.org'),
		('Dita', 'Possel', 'dpossel9@storify.com'),
		('Chad', 'Dumbar', 'cdumbara@unicef.org'),
		('Dene', 'Isakowicz', 'disakowiczb@omniture.com'),
		('Locke', 'Fillary', 'lfillaryc@google.cn'),
		('Wake', 'Wedgwood', 'wwedgwoodd@icio.us'),
		('Luz', 'Southern', 'lsoutherne@ask.com'),
		('Clarice', 'Bellward', 'cbellwardf@cornell.edu'),
		('Dorette', 'MacInerney', 'dmacinerneyg@mit.edu'),
		('Kaleena', 'Troop', 'ktrooph@hugedomains.com'),
		('Nell', 'Graver', 'ngraveri@patch.com'),
		('Martina', 'Boays', 'mboaysj@fda.gov'),
		('Remington', 'Axelbee', 'raxelbeek@blinklist.com'),
		('Genna', 'Deniske', 'gdeniskel@woothemes.com'),
		('Gerty', 'Willars', 'gwillarsm@youtube.com'),
		('Merilyn', 'Clougher', 'mcloughern@walmart.com'),
		('Shawn', 'Decort', 'sdecorto@mlb.com'),
		('Franklin', 'Lickorish', 'flickorishp@about.me'),
		('Clarie', 'Hof', 'chofq@walmart.com'),
		('Darrick', 'Brierley', 'dbrierleyr@cocolog-nifty.com'),
		('Lilia', 'Dunlop', 'ldunlops@bloglovin.com'),
		('Morgan', 'Heister', 'mheistert@cocolog-nifty.com'),
		('Herve', 'Alderton', 'haldertonu@nationalgeographic.com'),
		('Anallise', 'Goolden', 'agooldenv@artisteer.com'),
		('Clifford', 'Fonte', 'cfontew@earthlink.net'),
		('Remington', 'Malzard', 'rmalzardx@stanford.edu'),
		('Rivi', 'Langcaster', 'rlangcastery@hc360.com'),
		('Reggis', 'Syred', 'rsyredz@irs.gov'),
		('Lanny', 'Benedick', 'lbenedick10@opera.com'),
		('Erskine', 'Cornewall', 'ecornewall11@deviantart.com'),
		('Abel', 'Paz', 'apaz12@theglobeandmail.com'),
		('Darius', 'McGaffey', 'dmcgaffey13@is.gd'),
		('Rorke', 'Koomar', 'rkoomar14@hao123.com'),
		('Alexandre', 'Franseco', 'afranseco15@bandcamp.com'),
		('Anna-diane', 'Capun', 'acapun16@nasa.gov'),
		('Harp', 'Eagell', 'heagell17@free.fr'),
		('Yuri', 'Tander', 'ytander18@gnu.org'),
		('Marijn', 'Daymond', 'mdaymond19@godaddy.com'),
		('Jade', 'Jordin', 'jjordin1a@discovery.com'),
		('Nikos', 'Sheepy', 'nsheepy1b@discovery.com'),
		('Lilyan', 'Bleythin', 'lbleythin1c@gravatar.com'),
		('Joelynn', 'Marrion', 'jmarrion1d@usda.gov');
`
