package actor

type Merchant struct{
    Name string
    shopLocation string //<-- change to a node reference to a node on a graph of a city
    Comm chan
    Profession string
}

func (m *Merchant) eventLoop(){
    //TODO: listen to the comm channel for events
    for{
        //do stuff
    }
}

func (m *Merchant) Init() (Person,error) {
    //start the main event loop, so that the merchant will start its ai pattern
    m.Name = "random name"
    m.Comm = new(chan)
    go eventLoop();

    return m, nil
}
