/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_busser_service from './vo-service'; // ./client-ts/src/vo-validation.ts to ./client-ts/src/vo-service.ts
import * as github_com_foomo_busser_table from './vo-table'; // ./client-ts/src/vo-validation.ts to ./client-ts/src/vo-table.ts
import * as github_com_foomo_busser_table_validation from './vo-validation'; // ./client-ts/src/vo-validation.ts to ./client-ts/src/vo-validation.ts
// github.com/foomo/busser/table/validation.Cell
export interface Cell {
	valid:boolean;
}
// github.com/foomo/busser/table/validation.Row
export interface Row {
	valid:boolean;
	cells:Record<github_com_foomo_busser_table.ColumnName,github_com_foomo_busser_table_validation.Cell>;
}
// github.com/foomo/busser/table/validation.Table
export interface Table {
	valid:boolean;
	rows:github_com_foomo_busser_table_validation.Row[];
}
// end of common js