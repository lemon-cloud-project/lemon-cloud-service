package model

import (
	"crypto/sha1"
	"fmt"
	"github.com/jinzhu/inflection"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
	"unicode/utf8"
)

// CustomDbNamingStrategy tables, columi naming strategy
type CustomDbNamingStrategy struct {
	TablePrefix   string
	ColumnPrefix  string
	SingularTable bool
	NameReplacer  *strings.Replacer
}

// TableName convert string to table name
func (i *CustomDbNamingStrategy) TableName(str string) string {
	var tableName string
	if i.SingularTable {
		tableName = i.TablePrefix + i.toDBName(str)
	} else {
		tableName = i.TablePrefix + inflection.Plural(i.toDBName(str))
	}
	// 去掉表名中的_entity，然后转成大写
	return strings.ToUpper(strings.ReplaceAll(tableName, "_entity", ""))
}

// ColumnName convert string to column name
func (i *CustomDbNamingStrategy) ColumnName(table, column string) string {
	return strings.ToUpper(i.ColumnPrefix + i.toDBName(column))
}

// JoinTableName convert string to join table name
func (i *CustomDbNamingStrategy) JoinTableName(str string) string {
	var tableName string
	if strings.ToLower(str) == str {
		tableName = i.TablePrefix + str
	}
	if i.SingularTable {
		tableName = i.TablePrefix + i.toDBName(str)
	} else {
		tableName = i.TablePrefix + inflection.Plural(i.toDBName(str))
	}
	return strings.ToUpper(tableName)
}

// RelationshipFKName generate fk name for relation
func (i *CustomDbNamingStrategy) RelationshipFKName(rel schema.Relationship) string {
	return strings.Replace(fmt.Sprintf("fk_%s_%s", rel.Schema.Table, i.toDBName(rel.Name)), ".", "_", -1)
}

// CheckerName generate checker name
func (i *CustomDbNamingStrategy) CheckerName(table, column string) string {
	return strings.Replace(fmt.Sprintf("chk_%s_%s", table, column), ".", "_", -1)
}

// IndexName generate index name
func (i *CustomDbNamingStrategy) IndexName(table, column string) string {
	idxName := fmt.Sprintf("idx_%v_%v", table, i.toDBName(column))
	idxName = strings.Replace(idxName, ".", "_", -1)

	if utf8.RuneCountInString(idxName) > 64 {
		h := sha1.New()
		h.Write([]byte(idxName))
		bs := h.Sum(nil)

		idxName = fmt.Sprintf("idx%v%v", table, column)[0:56] + string(bs)[:8]
	}
	return strings.ToUpper(idxName)
}

var (
	smap sync.Map
	// https://github.com/golang/lint/blob/master/lint.go#L770
	commonInitialisms         = []string{"API", "ASCII", "CPU", "CSS", "Di", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
	commonInitialismsReplacer *strings.Replacer
)

func init() {
	var commonInitialismsForReplacer []string
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
}

func (i *CustomDbNamingStrategy) toDBName(name string) string {
	if name == "" {
		return ""
	} else if v, ok := smap.Load(name); ok {
		return v.(string)
	}

	if i.NameReplacer != nil {
		name = i.NameReplacer.Replace(name)
	}

	var (
		value                          = commonInitialismsReplacer.Replace(name)
		buf                            strings.Builder
		lastCase, nextCase, nextNumber bool // upper case == true
		curCase                        = value[0] <= 'Z' && value[0] >= 'A'
	)

	for i, v := range value[:len(value)-1] {
		nextCase = value[i+1] <= 'Z' && value[i+1] >= 'A'
		nextNumber = value[i+1] >= '0' && value[i+1] <= '9'

		if curCase {
			if lastCase && (nextCase || nextNumber) {
				buf.WriteRune(v + 32)
			} else {
				if i > 0 && value[i-1] != '_' && value[i+1] != '_' {
					buf.WriteByte('_')
				}
				buf.WriteRune(v + 32)
			}
		} else {
			buf.WriteRune(v)
		}

		lastCase = curCase
		curCase = nextCase
	}

	if curCase {
		if !lastCase && len(value) > 1 {
			buf.WriteByte('_')
		}
		buf.WriteByte(value[len(value)-1] + 32)
	} else {
		buf.WriteByte(value[len(value)-1])
	}
	ret := buf.String()
	smap.Store(name, ret)
	return ret
}
