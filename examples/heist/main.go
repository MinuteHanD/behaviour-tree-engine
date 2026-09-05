package main

import (
	"fmt"
	"time"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/blackboard"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/composite"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

type Actor struct {
	Name string
}

func main() {
	board := blackboard.New()

	
	board.Set("thief_distance", 100) 
	board.Set("cop_alerted", false)
	board.Set("cop_pos", 50)         
	board.Set("has_loot", false)
	board.Set("escaped", false)
	board.Set("busted", false)

	thief := &Actor{Name: "(Thief)"}
	cop := &Actor{Name: "(Cop)"}

	// THIEF BEHAVIOUR
	thiefTree := bt.NewTree(composite.NewSelector(
		
		leaf.NewCondition(func(ctx *bt.Context) bool {
			busted, _ := blackboard.Get[bool](ctx.Blackboard, "busted")
			if busted {
				fmt.Printf("[%s] 'Failed, I'm going to jail...'\n", thief.Name)
				return true 
			}
			return false
		}),

		
		leaf.NewCondition(func(ctx *bt.Context) bool {
			escaped, _ := blackboard.Get[bool](ctx.Blackboard, "escaped")
			if escaped {
				fmt.Printf("[%s] 'Success I'm rich'\n", thief.Name)
				return true
			}
			return false
		}),

		
		composite.NewSequence(
			leaf.NewCondition(func(ctx *bt.Context) bool {
				hasLoot, _ := blackboard.Get[bool](ctx.Blackboard, "has_loot")
				return hasLoot
			}),
			leaf.NewAction(func(ctx *bt.Context) bt.Status {
				dist, _ := blackboard.Get[int](ctx.Blackboard, "thief_distance")
				dist += 25
				ctx.Blackboard.Set("thief_distance", dist)
				fmt.Printf("[%s] *Sprints towards the exit* (Distance to exit: %d)\n", thief.Name, 100-dist)

				if dist >= 100 {
					ctx.Blackboard.Set("escaped", true)
					return bt.Success
				}
				return bt.Running
			}),
		),

		composite.NewSequence(
			leaf.NewCondition(func(ctx *bt.Context) bool {
				dist, _ := blackboard.Get[int](ctx.Blackboard, "thief_distance")
				return dist <= 0
			}),
			leaf.NewAction(func(ctx *bt.Context) bt.Status {
				fmt.Printf("[%s] *In process of Cracking the safe*\n", thief.Name)
				return bt.Success
			}),
			leaf.NewWait(1), // Takes 1 tick to crack
			leaf.NewAction(func(ctx *bt.Context) bt.Status {
				fmt.Printf("[%s] 'Got the loot' *Alarm is Triggered*\n", thief.Name)
				ctx.Blackboard.Set("has_loot", true)
				ctx.Blackboard.Set("cop_alerted", true)
				return bt.Success
			}),
		),

		composite.NewSequence(
			leaf.NewAction(func(ctx *bt.Context) bt.Status {
				dist, _ := blackboard.Get[int](ctx.Blackboard, "thief_distance")
				copPos, _ := blackboard.Get[int](ctx.Blackboard, "cop_pos")

				if abs(dist-copPos) <= 15 {
					fmt.Printf("[%s] 'cop is close' *Hides in shadows*\n", thief.Name)
					return bt.Running
				}

				
				dist -= 15
				ctx.Blackboard.Set("thief_distance", dist)
				fmt.Printf("[%s] *Crouch walks* (Distance to vault: %d)\n", thief.Name, dist)
				
				return bt.Running
			}),
		),
	))

	// COP BEHAVIOUR
	copTree := bt.NewTree(composite.NewSelector(
		// Arrest Thief (if close enough)
		composite.NewSequence(
			leaf.NewCondition(func(ctx *bt.Context) bool {
				busted, _ := blackboard.Get[bool](ctx.Blackboard, "busted")
				return !busted
			}),
			leaf.NewCondition(func(ctx *bt.Context) bool {
				thiefDist, _ := blackboard.Get[int](ctx.Blackboard, "thief_distance")
				copPos, _ := blackboard.Get[int](ctx.Blackboard, "cop_pos")
				return abs(thiefDist-copPos) < 10
			}),
			leaf.NewAction(func(ctx *bt.Context) bt.Status {
				fmt.Printf("[%s] 'Got you thief' *Tackles thief*\n", cop.Name)
				ctx.Blackboard.Set("busted", true)
				return bt.Success
			}),
		),

		// Chase/Investigate (If alerted)
		composite.NewSequence(
			leaf.NewCondition(func(ctx *bt.Context) bool {
				alerted, _ := blackboard.Get[bool](ctx.Blackboard, "cop_alerted")
				return alerted
			}),
			leaf.NewAction(func(ctx *bt.Context) bt.Status {
				thiefDist, _ := blackboard.Get[int](ctx.Blackboard, "thief_distance")
				copPos, _ := blackboard.Get[int](ctx.Blackboard, "cop_pos")

				
				if copPos < thiefDist {
					copPos += 30
				} else {
					copPos -= 30
				}
				ctx.Blackboard.Set("cop_pos", copPos)
				fmt.Printf("[%s] 'STOP RIGHT THERE SCUM!' *Runs after thief*\n", cop.Name)
				return bt.Running
			}),
		),

		// Normal Patrol
		leaf.NewAction(func(ctx *bt.Context) bt.Status {
			copPos, _ := blackboard.Get[int](ctx.Blackboard, "cop_pos")
			patrolDir, exists := blackboard.Get[int](ctx.Blackboard, "patrol_dir")
			if !exists {
				patrolDir = -10 
			}
			
			
			if copPos <= -20 {
				patrolDir = 10
			} else if copPos >= 80 {
				patrolDir = -10
			}
			
			copPos += patrolDir
			ctx.Blackboard.Set("cop_pos", copPos)
			ctx.Blackboard.Set("patrol_dir", patrolDir)
			
			fmt.Printf("[%s] *idle walk* 'Quiet night tonight' (Pos: %d)\n", cop.Name, copPos)
			return bt.Success
		}),
	))

	thiefCtx := bt.NewContext(thief, board)
	copCtx := bt.NewContext(cop, board)

	fmt.Println("--- STARTING THE HEIST SIMULATION ---")
	fmt.Println("Thief is trying to sneak to the vault (distance 0).")
	fmt.Println("Cop is patrolling nearby.")
	fmt.Println("-------------------------------------")

	// Game Loop
	for frame := 1; frame <= 30; frame++ {
		fmt.Printf("\n[Tick %d]\n", frame)

		
		thiefTree.Tick(thiefCtx)
		copTree.Tick(copCtx)

		
		busted, _ := blackboard.Get[bool](board, "busted")
		escaped, _ := blackboard.Get[bool](board, "escaped")

		if busted {
			fmt.Println("\n--- OUTCOME: THIEF WAS BUSTED! JUSTICE PREVAILS. ---")
			break
		}
		if escaped {
			fmt.Println("\n--- OUTCOME: THIEF ESCAPED WITH THE LOOT! THE PERFECT CRIME. ---")
			break
		}

		time.Sleep(400 * time.Millisecond) 
	}
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
