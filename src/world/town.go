package world

type event struct {
	Type    string
	Message string
}

//Town - the town structure
type Town struct {
	eventChannel chan event
	Name         string
	//TODO: add a graph to represent the town.
}

func (t *Town) newTown() error {
	return nil
}
