package runner

func New() *Runners {
	return &Runners{
		runners: make(map[string][]*Runner),
	}
}
