package idiom

type Message struct {
}

type Subscription struct {
	ch    chan<- Message
	Inbox chan Message
}

func (s *Subscription) Publish(msg Message) error {
	//if _, ok := <-s.ch; !ok {
	//
	//}

	s.ch <- msg
	return nil
}
