package loops

import "fmt"

func ForLoopCopyValue() error {
	type (
		account struct {
			balance float32
		}
	)

	var (
		accounts = []account{
			{balance: 100.},
			{balance: 200.},
			{balance: 300.},
		}
		p = func(accs []account) {
			for _, v := range accs {
				fmt.Printf("%v", v)
			}
			fmt.Println("")
		}
	)

	for _, v := range accounts {
		v.balance += 1000.0
	}

	//
	// which of the following two choices do you think shows the slice’s content?
	//
	// (1) {100}{200}{300}
	// (2) {1100}{1200}{1300}
	//
	p(accounts)

	//
	// The right way
	//
	for i := range accounts {
		accounts[i].balance += 1000.0
	}
	p(accounts)

	return nil
}
