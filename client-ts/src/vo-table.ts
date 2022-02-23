/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_busser_service from './vo-service'; // ./client-ts/src/vo-table.ts to ./client-ts/src/vo-service.ts
import * as github_com_foomo_busser_table from './vo-table'; // ./client-ts/src/vo-table.ts to ./client-ts/src/vo-table.ts
import * as github_com_foomo_busser_table_validation from './vo-validation'; // ./client-ts/src/vo-table.ts to ./client-ts/src/vo-validation.ts
// github.com/foomo/busser/table.ColumnName
export type ColumnName = string
// github.com/foomo/busser/table.ID
export type ID = string
// github.com/foomo/busser/table.Map
export type Map = Record<github_com_foomo_busser_table.ID,github_com_foomo_busser_table.TableSummary[]>
// github.com/foomo/busser/table.Row
export type Row = Record<github_com_foomo_busser_table.ColumnName,string>
// github.com/foomo/busser/table.Table
export interface Table {
	id:github_com_foomo_busser_table.ID;
	version:github_com_foomo_busser_table.Version;
	timestamp:number;
	rows:github_com_foomo_busser_table.Row[];
	readErrors:Array<string>|null;
}
// github.com/foomo/busser/table.TableSummary
export interface TableSummary {
	id:github_com_foomo_busser_table.ID;
	timestamp:number;
	version:github_com_foomo_busser_table.Version;
	valid:boolean;
	committed:boolean;
}
// github.com/foomo/busser/table.Version
export type Version = string
// end of common js