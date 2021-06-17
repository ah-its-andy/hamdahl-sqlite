package hamdahlsqlite

import "github.com/standardcore/hamdahl"

type SqliteDbProvider struct {
}

func (p *SqliteDbProvider) Get(tableName string, id string) (*hamdahl.DataSet, error) {

}
func (p *SqliteDbProvider) GetAll(tableName string, ids []string) (*hamdahl.DataSet, error) {

}
func (p *SqliteDbProvider) Query(tableName string, condition *hamdahl.QueryCondition) (*hamdahl.DataSet, error) {

}

func (p *SqliteDbProvider) Add(tableName string, row *hamdahl.DataRow) {

}
func (p *SqliteDbProvider) AddRange(tableName string, rows []*hamdahl.DataRow) {

}

func (p *SqliteDbProvider) Update(tableName string, row *hamdahl.DataRow) {

}
func (p *SqliteDbProvider) UpdateRange(tableName string, rows []*hamdahl.DataRow) {

}

func (p *SqliteDbProvider) Remove(tableName string, id string) {

}
func (p *SqliteDbProvider) RemoveRange(tableName string, ids []string) {

}
