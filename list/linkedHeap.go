package list

type Element struct {
	parent, left, right, next, prev *Element
	Value interface{}
}

type LinkedHeap struct {
	root Element
	len int
}

func (h*LinkedHeap) Init() *LinkedHeap {
	h.root.prev = &h.root
	h.root.next = &h.root
	h.root.parent = &h.root
	h.root.left = &h.root
	h.root.right = &h.root
	h.len = 0
	return h
}

func (h*LinkedHeap) Parent(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.parent
}
func (h*LinkedHeap) Left(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.left
}
func (h*LinkedHeap) Right(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.right
}

func (h*LinkedHeap) Prev(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.prev
}
func (h*LinkedHeap) Next(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.next
}

func (h*LinkedHeap) Last()(interface{}) {
	return h.root.prev
}
func (h*LinkedHeap) Head()(interface{}) {
	return h.root.next
}

func (h*LinkedHeap) Valid(i interface{})(bool){
	iE := i.(*Element)
	return iE != &h.root && iE != nil
}

func (h *LinkedHeap) Swap(i interface{}, j interface{}) () {
	iE := i.(*Element)
	jE := j.(*Element)
	iE.Value, jE.Value = jE.Value, iE.Value
}

func (h *LinkedHeap) Key(i interface{}) (int) {
	iE := i.(*Element)
	return iE.Value.(int)
}

func (h *LinkedHeap) Len() (int) {
	return h.len
}

func (h *LinkedHeap) Pop() (i interface{}) {
	last := h.Last().(*Element)
	lastParent := last.parent
	lastPrev := last.prev
	lastPrev.next = &h.root
	h.root.prev = lastPrev
	if lastParent.left == last {
		lastParent.left = nil
	} else {
		lastParent.right = nil
	}
	h.len--
	return last.Value
}

func (h *LinkedHeap) Append(i interface{}) {
	newE := Element{Value:i}
	newE.next = &h.root
	last := h.Last().(*Element)
	lastParent := last.parent
	last.next = &newE
	newE.prev = last
	h.root.prev = &newE
	if lastParent.right == nil{
		lastParent.right = &newE
		newE.parent = lastParent
	} else {
		lastParent.next.left = &newE
		newE.parent = lastParent.next
	}
	h.len++
}
