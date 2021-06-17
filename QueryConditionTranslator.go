package hamdahlsqlite

import "github.com/standardcore/hamdahl"

type QueryConditionTranslator struct {
}

func (t *QueryConditionTranslator) GetSql(condition *hamdahl.QueryCondition) (string, error) {

}

func (t *QueryConditionTranslator) GetSqlByGroup(group *hamdahl.QueryConditionGroup) (string, error) {

}

func (t *QueryConditionTranslator) VisitBinaryCondition(c *hamdahl.BinaryQueryCondition) (string, error) {
	operatorName := ""
}
