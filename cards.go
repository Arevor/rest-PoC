package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

func getCard(id string) (Card, error) {
	res := Card{}

	var cardid, priority, severity, casenumber, boardcol int
	var title, cardtype, desc1, desc2, desc3, assignedto string

	err := db.QueryRow(`SELECT cardid, title, cardtype, desc1, desc2, desc3, priority, severity, assignedto, 
			casenumber, boardcol FROM cards where cardid = $1`, id).Scan(&cardid, &title, &cardtype, &desc1, &desc2,
		&desc3, &priority, &severity, &assignedto, &casenumber, &boardcol)
	if err == nil {
		res = Card{CardID: cardid, Title: title, CardType: cardtype, DescriptionBox1: desc1, DescriptionBox2: desc2,
			DescriptionBox3: desc3, Priority: priority, Severity: severity, AssignedTo: assignedto,
			CaseNumber: casenumber, Column: boardcol}
	}

	return res, err
}

func getCards() ([]Card, error) {
	Cards := []Card{}

	rows, err := db.Query(`SELECT cardid, title, cardtype, desc1, desc2, desc3, priority, severity, assignedto, 
			casenumber, boardcol FROM cards`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var cardid, priority, severity, casenumber, boardcol int
			var title, cardtype, desc1, desc2, desc3, assignedto string

			err = rows.Scan(&cardid, &title, &cardtype, &desc1, &desc2, &desc3,
				&priority, &severity, &assignedto, &casenumber, &boardcol)
			if err == nil {
				currentCard := Card{CardID: cardid, Title: title, CardType: cardtype, DescriptionBox1: desc1,
					DescriptionBox2: desc2, DescriptionBox3: desc3, Priority: priority, Severity: severity,
					AssignedTo: assignedto, CaseNumber: casenumber, Column: boardcol}

				Cards = append(Cards, currentCard)
			} else {
				return Cards, err
			}
		}
	} else {
		return Cards, err
	}

	return Cards, err
}

func insertCard(title, ct, d1, d2, d3 string, prio, sev int, assto string, casenum, col int) (int, error) {
	var CardID int
	//TODO 'elegant' way to format this insert
	err := db.QueryRow(`INSERT INTO cards(title, cardtype, desc1, desc2, desc3, 
		priority, severity, assignedto, casenumber, boardcol)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING cardid`, title, ct, d1, d2, d3,
		prio, sev, assto, casenum, col).Scan(&CardID)

	if err != nil {
		fmt.Println("SQL error:", err)
		return 0, err
	}
	return CardID, err
}

func updateCard(cardid int, title, ct, d1, d2, d3 string, prio, sev int, assto string, casenum, col int) (int, error) {
	res, err := db.Exec(`UPDATE cards set title=$2, cardtype=$3 , desc1=$4, desc2=$5, desc3=$6, priority=$7,
		severity=$8, assignedto=$9, casenumber=$10, boardcol=$11 where cardid=$1`, cardid, title, ct, d1,
		d2, d3, prio, sev, assto, casenum, col)
	if err != nil {
		fmt.Println("SQL error:", err)
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsUpdated), err
}

func removeCard(id string) (int, error) {
	res, err := db.Exec(`delete from cards where cardid = $1`, id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		fmt.Println("SQL Error:", err)
		return 0, err
	}

	return int(rowsDeleted), nil
}
