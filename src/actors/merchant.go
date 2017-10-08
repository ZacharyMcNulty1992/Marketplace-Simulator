package actor

//Merchant - struct containing merchant values.
type Merchant struct {
	Name         string
	shopLocation string //<-- change to a node reference to a node on a graph of a city
	Comm         chan event
	Profession   string
	HungerRate   int //rate at which hunger lowers per action.
	HungerLevel  int //number between 1-100, 1 = dead from not eating, 100 = full
	IncomeRate   int // This number is 0 for merchants, since they buy/sell goods.
	CurrentFunds int // The current amount of money this actor has
}

func (m *Merchant) eventloop() error {
	//TODO: listen to the comm channel for events
	for {
		//TODO: use the global event listener to determine the next action to take.
		// this should implement some kind of time system.
		// This will allow the Merchant to determine what to do next.
		break
	}

	return nil
}

//Init - Will return an initialized merchant, and starts its event loop.
func (m *Merchant) Init() (Person, error) {
	//start the main event loop, so that the merchant will start its ai pattern
	m.Name = "random name"    //TODO: name generaton.
	m.Comm = make(chan event) //channel of structs.
	go m.eventloop()

	return m, nil
}
