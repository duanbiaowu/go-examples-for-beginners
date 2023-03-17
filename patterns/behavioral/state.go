package behavioral

type State interface {
	Approval(m *Machine)
	Reject(m *Machine)
	Name() string
}

type Machine struct {
	state State
}

type leaderApproveState struct{}

type financeApproveState struct{}

func (m *Machine) SetState(s State) {
	m.state = s
}

func (m *Machine) StateName() string {
	return m.state.Name()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

func (leaderApproveState) Approval(m *Machine) {
	m.SetState(GetFinanceApproveState())
}

func (leaderApproveState) Name() string {
	return "LeaderApproveState"
}

func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() State {
	return &leaderApproveState{}
}

func (financeApproveState) Approval(m *Machine) {

}

func (financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

func (financeApproveState) Name() string {
	return "FinanceApproveState"
}

func GetFinanceApproveState() State {
	return &financeApproveState{}
}
