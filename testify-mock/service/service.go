package service

type db struct{}

// DB is fake database interface
type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

type greeter struct {
	database DB
	lang     string
}

// GreeterService is service to greet your friends.
type GreeterService interface {
	Greet() string
	GreetInDefaultMsg() string
}

// NewDB is a factory function to return db struct
func NewDB() DB {
	return new(db)
}

// NewGreeter is a factory function to treturn greeter struct
func NewGreeter(db DB, lang string) GreeterService {
	return greeter{db, lang}
}

func (d *db) FetchMessage(lang string) (string, error) {
	// in real life, this code will call an external db
	// but for this sample we will just return the hardcoded example value
	if lang == "en" {
		return "hello", nil
	}
	if lang == "es" {
		return "holla", nil
	}
	return "bzzzz", nil
}

func (d *db) FetchDefaultMessage() (string, error) {
	return "default message", nil
}

func (g greeter) Greet() string {
	msg, _ := g.database.FetchMessage(g.lang) // call database to get the message based on the lang
	return "Message is: " + msg
}

func (g greeter) GreetInDefaultMsg() string {
	msg, _ := g.database.FetchDefaultMessage() // call database to get the default message
	return "Message is: " + msg
}
