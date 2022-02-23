/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_busser_service from './vo-service'; // ./client-ts/src/vo-validation.ts to ./client-ts/src/vo-service.ts
import * as github_com_foomo_busser_table from './vo-table'; // ./client-ts/src/vo-validation.ts to ./client-ts/src/vo-table.ts
import * as github_com_foomo_busser_table_validation from './vo-validation'; // ./client-ts/src/vo-validation.ts to ./client-ts/src/vo-validation.ts
// github.com/foomo/busser/table/validation.Cell
export interface Cell {
	valid:boolean;
	feedback:github_com_foomo_busser_table_validation.FeedbackEntry[];
}
// github.com/foomo/busser/table/validation.FeedbackEntry
export interface FeedbackEntry {
	level:github_com_foomo_busser_table_validation.FeedbackLevel;
	msg:string;
}
// github.com/foomo/busser/table/validation.FeedbackLevel
export enum FeedbackLevel {
	Error = "error",
	Valid = "valid",
	Warning = "warning",
}
// github.com/foomo/busser/table/validation.Row
export interface Row {
	valid:boolean;
	cells:Record<github_com_foomo_busser_table.ColumnName,github_com_foomo_busser_table_validation.Cell>|null;
	feedback:github_com_foomo_busser_table_validation.FeedbackEntry[];
}
// github.com/foomo/busser/table/validation.Table
export interface Table {
	valid:boolean;
	rows:github_com_foomo_busser_table_validation.Row[];
	feedback:github_com_foomo_busser_table_validation.FeedbackEntry[];
}
// end of common js