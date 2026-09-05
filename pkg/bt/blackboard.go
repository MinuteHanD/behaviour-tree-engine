package bt

type MutableBlackboard interface {
	Blackboard
	Delete(key string)
	Clear()
}
