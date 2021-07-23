package main

import (
	"fmt"
	"strconv"
	"strings"
)

// função para validar um CPF - antes da refatoração
func validate(str string) bool {
	if str != "" {
		if len(str) >= 11 || len(str) <= 14 {
			str = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(str, ".", ""), "-", ""), " ", "")
			if strings.Count(str, string(str[0])) == len(str) {
				return false
			}
			var d1 int
			var d2 int
			var dg1 int
			var dg2 int
			var rest int
			var digito int
			var nDigResult string
			d1 = 0
			d2 = 0
			dg1 = 0
			dg2 = 0
			rest = 0

			for nCont := 1; nCont < len(str)-1; nCont++ {
				// if (isNaN(parseInt(str.substring(nCount -1, nCount)))) {
				// 	return false;
				// } else {
					digito, _ = strconv.Atoi(str[nCont-1 : nCont])
					d1 = d1 + (11-nCont)*digito
					d2 = d2 + (12-nCont)*digito
				//}
			}

			rest = (d1 % 11)
			if rest < 2 {
				dg1 = 0
			} else {
				dg1 = 11 - rest
			}
			d2 += 2 * dg1
			rest = (d2 % 11)
			if rest < 2 {
				dg2 = 0
			} else {
				dg2 = 11 - rest
			}
			var nDigVerific = str[len(str)-2:]
			nDigResult = fmt.Sprintf("%d%d", dg1, dg2)
			return nDigVerific == nDigResult
		} else {
			return false
		}
	} else {
		return false
	}
}
