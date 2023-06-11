package main

import "fmt"

// Decorator pattern
// Scenerio:  Usually this scenario arises when we use a third party package, and we want to add extra functionalities to it
// For example, when using an http handler we may want to first have access to the databse to check if an user is authentitcated (middleware)
// This pattern can be used to pass that extra parameter that we need before actually executing the function of the third package
type DB interface {
	Store(string) error
}

type Store struct{}

func (store *Store) Store(val string) error {
	fmt.Println("Storing in DB, ", val)
	return nil
}

type LibraryFunc func(string)

func main() {
	store := &Store{}
	Execute(Printer(store))
	testFunctionalOptionsPattern()
}

func Printer(db DB) LibraryFunc {
	return func(s string) {
		db.Store(s)
		fmt.Println("Printer: ", s)
	}
}

func Execute(fn LibraryFunc) {
	fn("cat")
}

// Functional Options pattern
// Scenario: we have a struct that takes n number of parameters, but that eventaully can reach m parameters where n < m
// Having a constructor to provide all the parameters can be tedious for the user of the library and also much harder to maintain
// Solution; We use the functional options pattern to fix this issue

type Opts struct {
	addr           string
	maxConnections uint
	tls            bool
}

type Server struct {
	Opts
}

type OptsFunc func(*Opts)

func defaultOpts() Opts {
	return Opts{
		addr:           "4000",
		maxConnections: 100,
		tls:            false,
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func MaxConnections(n uint) OptsFunc {
	return func(opts *Opts) {
		opts.maxConnections = n
	}
}

func withAddr(addr string) OptsFunc {
	return func(o *Opts) {
		o.addr = addr
	}
}

func NewServer(opts ...OptsFunc) *Server {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		Opts: o,
	}
}

func testFunctionalOptionsPattern() {
	o := NewServer(MaxConnections(1000), withAddr("4000"), withTLS)
	fmt.Printf("%+v\n", o)
}
