package behavioral

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Adapter(t *testing.T) {
	m := &Machine{state: GetLeaderApproveState()}
	assert.Equal(t, "LeaderApproveState", m.StateName())

	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.StateName())

	m.Reject()
	assert.Equal(t, "LeaderApproveState", m.StateName())

	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.StateName())
	m.Approval()
}
