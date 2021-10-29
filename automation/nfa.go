package automation

import (
	"errors"
	"github.com/emirpasic/gods/sets/linkedhashset"
	"strconv"
)

type transitionInput struct{
	srcState int
	input interface{}
}
type destState struct {
	set *linkedhashset.Set
	reverse bool
} //int

type NFA struct{
	initState int
	currentState *linkedhashset.Set //int
	totalStates *linkedhashset.Set //int
	finalStates *linkedhashset.Set //int
	transition map[transitionInput]destState
	inputMap *linkedhashset.Set //string
}

func NewNFA(initState int, isFinal bool)(retNFA *NFA){

	retNFA = &NFA{
		transition: make(map[transitionInput]destState),
		inputMap: linkedhashset.New(),
		initState: initState,
		currentState: linkedhashset.New(),
	}
	retNFA.currentState.Add(initState)
	retNFA.AddState(initState, isFinal)
	return
}

func (d *NFA) AddState(state int, isFinal bool){
	d.totalStates.Add(state)
	if isFinal{
		d.finalStates.Add(state)
	}
}

func (d *NFA)AddTransition(srcState int, input interface{}, reverse bool, dstStateList ...int)( error){
	if !d.totalStates.Contains(srcState){
		return errors.New("No such state " + strconv.Itoa(srcState))
	}

	d.inputMap.Add(input)
	var temp destState
	temp.set = linkedhashset.New()
	for _, des := range dstStateList{
		temp.set.Add(des)
	}
	temp.reverse = reverse
	d.transition[transitionInput{srcState: srcState, input: input}] = temp
	return nil
}

func (d *NFA)removeTransition(srcState int, input interface{})(error){
	if !d.totalStates.Contains(srcState){
		return errors.New("No such state " + strconv.Itoa(srcState))
	}
	temp := transitionInput{srcState: srcState, input: input}
	if _ , ok := d.transition[temp]; ok{
		delete(d.transition, temp )
	}else{
		return errors.New("No such input")
	}

	return nil
}

func (d *NFA)Update(input interface{})(ret bool){
	//TODO reverse
	ret = false
	updateCurrentState := linkedhashset.New()
	for current, _ := range d.currentState.Values(){
		intputTrans := transitionInput{srcState: current, input: input}
		if valMap, ok := d.transition[intputTrans]; ok{

			for dst, _ := range valMap.set.Values(){
				if d.finalStates.Contains(dst){
					ret = true
				}else{
					updateCurrentState.Add(dst)
				}
			}
		}else{
			//remove in curren state
			//TODO
		}
		d.currentState = updateCurrentState

	}
	return
}