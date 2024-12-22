package llds

import (
	"fmt"

	"github.com/google/uuid"
)

//1. User add all people in its friend list.
//2. User creates group with people.
//3. User adds expense with people
//4. User adds expense with group.

type userInfo struct {
	id      string
	name    string
	address string
}

type expense struct {
	id          string
	name        string
	amount      float64
	expenseType string
}

type userGroup struct {
	id           string
	name         string
	members      []*userInfo
	balanceSheet map[string]*expense
}

type user struct {
	personalDetails userInfo
	friends         map[string]*userInfo
	groups          map[string]*userGroup
	expenseByType   map[string]map[string]map[string]*expense //PERSONAL, GROUP
}

type ExpenseByTypeDto struct {
}

/*
PERSONAL: {
"ID":
{
	"MEM1": "AMOUNT1"
}


GROUP : {
"ID" : {
"MEM1": "AMOUNT1",
},
{
"MEM2": "AMOUNT2",
}
}

*/

type splitWise struct {
	users []*user
}

type createGroupRequestDto struct {
	name    string
	friends []*userInfo
}

type createPersonalExpenseRequestDto struct {
	name        string
	expenseMap  map[string]*expense
	expenseType string
	groupName   string
}

func (u *user) createGroup(requestDto createGroupRequestDto) (*userGroup, error) {
	id := uuid.NewString()
	group := &userGroup{
		id:           id,
		name:         requestDto.name,
		members:      requestDto.friends,
		balanceSheet: map[string]*expense{},
	}

	u.groups[requestDto.name] = group
	u.expenseByType["GROUP"] = map[string]map[string]*expense{
		requestDto.name: map[string]*expense{},
	}
	return group, nil
}

func (u *user) createExpense(requestDto createPersonalExpenseRequestDto) (*userGroup, error) {
	if requestDto.expenseType == "GROUP" {
		for key, val := range requestDto.expenseMap {
			if val.expenseType == "DEBIT" {
				_, ok := u.groups[requestDto.groupName].balanceSheet[key]
				if !ok {
					u.groups[requestDto.groupName].balanceSheet[key] = &expense{
						amount: -val.amount,
					}
				} else {
					u.groups[requestDto.groupName].balanceSheet[key].amount = u.groups[requestDto.groupName].balanceSheet[key].amount - val.amount
				}

				_, ok = u.expenseByType["GROUP"][requestDto.groupName][key]
				if !ok {
					u.expenseByType["GROUP"][requestDto.groupName][key] = &expense{
						amount: -val.amount,
					}
				} else {
					u.expenseByType["GROUP"][requestDto.groupName][key].amount = u.expenseByType["GROUP"][requestDto.groupName][key].amount - val.amount
				}

			} else {
				_, ok := u.groups[requestDto.groupName].balanceSheet[key]
				if !ok {
					u.groups[requestDto.groupName].balanceSheet[key] = &expense{
						amount: val.amount,
					}
				} else {
					u.groups[requestDto.groupName].balanceSheet[key].amount = u.groups[requestDto.groupName].balanceSheet[key].amount + val.amount
				}

				_, ok = u.expenseByType["GROUP"][requestDto.groupName][key]
				if !ok {
					u.expenseByType["GROUP"][requestDto.groupName][key] = &expense{
						amount: -val.amount,
					}
				} else {
					u.expenseByType["GROUP"][requestDto.groupName][key].amount = u.expenseByType["GROUP"][requestDto.groupName][key].amount + val.amount
				}
			}
		}
	} else {
		for key, val := range requestDto.expenseMap {
			if val.expenseType == "DEBIT" {
				_, ok := u.expenseByType["PERSONAL"][key][key]
				if !ok {
					u.expenseByType["PERSONAL"][key][key] = &expense{
						amount: -val.amount,
					}
				} else {
					u.expenseByType["PERSONAL"][key][key].amount = u.expenseByType["PERSONAL"][requestDto.groupName][key].amount - val.amount
				}
			} else {
				_, ok := u.expenseByType["PERSONAL"][key][key]
				if !ok {
					u.expenseByType["PERSONAL"][key][key] = &expense{
						amount: val.amount,
					}
				} else {
					u.expenseByType["PERSONAL"][key][key].amount = u.expenseByType["PERSONAL"][requestDto.groupName][key].amount + val.amount
				}

			}
		}
	}
	return nil, nil
}

func Driver6() {
	id1 := uuid.NewString()
	id2 := uuid.NewString()
	id3 := uuid.NewString()
	id4 := uuid.NewString()
	userInfo2 := &userInfo{
		id:      id2,
		name:    "shuvam",
		address: "patiala",
	}
	userInfo3 := &userInfo{
		id:      id3,
		name:    "saryam",
		address: "patna",
	}

	userInfo4 := &userInfo{
		id:      id4,
		name:    "john",
		address: "patna",
	}
	//create user
	user := user{
		personalDetails: userInfo{
			id:      id1,
			name:    "shivam",
			address: "abhs",
		},
		friends: map[string]*userInfo{
			"shuvam": userInfo2,
			"saryam": userInfo3,
			"john":   userInfo4,
		},
		groups: map[string]*userGroup{},
		expenseByType: map[string]map[string]map[string]*expense{
			"GROUP": {},
			"PERSONAL": {
				"shuvam": {},
				"saryam": {},
				"john":   {},
			},
		},
	}

	user.createGroup(createGroupRequestDto{
		name: "Goa Trip",
		friends: []*userInfo{
			userInfo2, userInfo4,
		},
	})

	user.createExpense(createPersonalExpenseRequestDto{
		name: "Anjuna Ride",
		expenseMap: map[string]*expense{
			"john": {
				id:          uuid.NewString(),
				name:        "Anjuna Ride",
				amount:      89,
				expenseType: "DEBIT",
			},
			"shuvam": {
				id:          uuid.NewString(),
				name:        "Anjuna Ride",
				amount:      89,
				expenseType: "DEBIT",
			},
		},
		expenseType: "GROUP",
		groupName:   "Goa Trip",
	})

	user.createExpense(createPersonalExpenseRequestDto{
		name: "Lunch",
		expenseMap: map[string]*expense{
			"john": {
				id:          uuid.NewString(),
				name:        "Anjuna Ride",
				amount:      189,
				expenseType: "CREDIT",
			},
			"shuvam": {
				id:          uuid.NewString(),
				name:        "Anjuna Ride",
				amount:      189,
				expenseType: "DEBIT",
			},
		},
		expenseType: "GROUP",
		groupName:   "Goa Trip",
	})

	user.createExpense(createPersonalExpenseRequestDto{
		name: "Lunch",
		expenseMap: map[string]*expense{
			"john": {
				id:          uuid.NewString(),
				name:        "Anjuna Ride",
				amount:      1189,
				expenseType: "CREDIT",
			},
			"saryam": {
				id:          uuid.NewString(),
				name:        "Anjuna Ride",
				amount:      189,
				expenseType: "DEBIT",
			},
		},
		expenseType: "PERSONAL",
		groupName:   "Goa Trip",
	})

	fmt.Println("expense created successfully!!!")
}
