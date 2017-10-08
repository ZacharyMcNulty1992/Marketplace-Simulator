package actor

//Lumberjack - Struct containing lumberjack information.
type Lumberjack struct {
	Comm         chan event //This may only ever contain events for buying things,
	Name         string     //name of the actor
	Profession   string     //Lumberjack
	HungerRate   int        //rate at which hunger lowers per action.
	HungerLevel  int        //number between 1-100, 1 = dead from not eating, 100 = full
	IncomeRate   int        // This number is 0 for merchants, since they buy/sell goods.
	CurrentFunds int        // The current amount of money this actor has
}

func (l *Lumberjack) eventloop() error {
	//TODO: implenet the lumberjack ai.
	return nil
}

func (l *Lumberjack) Init() (Person, error) {
	l.Name = "Lumberjack"
	go l.eventloop()
	return l, nil
}
