package forum

const serialVersionUID = 1

type PostID struct {
	*AbstractId
}

func NewPostId(anId string) (rcvr PostID) {
	rcvr = &PostID{}
	rcvr.AbstractId = NewAbstractId(anId)
	return
}
func NewPostId2() (rcvr PostID) {
	rcvr = &PostID{}
	rcvr.AbstractId = NewAbstractId()
	return
}

func (rcvr PostID) hashOddValue() int {
	return 11735
}

func (rcvr PostID) hashPrimeValue() int {
	return 37
}
