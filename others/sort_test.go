package others

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"go.etcd.io/etcd/raft/v3/quorum"
)

// MajorityConfig is a set of IDs that uses majority quorums to make decisions.
type MajorityConfig map[uint64]struct{}

// Index is a Raft log position.
type Index uint64

// func (c MajorityConfig) sort() {

// 	type tup struct {
// 		id  uint64
// 		idx Index
// 		ok  bool // idx found?
// 		bar int  // length of bar displayed for this tup
// 	}

// 	n := len(c)
// 	info := make([]tup, 0, n)

// 	for id := range c {
// 		idx, ok := l.AckedIndex(id)
// 		info = append(info, tup{id: id, idx: idx, ok: ok})
// 	}

// }

// Config reflects the configuration tracked in a ProgressTracker.
type Config struct {
	Voters quorum.JointConfig
	// AutoLeave is true if the configuration is joint and a transition to the
	// incoming configuration should be carried out automatically by Raft when
	// this is possible. If false, the configuration will be joint until the
	// application initiates the transition manually.
	AutoLeave bool
	// Learners is a set of IDs corresponding to the learners active in the
	// current configuration.
	//
	// Invariant: Learners and Voters does not intersect, i.e. if a peer is in
	// either half of the joint config, it can't be a learner; if it is a
	// learner it can't be in either half of the joint config. This invariant
	// simplifies the implementation since it allows peers to have clarity about
	// its current role without taking into account joint consensus.
	Learners map[uint64]struct{}
	// When we turn a voter into a learner during a joint consensus transition,
	// we cannot add the learner directly when entering the joint state. This is
	// because this would violate the invariant that the intersection of
	// voters and learners is empty. For example, assume a Voter is removed and
	// immediately re-added as a learner (or in other words, it is demoted):
	//
	// Initially, the configuration will be
	//
	//   voters:   {1 2 3}
	//   learners: {}
	//
	// and we want to demote 3. Entering the joint configuration, we naively get
	//
	//   voters:   {1 2} & {1 2 3}
	//   learners: {3}
	//
	// but this violates the invariant (3 is both voter and learner). Instead,
	// we get
	//
	//   voters:   {1 2} & {1 2 3}
	//   learners: {}
	//   next_learners: {3}
	//
	// Where 3 is now still purely a voter, but we are remembering the intention
	// to make it a learner upon transitioning into the final configuration:
	//
	//   voters:   {1 2}
	//   learners: {3}
	//   next_learners: {}
	//
	// Note that next_learners is not used while adding a learner that is not
	// also a voter in the joint config. In this case, the learner is added
	// right away when entering the joint configuration, so that it is caught up
	// as soon as possible.
	LearnersNext map[uint64]struct{}
}

func TestConfig(t *testing.T) {

	// c := Config{
	// 	Voters: quorum.JointConfig{
	// 		quorum.MajorityConfig{},
	// 		nil, // only populated when used
	// 	},
	// 	Learners:     nil, // only populated when used
	// 	LearnersNext: nil, // only populated when used
	// }

	m := MajorityConfig{}
	m.String()
}

/// etcd magority.go
func (c MajorityConfig) String() string {
	sl := make([]uint64, 0, len(c))
	for id := range c {
		sl = append(sl, id)
	}
	sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
	var buf strings.Builder
	buf.WriteByte('(')
	for i := range sl {
		if i > 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprint(&buf, sl[i])
	}
	buf.WriteByte(')')
	return buf.String()
}
