package observer

type Observer interface {
	Update(string)
	GetState() string
}

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

type State interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll(state string)
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) Deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Item) NotifyAll(state string) {
	for _, observer := range i.observerList {
		observer.Update(state)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetState() == observer.GetState() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
