package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	pb "github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
	"google.golang.org/protobuf/proto"
)

func promptForAddress(r *io.Reader) (*pb.Person, error) {
	p := &pb.Person{}

	rd := bufio.NewReader(r)
	fmt.Print("Enter person ID: ")

	if _, err := fmt.scanf(rd, "%d\n", &p.ID); err != nil {
		return p, err
	}

	fmt.Print("Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}

	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
            break
        }

		pn := &pb.Person_PhoneNumber{
			Number: phone,
		}

		fmt.Print("Is this a mobile, home, or work phone? ")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}

		
	}

	rd := bufio.NewReader(r)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	addr, err := promptForAddress(os.Stdin)
	if err != nil {
		log.Fatalln("Error with address:", err)
	}

	book.People = append(book.People, addr)

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to marshal address book:", err)
	}

	if err := iouitl.WriteFile(fname, out, 0664); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}