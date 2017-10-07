package actor

type Person interface{
    Comm chan
    Name string
    Professtion string
    HungerLevel int //level of hunger 100 MAX, 0 and the character dies in the next loop iteration
    HungerRate int  //rate at which hunger will decrease wile working/performing actions
    IncomeRate int //rate at which the person will be gaining income. only affects non-merchant/non-chef type actors

    //functions -- methods in which are shared between actors
    //             Different types of actors will have several different functions
    //             these functions will relate to how that actor will act in the world

    //event loop - this is the handler loop that will control decisions that an actor can/will perform
    eventloop() error
    //Init -- will start the actors main event loop. This will cause the actor to
    Init() (Person, error)

}

func (p Person) NewPerson(Type string){

    if(Type == "Merchant"){
        p, err = Merchant.Init();
        if(err){
            log.fatal(err);
        }
    }else if(Type == "Lumberjack"){

    }else if (Type == "Miner"){

    }else if (Type == "non-Worker"){

    }else if (Type == "Chef"){

    }else if(Type == "other"){
        //TODO: this is a case where the actor is a drain on the economy of the market places.
    }
        //TODO: figure out more types
    return p;
}

