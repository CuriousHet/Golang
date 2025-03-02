package main

import (
	"fmt"
)

// SafetyPlacer interface defines safety mechanisms.
type SafetyPlacer interface {
	placeSafeties()
}

// RockClimber struct represents a climber and its dependencies.
type RockClimber struct {
	rocksClimbed int
	kind         string // "ice", "sand", "concrete"
	sp           SafetyPlacer
}

// NewRockClimber constructor initializes a RockClimber with a SafetyPlacer.
func NewRockClimber(kind string, sp SafetyPlacer) *RockClimber {
	return &RockClimber{
		kind: kind,
		sp:   sp,
	}
}

// IceSafetyPlacer implements SafetyPlacer for ice climbing.
type IceSafetyPlacer struct{}

func (sp IceSafetyPlacer) placeSafeties() {
	fmt.Println("Placing ICE safeties like ice screws and crampons...")
}

// SandSafetyPlacer implements SafetyPlacer for sand climbing.
type SandSafetyPlacer struct{}

func (sp SandSafetyPlacer) placeSafeties() {
	fmt.Println("Placing SAND safeties like pitons and hexes...")
}

// ConcreteSafetyPlacer implements SafetyPlacer for artificial surfaces.
type ConcreteSafetyPlacer struct{}

func (sp ConcreteSafetyPlacer) placeSafeties() {
	fmt.Println("Placing CONCRETE safeties like bolts and anchors...")
}

// NOPSafetyPlacer is a no-operation safety strategy.
type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("No safety measures in place. Free solo climbing!")
}

// climbRock simulates the climbing process.
func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	fmt.Printf("Climbed rock %d of type %s...\n", rc.rocksClimbed, rc.kind)

	// Place safety equipment every 5 climbs
	if rc.rocksClimbed%5 == 0 {
		rc.sp.placeSafeties()
	}
}

func main() {
	// Different climbing types with appropriate safety strategies
	climbers := []*RockClimber{
		NewRockClimber("ice", IceSafetyPlacer{}),
		NewRockClimber("sand", SandSafetyPlacer{}),
		NewRockClimber("concrete", ConcreteSafetyPlacer{}),
		NewRockClimber("free solo", NOPSafetyPlacer{}),
	}

	for _, climber := range climbers {
		fmt.Printf("\n--- Starting climb for %s climbing ---\n", climber.kind)
		for i := 0; i < 12; i++ {
			climber.climbRock()
		}
	}
}
